package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

const golangciLintRunTimeout = 120 * time.Second

// validateRunPath validates and cleans the path parameter for golangci_lint_run.
func validateRunPath(path string) (string, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", errors.New("parameter 'path' must not be empty")
	}
	if filepath.IsAbs(path) {
		return "", errors.New("parameter 'path' must be a relative path, got absolute path")
	}
	cleaned := filepath.Clean(path)
	if cleaned == ".." || strings.HasPrefix(cleaned, "../") {
		return "", errors.New("parameter 'path' must not traverse above the current directory")
	}
	return cleaned, nil
}

type lintRunResult struct {
	stdout    string
	stderr    string
	hadIssues bool // true if golangci-lint exited non-zero (issues found)
	notPath   bool // true if binary not found
	timedOut  bool // true if context deadline exceeded
	jsonErr   error
	parsed    lintJSONResult
}

// executeLint runs golangci-lint with JSON output and returns parsed results.
func executeLint(ctx context.Context, cleaned string) lintRunResult {
	binaryPath, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		return lintRunResult{
			stdout: "", stderr: "", hadIssues: false,
			timedOut: false,
			jsonErr:  nil, parsed: lintJSONResult{Issues: []lintIssue{}}, notPath: true,
		}
	}

	ctx, cancel := context.WithTimeout(ctx, golangciLintRunTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, binaryPath, "run", "--output.json.path", "stdout", cleaned)
	var outBuf strings.Builder
	var stderrBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &stderrBuf

	runErr := cmd.Run()
	stdout := outBuf.String()
	stderr := stderrBuf.String()
	hadIssues := false

	var exitErr *exec.ExitError
	if runErr != nil && errors.As(runErr, &exitErr) {
		hadIssues = true // non-zero exit with issues is expected
	}

	timedOut := ctx.Err() == context.DeadlineExceeded

	var parsed lintJSONResult
	jsonErr := json.NewDecoder(strings.NewReader(stdout)).Decode(&parsed)

	return lintRunResult{
		stdout:    stdout,
		stderr:    stderr,
		hadIssues: hadIssues,
		timedOut:  timedOut,
		jsonErr:   jsonErr,
		parsed:    parsed,
		notPath:   false,
	}
}

// buildFullProjectResponse builds the summary-only response for full-project scans.
func buildFullProjectResponse(
	path string,
	totalIssues int,
	unique []lintIssue,
	packages []packageEntry,
	strategyName, strategyReason string,
) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "## golangci-lint Results for %s\n\n", path)
	fmt.Fprintf(&builder, "- Total issues: %d\n", totalIssues)
	fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(unique))
	fmt.Fprintf(&builder, "- Packages affected: %d\n", len(packages))
	fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", strategyName, strategyReason)

	packageBreakdown := buildPackageBreakdown(packages)
	if packageBreakdown != "" {
		fmt.Fprintf(&builder, "\n## Package Breakdown\n\n%s\n", packageBreakdown)
	}

	fmt.Fprintf(&builder, "\n## Linter Breakdown\n\n%s\n", buildLinterBreakdown(unique))

	builder.WriteString(
		"\nCall golangci_lint_run with a specific package path " +
			"(e.g., \"./pkg/auth/...\") for detailed fix guidance.\n")
	return builder.String()
}

// buildPerPackageResponse builds the full-guidance response for per-package scans.
func buildPerPackageResponse(
	unique []lintIssue,
	packages []packageEntry,
	strategyName, strategyReason string,
	store *guides.Store,
	opts Options,
) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "## Summary\n\n")
	fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(unique))
	fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", strategyName, strategyReason)

	packageBreakdown := buildPackageBreakdown(packages)
	if packageBreakdown != "" {
		fmt.Fprintf(&builder, "\n## Package Breakdown\n\n%s\n", packageBreakdown)
	}

	fmt.Fprintf(&builder, "\n## Linter Breakdown\n\n%s\n", buildLinterBreakdown(unique))

	builder.WriteString("\n---\n\n")

	for idx, issue := range unique {
		if idx > 0 {
			builder.WriteString("\n---\n\n")
		}
		writeGuideForIssue(&builder, store, opts, issue)
	}

	relatedSection := buildRelatedContext(unique, store)
	if relatedSection != "" {
		builder.WriteString("\n\n" + relatedSection)
	}

	return builder.String()
}

func makeRunHandler(
	store *guides.Store,
	opts Options,
) func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path, err := req.RequireString("path")
		if err != nil {
			return mcp.NewToolResultError(
				fmt.Sprintf("missing required parameter 'path': %v", err)), nil
		}

		cleaned, validateErr := validateRunPath(path)
		if validateErr != nil {
			return mcp.NewToolResultError(validateErr.Error()), nil
		}

		isFullProject := cleaned == "./..." || cleaned == "..."

		result := executeLint(ctx, cleaned)

		if result.notPath {
			return mcp.NewToolResultError(
				"golangci-lint binary not found in PATH. " +
					"Install: https://golangci-lint.run/usage/install/"), nil
		}

		// Handle timeout
		if result.timedOut {
			return mcp.NewToolResultError(
				"golangci-lint timed out after " + golangciLintRunTimeout.String() +
					". Try scanning per-package instead of full-project."), nil
		}

		// JSON parse error with non-zero exit: return raw output
		jsonFailed := result.jsonErr != nil
		if jsonFailed && result.hadIssues {
			return mcp.NewToolResultError(
				"golangci-lint exited with error and output was not valid JSON.\n" +
					"Stdout: " + result.stdout +
					"\nStderr: " + result.stderr), nil
		}

		//nolint:nilerr // MCP handlers return errors via NewToolResultError (first return value), not Go error convention — second return is intentionally nil.
		if jsonFailed {
			return mcp.NewToolResultError(
				fmt.Sprintf("failed to parse golangci-lint JSON output: %v",
					result.jsonErr)), nil
		}

		// No issues found
		if len(result.parsed.Issues) == 0 {
			return mcp.NewToolResultText("No issues found in " + path + "."), nil
		}

		unique := deduplicateIssues(result.parsed.Issues)
		packages := extractPackagesFromIssues(unique)
		strategyName, strategyReason := recommendStrategy(len(unique), len(packages))

		if isFullProject {
			return mcp.NewToolResultText(
				buildFullProjectResponse(
					path, len(result.parsed.Issues), unique, packages,
					strategyName, strategyReason)), nil
		}

		return mcp.NewToolResultText(
			buildPerPackageResponse(
				unique, packages, strategyName, strategyReason,
				store, opts)), nil
	}
}
