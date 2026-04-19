package server

import (
	"context"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestServer(t *testing.T, opts ...Options) (*mcptest.Server, context.Context) {
	t.Helper()

	// Create a test FS with sample guides
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<examples>```go\nfile, _ := os.Open(\"f\")\n```</examples>"),
		},
		"guides/gocritic/badcall.md": &fstest.MapFile{
			Data: []byte("# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>"),
		},
		"guides/gocritic/appendassign.md": &fstest.MapFile{
			Data: []byte("# gocritic: appendAssign\n\n<instructions>Detects append result misassignment</instructions>"),
		},
		"guides/gocritic/commentedoutcode.md": &fstest.MapFile{
			Data: []byte("# gocritic: commentedOutCode\n\n<instructions>Detects commented-out code</instructions>"),
		},
		"guides/gosec/G101.md": &fstest.MapFile{
			Data: []byte("# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>"),
		},
		"guides/gosec/G201.md": &fstest.MapFile{
			Data: []byte("# G201\n\n<instructions>Detects SQL injection via string format</instructions>\n\n<examples>```go\nquery := fmt.Sprintf(\"SELECT * FROM users WHERE id = %s\", input)\n```</examples>"),
		},
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
		mcp.WithString("linter",
			mcp.Required(),
			mcp.Description("The linter name"),
		),
		mcp.WithString("rule",
			mcp.Description("Optional rule ID"),
		),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(tool, makeHandler(store, opt))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Known simple linter → returns guide text
func TestHandler_SimpleLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "errcheck"},
		},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

// Test 2: Known compound rule → returns guide text
func TestHandler_CompoundRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gocritic", "rule": "badcall"},
		},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "suspicious function calls")
}

// Test 3: Unknown linter → error with suggestion
func TestHandler_UnknownLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "errchek"}, // close misspelling of "errcheck"
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Unknown linter")
	assert.Contains(t, text, "errchek")
	// Should suggest errcheck as close match via Levenshtein
	assert.Contains(t, text, "errcheck")
}

// Test 4: Compound linter without rule → lists rules
func TestHandler_CompoundNoRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gocritic"},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gocritic")
	assert.Contains(t, text, "rule")
	// Should list available rules
	assert.Contains(t, text, "appendassign")
	assert.Contains(t, text, "badcall")
	assert.Contains(t, text, "commentedoutcode")
}

// Test 5: Missing linter parameter → error about missing parameter
func TestHandler_MissingLinter(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "missing")
	assert.Contains(t, strings.ToLower(text), "linter")
}

// Test 6: Simple linter with rule → error "does not have sub-rules"
func TestHandler_SimpleWithRule(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "errcheck", "rule": "anything"},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "errcheck")
	assert.Contains(t, strings.ToLower(text), "does not have sub-rules")
}

// Test 7: Gosec guide without --gosec-ai flag → no autofix section
func TestHandler_GosecWithoutAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t) // default: no options

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gosec", "rule": "G101"},
		},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.NotContains(t, text, "<autofix>")
	assert.NotContains(t, text, "-ai-api-provider")
}

// Test 8: Gosec guide with --gosec-ai flag → autofix section with MCP tool pointer
func TestHandler_GosecWithAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t, Options{GosecAI: true})

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gosec", "rule": "G101"},
		},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.Contains(t, text, "<autofix>")
	assert.Contains(t, text, "gosec_ai_autofix")
	assert.NotContains(t, text, "-ai-api-provider")
	assert.NotContains(t, text, "-ai-api-key")
	assert.NotContains(t, text, "YOUR_KEY")
}

// Test 9: Non-gosec linter with --gosec-ai flag → no autofix section
func TestHandler_NonGosecWithAIFlag(t *testing.T) {
	srv, ctx := setupTestServer(t, Options{GosecAI: true})

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gocritic", "rule": "badcall"},
		},
	})
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "suspicious function calls")
	assert.NotContains(t, text, "<autofix>")
}

// Test 10: Unknown rule for known compound linter → error listing available rules
func TestHandler_UnknownRuleForCompound(t *testing.T) {
	srv, ctx := setupTestServer(t)

	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "gocritic", "rule": "nonexistent"},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, `No rule "nonexistent" found for linter "gocritic"`)
	assert.Contains(t, text, "badcall") // should list available rules
}
