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

func testMapFileList(data string) *fstest.MapFile {
	return &fstest.MapFile{
		Data:    []byte(data),
		Mode:    0,
		ModTime: time.Time{},
		Sys:     nil,
	}
}

func setupListTestServer(t *testing.T) (*mcptest.Server, context.Context) {
	t.Helper()

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFileList(
			"# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<patterns>\n- Always check error return values\n</patterns>",
		),
		"guides/govet.md": testMapFileList(
			"# govet\n\n<instructions>Vet examines Go source code</instructions>\n\n<patterns>\n- Check Printf args\n</patterns>",
		),
		"guides/gocritic/badcall.md": testMapFileList(
			"# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>",
		),
		"guides/gocritic/dupSubExpr.md": testMapFileList(
			"# gocritic: dupSubExpr\n\n<instructions>Detects duplicate sub-expressions</instructions>",
		),
		"guides/gosec/G101.md": testMapFileList(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>",
		),
		"guides/gosec/G201.md": testMapFileList(
			"# G201\n\n<instructions>Detects SQL injection</instructions>",
		),
		"guides/staticcheck/SA1000.md": testMapFileList(
			"# staticcheck: SA1000\n\n<instructions>Invalid regex</instructions>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	listTool := mcp.NewTool("golangci_lint_list",
		mcp.WithDescription("List all supported linters"),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(listTool, makeListHandler(store))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Returns all linters from the store.
func TestListHandler_ReturnsAllLinters(t *testing.T) {
	srv, ctx := setupListTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "## Supported Linters")
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, text, "govet")
	assert.Contains(t, text, "gocritic")
	assert.Contains(t, text, "gosec")
	assert.Contains(t, text, "staticcheck")
}

// Test 2: Compound linters show rule count > 0.
func TestListHandler_CompoundLinterHasRules(t *testing.T) {
	srv, ctx := setupListTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "### Compound Linters")
	assert.Contains(t, text, "gocritic")
	assert.Contains(t, text, "rules")
	assert.Contains(t, text, "gosec")
	assert.Contains(t, text, "staticcheck")
}

// Test 3: Simple linters appear in simple section.
func TestListHandler_SimpleLinterNoRules(t *testing.T) {
	srv, ctx := setupListTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	assert.Contains(t, text, "### Simple Linters")
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, text, "govet")
	// Simple linters should NOT have "| errcheck |" table format
	assert.NotContains(t, text, "| errcheck |")
}

// Test 4: Linters are sorted alphabetically.
func TestListHandler_SortedByName(t *testing.T) {
	srv, ctx := setupListTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text

	// errcheck should appear before govet
	errcheckIdx := strings.Index(text, "errcheck")
	govetIdx := strings.Index(text, "govet")
	assert.Greater(t, govetIdx, errcheckIdx, "govet should appear after errcheck alphabetically")
}

// Test 5: Handler works with no arguments.
func TestListHandler_NoParameters(t *testing.T) {
	srv, ctx := setupListTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_list", map[string]any{}))
	require.NoError(t, err)
	require.False(t, result.IsError, "should not return error")
}
