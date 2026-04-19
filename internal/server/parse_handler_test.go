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

func setupParseTestServer(t *testing.T, opts ...Options) (*mcptest.Server, context.Context) {
	t.Helper()
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<examples>```go\nfile, _ := os.Open(\"f\")\n```</examples>"),
		},
		"guides/gocritic/badCall.md": &fstest.MapFile{
			Data: []byte("# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>"),
		},
		"guides/gocritic/dupSubExpr.md": &fstest.MapFile{
			Data: []byte("# gocritic: dupSubExpr\n\n<instructions>Detects duplicate sub-expressions</instructions>"),
		},
		"guides/gosec/G101.md": &fstest.MapFile{
			Data: []byte("# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>"),
		},
		"guides/gosec/G201.md": &fstest.MapFile{
			Data: []byte("# G201\n\n<instructions>Detects SQL injection via string format</instructions>"),
		},
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
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
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
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Equal(t, 1, strings.Count(text, "## errcheck"))
}

func TestParseHandler_CompoundRuleExtraction(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"gocritic","Text":"dupSubExpr: suspicious identical LHS and RHS","Pos":{"Filename":"main.go","Line":10,"Column":5}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gocritic: dupSubExpr")
	assert.Contains(t, text, "duplicate sub-expressions")
}

func TestParseHandler_InvalidJSON(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": "not json at all"},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "invalid json")
}

func TestParseHandler_EmptyOutput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": ""},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError)
}

func TestParseHandler_EmptyIssues(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "no issues found")
}

func TestParseHandler_UnknownLinter(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"typolinter","Text":"Some issue","Pos":{"Filename":"main.go","Line":1,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Unknown linter")
	assert.Contains(t, text, "typolinter")
}

func TestParseHandler_GosecWithAIFlag(t *testing.T) {
	srv, ctx := setupParseTestServer(t, Options{GosecAI: true})
	json := `{"Issues":[{"FromLinter":"gosec","Text":"G101: Potential hardcoded credentials","Pos":{"Filename":"main.go","Line":5,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.Contains(t, text, "<autofix>")
}

func TestParseHandler_GosecWithoutAIFlag(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := `{"Issues":[{"FromLinter":"gosec","Text":"G101: Potential hardcoded credentials","Pos":{"Filename":"main.go","Line":5,"Column":1}}],"Report":{}}`
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "hardcoded credentials")
	assert.NotContains(t, text, "<autofix>")
}

func TestParseHandler_MultiLineOutput(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	json := "{\"Issues\":[{\"FromLinter\":\"errcheck\",\"Text\":\"Error return value is not checked\",\"Pos\":{\"Filename\":\"main.go\",\"Line\":10,\"Column\":5}}],\"Report\":{}}\n2 issues:\n* errcheck: 2\n"
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_parse",
			Arguments: map[string]any{"output": json},
		},
	})
	require.NoError(t, err)
	require.False(t, result.IsError)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}

func TestParseHandler_ExistingGuideToolUnchanged(t *testing.T) {
	srv, ctx := setupParseTestServer(t)
	result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "golangci_lint_guide",
			Arguments: map[string]any{"linter": "errcheck"},
		},
	})
	require.NoError(t, err)
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "Errcheck detects unchecked errors")
}
