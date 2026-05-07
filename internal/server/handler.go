package server

import (
	"context"
	"errors"
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

// guideResultOption configures a guideResult.
type guideResultOption func(*guideResult)

// newGuideResult creates a guideResult with the provided options.
func newGuideResult(opts ...guideResultOption) guideResult {
	r := guideResult{Body: "", Related: nil, IsError: false}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

func withBody(body string) guideResultOption {
	return func(r *guideResult) { r.Body = body }
}

func withRelated(related []string) guideResultOption {
	return func(r *guideResult) { r.Related = related }
}

func withError() guideResultOption {
	return func(r *guideResult) { r.IsError = true }
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
			return newGuideResult(withBody(body), withRelated(guide.Related))
		}
		_, linterExists := store.Lookup(linter, "")
		if !linterExists && len(store.ListRules(linter)) == 0 {
			return newGuideResult(withBody(unknownLinterMessage(linter, store)), withError())
		}
		rules := store.ListRules(linter)
		if len(rules) > 0 {
			return newGuideResult(withBody(fmt.Sprintf(
				"No rule %q found for linter %q. Available rules: %s",
				rule,
				linter,
				strings.Join(rules, ", "),
			)), withError())
		}
		return newGuideResult(withBody(
			fmt.Sprintf("Linter %q does not have sub-rules. Query it without the 'rule' parameter.", linter),
		), withError())
	}
	// No rule provided
	guide, found := store.Lookup(linter, "")
	if found {
		body := stripRelatedTag(maybeAppendGosecAI(guide.RawBody, opts, linter))
		return newGuideResult(withBody(body), withRelated(guide.Related))
	}
	rules := store.ListRules(linter)
	if len(rules) > 0 {
		return newGuideResult(withBody(fmt.Sprintf(
			"Linter %q has %d rules. Specify a rule to get specific guidance. Available rules: %s",
			linter,
			len(rules),
			strings.Join(rules, ", "),
		)), withError())
	}
	return newGuideResult(withBody(unknownLinterMessage(linter, store)), withError())
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

// parseGuideQueries extracts, validates, and deduplicates guide queries from raw MCP arguments.
func parseGuideQueries(queriesRaw []any, batchMax int) ([]guideQuery, error) {
	if len(queriesRaw) > batchMax {
		return nil, fmt.Errorf("batch size %d exceeds maximum of %d; "+
			"use golangci_lint_parse or golangci_lint_run for bulk analysis",
			len(queriesRaw), batchMax)
	}

	var queries []guideQuery
	seen := make(map[guideQuery]bool)
	for _, raw := range queriesRaw {
		rawMap, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		linter, _ := rawMap["linter"].(string)
		linter = strings.TrimSpace(linter)
		if linter == "" {
			continue
		}
		rule, _ := rawMap["rule"].(string)
		rule = strings.TrimSpace(rule)
		gq := guideQuery{Linter: linter, Rule: rule}
		if !seen[gq] {
			seen[gq] = true
			queries = append(queries, gq)
		}
	}

	if len(queries) == 0 {
		return nil, errors.New("no valid queries after parsing")
	}
	return queries, nil
}

// formatGuideSection formats a single query result into a guide or error section string.
func formatGuideSection(query guideQuery, result guideResult) string {
	if result.IsError {
		if query.Rule != "" {
			return fmt.Sprintf("<error linter=%q rule=%q>%s</error>",
				query.Linter, query.Rule, result.Body)
		}
		return fmt.Sprintf("<error linter=%q>%s</error>",
			query.Linter, result.Body)
	}
	if query.Rule != "" {
		return fmt.Sprintf("<guide linter=%q rule=%q>\n%s\n</guide>",
			query.Linter, query.Rule, result.Body)
	}
	return fmt.Sprintf("<guide linter=%q>\n%s\n</guide>",
		query.Linter, result.Body)
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

		queries, err := parseGuideQueries(queriesRaw, batchMaxFromEnv())
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		var sections []string
		var allRelated []string
		allFailed := true

		for _, query := range queries {
			result := resolveGuideBody(store, opts, query.Linter, query.Rule)
			sections = append(sections, formatGuideSection(query, result))
			if !result.IsError {
				allFailed = false
				allRelated = append(allRelated, result.Related...)
			}
		}

		if allFailed {
			return mcp.NewToolResultError(strings.Join(sections, "\n\n")), nil
		}

		relatedSection := buildBatchRelatedContext(allRelated, store)
		response := strings.Join(sections, "\n\n")
		if relatedSection != "" {
			response += "\n\n" + relatedSection
		}

		return mcp.NewToolResultText(response), nil
	}
}
