package server

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

// Options configures server behavior.
type Options struct {
	// GosecAI enables appending an <autofix> section to gosec guide responses,
	// informing agents about gosec's built-in AI autofix capabilities.
	GosecAI bool
	// GosecAIProvider is the AI provider for gosec autofix (e.g., "gemini-2.0-flash").
	GosecAIProvider string
	// GosecAIKey is the API key for the AI provider. Read from GOSEC_AI_API_KEY env var.
	GosecAIKey string
	// GosecAIBaseURL is an optional custom base URL for the AI provider.
	GosecAIBaseURL string
	// GosecAISkipSSL skips SSL verification for the AI provider connection.
	GosecAISkipSSL bool
}

// GosecAIConfigured returns true when gosec AI is enabled and an API key is available.
func (o Options) GosecAIConfigured() bool {
	return o.GosecAI && o.GosecAIKey != ""
}

// NewServer creates an MCP server with the linter guide tool registered.
func NewServer(store *guides.Store, opts ...Options) *server.MCPServer {
	var opt Options
	if len(opts) > 0 {
		opt = opts[0]
	}

	s := server.NewMCPServer(
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

	s.AddTool(tool, makeHandler(store, opt))

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

	s.AddTool(parseTool, makeParseHandler(store, opt))

	// Conditionally register the gosec AI autofix tool when both --gosec-ai and API key are configured.
	if opt.GosecAIConfigured() {
		autofixTool := mcp.NewTool("gosec_ai_autofix",
			mcp.WithDescription(
				"Run gosec with AI-powered autofix on a file or directory. "+
					"Returns AI-generated fix suggestions for gosec findings. "+
					"The API key is handled server-side — do not pass credentials."),
			mcp.WithString("path",
				mcp.Required(),
				mcp.Description("File or directory path to scan with gosec per-package (e.g., './pkg/auth/...', 'main.go'). Do NOT pass './...' — use individual package paths for correct type resolution."),
			),
		)
		s.AddTool(autofixTool, makeGosecAutofixHandler(opt))
	}

	return s
}
