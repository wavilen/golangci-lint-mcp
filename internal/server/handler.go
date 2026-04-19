package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

const gosecAISection = `

<autofix>
gosec supports AI-powered autofix suggestions. Call the ` + "`gosec_ai_autofix`" + ` MCP tool per package directory — gosec requires Go package context for type resolution, so batch issues by package rather than running on the entire project at once.

Identify gosec-bearing packages from the golangci-lint JSON output (group diagnostics with FromLinter "gosec" by their Pos.Filename directory). Then for each package:

  gosec_ai_autofix(path="./pkg/auth/...")

Do NOT call gosec_ai_autofix(path="./...") on the whole project. If the tool times out or fails, fall back to this guide's <instructions> and <examples> for manual fixes.
Review AI suggestions carefully before committing.
</autofix>`

// makeHandler creates a tool handler that resolves linter guides using the store.
func makeHandler(store *guides.Store, opts Options) func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		linter, err := req.RequireString("linter")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'linter': %v", err)), nil
		}
		rule := req.GetString("rule", "")

		// Validate: reject empty linter name
		linter = strings.TrimSpace(linter)
		rule = strings.TrimSpace(rule)
		if linter == "" {
			return mcp.NewToolResultError("parameter 'linter' must not be empty"), nil
		}

		// Try lookup with rule first
		if rule != "" {
			guide, found := store.Lookup(linter, rule)
			if found {
				body := guide.RawBody
				if opts.GosecAI && linter == "gosec" {
					body += gosecAISection
				}
				return mcp.NewToolResultText(body), nil
			}

			// Check if linter itself exists (as a simple guide or compound linter)
			_, linterExists := store.Lookup(linter, "")
			if !linterExists && len(store.ListRules(linter)) == 0 {
				suggestion := store.Suggest(linter)
				msg := fmt.Sprintf("Unknown linter %q.", linter)
				if suggestion != "" {
					msg = fmt.Sprintf("Unknown linter %q. Did you mean %q?", linter, suggestion)
				}
				return mcp.NewToolResultError(msg), nil
			}

			// Linter exists but rule doesn't — check if it's a compound linter
			rules := store.ListRules(linter)
			if len(rules) > 0 {
				sort.Strings(rules)
				return mcp.NewToolResultError(
					fmt.Sprintf("No rule %q found for linter %q. Available rules: %s",
						rule, linter, strings.Join(rules, ", "))), nil
			}
			// Simple linter — doesn't have rules
			return mcp.NewToolResultError(
				fmt.Sprintf("Linter %q does not have sub-rules. Query it without the 'rule' parameter.", linter)), nil
		}

		// No rule specified
		guide, found := store.Lookup(linter, "")
		if found {
			// Simple linter — return its guide
			body := guide.RawBody
			if opts.GosecAI && linter == "gosec" {
				body += gosecAISection
			}
			return mcp.NewToolResultText(body), nil
		}

		// Not a simple linter — check if it's a compound linter
		rules := store.ListRules(linter)
		if len(rules) > 0 {
			sort.Strings(rules)
			return mcp.NewToolResultError(
				fmt.Sprintf("Linter %q has %d rules. Specify a rule to get specific guidance. Available rules: %s",
					linter, len(rules), strings.Join(rules, ", "))), nil
		}

		// Unknown linter — suggest alternatives
		suggestion := store.Suggest(linter)
		msg := fmt.Sprintf("Unknown linter %q.", linter)
		if suggestion != "" {
			msg = fmt.Sprintf("Unknown linter %q. Did you mean %q?", linter, suggestion)
		}
		return mcp.NewToolResultError(msg), nil
	}
}
