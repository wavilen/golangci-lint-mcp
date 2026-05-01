package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
	"github.com/wavilen/golangci-lint-mcp/internal/linttypes"

	"github.com/mark3labs/mcp-go/mcp"
)

// Type aliases — single source of truth in internal/linttypes.
type LintIssue = linttypes.LintIssue
type LintJSONResult = linttypes.LintJSONResult

// Function redirects — delegate to linttypes implementations.
// Kept for backward compatibility with existing test code.

// ExtractRule extracts the rule name from a lint issue text.
func ExtractRule(text string) string { return linttypes.ExtractRule(text) }

// DeduplicateIssues removes duplicate (linter, rule) pairs from the issue slice.
func DeduplicateIssues(issues []LintIssue) []LintIssue {
	return linttypes.DeduplicateIssues(issues)
}

func buildLinterBreakdown(unique []LintIssue) string {
	linterCounts := make(map[string]int)
	for _, issue := range unique {
		linterCounts[issue.FromLinter]++
	}

	type linterEntry struct {
		name  string
		count int
	}
	var sortedEntries = make([]linterEntry, 0, len(linterCounts))
	for name, count := range linterCounts {
		sortedEntries = append(sortedEntries, linterEntry{name, count})
	}
	sort.Slice(sortedEntries, func(left, right int) bool {
		if sortedEntries[left].count != sortedEntries[right].count {
			return sortedEntries[left].count > sortedEntries[right].count
		}
		return sortedEntries[left].name < sortedEntries[right].name
	})

	parts := make([]string, 0, len(sortedEntries))
	for _, entry := range sortedEntries {
		parts = append(parts, fmt.Sprintf("%s (%d)", entry.name, entry.count))
	}
	return strings.Join(parts, ", ")
}

func makeParseHandler(
	store *guides.Store,
	opts Options,
) func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output, err := req.RequireString("output")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'output': %v", err)), nil
		}
		output = strings.TrimSpace(output)
		if output == "" {
			return mcp.NewToolResultError("parameter 'output' must not be empty"), nil
		}

		result, parseErr := linttypes.ParseLintOutput(output)
		if parseErr != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf(
					"invalid JSON: %v. Try golangci_lint_run(path=\"./pkg/...\") "+
						"for pre-parsed results, or golangci_lint_guide(linter=\"<name>\") "+
						"for individual diagnostics.",
					parseErr,
				)), nil
		}

		if len(result.Issues) == 0 {
			return mcp.NewToolResultText("No issues found in the golangci-lint output."), nil
		}

		// Unified pipeline: analyze → build response (D-03)
		strategyResult := AnalyzeStrategy(result.Issues)
		return mcp.NewToolResultText(
			BuildResponse(strategyResult, ResponseConfig{
				Store: store, Opts: opts, IncludeGuidance: true,
			})), nil
	}
}

// relatedEntry tracks a candidate related linter with its best fix hint and score.
type relatedEntry struct {
	ref   string // original ref string (e.g., "govet" or "gosec/G304")
	hint  string
	score int
}

// buildRelatedContext builds a consolidated Related Context section from all
// unique issues. It collects related refs, deduplicates against primary
// diagnostics, keeps best keyword-overlap hint per related linter, and
// enforces max 5 entries / ~500 byte budget.
//
//nolint:gocognit,funlen // Single-pass related context builder: collect→dedup→sort→budget. Splitting would add indirection without reducing actual complexity.
func buildRelatedContext(unique []LintIssue, store *guides.Store) string {
	// Build primary diagnostic set for exclusion
	primarySet := make(map[string]bool)
	for _, issue := range unique {
		rule := linttypes.ExtractRule(issue.Text)
		if rule != "" {
			primarySet[issue.FromLinter+"/"+rule] = true
		}
		primarySet[issue.FromLinter] = true
	}

	// Collect related entries with best hint per ref
	bestEntries := make(map[string]*relatedEntry)

	for _, issue := range unique {
		linter := issue.FromLinter
		rule := linttypes.ExtractRule(issue.Text)

		var guide *guides.Guide
		if rule != "" {
			if g, found := store.Lookup(linter, rule); found {
				guide = g
			}
		}
		if guide == nil {
			if g, found := store.Lookup(linter, ""); found {
				guide = g
			}
		}
		if guide == nil || len(guide.Related) == 0 {
			continue
		}

		for _, ref := range guide.Related {
			// Skip if ref is in primary diagnostic set
			if primarySet[ref] {
				continue
			}
			refLinter, refRule := parseRelatedRef(ref)
			relatedGuide, found := store.Lookup(refLinter, refRule)
			if !found && refRule != "" {
				relatedGuide, found = store.Lookup(refLinter, "")
			}
			if !found {
				continue
			}

			fixHint := guides.BestPatternBullet(relatedGuide.Patterns, issue.Text)
			if fixHint == "" {
				continue
			}

			score := guides.KeywordOverlapScore(issue.Text, fixHint)
			if existing, ok := bestEntries[ref]; ok {
				if score > existing.score {
					existing.hint = fixHint
					existing.score = score
				}
			} else {
				bestEntries[ref] = &relatedEntry{ref: ref, hint: fixHint, score: score}
			}
		}
	}

	if len(bestEntries) == 0 {
		return ""
	}

	// Sort by score descending, take top 5
	sorted := make([]*relatedEntry, 0, len(bestEntries))
	for _, entry := range bestEntries {
		sorted = append(sorted, entry)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].score != sorted[j].score {
			return sorted[i].score > sorted[j].score
		}
		return sorted[i].ref < sorted[j].ref
	})
	if len(sorted) > maxRelatedEntries {
		sorted = sorted[:maxRelatedEntries]
	}

	// Build section
	lines := make([]string, 0, len(sorted))
	for _, entry := range sorted {
		lines = append(lines, fmt.Sprintf("- %s: %s", entry.ref, entry.hint))
	}
	section := "<related_context>\n" + strings.Join(lines, "\n") + "\n</related_context>"

	// Enforce byte budget
	for len(section) > maxRelatedBytes && len(lines) > 0 {
		lines = lines[:len(lines)-1]
		section = "<related_context>\n" + strings.Join(lines, "\n") + "\n</related_context>"
	}

	if len(lines) == 0 {
		return ""
	}

	return section
}

func resolveGuide(store *guides.Store, linter, rule string) (string, bool) {
	if rule != "" {
		guide, found := store.Lookup(linter, rule)
		if found {
			return guide.RawBody, true
		}
	}

	guide, found := store.Lookup(linter, "")
	if found {
		return guide.RawBody, true
	}

	return "", false
}

// stripGuideHeading removes the leading markdown h1 heading (e.g., "# errcheck" or
// "# gocritic: badCall") from the guide body. The writeGuideForIssue function writes
// its own h2 heading ("## linter: rule"), so the guide's h1 is redundant and produces
// duplicate headers in the output.
func stripGuideHeading(body string) string {
	if !strings.HasPrefix(body, "# ") {
		return body
	}
	// Find end of first line using strings.Cut (modernize:stringscut)
	_, after, ok := strings.Cut(body, "\n")
	if !ok {
		return ""
	}
	// Skip any trailing blank lines after the heading
	return strings.TrimLeft(after, "\n")
}

func writeGuideForIssue(builder *strings.Builder, store *guides.Store, opts Options, issue LintIssue) {
	linter := issue.FromLinter
	rule := linttypes.ExtractRule(issue.Text)

	body, ok := resolveGuide(store, linter, rule)
	if ok {
		if rule != "" {
			fmt.Fprintf(builder, "## %s: %s\n\n", linter, rule)
		} else {
			fmt.Fprintf(builder, "## %s\n\n", linter)
		}
		body = stripGuideHeading(body)
		body = stripRelatedTag(maybeAppendGosecAI(body, opts, linter))
		builder.WriteString(body)
		return
	}

	rules := store.ListRules(linter)
	if len(rules) > 0 {
		fmt.Fprintf(builder, "## %s: %s\n\nNo guide found for rule %q of linter %q. Available rules: %s",
			linter, rule, rule, linter, strings.Join(rules, ", "))
		return
	}

	msg := fmt.Sprintf("## %s\n\nUnknown linter %q.", linter, linter)
	suggestion := store.Suggest(linter)
	if suggestion != "" {
		msg = fmt.Sprintf("## %s\n\nUnknown linter %q. Did you mean %q?", linter, linter, suggestion)
	}
	builder.WriteString(msg)
}
