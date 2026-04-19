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
	idx := strings.Index(text, ": ")
	if idx == -1 {
		return ""
	}
	return strings.TrimSpace(text[:idx])
}

func makeParseHandler(store *guides.Store, opts Options) func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output, err := req.RequireString("output")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'output': %v", err)), nil
		}
		output = strings.TrimSpace(output)
		if output == "" {
			return mcp.NewToolResultError("parameter 'output' must not be empty"), nil
		}

		var result lintJSONResult
		firstLine := output
		if idx := strings.Index(output, "\n"); idx != -1 {
			firstLine = output[:idx]
		}
		if err := json.Unmarshal([]byte(firstLine), &result); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid JSON: %v", err)), nil
		}

		if len(result.Issues) == 0 {
			return mcp.NewToolResultText("No issues found in the golangci-lint output."), nil
		}

		seen := make(map[diagnostic]bool)
		var unique []lintIssue
		for _, issue := range result.Issues {
			rule := extractRule(issue.Text)
			key := diagnostic{linter: issue.FromLinter, rule: rule}
			if !seen[key] {
				seen[key] = true
				unique = append(unique, issue)
			}
		}

		var sb strings.Builder
		for i, issue := range unique {
			if i > 0 {
				sb.WriteString("\n---\n\n")
			}

			linter := issue.FromLinter
			rule := extractRule(issue.Text)

			if rule != "" {
				guide, found := store.Lookup(linter, rule)
				if found {
					fmt.Fprintf(&sb, "## %s: %s\n\n", linter, rule)
					body := guide.RawBody
					if opts.GosecAI && linter == "gosec" {
						body += gosecAISection
					}
					sb.WriteString(body)
					continue
				}
			}

			guide, found := store.Lookup(linter, "")
			if found {
			fmt.Fprintf(&sb, "## %s\n\n", linter)
				body := guide.RawBody
				if opts.GosecAI && linter == "gosec" {
					body += gosecAISection
				}
				sb.WriteString(body)
				continue
			}

			rules := store.ListRules(linter)
			if len(rules) > 0 {
				sort.Strings(rules)
				fmt.Fprintf(&sb, "## %s: %s\n\nNo guide found for rule %q of linter %q. Available rules: %s",
					linter, rule, rule, linter, strings.Join(rules, ", "))
				continue
			}

			suggestion := store.Suggest(linter)
			msg := fmt.Sprintf("## %s\n\nUnknown linter %q.", linter, linter)
			if suggestion != "" {
				msg = fmt.Sprintf("## %s\n\nUnknown linter %q. Did you mean %q?", linter, linter, suggestion)
			}
			sb.WriteString(msg)
		}

		return mcp.NewToolResultText(sb.String()), nil
	}
}
