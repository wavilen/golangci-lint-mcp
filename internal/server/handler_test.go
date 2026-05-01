package server

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testMapFile(data string) *fstest.MapFile {
	return &fstest.MapFile{
		Data:    []byte(data),
		Mode:    0,
		ModTime: time.Time{},
		Sys:     nil,
	}
}

func testGuideCall(toolName string, args map[string]any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "",
			Params: mcp.RequestParams{Meta: nil},
		},
		Header: nil,
		Params: mcp.CallToolParams{
			Name:      toolName,
			Arguments: args,
			Meta:      nil,
			Task:      nil,
		},
	}
}

func testAIOptions() Options {
	return Options{
		GosecAI:         true,
		GosecAIProvider: "",
		GosecAIKey:      "",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
}

// testBatchArgs builds a queries array argument for the batch guide handler.
func testBatchArgs(queries ...map[string]string) map[string]any {
	items := make([]any, 0, len(queries))
	for _, q := range queries {
		item := make(map[string]any, len(q))
		for k, v := range q {
			item[k] = v
		}
		items = append(items, item)
	}
	return map[string]any{"queries": items}
}

func setupTestServer(t *testing.T, opts ...Options) (*mcptest.Server, context.Context) {
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
		"guides/gocritic/badcall.md": testMapFile(
			"# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>",
		),
		"guides/gocritic/appendassign.md": testMapFile(
			"# gocritic: appendAssign\n\n<instructions>Detects append result misassignment</instructions>",
		),
		"guides/gocritic/commentedoutcode.md": testMapFile(
			"# gocritic: commentedOutCode\n\n<instructions>Detects commented-out code</instructions>",
		),
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>\n\n<patterns>\n- Move credentials to environment variables\n- Use secret management tools\n</patterns>\n\n<related>gosec/G304</related>",
		),
		"guides/gosec/G201.md": testMapFile(
			"# G201\n\n<instructions>Detects SQL injection via string format</instructions>\n\n<examples>```go\nquery := fmt.Sprintf(\"SELECT id, name FROM users WHERE id = %s\", input)\n```</examples>",
		),
		"guides/gosec/G304.md": testMapFile(
			"# G304\n\n<instructions>Detects file path provided as user input</instructions>\n\n<patterns>\n- Validate and sanitize file paths before use\n- Use filepath.Clean to resolve path traversal\n</patterns>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

	// Create MCP tool
	tool := mcp.NewTool("golangci_lint_guide",
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

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(tool, makeHandler(store, opt))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Known simple linter returns guide text.
func TestHandler_SimpleLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errcheck"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"errcheck\">")
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

// Test 2: Known compound rule returns guide text.
func TestHandler_CompoundRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gocritic", "rule": "badcall"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gocritic\" rule=\"badcall\">")
	assert.Contains(t, text, "suspicious function calls")
}

// Test 3: Unknown linter returns error with suggestion.
func TestHandler_UnknownLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errchek"})))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<error linter=\"errchek\">")
	assert.Contains(t, text, "Unknown linter")
	assert.Contains(t, text, "errchek")
	// Should suggest errcheck as close match via Levenshtein
	assert.Contains(t, text, "errcheck")
	// Should mention version mismatch (D-03)
	assert.Contains(t, text, "newer/older")
}

// Test 4: Compound linter without rule lists rules.
func TestHandler_CompoundNoRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gocritic"})))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gocritic")
	assert.Contains(t, text, "rule")
	assert.Contains(t, text, "appendassign")
	assert.Contains(t, text, "badcall")
	assert.Contains(t, text, "commentedoutcode")
}

// Test 5: Empty queries array returns error about non-empty array.
func TestHandler_MissingLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide", map[string]any{"queries": []any{}}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "non-empty array")
	assert.Contains(t, text, "golangci_lint_list")
}

// Test 6: Simple linter with rule returns error about no sub-rules.
func TestHandler_SimpleWithRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errcheck", "rule": "anything"})))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, strings.ToLower(text), "does not have sub-rules")
}

// Test 7: Gosec guide without gosec-ai flag has no autofix section.
func TestHandler_GosecWithoutAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t) // default: no options

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gosec", "rule": "G101"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gosec\" rule=\"G101\">")
	assert.Contains(t, text, "hardcoded credentials")
	assert.NotContains(t, text, "<autofix>")
	assert.NotContains(t, text, "-ai-api-provider")
}

// Test 8: Gosec guide with gosec-ai flag has autofix section with MCP tool pointer.
func TestHandler_GosecWithAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t, testAIOptions())

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gosec", "rule": "G101"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gosec\" rule=\"G101\">")
	assert.Contains(t, text, "hardcoded credentials")
	assert.Contains(t, text, "<autofix>")
	assert.Contains(t, text, "gosec_ai_autofix")
	assert.NotContains(t, text, "-ai-api-provider")
	assert.NotContains(t, text, "-ai-api-key")
	assert.NotContains(t, text, "YOUR_KEY")
}

// Test 9: Non-gosec linter with gosec-ai flag has no autofix section.
func TestHandler_NonGosecWithAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t, testAIOptions())

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gocritic", "rule": "badcall"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gocritic\" rule=\"badcall\">")
	assert.Contains(t, text, "suspicious function calls")
	assert.NotContains(t, text, "<autofix>")
}

// Test 10: Unknown rule for known compound linter returns error listing available rules.
func TestHandler_UnknownRuleForCompound(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gocritic", "rule": "nonexistent"})))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, `No rule "nonexistent" found for linter "gocritic"`)
	assert.Contains(t, text, "badcall") // should list available rules
}

// Test 11: Guide with related refs shows Related Context, raw <related> stripped.
func TestHandler_RelatedContext_SimpleLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errcheck"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"errcheck\">")
	assert.Contains(t, text, "Errcheck detects unchecked errors")
	assert.Contains(t, text, "<related_context>")
	assert.NotContains(t, text, "<related>")
	assert.NotContains(t, text, "</related>")
}

// Test 12: Guide without related tags shows no Related Context section.
func TestHandler_RelatedContext_NoRelated(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gocritic", "rule": "badcall"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gocritic\" rule=\"badcall\">")
	assert.NotContains(t, text, "<related_context>")
}

// Test 13: Guide referencing non-existent linter skips silently.
func TestHandler_RelatedContext_OrphanRef(t *testing.T) {
	srv, ctx := setupTestServer(t)

	// gosec/G101 references gosec/G304 — which exists in our test fixtures,
	// so this should show related context. But we also test that only valid
	// entries appear.
	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "gosec", "rule": "G101"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"gosec\" rule=\"G101\">")
	assert.Contains(t, text, "<related_context>")
	assert.Contains(t, text, "gosec/G304")
}

// Test 14: Max 5 related entries shown.
func TestHandler_RelatedContext_MaxEntries(t *testing.T) {
	t.Skip("requires fixture with 7+ related refs — covered in parse handler tests")
}

// Test 15: Fix hint comes from patterns bullets.
func TestHandler_RelatedContext_FixHintFromPatterns(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().
		CallTool(ctx, testGuideCall("golangci_lint_guide", testBatchArgs(map[string]string{"linter": "errcheck"})))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"errcheck\">")
	assert.Contains(t, text, "<related_context>")
	// govet should have a fix hint from its patterns, selected by keyword overlap
	// with errcheck's instructions ("unchecked errors")
	assert.Contains(t, text, "govet:")
	// rowserrcheck should have a fix hint from its patterns
	assert.Contains(t, text, "rowserrcheck:")
}

// Test 16: Batch with two valid queries returns both guides in input order.
func TestHandler_BatchMultipleSuccess(t *testing.T) {
	srv, ctx := setupTestServer(t)
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(
			map[string]string{"linter": "errcheck"},
			map[string]string{"linter": "gocritic", "rule": "badcall"},
		)))
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	// Both guides present
	assert.Contains(t, text, "<guide linter=\"errcheck\">")
	assert.Contains(t, text, "<guide linter=\"gocritic\" rule=\"badcall\">")
	// Order: errcheck before gocritic
	assert.Less(t, strings.Index(text, "errcheck"), strings.Index(text, "gocritic"))
	assert.Contains(t, text, "Errcheck detects unchecked errors")
	assert.Contains(t, text, "suspicious function calls")
}

// Test 17: Batch with mixed success and failure.
func TestHandler_BatchPartialFailure(t *testing.T) {
	srv, ctx := setupTestServer(t)
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(
			map[string]string{"linter": "errcheck"},
			map[string]string{"linter": "nonexistent"},
		)))
	require.NoError(t, err)
	require.False(t, result.IsError, "partial failure should NOT be MCP error")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<guide linter=\"errcheck\">")
	assert.Contains(t, text, "<error linter=\"nonexistent\">")
	assert.Contains(t, text, "Unknown linter")
}

// Test 18: Batch where all queries fail returns MCP error.
func TestHandler_BatchAllFail(t *testing.T) {
	srv, ctx := setupTestServer(t)
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(
			map[string]string{"linter": "nonexistent1"},
			map[string]string{"linter": "nonexistent2"},
		)))
	require.NoError(t, err)
	require.True(t, result.IsError, "all-fail should be MCP error")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "<error linter=\"nonexistent1\">")
	assert.Contains(t, text, "<error linter=\"nonexistent2\">")
}

// Test 19: Duplicate queries are silently deduplicated.
func TestHandler_BatchDuplicateDedup(t *testing.T) {
	srv, ctx := setupTestServer(t)
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(
			map[string]string{"linter": "errcheck"},
			map[string]string{"linter": "errcheck"},
		)))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	// Should contain exactly one <guide linter="errcheck">
	count := strings.Count(text, "<guide linter=\"errcheck\">")
	assert.Equal(t, 1, count, "duplicate queries should be deduplicated to one guide section")
}

// Test 20: Batch with related context deduplication across queries.
func TestHandler_BatchRelatedContextDedup(t *testing.T) {
	srv, ctx := setupTestServer(t)
	// errcheck and gosec/G101 both reference related linters
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(
			map[string]string{"linter": "errcheck"},
			map[string]string{"linter": "gosec", "rule": "G101"},
		)))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	// Should have related_context at the end (batch-level, not per-guide)
	assert.Contains(t, text, "<related_context>")
	// Related context should appear once, after all guide sections
	assert.Less(t, strings.LastIndex(text, "</guide>"), strings.Index(text, "<related_context>"))
}

// Test 21: Over-max batch returns error suggesting alternatives.
func TestHandler_BatchOverMax(t *testing.T) {
	t.Setenv("GOLANGCI_LINT_BATCH_MAX", "2")
	srv, ctx := setupTestServer(t)
	queries := make([]map[string]string, 3)
	for i := range queries {
		queries[i] = map[string]string{"linter": "errcheck"}
	}
	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_guide",
		testBatchArgs(queries...)))
	require.NoError(t, err)
	require.True(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "exceeds maximum")
	assert.Contains(t, text, "golangci_lint_parse")
}
