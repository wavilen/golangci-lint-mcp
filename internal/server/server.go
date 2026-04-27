package server

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/wavilen/golangci-lint-mcp/internal/guides"
	"github.com/wavilen/golangci-lint-mcp/internal/version"
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
		version.Server,
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

	listTool := mcp.NewTool("golangci_lint_list",
		mcp.WithDescription(
			"List all supported linters, their rules, and compound/simple classification. "+
				"Call this to discover available linters and determine which require a rule parameter. "+
				"Returns compound linters with rule counts and simple linter names."),
	)
	mcpServer.AddTool(listTool, makeListHandler(store))

	summarizeTool := mcp.NewTool("golangci_lint_summarize",
		mcp.WithDescription(
			"Summarize raw golangci-lint JSON output with package-level breakdown and strategy recommendation. "+
				"Call this when you need issue distribution data without full fix guidance. "+
				"Returns total issue counts, package breakdown sorted by count, and recommended approach."),
		mcp.WithString("output",
			mcp.Required(),
			mcp.Description("The raw golangci-lint JSON output string"),
		),
	)
	mcpServer.AddTool(summarizeTool, makeSummarizeHandler(store))

	runTool := mcp.NewTool("golangci_lint_run",
		mcp.WithDescription(
			"Run golangci-lint on a path and return parsed results with fix guidance. "+
				"For per-package paths (e.g., './pkg/auth/...'), returns full guidance with fix directions. "+
				"For full-project ('./...'), returns summary with package breakdown and strategy recommendation. "+
				"Uses the project's .golangci.yml configuration."),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("File or directory path to scan (e.g., './pkg/auth/...', './cmd/main.go', './...')"),
		),
	)
	mcpServer.AddTool(runTool, makeRunHandler(store, opt))

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
