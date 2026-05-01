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

// IssueCountThreshold controls routing: ≤threshold → full guidance, >threshold → summary-only.
const IssueCountThreshold = 30

// ValidateRunPath validates and cleans the path parameter for golangci_lint_run.
func ValidateRunPath(path string) (string, error) {
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

// ParsePartialOutput attempts to extract lint issues from partial stdout
// collected before a timeout. Tries full JSON parse first, then NDJSON fallback.
func ParsePartialOutput(stdout string) []LintIssue {
	// Try full JSON parse
	var result LintJSONResult
	if json.Unmarshal([]byte(stdout), &result) == nil && len(result.Issues) > 0 {
		return result.Issues
	}
	// Try NDJSON fallback
	ndjsonResult := parseNDJSON(stdout)
	return ndjsonResult.Issues
}

// BuildTimeoutMessage constructs the timeout error message with partial issue count.
func BuildTimeoutMessage(timeout time.Duration, partialIssues []LintIssue) string {
	msg := fmt.Sprintf("golangci-lint timed out after %s.", timeout.String())
	if len(partialIssues) > 0 {
		msg += fmt.Sprintf(" %d issues collected before timeout.", len(partialIssues))
	} else {
		msg += " 0 issues collected before timeout."
	}
	msg += " Try scanning a specific package path (e.g., './pkg/auth/...') for faster results."
	return msg
}

// LintRunResult holds the result of a golangci-lint execution.
type LintRunResult struct {
	Stdout    string         // Stdout holds raw stdout from golangci-lint.
	Stderr    string         // Stderr holds raw stderr from golangci-lint.
	HadIssues bool           // HadIssues is true if golangci-lint exited non-zero (issues found).
	NotPath   bool           // NotPath is true if golangci-lint binary not found in PATH.
	TimedOut  bool           // TimedOut is true if context deadline exceeded.
	JsonErr   error          // JsonErr is non-nil if JSON parsing of stdout failed.
	Parsed    LintJSONResult // Parsed holds the decoded issues from golangci-lint JSON output.
}

// ExecuteLint runs golangci-lint with JSON output and returns parsed results.
func ExecuteLint(ctx context.Context, cleaned string, timeout time.Duration) LintRunResult {
	binaryPath, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		return LintRunResult{
			Stdout: "", Stderr: "", HadIssues: false,
			TimedOut: false,
			JsonErr:  nil, Parsed: LintJSONResult{Issues: []LintIssue{}}, NotPath: true,
		}
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, binaryPath, "run", "--fix", "--output.json.path", "stdout", cleaned)
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

	var parsed LintJSONResult
	jsonErr := json.NewDecoder(strings.NewReader(stdout)).Decode(&parsed)

	return LintRunResult{
		Stdout:    stdout,
		Stderr:    stderr,
		HadIssues: hadIssues,
		TimedOut:  timedOut,
		JsonErr:   jsonErr,
		Parsed:    parsed,
		NotPath:   false,
	}
}

// BuildPanicResponse builds a detailed error message from golangci-lint panic output.
func BuildPanicResponse(stderr string) string {
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

// BuildResponse is the unified response builder that handles all strategy routing.
// Per D-08: replaces BuildFullProjectResponse and BuildPerPackageResponse.
// Per D-09: produces the appropriate response based on StrategyResult:
//   - SummaryOnly (high-volume): summary + strategy instructions, no guidance
//   - Subagent strategy: summary + strategy instructions, no guidance
//   - Single-agent: summary + guidance + related context
//
// Parameters:
//   - result: StrategyResult from AnalyzeStrategy
//   - path: lint target path ("./...", "./pkg/auth/...", etc.) — included in summary for run/intercept, "" for parse/summarize
//   - store: guides store for guidance generation — nil for summarize handler
//   - opts: server options (GosecAI, etc.) — zero value for intercept
//   - includeGuidance: false for summarize (never shows guidance), true for all others
//   - autoFixApplied: true when golangci-lint was run with --fix (run/intercept), false for parse/summarize
func BuildResponse(
	result StrategyResult,
	path string,
	store *guides.Store,
	opts Options,
	includeGuidance bool,
	autoFixApplied bool,
) string {
	var builder strings.Builder

	// Summary section
	fmt.Fprintf(&builder, "<summary>\n\n")
	if path != "" {
		fmt.Fprintf(&builder, "golangci-lint results for %s\n\n", path)
	}
	fmt.Fprintf(&builder, "- Total issues: %d\n", result.TotalRawIssues)
	fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(result.UniqueIssues))
	fmt.Fprintf(&builder, "- Packages affected: %d\n", len(result.Packages))
	fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", result.StrategyName, result.StrategyReason)

	// Auto-fix status (per D-05, D-06, D-07) — only for run/intercept where --fix was used
	if autoFixApplied {
		if len(result.UniqueIssues) == 0 {
			fmt.Fprintf(&builder, "\nAuto-fix applied. No issues remain.\n")
		} else {
			fmt.Fprintf(&builder, "\nAuto-fix applied. %d issues remain.\n", len(result.UniqueIssues))
		}
	}

	packageBreakdown := buildPackageBreakdown(result.Packages)
	if packageBreakdown != "" {
		fmt.Fprintf(&builder, "\n%s\n", packageBreakdown)
	}

	fmt.Fprintf(&builder, "\n%s\n", buildLinterBreakdown(result.UniqueIssues))

	// Summary-only suggestion for full-project scans
	if result.SummaryOnly && path != "" {
		builder.WriteString(
			"\nCall golangci_lint_run with a specific package path " +
				"(e.g., \"./pkg/auth/...\") for detailed fix guidance.\n")
	}

	fmt.Fprintf(&builder, "\n</summary>")

	// Route by strategy
	if result.SummaryOnly || !includeGuidance || isSubagentStrategy(result.StrategyName) {
		// Summary-only or subagent: strategy instructions, no guidance
		builder.WriteString(buildStrategyInstructions(result))
		return builder.String()
	}

	// Single-agent with guidance
	fmt.Fprintf(&builder, "\n\n<guidance>\n\n")

	if store != nil {
		for idx, issue := range result.UniqueIssues {
			if idx > 0 {
				builder.WriteString("\n---\n\n")
			}
			writeGuideForIssue(&builder, store, opts, issue)
		}
	}

	builder.WriteString("\n</guidance>")

	relatedSection := buildRelatedContext(result.UniqueIssues, store)
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

		cleaned, validateErr := ValidateRunPath(path)
		if validateErr != nil {
			return mcp.NewToolResultError(validateErr.Error()), nil
		}

		timeout := opts.Timeout
		if timeout == 0 {
			timeout = golangciLintRunDefaultTimeout
		}

		result := ExecuteLint(ctx, cleaned, timeout)

		if result.NotPath {
			return mcp.NewToolResultError(
				"golangci-lint binary not found in PATH. " +
					"Install: https://golangci-lint.run/usage/install/ " +
					"Alternatively, use golangci_lint_guide(linter=\"<name>\") for per-diagnostic fix guidance."), nil
		}

		// Handle timeout
		if result.TimedOut {
			partialIssues := ParsePartialOutput(result.Stdout)
			return mcp.NewToolResultError(BuildTimeoutMessage(timeout, partialIssues)), nil
		}

		// Detect golangci-lint panics in stderr
		if strings.Contains(result.Stderr, "panic:") {
			return mcp.NewToolResultError(BuildPanicResponse(result.Stderr)), nil
		}

		// JSON parse error with non-zero exit: return raw output
		jsonFailed := result.JsonErr != nil
		if jsonFailed && result.HadIssues {
			return mcp.NewToolResultError(
				"golangci-lint exited with error and output was not valid JSON.\n" +
					"Stdout: " + result.Stdout +
					"\nStderr: " + result.Stderr), nil
		}

		//nolint:nilerr // MCP handlers return errors via NewToolResultError (first return value), not Go error convention — second return is intentionally nil.
		if jsonFailed {
			return mcp.NewToolResultError(
				fmt.Sprintf("failed to parse golangci-lint JSON output: %v",
					result.JsonErr)), nil
		}

		// No issues found
		if len(result.Parsed.Issues) == 0 {
			return mcp.NewToolResultText("Auto-fix applied. No issues remain."), nil
		}

		// Unified pipeline: analyze → build response (D-03)
		strategyResult := AnalyzeStrategy(result.Parsed.Issues)
		return mcp.NewToolResultText(
			BuildResponse(strategyResult, path, store, opts, true, true)), nil
	}
}
