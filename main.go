package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
	"github.com/wavilen/golangci-lint-mcp/internal/server"
	"github.com/wavilen/golangci-lint-mcp/internal/version"

	mcpserver "github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	gosecAI := flag.Bool("gosec-ai", false, "append AI autofix hints to gosec guide responses")
	flag.Parse()

	version.Check()
	log.Printf("golangci-lint-mcp version %s", version.Server)

	err := run(guideFS, *gosecAI, os.Getenv, mcpserver.ServeStdio)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type envGetter func(string) string
type serveFunc func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error

func run(fsys fs.FS, gosecAI bool, getenv envGetter, serve serveFunc) error {
	store, err := guides.NewStore(fsys)
	if err != nil {
		return fmt.Errorf("error loading guides: %w", err)
	}

	log.Printf("loaded %d linters with guides", len(store.LinterNames()))

	opts := server.Options{
		GosecAI:         gosecAI,
		GosecAIProvider: getenv("GOSEC_AI_API_PROVIDER"),
		GosecAIKey:      getenv("GOSEC_AI_API_KEY"),
		GosecAIBaseURL:  getenv("GOSEC_AI_BASE_URL"),
		GosecAISkipSSL:  getenv("GOSEC_AI_SKIP_SSL") == "true",
	}
	if opts.GosecAI && opts.GosecAIKey == "" {
		log.Printf(
			"warning: --gosec-ai enabled but GOSEC_AI_API_KEY not set; gosec_ai_autofix tool will not be available",
		)
	}
	mcpSrv := server.NewServer(store, opts)

	serveErr := serve(mcpSrv)
	if serveErr != nil {
		return fmt.Errorf("server error: %w", serveErr)
	}
	return nil
}
