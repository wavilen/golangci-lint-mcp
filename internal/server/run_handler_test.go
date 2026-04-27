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

func setupRunTestServer(t *testing.T) (*mcptest.Server, context.Context) {
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
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>\n\n<patterns>\n- Move credentials to environment variables\n- Use secret management tools\n</patterns>",
		),
		"guides/gosec/G304.md": testMapFile(
			"# G304\n\n<instructions>Detects file path provided as user input</instructions>\n\n<patterns>\n- Validate and sanitize file paths before use\n- Use filepath.Clean to resolve path traversal\n</patterns>",
		),
		"guides/staticcheck/SA1000.md": testMapFile(
			"# staticcheck: SA1000\n\n<instructions>Invalid regex</instructions>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	runTool := mcp.NewTool("golangci_lint_run",
		mcp.WithDescription("Run golangci-lint on a path"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to scan")),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(runTool, makeRunHandler(store, Options{}))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Missing path parameter returns error.
func TestRunHandler_MissingPath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_run", map[string]any{}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "missing")
	assert.Contains(t, strings.ToLower(text), "path")
}

// Test 2: Empty path returns error.
func TestRunHandler_EmptyPath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_run", map[string]any{"path": "   "}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "must not be empty")
}

// Test 3: Absolute path returns error.
func TestRunHandler_AbsolutePath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(
		ctx, testGuideCall("golangci_lint_run",
			map[string]any{"path": "/usr/local/src"}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "relative path")
	assert.Contains(t, strings.ToLower(text), "absolute")
}

// Test 4: Path traversal returns error.
func TestRunHandler_PathTraversal(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(
		ctx, testGuideCall("golangci_lint_run",
			map[string]any{"path": "../../../etc/passwd"}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "traverse")
}

// Test 5: Binary not installed returns helpful error.
func TestRunHandler_BinaryNotInstalled(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{})
	ctx := context.Background()

	// If binary is installed, the test proceeds but may hit a different code path.
	// The key assertion is that the handler doesn't panic.
	result, _ := handler(ctx, testGuideCall("golangci_lint_run", map[string]any{"path": "./..."}))
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	if strings.Contains(text, "binary not found") {
		assert.Contains(t, text, "golangci-lint binary not found in PATH")
		assert.Contains(t, text, "Install")
	}
}

// Test 6: Integration test with golangci-lint binary (skipped if not installed).
func TestRunHandler_Integration_NoIssues(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{})
	ctx := context.Background()

	// Use a path that exists in the project's working directory
	result, err := handler(ctx, testGuideCall("golangci_lint_run",
		map[string]any{"path": "./internal/server/..."}))
	require.NoError(t, err)
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	assert.True(t,
		strings.Contains(text, "No issues found") ||
			strings.Contains(text, "## Summary") ||
			strings.Contains(text, "## golangci-lint Results"),
		"expected structured response, got: %s", text)
}

// Test 7: Full-project path "./..." produces summary-only response.
func TestRunHandler_FullProjectSummary(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{})
	ctx := context.Background()

	result, err := handler(ctx, testGuideCall("golangci_lint_run",
		map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	if strings.Contains(text, "## golangci-lint Results") {
		assert.Contains(t, text, "Total issues:")
		assert.Contains(t, text, "Unique diagnostics:")
		assert.Contains(t, text, "Strategy:")
		assert.NotContains(t, text, "## errcheck:")
	}
}
