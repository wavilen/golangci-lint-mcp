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

const golangciLintRunDefaultTimeout = 300 * time.Second

// issueCountThreshold controls routing: ≤threshold → full guidance, >threshold → summary-only.
const issueCountThreshold = 30

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
	if strings.HasPrefix(cleaned, "...") {
		cleaned = "./" + cleaned
	}
	return cleaned, nil
}

// parsePartialOutput attempts to extract lint issues from partial stdout
// collected before a timeout. Tries full JSON parse first, then NDJSON fallback.
func parsePartialOutput(stdout string) []lintIssue {
	// Try full JSON parse
	var result lintJSONResult
	if json.Unmarshal([]byte(stdout), &result) == nil && len(result.Issues) > 0 {
		return result.Issues
	}
	// Try NDJSON fallback
	ndjsonResult := parseNDJSON(stdout)
	return ndjsonResult.Issues
}

// buildTimeoutMessage constructs the timeout error message with partial issue count.
func buildTimeoutMessage(timeout time.Duration, partialIssues []lintIssue) string {
	msg := fmt.Sprintf("golangci-lint timed out after %s.", timeout.String())
	if len(partialIssues) > 0 {
		msg += fmt.Sprintf(" %d issues collected before timeout.", len(partialIssues))
	} else {
		msg += " 0 issues collected before timeout."
	}
	msg += " Try scanning a specific package path (e.g., './pkg/auth/...') for faster results."
	return msg
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
func executeLint(ctx context.Context, cleaned string, timeout time.Duration) lintRunResult {
	binaryPath, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		return lintRunResult{
			stdout: "", stderr: "", hadIssues: false,
			timedOut: false,
			jsonErr:  nil, parsed: lintJSONResult{Issues: []lintIssue{}}, notPath: true,
		}
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
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

// buildPanicResponse builds a detailed error message from golangci-lint panic output.
func buildPanicResponse(stderr string) string {
	panicLine := ""
	for line := range strings.SplitSeq(stderr, "\n") {
		trimmed := strings.TrimSpace(line)
		if after, ok := strings.CutPrefix(trimmed, "panic:"); ok {
			panicLine = strings.TrimSpace(after)
			break
		}
	}

	linterName := extractPanicLinter(stderr)

	var builder strings.Builder
	builder.WriteString("<summary>\n\ngolangci-lint crashed (panic detected)\n\n")
	builder.WriteString("This is a golangci-lint bug, not an issue with your code.\n\n")

	if panicLine != "" {
		fmt.Fprintf(&builder, "### Panic message\n```\n%s\n```\n\n", panicLine)
	}

	firstStackLine := extractFirstStackLine(stderr)
	if firstStackLine != "" {
		fmt.Fprintf(&builder, "### Location\n```\n%s\n```\n\n", firstStackLine)
	}

	builder.WriteString("### Suggested fixes\n\n")

	if linterName != "" {
		fmt.Fprintf(&builder,
			"1. **Disable the crashing linter** in `.golangci.yml`:\n")
		fmt.Fprintf(&builder,
			"   ```yaml\n   linters:\n     disable:\n       - %s\n   ```\n\n", linterName)
		builder.WriteString(
			"2. **Update golangci-lint** — the panic may be fixed in a newer version:\n" +
				"   ```bash\n" +
				"   golangci-lint cache status  # check current version\n" +
				"   go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest\n" +
				"   ```\n\n")
		builder.WriteString(
			"3. **Narrow the scope** — run on a specific package instead of `./...`:\n" +
				"   ```\n" +
				"   golangci_lint_run with path: \"./pkg/auth/...\"\n" +
				"   ```\n")
	} else {
		builder.WriteString(
			"1. **Update golangci-lint** — the panic may be fixed in a newer version:\n" +
				"   ```bash\n" +
				"   go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest\n" +
				"   ```\n\n")
		builder.WriteString(
			"2. **Narrow the scope** — run on a specific package instead of `./...`:\n" +
				"   ```\n" +
				"   golangci_lint_run with path: \"./pkg/auth/...\"\n" +
				"   ```\n\n")
		builder.WriteString(
			"3. **Check stderr** for the offending linter name in the stack trace," +
				" then disable it in `.golangci.yml`.\n")
	}

	builder.WriteString("\n</summary>")
	return builder.String()
}

// extractPanicLinter attempts to identify the linter that caused the panic from the stack trace.
func extractPanicLinter(stderr string) string {
	knownLinters := []string{
		"exhaustruct", "gocritic", "revive", "staticcheck", "errcheck",
		"gosec", "govet", "ineffassign", "typecheck", "unconvert",
		"unparam", "unused", "varcheck", "deadcode", "gofmt", "goimports",
		"goconst", "gocyclo", "gocognit", "nestif", "prealloc",
		"misspell", "lll", "dupl", "gochecknoinits", "gochecknoglobals",
		"whitespace", "wsl", "nlreturn", "godot", "godox", "depguard",
		"forbidigo", "funlen", "mnd", "interfacebloat", "ireturn",
		"maintidx", "tagliatelle", "tenv", "testpackage", "thelper",
		"wrapcheck", "exptostd", "fatcontext", "perfsprint", "protogetter",
		"sloglint", "spancheck", "intrange",
	}

	for line := range strings.SplitSeq(stderr, "\n") {
		lower := strings.ToLower(strings.TrimSpace(line))
		for _, linter := range knownLinters {
			if strings.Contains(lower, "golinters/"+linter) ||
				strings.Contains(lower, "/"+linter+".") ||
				strings.Contains(lower, linter+"{") {
				return linter
			}
		}
	}

	return ""
}

// extractFirstStackLine returns the first goroutine stack line from panic output.
func extractFirstStackLine(stderr string) string {
	lines := strings.Split(stderr, "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "goroutine ") {
			if i+1 < len(lines) {
				return strings.TrimSpace(lines[i+1])
			}
		}
	}
	return ""
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
	fmt.Fprintf(&builder, "<summary>\n\n")
	fmt.Fprintf(&builder, "golangci-lint results for %s\n\n", path)
	fmt.Fprintf(&builder, "- Total issues: %d\n", totalIssues)
	fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(unique))
	fmt.Fprintf(&builder, "- Packages affected: %d\n", len(packages))
	fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", strategyName, strategyReason)

	packageBreakdown := buildPackageBreakdown(packages)
	if packageBreakdown != "" {
		fmt.Fprintf(&builder, "\n%s\n", packageBreakdown)
	}

	fmt.Fprintf(&builder, "\n%s\n", buildLinterBreakdown(unique))

	builder.WriteString(
		"\nCall golangci_lint_run with a specific package path " +
			"(e.g., \"./pkg/auth/...\") for detailed fix guidance.\n")
	fmt.Fprintf(&builder, "\n</summary>")

	builder.WriteString(buildStrategyInstructions(strategyName, packages, totalIssues))

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
	fmt.Fprintf(&builder, "<summary>\n\n")
	fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(unique))
	fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", strategyName, strategyReason)

	packageBreakdown := buildPackageBreakdown(packages)
	if packageBreakdown != "" {
		fmt.Fprintf(&builder, "\n%s\n", packageBreakdown)
	}

	fmt.Fprintf(&builder, "\n%s\n", buildLinterBreakdown(unique))
	fmt.Fprintf(&builder, "\n</summary>\n\n")

	builder.WriteString("<guidance>\n\n")

	for idx, issue := range unique {
		if idx > 0 {
			builder.WriteString("\n---\n\n")
		}
		writeGuideForIssue(&builder, store, opts, issue)
	}

	builder.WriteString("\n</guidance>")

	relatedSection := buildRelatedContext(unique, store)
	if relatedSection != "" {
		builder.WriteString("\n\n" + relatedSection)
	}

	builder.WriteString(buildStrategyInstructions(strategyName, packages, len(unique)))

	return builder.String()
}

//nolint:gocognit // Dispatch logic is inherently multi-branch; extracting would reduce clarity.
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

		timeout := opts.Timeout
		if timeout == 0 {
			timeout = golangciLintRunDefaultTimeout
		}

		result := executeLint(ctx, cleaned, timeout)

		if result.notPath {
			return mcp.NewToolResultError(
				"golangci-lint binary not found in PATH. " +
					"Install: https://golangci-lint.run/usage/install/ " +
					"Alternatively, use golangci_lint_guide(linter=\"<name>\") for per-diagnostic fix guidance."), nil
		}

		// Handle timeout
		if result.timedOut {
			partialIssues := parsePartialOutput(result.stdout)
			return mcp.NewToolResultError(buildTimeoutMessage(timeout, partialIssues)), nil
		}

		// Detect golangci-lint panics in stderr
		if strings.Contains(result.stderr, "panic:") {
			return mcp.NewToolResultError(buildPanicResponse(result.stderr)), nil
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
		strategyName, strategyReason := recommendStrategy(len(result.parsed.Issues), len(packages))

		// Route by issue count, not path syntax
		if len(unique) > issueCountThreshold {
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
