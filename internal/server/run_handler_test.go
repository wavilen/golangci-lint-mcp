package server

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/mcptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// chdirProjectRoot changes to the Go module root (directory containing go.mod).
// Tests that use file-system-relative paths need this since go test runs from
// the package directory, not the project root.
func chdirProjectRoot(t *testing.T) {
	t.Helper()
	dir, err := os.Getwd()
	require.NoError(t, err)
	for {
		if _, statErr := os.Stat(filepath.Join(dir, "go.mod")); statErr == nil {
			t.Chdir(dir)
			return
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find project root (go.mod)")
		}
		dir = parent
	}
}

func setupRunTestServer(t *testing.T) (*mcptest.Server, context.Context) {
	t.Helper()

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile(
			"# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>\n\n<examples>```go\nfile, _ := os.Open(\"f\")\n```</examples>\n\n<patterns>\n- Always check error return values\n- Use comma-ok for type assertions\n</patterns>\n\n<related>govet, rowserrcheck</related>",
		),
		"guides/rowserrcheck.md": testMapFile(
			"# rowserrcheck\n\n<instructions>Checks whether Rows.Err is checked</instructions>\n\n<patterns>\n- Always check rows.Err after iterating with rows.Next\n- Use defer rows.Close() before iterating\n</patterns>",
		),
		"guides/govet.md": testMapFile(
			"# govet\n\n<instructions>Vet examines Go source code and reports suspicious constructs</instructions>\n\n<patterns>\n- Check Printf argument count matches format verbs\n- Verify composite literal field keys\n</patterns>",
		),
		"guides/gocritic/badcall.md": testMapFile(
			"# gocritic: badCall\n\n<instructions>Detects suspicious function calls</instructions>",
		),
		"guides/gosec/G101.md": testMapFile(
			"# G101\n\n<instructions>Detects hardcoded credentials</instructions>\n\n<examples>```go\npassword := \"secret123\"\n```</examples>\n\n<patterns>\n- Move credentials to environment variables\n- Use secret management tools\n</patterns>",
		),
		"guides/gosec/G304.md": testMapFile(
			"# G304\n\n<instructions>Detects file path provided as user input</instructions>\n\n<patterns>\n- Validate and sanitize file paths before use\n- Use filepath.Clean to resolve path traversal\n</patterns>",
		),
		"guides/staticcheck/SA1000.md": testMapFile(
			"# staticcheck: SA1000\n\n<instructions>Invalid regex</instructions>",
		),
	}

	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	runTool := mcp.NewTool("golangci_lint_run",
		mcp.WithDescription("Run golangci-lint on a path"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to scan")),
	)

	mcpServer := mcptest.NewUnstartedServer(t)
	mcpServer.AddTool(runTool, makeRunHandler(store, Options{Timeout: 30 * time.Second}))
	ctx := context.Background()
	require.NoError(t, mcpServer.Start(ctx))
	t.Cleanup(mcpServer.Close)

	return mcpServer, ctx
}

// Test 1: Missing path parameter returns error.
func TestRunHandler_MissingPath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_run", map[string]any{}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "missing")
	assert.Contains(t, strings.ToLower(text), "path")
}

// Test 2: Empty path returns error.
func TestRunHandler_EmptyPath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(ctx, testGuideCall("golangci_lint_run", map[string]any{"path": "   "}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "must not be empty")
}

// Test 3: Absolute path returns error.
func TestRunHandler_AbsolutePath(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(
		ctx, testGuideCall("golangci_lint_run",
			map[string]any{"path": "/usr/local/src"}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "relative path")
	assert.Contains(t, strings.ToLower(text), "absolute")
}

// Test 4: Path traversal returns error.
func TestRunHandler_PathTraversal(t *testing.T) {
	srv, ctx := setupRunTestServer(t)

	result, err := srv.Client().CallTool(
		ctx, testGuideCall("golangci_lint_run",
			map[string]any{"path": "../../../etc/passwd"}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")

	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, strings.ToLower(text), "traverse")
}

// Test 5: Binary not installed returns helpful error.
func TestRunHandler_BinaryNotInstalled(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{Timeout: 30 * time.Second})
	ctx := context.Background()

	// If binary is installed, the test proceeds but may hit a different code path.
	// The key assertion is that the handler doesn't panic.
	result, _ := handler(ctx, testGuideCall("golangci_lint_run", map[string]any{"path": "./..."}))
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	if strings.Contains(text, "binary not found") {
		assert.Contains(t, text, "golangci-lint binary not found in PATH")
		assert.Contains(t, text, "Install")
	}
}

// Test 6: Integration test with golangci-lint binary (skipped if not installed).
func TestRunHandler_Integration_NoIssues(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{Timeout: 30 * time.Second})
	ctx := context.Background()

	// Use a path that exists in the project's working directory
	result, err := handler(ctx, testGuideCall("golangci_lint_run",
		map[string]any{"path": "./internal/server/..."}))
	require.NoError(t, err)
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	assert.True(t,
		strings.Contains(text, "Auto-fix applied") ||
			strings.Contains(text, "<summary>"),
		"expected structured response, got: %s", text)
}

// Test 7: Full-project path "./..." produces structured response (routing is issue-count-based).
func TestRunHandler_FullProjectPath(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	handler := makeRunHandler(store, Options{Timeout: 30 * time.Second})
	ctx := context.Background()

	result, err := handler(ctx, testGuideCall("golangci_lint_run",
		map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.NotNil(t, result)

	text := result.Content[0].(mcp.TextContent).Text
	if strings.Contains(text, "<summary>") {
		// Routing is issue-count-based: >30 unique → summary-only, ≤30 → per-issue guidance.
		if strings.Contains(text, "Total issues:") {
			// Full-project response (>30 unique issues)
			assert.Contains(t, text, "Unique diagnostics:")
			assert.Contains(t, text, "Strategy:")
			assert.NotContains(t, text, "## errcheck:")
		} else {
			// Per-package response (≤30 unique issues)
			assert.Contains(t, text, "Unique diagnostics:")
			assert.Contains(t, text, "Strategy:")
		}
	}
}

// Test 8: Panic detection validates the panic string check pattern.
func TestPanicDetection(t *testing.T) {
	stderr := "runtime error: invalid memory address\npanic: runtime error: invalid memory address"
	assert.Contains(t, stderr, "panic:")

	noPanic := "some normal stderr output"
	assert.NotContains(t, noPanic, "panic:")
}

func TestBuildPanicResponse_WithLinter(t *testing.T) {
	stderr := "panic: runtime error: index out of range\n\ngoroutine 1 [running]:\ngithub.com/golangci/golangci-lint/pkg/golinters/exhaustruct.analyze(0x0)\n\t/build/pkg/golinters/exhaustruct/exhaustruct.go:42 +0x123"
	resp := BuildPanicResponse(stderr)
	assert.Contains(t, resp, "<summary>")
	assert.Contains(t, resp, "golangci-lint crashed")
	assert.Contains(t, resp, "runtime error: index out of range")
	assert.Contains(t, resp, "exhaustruct")
	assert.Contains(t, resp, "disable:\n       - exhaustruct")
	assert.Contains(t, resp, "golangci_lint_run with path:")
}

func TestBuildPanicResponse_WithoutLinter(t *testing.T) {
	stderr := "panic: something broke\nno useful stack trace"
	resp := BuildPanicResponse(stderr)
	assert.Contains(t, resp, "something broke")
	assert.NotContains(t, resp, "disable:")
	assert.Contains(t, resp, "Update golangci-lint")
}

func TestBuildPanicResponse_NoPanicMessage(t *testing.T) {
	stderr := "panic:\ngoroutine 1 [running]:\nsome/path.go:1"
	resp := BuildPanicResponse(stderr)
	assert.Contains(t, resp, "<summary>")
	assert.Contains(t, resp, "golangci-lint crashed")
}

func TestExtractPanicLinter(t *testing.T) {
	tests := []struct {
		name     string
		stderr   string
		expected string
	}{
		{
			"gocritic in golinters path",
			"panic: foo\n\tgithub.com/golangci/golangci-lint/pkg/golinters/gocritic.Wrap(0x)",
			"gocritic",
		},
		{
			"revive in golinters path",
			"panic: foo\n\tgithub.com/golangci/golangci-lint/pkg/golinters/revive.New(0x)",
			"revive",
		},
		{
			"no known linter",
			"panic: foo\n\tsome/random/package.Func()",
			"",
		},
		{
			"forbidigo",
			"panic: foo\n\tgithub.com/golangci/golangci-lint/pkg/golinters/forbidigo.run(0x)",
			"forbidigo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, extractPanicLinter(tt.stderr))
		})
	}
}

// parsePartialOutput tests

// Test: Full wrapped JSON returns issues from partial output.
func TestParsePartialOutput_ValidJSON(t *testing.T) {
	stdout := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}},{"FromLinter":"govet","Text":"Printf issue","Pos":{"Filename":"main.go","Line":2,"Column":1}}],"Report":{}}`
	issues := ParsePartialOutput(stdout)
	assert.Len(t, issues, 2)
	assert.Equal(t, "errcheck", issues[0].FromLinter)
	assert.Equal(t, "govet", issues[1].FromLinter)
}

// Test: NDJSON lines parse into issues from partial output.
func TestParsePartialOutput_NDJSONFallback(t *testing.T) {
	stdout := `{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}}
{"FromLinter":"gosec","Text":"G101: hardcoded credentials","Pos":{"Filename":"main.go","Line":2,"Column":1}}`
	issues := ParsePartialOutput(stdout)
	assert.Len(t, issues, 2)
	assert.Equal(t, "errcheck", issues[0].FromLinter)
	assert.Equal(t, "gosec", issues[1].FromLinter)
}

// Test: Garbage input returns empty slice.
func TestParsePartialOutput_Unparseable(t *testing.T) {
	stdout := "some random garbage output\nnot json at all"
	issues := ParsePartialOutput(stdout)
	assert.Empty(t, issues)
}

// Test: Timeout message includes partial issue count.
func TestBuildTimeoutMessage_WithPartialIssues(t *testing.T) {
	stdout := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}}],"Report":{}}`
	issues := ParsePartialOutput(stdout)
	msg := BuildTimeoutMessage(30*time.Second, issues)
	assert.Contains(t, msg, "timed out")
	assert.Contains(t, msg, "1 issues collected before timeout")
	assert.Contains(t, msg, "specific package path")
}

// Test: Timeout message with 0 issues.
func TestBuildTimeoutMessage_NoPartialIssues(t *testing.T) {
	msg := BuildTimeoutMessage(30*time.Second, nil)
	assert.Contains(t, msg, "timed out")
	assert.Contains(t, msg, "0 issues collected before timeout")
}

// Test: BuildResponse with subagent-per-package strategy skips guidance.
func TestBuildResponse_SubagentPerPackage_SkipsGuidance(t *testing.T) {
	issues := make([]LintIssue, 0, 35)
	for idx := range 35 {
		issues = append(issues, LintIssue{
			FromLinter: "errcheck",
			Text:       fmt.Sprintf("RULE%02d: unchecked error", idx),
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: fmt.Sprintf("pkg/%c/file.go", 'a'+idx%5), Line: idx, Column: 1},
		})
	}

	// 5 packages → subagent-per-package strategy
	strategyResult := AnalyzeStrategy(issues)
	assert.Equal(t, "subagent-per-package", strategyResult.StrategyName)

	result := BuildResponse(strategyResult, ResponseConfig{AutoFixApplied: true})

	assert.Contains(t, result, "<summary>")
	assert.Contains(t, result, "<strategy_instructions>")
	assert.NotContains(t, result, "<guidance>")
	assert.NotContains(t, result, "<related_context>")
}

// BuildResponse tests

// Test: BuildResponse with SummaryOnly (high-volume) produces summary + strategy, no guidance.
func TestBuildResponse_SummaryOnly(t *testing.T) {
	// 35 unique issues → SummaryOnly=true (>30 threshold)
	issues := make([]LintIssue, 0, 35)
	for idx := range 35 {
		issues = append(issues, LintIssue{
			FromLinter: fmt.Sprintf("linter_%d", idx),
			Text:       fmt.Sprintf("RULE%02d: error", idx),
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: idx, Column: 1},
		})
	}

	strategyResult := AnalyzeStrategy(issues)
	result := BuildResponse(strategyResult, ResponseConfig{Path: "./...", AutoFixApplied: true})

	assert.Contains(t, result, "<summary>")
	assert.Contains(t, result, "golangci-lint results for ./...")
	assert.Contains(t, result, "Total issues: 35")
	assert.Contains(t, result, "Unique diagnostics: 35")
	assert.Contains(t, result, "Call golangci_lint_run with a specific package path")
	assert.Contains(t, result, "</summary>")
	assert.NotContains(t, result, "<guidance>")
	assert.True(t, strategyResult.SummaryOnly)
}

// Test: BuildResponse with single-agent strategy produces guidance.
func TestBuildResponse_SingleAgentWithGuidance(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "SA1001: unchecked error",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 1, Column: 1}},
		{FromLinter: "govet", Text: "SA1002: printf mismatch",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 2, Column: 1}},
	}

	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile("# errcheck\n\n<instructions>test</instructions>"),
		"guides/govet.md":    testMapFile("# govet\n\n<instructions>test</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)

	strategyResult := AnalyzeStrategy(issues)
	result := BuildResponse(
		strategyResult,
		ResponseConfig{Path: "./pkg/main.go", Store: store, IncludeGuidance: true, AutoFixApplied: true},
	)

	assert.Contains(t, result, "<summary>")
	assert.Contains(t, result, "golangci-lint results for ./pkg/main.go")
	assert.Contains(t, result, "<guidance>")
	assert.Contains(t, result, "errcheck")
	assert.Contains(t, result, "govet")
	assert.Contains(t, result, "</guidance>")
	assert.NotContains(t, result, "<strategy_instructions>")
}

// Test: BuildResponse with includeGuidance=false skips guidance (summarize handler).
func TestBuildResponse_IncludeGuidanceFalse(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "SA1001: error",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 1, Column: 1}},
	}

	strategyResult := AnalyzeStrategy(issues)
	result := BuildResponse(strategyResult, ResponseConfig{})

	assert.Contains(t, result, "<summary>")
	assert.Contains(t, result, "Unique diagnostics: 1")
	assert.Contains(t, result, "Total issues: 1")
	assert.NotContains(t, result, "<guidance>")
}

// Test: BuildResponse with subagent-per-package strategy, no guidance.
func TestBuildResponse_SubagentPerPackage(t *testing.T) {
	issues := make([]LintIssue, 0, 5)
	for idx := range 5 {
		issues = append(issues, LintIssue{
			FromLinter: fmt.Sprintf("linter_%d", idx),
			Text:       fmt.Sprintf("RULE%02d: error", idx),
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: fmt.Sprintf("pkg_%d/file.go", idx), Line: idx, Column: 1},
		})
	}

	strategyResult := AnalyzeStrategy(issues)
	// 5 packages → subagent-per-package
	assert.Equal(t, "subagent-per-package", strategyResult.StrategyName)

	result := BuildResponse(strategyResult, ResponseConfig{Path: "./...", AutoFixApplied: true})
	assert.Contains(t, result, "<summary>")
	assert.Contains(t, result, "<strategy_instructions>")
	assert.NotContains(t, result, "<guidance>")
}

// Test: BuildResponse without path still shows total issues and packages.
func TestBuildResponse_NoPathStillShowsTotalAndPackages(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "SA1001: error",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 1, Column: 1}},
	}

	strategyResult := AnalyzeStrategy(issues)
	result := BuildResponse(strategyResult, ResponseConfig{})

	assert.Contains(t, result, "<summary>")
	assert.NotContains(t, result, "golangci-lint results for")
	assert.Contains(t, result, "Total issues: 1")
	assert.Contains(t, result, "Unique diagnostics: 1")
	assert.Contains(t, result, "Packages affected:")
}

// Test: BuildResponse with issues shows "Auto-fix applied. N issues remain.".
func TestBuildResponse_AutoFixApplied_IssuesRemain(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "SA1001: unchecked error",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 1, Column: 1}},
		{FromLinter: "govet", Text: "SA1002: printf mismatch",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/main.go", Line: 2, Column: 1}},
	}

	strategyResult := AnalyzeStrategy(issues)
	// Use includeGuidance=false since store is nil; autoFixApplied=true to test auto-fix message
	result := BuildResponse(strategyResult, ResponseConfig{Path: "./pkg/main.go", AutoFixApplied: true})

	assert.Contains(t, result, "Auto-fix applied.")
	assert.Contains(t, result, "2 issues remain.")
	assert.Contains(t, result, "<summary>")
}

// Test: BuildResponse with zero issues shows "Auto-fix applied. No issues remain.".
func TestBuildResponse_AutoFixApplied_NoIssuesRemain(t *testing.T) {
	// Edge case: empty issues slice passed to BuildResponse
	// (normally caught by early return in handlers, but BuildResponse should handle it)
	issues := []LintIssue{}

	strategyResult := AnalyzeStrategy(issues)
	result := BuildResponse(strategyResult, ResponseConfig{Path: "./...", AutoFixApplied: true})

	assert.Contains(t, result, "Auto-fix applied.")
	assert.Contains(t, result, "No issues remain.")
	assert.Contains(t, result, "<summary>")
}

// --- isExternalModule tests (D-01, D-02) ---

func TestIsExternalModule(t *testing.T) {
	chdirProjectRoot(t)
	tests := []struct {
		name string
		path string
		want bool
	}{
		{"root path ./...", "./...", false},
		{"dot path", ".", false},
		{"empty path", "", false},
		{"normal subdirectory no go.mod", "./internal/server", false},
		{"normal subdirectory with trailing /...", "./internal/server/...", false},
		{"external module e2e/testdata/large", "e2e/testdata/large", true},
		{"external module with leading ./", "./e2e/testdata/large", true},
		{"external module with trailing /", "./e2e/testdata/large/", true},
		{"external module with trailing /...", "./e2e/testdata/large/...", true},
		{"external module e2e/testdata/simple", "e2e/testdata/simple", true},
		{"external module e2e/testdata/multipkg", "e2e/testdata/multipkg", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isExternalModule(tt.path))
		})
	}
}

func TestExecuteLint_ExternalModule(t *testing.T) {
	chdirProjectRoot(t)
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	_, lookErr := exec.LookPath("golangci-lint")
	if lookErr != nil {
		t.Skip("golangci-lint not installed — skipping integration test")
	}

	ctx := context.Background()
	// This path is an external module (e2e/testdata/large has its own go.mod)
	result := ExecuteLint(ctx, "e2e/testdata/large", 60*time.Second)

	assert.False(t, result.NotPath, "golangci-lint binary should be found")
	assert.False(t, result.TimedOut, "should not time out")
	assert.NoError(t, result.JSONErr, "JSON parse should succeed: %v", result.JSONErr)
	// The large fixture has known issues — verify some were found
	assert.NotEmpty(t, result.Parsed.Issues, "expected issues in large fixture")
}
