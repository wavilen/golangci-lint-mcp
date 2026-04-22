package server

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

type Options struct {
	GosecAI         bool
	GosecAIProvider string
	GosecAIKey      string
	GosecAIBaseURL  string
	GosecAISkipSSL  bool
}

func (o Options) GosecAIConfigured() bool {
	return o.GosecAI && o.GosecAIKey != ""
}

func NewServer(store *guides.Store, opts ...Options) *server.MCPServer {
	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

	mcpServer := server.NewMCPServer(
		"golangci-lint-mcp",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tool := mcp.NewTool("golangci_lint_guide",
		mcp.WithDescription(
			"Get concise, actionable guidance for fixing golangci-lint issues. "+
				"Call this when you encounter golangci-lint diagnostics such as errcheck, "+
				"govet, staticcheck SA*, gosec G*, gocritic, or revive warnings. "+
				"Returns what the issue means, how to fix it, and a code example."),
		mcp.WithString("linter",
			mcp.Required(),
			mcp.Description("The linter name (e.g., errcheck, gocritic, gosec, revive, staticcheck, govet)"),
		),
		mcp.WithString("rule",
			mcp.Description("Optional rule ID for compound linters "+
				"(e.g., 'badcall' for gocritic, 'G101' for gosec, 'SA1000' for staticcheck)"),
		),
	)

	mcpServer.AddTool(tool, makeHandler(store, opt))

	parseTool := mcp.NewTool("golangci_lint_parse",
		mcp.WithDescription(
			"Parse raw golangci-lint JSON output and return fix guidance for all diagnostics at once. "+
				"Call this when you have golangci-lint JSON output (from `golangci-lint run --output.json.path stdout`). "+
				"Returns guidance for each unique (linter, rule) pair found in the output."),
		mcp.WithString("output",
			mcp.Required(),
			mcp.Description("The raw golangci-lint JSON output string"),
		),
	)

	mcpServer.AddTool(parseTool, makeParseHandler(store, opt))

	if opt.GosecAIConfigured() {
		autofixTool := mcp.NewTool("gosec_ai_autofix",
			mcp.WithDescription(
				"Run gosec with AI-powered autofix on a file or directory. "+
					"Returns AI-generated fix suggestions for gosec findings. "+
					"The API key is handled server-side — do not pass credentials."),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description(
					"File or directory path to scan with gosec per-package (e.g., './pkg/auth/...', 'main.go'). Do NOT pass './...' — use individual package paths for correct type resolution.",
				),
			),
		)
		mcpServer.AddTool(autofixTool, makeGosecAutofixHandler(opt))
	}

	return mcpServer
}
