package server

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupAllToolsServer creates a single test MCP server with all 5 tools registered.
// This avoids duplicating fixture setup per test group.
func setupAllToolsServer(t *testing.T) (*mcptest.Server, context.Context) {
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
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>\n\n<patterns>\n- Move credentials to environment variables\n- Use secret management tools\n</patterns>\n\n<related>gosec/G304</related>",
		),
		"guides/gosec/G201.md": testMapFile(
			"# G201\n\n<instructions>Detects SQL injection via string format</instructions>\n\n<examples>```go\nquery := fmt.Sprintf(\"SELECT id, name FROM users WHERE id = %s\", input)\n```</examples>",
		),
		"guides/gosec/G304.md": testMapFile(
			"# G304\n\n<instructions>Detects file path provided as user input</instructions>\n\n<patterns>\n- Validate and sanitize file paths before use\n- Use filepath.Clean to resolve path traversal\n</patterns>",
		),
		"guides/staticcheck/SA1000.md": testMapFile(
			"# staticcheck: SA1000\n\n<instructions>Detects invalid regex patterns</instructions>",
		),
		"guides/staticcheck/SA2000.md": testMapFile(
			"# staticcheck: SA2000\n\n<instructions>Detects sync.WaitGroup misuse</instructions>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	guideTool := mcp.NewTool("golangci_lint_guide",
		mcp.WithDescription("Get concise guidance for fixing golangci-lint issues"),
		mcp.WithString("linter", mcp.Required(), mcp.Description("The linter name")),
		mcp.WithString("rule", mcp.Description("Optional rule ID")),
	)
	parseTool := mcp.NewTool("golangci_lint_parse",
		mcp.WithDescription("Parse golangci-lint JSON and return fix guidance"),
		mcp.WithString("output", mcp.Required(), mcp.Description("Raw golangci-lint JSON")),
	)
	listTool := mcp.NewTool("golangci_lint_list",
		mcp.WithDescription("List all supported linters"),
	)
	summarizeTool := mcp.NewTool("golangci_lint_summarize",
		mcp.WithDescription("Summarize golangci-lint JSON"),
		mcp.WithString("output", mcp.Required(), mcp.Description("Raw JSON")),
	)
	runTool := mcp.NewTool("golangci_lint_run",
		mcp.WithDescription("Run golangci-lint on a path"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to scan")),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(guideTool, makeHandler(store, Options{}))
	mcpServer.AddTool(parseTool, makeParseHandler(store, Options{}))
	mcpServer.AddTool(listTool, makeListHandler(store))
	mcpServer.AddTool(summarizeTool, makeSummarizeHandler(store))
	mcpServer.AddTool(runTool, makeRunHandler(store, Options{}))

	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// ---------- golangci_lint_guide ----------

func TestSmoke_Guide_Valid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	t.Run("simple_linter", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_guide", map[string]any{"linter": "errcheck"}))
		require.NoError(t, err)
		require.False(t, result.IsError, "expected non-error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, text, "unchecked errors")
	})

	t.Run("compound_linter_with_rule", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_guide", map[string]any{"linter": "gocritic", "rule": "badcall"}))
		require.NoError(t, err)
		require.False(t, result.IsError, "expected non-error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, text, "suspicious function calls")
	})
}

func TestSmoke_Guide_Invalid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	t.Run("missing_linter", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_guide", map[string]any{}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "missing")
	})

	t.Run("unknown_linter", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_guide", map[string]any{"linter": "nonexistent_linter_xyz"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, text, "Unknown linter")
	})

	t.Run("simple_linter_with_bogus_rule", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_guide", map[string]any{"linter": "errcheck", "rule": "bogus"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "does not have sub-rules")
	})
}

// ---------- golangci_lint_parse ----------

func TestSmoke_Parse_Valid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	validJSON := `{"Issues":[{"FromLinter":"errcheck","Text":"unchecked error","Pos":{"Filename":"main.go","Line":1,"Column":1}}],"Report":{}}`

	result, err := srv.Client().CallTool(ctx,
		testGuideCall("golangci_lint_parse", map[string]any{"output": validJSON}))
	require.NoError(t, err)
	require.False(t, result.IsError, "expected non-error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "errcheck")
}

func TestSmoke_Parse_Invalid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	t.Run("missing_output", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_parse", map[string]any{}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "missing")
	})

	t.Run("empty_output", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_parse", map[string]any{"output": ""}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
	})

	t.Run("malformed_output", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_parse", map[string]any{"output": "not json"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
	})
}

// ---------- golangci_lint_list ----------

func TestSmoke_List_Valid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	result, err := srv.Client().CallTool(ctx,
		testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	require.False(t, result.IsError, "expected non-error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Supported Linters")
	assert.Contains(t, text, "errcheck")
}

// ---------- golangci_lint_summarize ----------

func TestSmoke_Summarize_Valid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	validJSON := `{"Issues":[` +
		`{"FromLinter":"errcheck","Text":"unchecked error in auth","Pos":{"Filename":"pkg/auth/handler.go","Line":1,"Column":1}},` +
		`{"FromLinter":"govet","Text":"Printf issue","Pos":{"Filename":"pkg/db/conn.go","Line":2,"Column":1}}` +
		`],"Report":{}}`

	result, err := srv.Client().CallTool(ctx,
		testGuideCall("golangci_lint_summarize", map[string]any{"output": validJSON}))
	require.NoError(t, err)
	require.False(t, result.IsError, "expected non-error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Summary")
	assert.Contains(t, text, "Total issues")
}

func TestSmoke_Summarize_Invalid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	t.Run("missing_output", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_summarize", map[string]any{}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
	})

	t.Run("malformed_output", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_summarize", map[string]any{"output": "bad"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
	})
}

// ---------- golangci_lint_run ----------

func TestSmoke_Run_Valid(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	srv, ctx := setupAllToolsServer(t)

	// Call with valid path — may succeed with issues or no-issues
	result, err := srv.Client().CallTool(ctx,
		testGuideCall("golangci_lint_run", map[string]any{"path": "./internal/server/..."}))
	require.NoError(t, err)
	require.NotNil(t, result)
	// Result may be error (issues found) or success (no issues) — both are valid
	text := result.Content[0].(mcp.TextContent).Text
	assert.True(t,
		strings.Contains(text, "No issues found") ||
			strings.Contains(text, "## Summary") ||
			strings.Contains(text, "## golangci-lint Results"),
		"expected structured response, got: %s", text)
}

func TestSmoke_Run_Invalid(t *testing.T) {
	srv, ctx := setupAllToolsServer(t)

	t.Run("missing_path", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_run", map[string]any{}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "missing")
	})

	t.Run("empty_path", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_run", map[string]any{"path": "   "}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "must not be empty")
	})

	t.Run("absolute_path", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_run", map[string]any{"path": "/tmp"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "absolute")
	})

	t.Run("traversal_path", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx,
			testGuideCall("golangci_lint_run", map[string]any{"path": "../../../etc"}))
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result")
		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "traverse")
	})
}
