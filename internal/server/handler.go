package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
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

func unknownLinterMessage(linter string, store *guides.Store) string {
	suggestion := store.Suggest(linter)
	msg := fmt.Sprintf("Unknown linter %q.", linter)
	if suggestion != "" {
		msg = fmt.Sprintf("Unknown linter %q. Did you mean %q?", linter, suggestion)
	}
	return msg
}

func handleRuleQuery(store *guides.Store, opts Options, linter, rule string) (*mcp.CallToolResult, error) {
	guide, found := store.Lookup(linter, rule)
	if found {
		return mcp.NewToolResultText(maybeAppendGosecAI(guide.RawBody, opts, linter)), nil
	}

	_, linterExists := store.Lookup(linter, "")
	if !linterExists && len(store.ListRules(linter)) == 0 {
		return mcp.NewToolResultError(unknownLinterMessage(linter, store)), nil
	}

	rules := store.ListRules(linter)
	if len(rules) > 0 {
		sort.Strings(rules)
		return mcp.NewToolResultError(
			fmt.Sprintf("No rule %q found for linter %q. Available rules: %s",
				rule, linter, strings.Join(rules, ", "))), nil
	}
	return mcp.NewToolResultError(
		fmt.Sprintf("Linter %q does not have sub-rules. Query it without the 'rule' parameter.", linter)), nil
}

func handleNoRuleQuery(store *guides.Store, opts Options, linter string) (*mcp.CallToolResult, error) {
	guide, found := store.Lookup(linter, "")
	if found {
		return mcp.NewToolResultText(maybeAppendGosecAI(guide.RawBody, opts, linter)), nil
	}

	rules := store.ListRules(linter)
	if len(rules) > 0 {
		sort.Strings(rules)
		return mcp.NewToolResultError(
			fmt.Sprintf("Linter %q has %d rules. Specify a rule to get specific guidance. Available rules: %s",
				linter, len(rules), strings.Join(rules, ", "))), nil
	}

	return mcp.NewToolResultError(unknownLinterMessage(linter, store)), nil
}

func makeHandler(
	store *guides.Store,
	opts Options,
) func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		linter, err := req.RequireString("linter")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'linter': %v", err)), nil
		}
		rule := req.GetString("rule", "")

		linter = strings.TrimSpace(linter)
		rule = strings.TrimSpace(rule)
		if linter == "" {
			return mcp.NewToolResultError("parameter 'linter' must not be empty"), nil
		}

		if rule != "" {
			return handleRuleQuery(store, opts, linter, rule)
		}
		return handleNoRuleQuery(store, opts, linter)
	}
}
