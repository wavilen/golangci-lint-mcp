package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"time"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
	"github.com/wavilen/golangci-lint-mcp/internal/server"
)

const defaultInterceptTimeout = 300 * time.Second

// filterStderrNoise removes noisy golangci-lint internal log lines from stderr.
// Lines containing "[runner/exclusion_paths]" or "[runner/exclusion_rules]" are
// suppressed; all other output (real errors, panics, warnings) passes through.
func filterStderrNoise(stderr string) string {
	if stderr == "" {
		return ""
	}
	lines := strings.Split(stderr, "\n")
	filtered := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.Contains(line, "[runner/exclusion_paths]") || strings.Contains(line, "[runner/exclusion_rules]") {
			continue
		}
		filtered = append(filtered, line)
	}
	result := strings.Join(filtered, "\n")
	return result
}

// lintRunFunc allows tests to override the lint execution.
//
//nolint:gochecknoglobals // Package-level function variable for testability — same pattern as main.go serveFunc.
var lintRunFunc = server.ExecuteLint

// parseRawFlag removes --raw/-raw flags from args and reports whether --raw was present.
func parseRawFlag(args []string) ([]string, bool) {
	raw := false
	cleaned := make([]string, 0, len(args))
	for _, arg := range args {
		if arg == "--raw" || arg == "-raw" {
			raw = true
		} else {
			cleaned = append(cleaned, arg)
		}
	}
	return cleaned, raw
}

// writeUsage prints the intercept subcommand usage to stderr.
func writeUsage(stderr io.Writer) {
	fmt.Fprintln(stderr, "Usage: golangci-lint-mcp intercept [--raw] <path>")
	fmt.Fprintln(stderr, "")
	fmt.Fprintln(stderr, "Run golangci-lint and output enriched guidance.")
	fmt.Fprintln(stderr, "Flags:")
	fmt.Fprintln(stderr, "  --raw   Output raw golangci-lint JSON to stdout without parsing or summarization")
	fmt.Fprintln(stderr, "Example: golangci-lint-mcp intercept ./...")
	fmt.Fprintln(stderr, "Error: missing required argument: path")
}

// RunIntercept executes the intercept subcommand: runs golangci-lint on the
// given path, enriches output with guide content, and writes structured
// guidance to stdout. Stderr from golangci-lint passes through to stderr.
func RunIntercept(fsys fs.FS, args []string, stdout, stderr io.Writer) error {
	// Parse --raw flag from args before processing path
	args, raw := parseRawFlag(args)

	if len(args) == 0 {
		writeUsage(stderr)
		return errors.New("missing required argument: path")
	}

	path := args[0]

	cleaned, validateErr := server.ValidateRunPath(path)
	if validateErr != nil {
		fmt.Fprintln(stderr, validateErr.Error())
		return fmt.Errorf("path validation failed: %w", validateErr)
	}

	store, storeErr := guides.NewStore(fsys)
	if storeErr != nil {
		return fmt.Errorf("error loading guides: %w", storeErr)
	}

	result := lintRunFunc(context.Background(), cleaned, defaultInterceptTimeout)

	// Filter noisy internal log lines from golangci-lint stderr
	filtered := filterStderrNoise(result.Stderr)

	// Forward golangci-lint stderr to user stderr
	if filtered != "" {
		fmt.Fprint(stderr, filtered)
	}

	// Raw mode: output raw golangci-lint JSON without parsing/summarization
	if raw {
		if result.NotPath {
			fmt.Fprintln(stderr, "golangci-lint binary not found in PATH.")
			return errors.New("golangci-lint binary not found")
		}
		fmt.Fprint(stdout, result.Stdout)
		return nil
	}

	// Handle special cases
	if result.NotPath {
		fmt.Fprintln(stderr, "golangci-lint binary not found in PATH. "+
			"Install: https://golangci-lint.run/usage/install/ "+
			"Alternatively, use golangci_lint_guide(linter=\"<name>\") for per-diagnostic fix guidance.")
		return errors.New("golangci-lint binary not found")
	}

	if result.TimedOut {
		partialIssues := server.ParsePartialOutput(result.Stdout)
		msg := server.BuildTimeoutMessage(defaultInterceptTimeout, partialIssues)
		fmt.Fprintln(stderr, msg)
		return errors.New(msg)
	}

	// Detect panics in stderr
	if strings.Contains(filtered, "panic:") {
		resp := server.BuildPanicResponse(result.Stderr)
		fmt.Fprintln(stdout, resp)
		return nil
	}

	// JSON parse errors
	if result.JSONErr != nil {
		if result.HadIssues {
			fmt.Fprintln(stderr, "golangci-lint exited with error and output was not valid JSON.")
			fmt.Fprintln(stderr, "Stdout:", result.Stdout)
			fmt.Fprintln(stderr, "Stderr:", result.Stderr)
		} else {
			fmt.Fprintf(stderr, "failed to parse golangci-lint JSON output: %v\n", result.JSONErr)
		}
		return errors.New("failed to parse golangci-lint output")
	}

	// No issues found
	if len(result.Parsed.Issues) == 0 {
		fmt.Fprintln(stdout, "Auto-fix applied. No issues remain.")
		return nil
	}

	// Unified pipeline: analyze → build response (D-03)
	strategyResult := server.AnalyzeStrategy(result.Parsed.Issues)
	response := server.BuildResponse(strategyResult, server.ResponseConfig{
		Path:  cleaned,
		Store: store,
		Opts: server.Options{
			GosecAI:         false,
			GosecAIProvider: "",
			GosecAIKey:      "",
			GosecAIBaseURL:  "",
			GosecAISkipSSL:  false,
			Timeout:         0,
		},
		IncludeGuidance: true,
		AutoFixApplied:  true,
	})

	fmt.Fprintln(stdout, response)
	return nil
}
