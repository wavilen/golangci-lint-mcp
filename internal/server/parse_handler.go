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

const largeOutputThreshold = 10

func buildStrategy(totalCount int) (string, string) {
	if totalCount > largeOutputThreshold {
		return "B", ">10 diagnostics — summarize first"
	}
	return "A", "≤10 diagnostics — present inline"
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
		strategy, strategyReason := buildStrategy(len(unique))

		var builder strings.Builder
		fmt.Fprintf(
			&builder,
			"## Summary\n\n- Unique diagnostics: %d\n- Strategy: %s (%s)\n- Breakdown: %s\n\n---\n\n",
			len(unique),
			strategy,
			strategyReason,
			buildLinterBreakdown(unique),
		)

		for idx, issue := range unique {
			if idx > 0 {
				builder.WriteString("\n---\n\n")
			}
			writeGuideForIssue(&builder, store, opts, issue)
		}

		return mcp.NewToolResultText(builder.String()), nil
	}
}

func writeGuideForIssue(builder *strings.Builder, store *guides.Store, opts Options, issue lintIssue) {
	linter := issue.FromLinter
	rule := extractRule(issue.Text)

	if rule != "" {
		guide, found := store.Lookup(linter, rule)
		if found {
			fmt.Fprintf(builder, "## %s: %s\n\n", linter, rule)
			builder.WriteString(maybeAppendGosecAI(guide.RawBody, opts, linter))
			return
		}
	}

	guide, found := store.Lookup(linter, "")
	if found {
		fmt.Fprintf(builder, "## %s\n\n", linter)
		builder.WriteString(maybeAppendGosecAI(guide.RawBody, opts, linter))
		return
	}

	rules := store.ListRules(linter)
	if len(rules) > 0 {
		sort.Strings(rules)
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
