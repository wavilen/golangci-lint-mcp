package server

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	maxRelatedEntries = 5
	maxRelatedBytes   = 500
	defaultBatchMax   = 100
)

const gosecLinterName = "gosec"

const gosecAISection = `

<autofix>
gosec supports AI-powered autofix suggestions. Call the ` + "`gosec_ai_autofix`" + ` MCP tool per package directory — gosec requires Go package context for type resolution, so batch issues by package rather than running on the entire project at once.

Identify gosec-bearing packages from the golangci-lint JSON output (group diagnostics with FromLinter "gosec" by their Pos.Filename directory). Then for each package:

  gosec_ai_autofix(path="./pkg/auth/...")

Do NOT call gosec_ai_autofix(path="./...") on the whole project. If the tool times out or fails, fall back to this guide's <instructions> and <examples> for manual fixes.
Review AI suggestions carefully before committing.
</autofix>`

func maybeAppendGosecAI(body string, opts Options, linter string) string {
	if opts.GosecAI && linter == gosecLinterName {
		return body + gosecAISection
	}
	return body
}

// guideQuery represents a single {linter, rule} query from the batch request.
type guideQuery struct {
	Linter string
	Rule   string
}

// guideResult holds the outcome of processing a single query.
type guideResult struct {
	Body    string   // guide body (stripped of <related> tag, with gosec AI if applicable) or error message
	Related []string // related refs from the guide (empty on error)
	IsError bool     // true when the query failed (unknown linter, bad rule, etc.)
}

// batchMaxFromEnv returns the maximum batch size from GOLANGCI_LINT_BATCH_MAX
// env var, defaulting to defaultBatchMax. Invalid values are silently ignored.
func batchMaxFromEnv() int {
	val := os.Getenv("GOLANGCI_LINT_BATCH_MAX")
	if val != "" {
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 {
			return n
		}
	}
	return defaultBatchMax
}

// resolveGuideBody looks up a single (linter, rule) pair and returns the guide body
// with related refs, or an error message. It reuses the same error message patterns
// as the old handleRuleQuery/handleNoRuleQuery.
func resolveGuideBody(store *guides.Store, opts Options, linter, rule string) guideResult {
	if rule != "" {
		guide, found := store.Lookup(linter, rule)
		if found {
			body := stripRelatedTag(maybeAppendGosecAI(guide.RawBody, opts, linter))
			return guideResult{Body: body, Related: guide.Related}
		}
		_, linterExists := store.Lookup(linter, "")
		if !linterExists && len(store.ListRules(linter)) == 0 {
			return guideResult{Body: unknownLinterMessage(linter, store), IsError: true}
		}
		rules := store.ListRules(linter)
		if len(rules) > 0 {
			return guideResult{
				Body: fmt.Sprintf(
					"No rule %q found for linter %q. Available rules: %s",
					rule,
					linter,
					strings.Join(rules, ", "),
				),
				IsError: true,
			}
		}
		return guideResult{
			Body:    fmt.Sprintf("Linter %q does not have sub-rules. Query it without the 'rule' parameter.", linter),
			IsError: true,
		}
	}
	// No rule provided
	guide, found := store.Lookup(linter, "")
	if found {
		body := stripRelatedTag(maybeAppendGosecAI(guide.RawBody, opts, linter))
		return guideResult{Body: body, Related: guide.Related}
	}
	rules := store.ListRules(linter)
	if len(rules) > 0 {
		return guideResult{
			Body: fmt.Sprintf(
				"Linter %q has %d rules. Specify a rule to get specific guidance. Available rules: %s",
				linter,
				len(rules),
				strings.Join(rules, ", "),
			),
			IsError: true,
		}
	}
	return guideResult{Body: unknownLinterMessage(linter, store), IsError: true}
}

// stripRelatedTag removes <related>...</related> from body and cleans up
// resulting blank lines.
func stripRelatedTag(body string) string {
	for {
		start := strings.Index(body, "<related>")
		if start == -1 {
			break
		}
		end := strings.Index(body, "</related>")
		if end == -1 {
			break
		}
		body = body[:start] + body[end+len("</related>"):]
	}
	// Clean up multiple blank lines left behind
	for strings.Contains(body, "\n\n\n") {
		body = strings.ReplaceAll(body, "\n\n\n", "\n\n")
	}
	body = strings.TrimRight(body, "\n")
	return body
}

// parseRelatedRef splits a related reference into linter and rule parts.
func parseRelatedRef(ref string) (string, string) {
	const pathParts = 2

	parts := strings.SplitN(ref, "/", pathParts)
	if len(parts) == pathParts {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func unknownLinterMessage(linter string, store *guides.Store) string {
	suggestion := store.Suggest(linter)
	msg := fmt.Sprintf(
		"Unknown linter %q. This may be from a newer/older golangci-lint version.",
		linter,
	)
	if suggestion != "" {
		msg = fmt.Sprintf(
			"Unknown linter %q. Did you mean %q? This may be from a newer/older golangci-lint version.",
			linter,
			suggestion,
		)
	}
	return msg
}

// buildBatchRelatedContext builds a deduplicated <related_context> section from
// all related refs collected across successful queries in the batch.
func buildBatchRelatedContext(allRelated []string, store *guides.Store) string {
	if len(allRelated) == 0 {
		return ""
	}

	// Deduplicate related refs
	seen := make(map[string]bool)
	var unique []string
	for _, ref := range allRelated {
		if !seen[ref] {
			seen[ref] = true
			unique = append(unique, ref)
		}
	}

	var entries []string
	for _, ref := range unique {
		if len(entries) >= maxRelatedEntries {
			break
		}
		linter, rule := parseRelatedRef(ref)
		relatedGuide, found := store.Lookup(linter, rule)
		if !found && rule != "" {
			relatedGuide, found = store.Lookup(linter, "")
		}
		if !found {
			continue
		}
		fixHint := guides.BestPatternBullet(relatedGuide.Patterns, ref)
		if fixHint == "" {
			continue
		}
		entries = append(entries, fmt.Sprintf("- %s: %s", ref, fixHint))
	}

	if len(entries) == 0 {
		return ""
	}

	section := "<related_context>\n" + strings.Join(entries, "\n") + "\n</related_context>"
	// Enforce byte budget: trim entries from bottom if too long
	for len(section) > maxRelatedBytes && len(entries) > 0 {
		entries = entries[:len(entries)-1]
		section = "<related_context>\n" + strings.Join(entries, "\n") + "\n</related_context>"
	}

	if len(entries) == 0 {
		return ""
	}

	return section
}

func makeHandler(
	store *guides.Store,
	opts Options,
) func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := req.GetArguments()
		queriesRaw, ok := args["queries"].([]any)
		if !ok || len(queriesRaw) == 0 {
			return mcp.NewToolResultError(
				"parameter 'queries' must be a non-empty array of {linter, rule} objects. " +
					"Use golangci_lint_list to discover available linters."), nil
		}

		batchMax := batchMaxFromEnv()
		if len(queriesRaw) > batchMax {
			return mcp.NewToolResultError(
				fmt.Sprintf("batch size %d exceeds maximum of %d. "+
					"Use golangci_lint_parse or golangci_lint_run for bulk analysis.",
					len(queriesRaw), batchMax)), nil
		}

		// Parse and deduplicate queries
		var queries []guideQuery
		seen := make(map[guideQuery]bool)
		for _, q := range queriesRaw {
			qMap, ok := q.(map[string]any)
			if !ok {
				continue
			}
			linter, _ := qMap["linter"].(string)
			linter = strings.TrimSpace(linter)
			if linter == "" {
				continue
			}
			rule, _ := qMap["rule"].(string)
			rule = strings.TrimSpace(rule)
			gq := guideQuery{Linter: linter, Rule: rule}
			if !seen[gq] {
				seen[gq] = true
				queries = append(queries, gq)
			}
		}

		if len(queries) == 0 {
			return mcp.NewToolResultError("no valid queries after parsing"), nil
		}

		// Process each query
		var sections []string
		var allRelated []string
		allFailed := true

		for _, q := range queries {
			result := resolveGuideBody(store, opts, q.Linter, q.Rule)
			if result.IsError {
				if q.Rule != "" {
					sections = append(sections, fmt.Sprintf("<error linter=%q rule=%q>%s</error>",
						q.Linter, q.Rule, result.Body))
				} else {
					sections = append(sections, fmt.Sprintf("<error linter=%q>%s</error>",
						q.Linter, result.Body))
				}
			} else {
				allFailed = false
				if q.Rule != "" {
					sections = append(sections, fmt.Sprintf("<guide linter=%q rule=%q>\n%s\n</guide>",
						q.Linter, q.Rule, result.Body))
				} else {
					sections = append(sections, fmt.Sprintf("<guide linter=%q>\n%s\n</guide>",
						q.Linter, result.Body))
				}
				allRelated = append(allRelated, result.Related...)
			}
		}

		// If ALL queries failed, return MCP error result
		if allFailed {
			return mcp.NewToolResultError(strings.Join(sections, "\n\n")), nil
		}

		// Batch-level related context deduplication
		relatedSection := buildBatchRelatedContext(allRelated, store)

		response := strings.Join(sections, "\n\n")
		if relatedSection != "" {
			response += "\n\n" + relatedSection
		}

		return mcp.NewToolResultText(response), nil
	}
}
