package server

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

func getProjectRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	// filename: .../golangci-lint-mcp/internal/server/compound_integration_test.go
	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func setupRealGuideServer(t *testing.T, opts ...Options) (*mcptest.Server, context.Context) {
	t.Helper()
	root := getProjectRoot(t)
	fsys := os.DirFS(root)
	store, err := guides.NewStore(fsys)
	require.NoError(t, err, "failed to create store from real guides at %s", root)

	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

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

// TestCompoundMCPHandlerIntegration verifies ~18 compound rules return correct
// content through the full MCP pipeline (Store → handler → mcptest client).
func TestCompoundMCPHandlerIntegration(t *testing.T) {
	srv, ctx := setupRealGuideServer(t)

	samples := []struct {
		linter string
		rule   string
	}{
		{"gocritic", "appendAssign"}, {"gocritic", "badCond"}, {"gocritic", "elseif"},
		{"staticcheck", "SA1019"}, {"staticcheck", "S1003"}, {"staticcheck", "QF1001"},
		{"revive", "add-constant"}, {"revive", "exported"},
		{"gosec", "G101"}, {"gosec", "G201"}, {"gosec", "G401"},
		{"govet", "atomic"}, {"govet", "copylocks"},
		{"modernize", "errorf"},
		{"testifylint", "bool-compare"},
		{"ginkgolinter", "async-assertion"},
		{"errorlint", "asserts"},
		{"grouper", "const"},
	}

	for _, sample := range samples {
		t.Run(sample.linter+"/"+sample.rule, func(t *testing.T) {
			result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
				Params: mcp.CallToolParams{
					Name: "golangci_lint_guide",
					Arguments: map[string]any{
						"linter": sample.linter,
						"rule":   sample.rule,
					},
				},
			})
			require.NoError(t, err)
			require.Len(t, result.Content, 1)

			text := result.Content[0].(mcp.TextContent).Text
			assert.NotEmpty(t, text, "should return content for %s/%s", sample.linter, sample.rule)
			assert.NotContains(t, text, "Unknown linter",
				"%s should be a known linter", sample.linter)
			assert.NotContains(t, text, "No rule",
				"%s/%s should be found", sample.linter, sample.rule)
			assert.True(t, len(text) > 50,
				"%s/%s should return substantial guide content (got %d bytes)",
				sample.linter, sample.rule, len(text))
		})
	}

	// Error path: compound linter without rule
	t.Run("compound_no_rule", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
			Params: mcp.CallToolParams{
				Name: "golangci_lint_guide",
				Arguments: map[string]any{
					"linter": "gocritic",
				},
			},
		})
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result for compound linter without rule")

		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, strings.ToLower(text), "has",
			"should mention that gocritic has rules")
		assert.Contains(t, strings.ToLower(text), "rule",
			"should mention rules")
	})

	// Error path: compound linter with unknown rule
	t.Run("compound_unknown_rule", func(t *testing.T) {
		result, err := srv.Client().CallTool(ctx, mcp.CallToolRequest{
			Params: mcp.CallToolParams{
				Name: "golangci_lint_guide",
				Arguments: map[string]any{
					"linter": "gocritic",
					"rule":   "nonexistent",
				},
			},
		})
		require.NoError(t, err)
		require.True(t, result.IsError, "expected error result for unknown rule")

		text := result.Content[0].(mcp.TextContent).Text
		assert.Contains(t, text, "No rule",
			"should report no rule found")
	})
}
