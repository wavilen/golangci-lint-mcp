package server

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupSummarizeTestServer(t *testing.T) (*mcptest.Server, context.Context) {
	t.Helper()

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile(
			"# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<patterns>\n- Always check error return values\n</patterns>",
		),
		"guides/govet.md": testMapFile(
			"# govet\n\n<instructions>Vet examines Go source code</instructions>\n\n<patterns>\n- Check Printf args\n</patterns>",
		),
		"guides/gocritic/badcall.md": testMapFile(
			"# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>",
		),
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>",
		),
		"guides/staticcheck/SA1000.md": testMapFile(
			"# staticcheck: SA1000\n\n<instructions>Invalid regex</instructions>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	summarizeTool := mcp.NewTool("golangci_lint_summarize",
		mcp.WithDescription("Summarize golangci-lint JSON"),
		mcp.WithString("output", mcp.Required(), mcp.Description("Raw JSON")),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(summarizeTool, makeSummarizeHandler(store))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Multiple packages shows correct breakdown sorted by count descending.
func TestSummarizeHandler_MultiplePackages(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	// Use different (linter, rule) pairs to avoid dedup
	json := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"R1: unchecked error in handler","Pos":{"Filename":"pkg/auth/handler.go","Line":1,"Column":1}},` +
		`{"FromLinter":"errcheck","Text":"R2: unchecked error in middleware","Pos":{"Filename":"pkg/auth/middleware.go","Line":2,"Column":1}},` +
		`{"FromLinter":"govet","Text":"R3: Printf arg issue","Pos":{"Filename":"pkg/db/conn.go","Line":3,"Column":1}},` +
		`{"FromLinter":"gosec","Text":"G101: hardcoded credentials","Pos":{"Filename":"main.go","Line":4,"Column":1}}` +
		`],"Report":{}}`

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_summarize", map[string]any{"output": json}))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "## Summary")
	assert.Contains(t, text, "Total issues: 4")
	assert.Contains(t, text, "Unique diagnostics: 4")
	assert.Contains(t, text, "## Package Breakdown")
	assert.Contains(t, text, "pkg/auth: 2 issues")
	assert.Contains(t, text, "pkg/db: 1 issues")
	assert.Contains(t, text, "TOTAL: 4 issues across 3 packages")
}

// Test 2: ≤30 issues → single-agent strategy.
func TestSummarizeHandler_StrategySingleAgent(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	json := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}},` +
		`{"FromLinter":"govet","Text":"Printf arg issue","Pos":{"Filename":"main.go","Line":2,"Column":1}}` +
		`],"Report":{}}`

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_summarize", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "Strategy: single-agent")
	assert.Contains(t, text, "single-agent flow")
}

// Test 3: >30 issues → subagent-per-package strategy.
func TestSummarizeHandler_StrategySubagent(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	// Build 31 unique (linter, rule) pairs to avoid dedup
	linters := []string{"errcheck", "govet", "gocritic", "gosec"}
	pkgs := []string{"pkg/a", "pkg/b", "pkg/c", "pkg/d", "pkg/e"}
	issues := make([]string, 0, 31)
	for i := range 31 {
		linter := linters[i%len(linters)]
		pkg := pkgs[i%len(pkgs)]
		rule := fmt.Sprintf("RULE%02d: issue desc", i)
		issues = append(issues,
			fmt.Sprintf(`{"FromLinter":"%s","Text":"%s","Pos":{"Filename":"%s/file.go","Line":1,"Column":1}}`,
				linter, rule, pkg))
	}
	json := `{"Issues":[` + strings.Join(issues, ",") + `],"Report":{}}`

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_summarize", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "Strategy: subagent-per-package")
}

// Test 4: Empty issues array → "No issues found".
func TestSummarizeHandler_EmptyOutput(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	json := `{"Issues":[],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_summarize", map[string]any{"output": json}))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "No issues found")
}

// Test 5: Invalid JSON → error message.
func TestSummarizeHandler_InvalidJSON(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	result, err := srv.Client().CallTool(ctx,
		testGuideCall("golangci_lint_summarize", map[string]any{"output": "not json"}))
	require.NoError(t, err)
	require.True(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "invalid json")
}

// Test 6: Linter breakdown section is present.
func TestSummarizeHandler_LinterBreakdown(t *testing.T) {
	srv, ctx := setupSummarizeTestServer(t)

	json := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}},` +
		`{"FromLinter":"gosec","Text":"G101: hardcoded credentials","Pos":{"Filename":"main.go","Line":2,"Column":1}}` +
		`],"Report":{}}`

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_summarize", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "## Linter Breakdown")
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, text, "gosec")
}
