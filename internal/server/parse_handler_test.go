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

func setupParseTestServer(t *testing.T, opts ...Options) (*mcptest.Server, context.Context) {
	t.Helper()
	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile(
			"# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<examples>```go\nfile, _ := os.Open(\"f\")\n```</examples>\n\n<patterns>\n- Always check error return values\n- Use comma-ok for type assertions\n</patterns>\n\n<related>govet, rowserrcheck</related>",
		),
		"guides/rowserrcheck.md": testMapFile(
			"# rowserrcheck\n\n<instructions>Checks whether Rows.Err is checked</instructions>\n\n<patterns>\n- Always check rows.Err after iterating with rows.Next\n- Use defer rows.Close() before iterating\n</patterns>",
		),
		"guides/govet.md": testMapFile(
			"# govet\n\n<instructions>Vet examines Go source code and reports suspicious constructs</instructions>\n\n<patterns>\n- Check Printf argument count matches format verbs\n- Verify composite literal field keys\n</patterns>",
		),
		"guides/gocritic/badCall.md": testMapFile(
			"# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>",
		),
		"guides/gocritic/dupSubExpr.md": testMapFile(
			"# gocritic: dupSubExpr\n\n<instructions>Detects duplicate sub-expressions</instructions>\n\n<patterns>\n- Compare operands for intentional symmetry\n- Remove duplicate conditions in boolean expressions\n</patterns>\n\n<related>gocritic/badCall</related>",
		),
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>\n\n<patterns>\n- Move credentials to environment variables\n- Use secret management tools\n</patterns>\n\n<related>gosec/G201</related>",
		),
		"guides/gosec/G201.md": testMapFile(
			"# G201\n\n<instructions>Detects SQL injection via string format</instructions>\n\n<patterns>\n- Use parameterized queries instead of string formatting\n- Validate user input before using in SQL\n</patterns>",
		),
		"guides/staticcheck/SA1000.md": testMapFile(
			"# staticcheck: SA1000\n\n<instructions>Detects invalid regex patterns</instructions>",
		),
		"guides/staticcheck/SA2000.md": testMapFile(
			"# staticcheck: SA2000\n\n<instructions>Detects sync.WaitGroup misuse</instructions>",
		),
		"guides/staticcheck/SA3000.md": testMapFile(
			"# staticcheck: SA3000\n\n<instructions>Detects test division by zero</instructions>",
		),
		"guides/staticcheck/SA4000.md": testMapFile(
			"# staticcheck: SA4000\n\n<instructions>Detects identical binary expressions</instructions>",
		),
		"guides/govet/assign.md": testMapFile(
			"# govet: assign\n\n<instructions>Detects useless assignments</instructions>",
		),
		"guides/govet/composite.md": testMapFile(
			"# govet: composite\n\n<instructions>Detects unkeyed composite literals</instructions>",
		),
		"guides/govet/copylocks.md": testMapFile(
			"# govet: copylocks\n\n<instructions>Detects copies of lock values</instructions>",
		),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

	guideTool := mcp.NewTool("golangci_lint_guide",
		mcp.WithDescription("Get concise guidance for fixing golangci-lint issues"),
		mcp.WithString("linter", mcp.Required(), mcp.Description("The linter name")),
		mcp.WithString("rule", mcp.Description("Optional rule ID")),
	)
	parseTool := mcp.NewTool("golangci_lint_parse",
		mcp.WithDescription("Parse golangci-lint JSON and return fix guidance"),
		mcp.WithString("output", mcp.Required(), mcp.Description("Raw golangci-lint JSON")),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(guideTool, makeHandler(store, opt))
	mcpServer.AddTool(parseTool, makeParseHandler(store, opt))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

func TestParseHandler_MultipleLinters(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}},{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":15,"Column":8}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
	assert.Contains(t, text, "duplicate sub-expressions")
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, text, "gocritic")
}

func TestParseHandler_Deduplication(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"a.go","Line":10,"Column":5}},{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"b.go","Line":20,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Equal(t, 1, strings.Count(text, "## errcheck"))
}

func TestParseHandler_CompoundRuleExtraction(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gocritic: dupSubExpr")
	assert.Contains(t, text, "duplicate sub-expressions")
}

func TestParseHandler_InvalidJSON(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": "not json at all"}))
	require.NoError(t, err)
	require.True(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "invalid json")
}

func TestParseHandler_EmptyOutput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": ""}))
	require.NoError(t, err)
	require.True(t, result.IsError)
}

func TestParseHandler_EmptyIssues(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "no issues found")
}

func TestParseHandler_UnknownLinter(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"typolinter","Text":"Some issue","Pos":{"Filename":"main.go","Line":1,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Unknown linter")
	assert.Contains(t, text, "typolinter")
}

func TestParseHandler_GosecWithAIFlag(t *testing.T) {
	srv, ctx := setupParseTestServer(t, testAIOptions())
	json := `{"Issues":[{"FromLinter":"gosec","Text":"G101: Potential hardcoded credentials","Pos":{"Filename":"main.go","Line":5,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.Contains(t, text, "<autofix>")
}

func TestParseHandler_GosecWithoutAIFlag(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"gosec","Text":"G101: Potential hardcoded credentials","Pos":{"Filename":"main.go","Line":5,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.NotContains(t, text, "<autofix>")
}

func TestParseHandler_MultiLineOutput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := "{\"Issues\":[{\"FromLinter\":\"errcheck\",\"Text\":\"Error return value is not checked\",\"Pos\":{\"Filename\":\"main.go\",\"Line\":10,\"Column\":5}}],\"Report\":{}}\n2 issues:\n* errcheck: 2\n"
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

func TestParseHandler_SummaryBlock_SingleDiagnostic(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "## Summary")
	assert.Contains(t, text, "Unique diagnostics: 1")
	assert.Contains(t, text, "Strategy: single-agent")
	assert.Contains(t, text, "errcheck (1)")
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

func TestParseHandler_SummaryBlock_StrategyAThreshold(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"a.go","Line":1,"Column":1}},` +
		`{"FromLinter":"gocritic","Text":"badCall: something","Pos":{"Filename":"a.go","Line":2,"Column":1}},` +
		`{"FromLinter":"gocritic","Text":"dupSubExpr: something","Pos":{"Filename":"a.go","Line":3,"Column":1}},` +
		`{"FromLinter":"gosec","Text":"G101: hardcoded credentials","Pos":{"Filename":"a.go","Line":4,"Column":1}},` +
		`{"FromLinter":"gosec","Text":"G201: SQL injection","Pos":{"Filename":"a.go","Line":5,"Column":1}},` +
		`{"FromLinter":"staticcheck","Text":"SA1000: something","Pos":{"Filename":"a.go","Line":6,"Column":1}},` +
		`{"FromLinter":"staticcheck","Text":"SA2000: something","Pos":{"Filename":"a.go","Line":7,"Column":1}},` +
		`{"FromLinter":"staticcheck","Text":"SA3000: something","Pos":{"Filename":"a.go","Line":8,"Column":1}},` +
		`{"FromLinter":"govet","Text":"assign: something","Pos":{"Filename":"a.go","Line":9,"Column":1}},` +
		`{"FromLinter":"govet","Text":"composite: something","Pos":{"Filename":"a.go","Line":10,"Column":1}}` +
		`],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Unique diagnostics: 10")
	assert.Contains(t, text, "Strategy: single-agent")
	assert.Contains(t, text, "single-agent flow")
}

func TestParseHandler_SummaryBlock_StrategyB(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// Build 31 unique issues to exceed the 30-issue subagent threshold
	linters := []string{"errcheck", "gocritic", "gosec", "staticcheck", "govet"}
	issueParts := make([]string, 0, 31)
	for i := range 31 {
		linter := linters[i%len(linters)]
		rule := fmt.Sprintf("RULE%02d: something", i)
		issueParts = append(issueParts,
			fmt.Sprintf(`{"FromLinter":"%s","Text":"%s","Pos":{"Filename":"a.go","Line":%d,"Column":1}}`,
				linter, rule, i+1))
	}
	json := `{"Issues":[` + strings.Join(issueParts, ",") + `],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "## Summary")
	assert.Contains(t, text, "Unique diagnostics: 31")
	assert.Contains(t, text, "Strategy: subagent-per-package")
	assert.Contains(t, text, "subagent-per-package strategy")
}

func TestParseHandler_ExistingGuideToolUnchanged(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", map[string]any{"linter": "errcheck"}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

// Test: Parse response includes Related Context after all issues.
func TestParseHandler_RelatedContext_MultipleLinters(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}},{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":15,"Column":8}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "### Related Context")
	// errcheck's related: govet, rowserrcheck
	// gocritic/dupSubExpr's related: gocritic/badCall
	// Should see at least govet and rowserrcheck as related entries
}

// Test: Linters in primary diagnostic set excluded from Related Context.
func TestParseHandler_RelatedContext_DeduplicationAgainstPrimary(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// errcheck has related: govet, rowserrcheck
	// If errcheck is in the primary set, it should NOT appear in related
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	// errcheck is in the primary set, it should not be listed in Related Context
	// The section should contain govet and rowserrcheck but NOT errcheck
	if strings.Contains(text, "### Related Context") {
		assert.NotContains(t, text, "- errcheck:", "primary linter should not appear in Related Context")
	}
}

// Test: Same related linter for multiple issues appears once with best hint.
func TestParseHandler_RelatedContext_DeduplicationWithinRelated(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// Both errcheck and gocritic/dupSubExpr have related refs
	// but they should not duplicate in the related section
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}},{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":15,"Column":8}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "### Related Context")
}

// Test: Max 5 entries in Related Context.
func TestParseHandler_RelatedContext_MaxEntries(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// Use all Strategy B issues — many linters, many potential related refs
	json := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"a.go","Line":1,"Column":1}},` +
		`{"FromLinter":"gosec","Text":"G101: hardcoded credentials","Pos":{"Filename":"a.go","Line":4,"Column":1}}` +
		`],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "### Related Context")
	// Count lines starting with "- " in the Related Context section
	section, found := cutAfter(text, "### Related Context")
	assert.True(t, found, "should find Related Context section")
	lines := strings.Split(section, "\n")
	var entryLines []string
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "- ") {
			entryLines = append(entryLines, line)
		}
	}
	assert.LessOrEqual(t, len(entryLines), 5, "should have at most 5 related entries")
}

// Test: No related refs → no Related Context section.
func TestParseHandler_RelatedContext_NoRelated(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// gocritic/badCall has no <related> tag
	json := `{"Issues":[{"FromLinter":"gocritic","Text":"badCall: something","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.NotContains(t, text, "### Related Context")
}

// Test: Fix hint comes from pattern bullets.
func TestParseHandler_RelatedContext_FixHintFromPatterns(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "### Related Context")
	// govet should appear as related with a pattern-based hint
	assert.Contains(t, text, "govet:")
	// rowserrcheck should appear with a hint about rows.Err
	assert.Contains(t, text, "rowserrcheck:")
}

// cutAfter returns the substring after the first occurrence of marker.
func cutAfter(s, marker string) (string, bool) {
	_, after, ok := strings.Cut(s, marker)
	if !ok {
		return "", false
	}
	return after, true
}
