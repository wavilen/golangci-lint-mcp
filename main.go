package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/wavilen/golangci-lint-mcp/cmd"
	"github.com/wavilen/golangci-lint-mcp/internal/guides"
	"github.com/wavilen/golangci-lint-mcp/internal/server"
	"github.com/wavilen/golangci-lint-mcp/internal/version"

	mcpserver "github.com/mark3labs/mcp-go/server"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	// Subcommand routing: intercept → RunIntercept, otherwise MCP stdio server.
	if len(os.Args) > 1 && os.Args[1] == "intercept" {
		if err := cmd.RunIntercept(guideFS, os.Args[2:], os.Stdout, os.Stderr); err != nil {
			os.Exit(1)
		}
		return
	}

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
		Timeout:         parseTimeout(getenv("GOLANGCI_LINT_TIMEOUT")),
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

const defaultTimeout = 300 * time.Second

func parseTimeout(raw string) time.Duration {
	if raw == "" {
		return defaultTimeout
	}
	d, err := time.ParseDuration(raw)
	if err == nil {
		return d
	}
	secs, err := strconv.Atoi(raw)
	if err == nil && secs > 0 {
		return time.Duration(secs) * time.Second
	}
	return defaultTimeout
}
