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
		mcp.WithArray("queries",
			mcp.Required(),
			mcp.Description("Array of query objects"),
			mcp.Items(map[string]any{
				"type": "object",
				"properties": map[string]any{
					"linter": map[string]any{"type": "string"},
					"rule":   map[string]any{"type": "string"},
				},
				"required": []string{"linter"},
			}),
		),
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

// Test: No duplicate headers in guidance output. The synthetic h2 heading
// ("## linter: rule") must not be followed by the guide body's own h1 heading
// ("# linter: rule"), which was the bug in writeGuideForIssue.
func TestParseHandler_NoDuplicateHeaders(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	// Should have the synthetic h2 heading
	assert.Contains(t, text, "## errcheck")

	// Should NOT have the guide's original h1 heading (which was "# errcheck\n\n")
	// after the h2. The guide body starts with "# errcheck\n\n<instructions>..."
	// so the h1 should be stripped.
	assert.NotContains(t, text, "## errcheck\n\n# errcheck", "duplicate h1 after h2 heading")

	// Also check compound linter:rule case
	t.Run("compound_linter_rule", func(t *testing.T) {
		json2 := `{"Issues":[{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
		call := testGuideCall("golangci_lint_parse", map[string]any{"output": json2})
		result2, err2 := srv.Client().CallTool(ctx, call)
		require.NoError(t, err2)
		text2 := result2.Content[0].(mcp.TextContent).Text

		assert.Contains(t, text2, "## gocritic: dupSubExpr")
		assert.NotContains(t, text2, "## gocritic: dupSubExpr\n\n# gocritic: dupSubExpr",
			"duplicate h1 after h2 heading for compound linter:rule")
	})
}

func TestStripGuideHeading(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple linter heading",
			input:    "# errcheck\n\n<instructions>Check errors</instructions>",
			expected: "<instructions>Check errors</instructions>",
		},
		{
			name:     "compound linter:rule heading",
			input:    "# gocritic: badCall\n\n<instructions>Detect suspicious calls</instructions>",
			expected: "<instructions>Detect suspicious calls</instructions>",
		},
		{
			name:     "rule-only heading",
			input:    "# G101\n\n<instructions>Hardcoded creds</instructions>",
			expected: "<instructions>Hardcoded creds</instructions>",
		},
		{
			name:     "no heading returns as-is",
			input:    "<instructions>No heading here</instructions>",
			expected: "<instructions>No heading here</instructions>",
		},
		{
			name:     "heading only no body",
			input:    "# errcheck\n",
			expected: "",
		},
		{
			name:     "h2 heading not stripped",
			input:    "## errcheck\n\n<instructions>Not touched</instructions>",
			expected: "## errcheck\n\n<instructions>Not touched</instructions>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripGuideHeading(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
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
	// Should suggest golangci_lint_run as primary and golangci_lint_guide as secondary (D-02)
	assert.Contains(t, text, "golangci_lint_run")
	assert.Contains(t, text, "golangci_lint_guide")
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
	assert.Contains(t, text, "<summary>")
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
	assert.Contains(t, text, "Unique diagnostics: 31")
	assert.Contains(t, text, "Strategy: subagent-per-file")
	assert.Contains(t, text, "subagent-per-file strategy")
	// Per D-01: subagent strategy responses should NOT contain guidance or related_context
	assert.NotContains(t, text, "<guidance>")
	assert.NotContains(t, text, "<related_context>")
	// Should contain strategy_instructions with guide call references
	assert.Contains(t, text, "<strategy_instructions>")
	assert.Contains(t, text, `golangci_lint_guide(queries=`)
}

// Test: Subagent strategy responses skip guidance and related_context blocks.
func TestParseHandler_SubagentStrategy_NoGuidance(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	// 31 unique issues → subagent-per-file strategy
	linters := []string{"errcheck", "gocritic", "gosec", "staticcheck", "govet"}
	issueParts := make([]string, 0, 31)
	for i := range 31 {
		linter := linters[i%len(linters)]
		rule := fmt.Sprintf("RULE%02d: issue", i)
		issueParts = append(issueParts,
			fmt.Sprintf(`{"FromLinter":"%s","Text":"%s","Pos":{"Filename":"pkg/file%d.go","Line":%d,"Column":1}}`,
				linter, rule, i, i+1))
	}
	json := `{"Issues":[` + strings.Join(issueParts, ",") + `],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	// Must have summary and strategy_instructions
	assert.Contains(t, text, "<summary>")
	assert.Contains(t, text, "<strategy_instructions>")

	// Must NOT have guidance or related_context for subagent strategy
	assert.NotContains(t, text, "<guidance>", "subagent response should not contain <guidance>")
	assert.NotContains(t, text, "<related_context>", "subagent response should not contain <related_context>")

	// Strategy instructions must contain guide call references
	assert.Contains(t, text, `golangci_lint_guide(queries=`, "strategy must include guide call references")
}

func TestParseHandler_ExistingGuideToolUnchanged(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errcheck"})))
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
	assert.Contains(t, text, "<related_context>")
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
	if strings.Contains(text, "<related_context>") {
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
	assert.Contains(t, text, "<related_context>")
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
	assert.Contains(t, text, "<related_context>")
	// Count lines starting with "- " in the Related Context section
	section, found := cutAfter(text, "<related_context>")
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
	assert.NotContains(t, text, "<related_context>")
}

// Test: Fix hint comes from pattern bullets.
func TestParseHandler_RelatedContext_FixHintFromPatterns(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<related_context>")
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

// NDJSON Tests

// Test: Newline-delimited individual issue objects parse correctly.
func TestParseHandler_NDJSONInput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	ndjson := `{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}
{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":15,"Column":8}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": ndjson}))
	require.NoError(t, err)
	require.False(t, result.IsError, "NDJSON input should parse successfully")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
	assert.Contains(t, text, "duplicate sub-expressions")
	assert.Contains(t, text, "Unique diagnostics: 2")
}

// Test: Existing wrapped JSON still works (backward compatible).
func TestParseHandler_WrappedJSONStillWorks(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": json}))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
	assert.Contains(t, text, "Unique diagnostics: 1")
}

// Test: Non-JSON lines in NDJSON input are skipped silently.
func TestParseHandler_MixedInvalidLinesSkipped(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	ndjson := `some garbage line
{"FromLinter":"errcheck","Text":"Error return value is not checked","Pos":{"Filename":"main.go","Line":10,"Column":5}}
another garbage line
{"FromLinter":"gosec","Text":"G101: Potential hardcoded credentials","Pos":{"Filename":"main.go","Line":5,"Column":1}}`
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": ndjson}))
	require.NoError(t, err)
	require.False(t, result.IsError, "NDJSON with garbage lines should parse valid lines only")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Unique diagnostics: 2")
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

// Test: Empty lines only input returns error.
func TestParseHandler_NDJSONEmptyInput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	ndjson := "\n\n  \n\n"
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_parse", map[string]any{"output": ndjson}))
	require.NoError(t, err)
	require.True(t, result.IsError, "empty lines only should return error")
	text := result.Content[0].(mcp.TextContent).Text
	// After TrimSpace, whitespace-only input triggers "must not be empty" check
	assert.Contains(t, strings.ToLower(text), "empty")
}
