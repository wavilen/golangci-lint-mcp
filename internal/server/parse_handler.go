package server

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

type lintJSONResult struct {
	Issues []lintIssue `json:"Issues"`
}

type lintIssue struct {
	FromLinter string `json:"FromLinter"`
	Text       string `json:"Text"`
	Pos        struct {
		Filename string `json:"Filename"`
		Line     int    `json:"Line"`
		Column   int    `json:"Column"`
	} `json:"Pos"`
}

type diagnostic struct {
	linter string
	rule   string
}

func extractRule(text string) string {
	before, _, found := strings.Cut(text, ": ")
	if !found {
		return ""
	}
	return strings.TrimSpace(before)
}

func deduplicateIssues(issues []lintIssue) []lintIssue {
	seen := make(map[diagnostic]bool)
	var unique []lintIssue
	for _, issue := range issues {
		rule := extractRule(issue.Text)
		key := diagnostic{linter: issue.FromLinter, rule: rule}
		if !seen[key] {
			seen[key] = true
			unique = append(unique, issue)
		}
	}
	return unique
}

func buildLinterBreakdown(unique []lintIssue) string {
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

		firstLine := output
		if before, _, found := strings.Cut(output, "\n"); found {
			firstLine = before
		}
		var result lintJSONResult
		err = json.Unmarshal([]byte(firstLine), &result)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid JSON: %v", err)), nil
		}

		if len(result.Issues) == 0 {
			return mcp.NewToolResultText("No issues found in the golangci-lint output."), nil
		}

		unique := deduplicateIssues(result.Issues)
		packages := extractPackagesFromIssues(unique)
		strategyName, strategyReason := recommendStrategy(len(unique), len(packages))

		var builder strings.Builder
		fmt.Fprintf(&builder, "## Summary\n\n- Unique diagnostics: %d\n- Strategy: %s (%s)\n",
			len(unique), strategyName, strategyReason)

		if len(packages) > 1 {
			fmt.Fprintf(&builder, "\n## Package Breakdown\n\n%s\n", buildPackageBreakdown(packages))
		}

		fmt.Fprintf(&builder, "\n- Breakdown: %s\n\n---\n\n", buildLinterBreakdown(unique))

		for idx, issue := range unique {
			if idx > 0 {
				builder.WriteString("\n---\n\n")
			}
			writeGuideForIssue(&builder, store, opts, issue)
		}

		relatedSection := buildRelatedContext(unique, store)
		if relatedSection != "" {
			builder.WriteString("\n\n" + relatedSection)
		}

		return mcp.NewToolResultText(builder.String()), nil
	}
}

// resolveRelatedRef splits a related reference into linter and rule parts.
func resolveRelatedRef(ref string) (string, string) {
	const pathParts = 2

	parts := strings.SplitN(ref, "/", pathParts)
	if len(parts) == pathParts {
		return parts[0], parts[1]
	}
	return parts[0], ""
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
func buildRelatedContext(unique []lintIssue, store *guides.Store) string {
	// Build primary diagnostic set for exclusion
	primarySet := make(map[string]bool)
	for _, issue := range unique {
		rule := extractRule(issue.Text)
		if rule != "" {
			primarySet[issue.FromLinter+"/"+rule] = true
		}
		primarySet[issue.FromLinter] = true
	}

	// Collect related entries with best hint per ref
	bestEntries := make(map[string]*relatedEntry)

	for _, issue := range unique {
		linter := issue.FromLinter
		rule := extractRule(issue.Text)

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
			refLinter, refRule := resolveRelatedRef(ref)
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
	section := "### Related Context\n" + strings.Join(lines, "\n")

	// Enforce byte budget
	for len(section) > maxRelatedBytes && len(lines) > 0 {
		lines = lines[:len(lines)-1]
		section = "### Related Context\n" + strings.Join(lines, "\n")
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

func writeGuideForIssue(builder *strings.Builder, store *guides.Store, opts Options, issue lintIssue) {
	linter := issue.FromLinter
	rule := extractRule(issue.Text)

	body, ok := resolveGuide(store, linter, rule)
	if ok {
		if rule != "" {
			fmt.Fprintf(builder, "## %s: %s\n\n", linter, rule)
		} else {
			fmt.Fprintf(builder, "## %s\n\n", linter)
		}
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
