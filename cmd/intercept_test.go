package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"
	"testing/fstest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/server"
)

// testInterceptFS creates a minimal guide filesystem for testing.
func testInterceptFS() fstest.MapFS {
	return fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte(
				"# errcheck\n\n" +
					"<instructions>Errcheck detects unchecked errors</instructions>\n\n" +
					"<examples>```go\nfile, _ := os.Open(\"f\")\n```</examples>\n\n" +
					"<patterns>\n- Always check error return values\n</patterns>",
			),
		},
		"guides/govet.md": &fstest.MapFile{
			Data: []byte(
				"# govet\n\n" +
					"<instructions>Vet examines Go source code</instructions>\n\n" +
					"<patterns>\n- Check Printf args\n</patterns>",
			),
		},
	}
}

func TestIntercept_NoArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), nil, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "Usage:")
	assert.Contains(t, stderr.String(), "missing required argument: path")
	assert.Empty(t, stdout.String())
}

func TestIntercept_PathTraversal(t *testing.T) {
	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"../../../etc/passwd"}, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "traverse")
}

func TestIntercept_BinaryNotFound(t *testing.T) {
	// Save and restore the original lint function
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{NotPath: true}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "binary not found")
}

func TestIntercept_NoIssues(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, cleaned string, _ time.Duration) server.LintRunResult {
		assert.Equal(t, "./...", cleaned)
		return server.LintRunResult{
			Parsed: server.LintJSONResult{Issues: []server.LintIssue{}},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)
	assert.Contains(t, stdout.String(), "Auto-fix applied. No issues remain.")
}

func TestIntercept_FewIssues(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Parsed: server.LintJSONResult{
				Issues: []server.LintIssue{
					{FromLinter: "errcheck", Text: "Error return value not checked"},
					{FromLinter: "govet", Text: "Printf argument mismatch"},
				},
			},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)

	output := stdout.String()
	assert.Contains(t, output, "<summary>")
	assert.Contains(t, output, "<guidance>")
	assert.Contains(t, output, "errcheck")
}

func TestIntercept_ManyIssues(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	// Create >30 unique issues (distinct linter+rule) to trigger summary-only response.
	// Deduplication groups by linter+rule, so each must differ.
	issues := make([]server.LintIssue, 0, 35)
	for idx := range 35 {
		issues = append(issues, server.LintIssue{
			FromLinter: "linter_" + string(rune('A'+idx%26)) + string(rune('0'+idx/26)),
			Text:       fmt.Sprintf("R%03d: some issue", idx),
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/auth/handler.go", Line: idx, Column: 1},
		})
	}

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Parsed: server.LintJSONResult{Issues: issues},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)

	output := stdout.String()
	assert.Contains(t, output, "<summary>")
	assert.NotContains(t, output, "<guidance>")
	assert.Contains(t, output, "Total issues:")
}

func TestIntercept_PanicDetection(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Stderr: "panic: runtime error: index out of range\ngoroutine 1 [running]:",
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)

	output := stdout.String()
	assert.Contains(t, output, "<summary>")
	assert.Contains(t, output, "golangci-lint crashed")
}

func TestIntercept_TimedOut(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{TimedOut: true}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "timed out")
}

func TestIntercept_JSONParseError(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			JSONErr:   errors.New("invalid character"),
			HadIssues: false,
			Stdout:    "not json",
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "failed to parse")
}

func TestFilterStderrNoise_RemovesExclusionPaths(t *testing.T) {
	input := "level=warning msg=\"[runner/exclusion_paths] skipped 3 paths\"\nreal error here\nlevel=warning msg=\"[runner/exclusion_paths] skipped 5 paths\"\n"
	result := filterStderrNoise(input)
	assert.NotContains(t, result, "runner/exclusion_paths")
	assert.Contains(t, result, "real error here")
}

func TestFilterStderrNoise_RemovesExclusionRules(t *testing.T) {
	input := "level=warning msg=\"[runner/exclusion_rules] Skipped 0 issues by rules: [Source: \\\"TODO\\\", Linters: \\\"godot\\\"]\"\nreal error here\n"
	result := filterStderrNoise(input)
	assert.NotContains(t, result, "runner/exclusion_rules")
	assert.Contains(t, result, "real error here")
}

func TestFilterStderrNoise_MixedExclusionNoise(t *testing.T) {
	input := "level=warning msg=\"[runner/exclusion_paths] skipped 3 paths\"\nlevel=warning msg=\"[runner/exclusion_rules] Skipped 0 issues\"\nlevel=warning msg=\"[runner/exclusion_rules] Skipped 0 issues by rules: [Path: \\\"cmd/\\\"]\"\nreal error\n"
	result := filterStderrNoise(input)
	assert.NotContains(t, result, "runner/exclusion_paths")
	assert.NotContains(t, result, "runner/exclusion_rules")
	assert.Contains(t, result, "real error")
}

func TestFilterStderrNoise_PreservesRealErrors(t *testing.T) {
	input := "panic: runtime error\nsome real warning\n"
	result := filterStderrNoise(input)
	assert.Equal(t, input, result)
}

func TestFilterStderrNoise_EmptyInput(t *testing.T) {
	result := filterStderrNoise("")
	assert.Empty(t, result)
}

func TestIntercept_StderrNoiseFiltered(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Stderr: "level=warning msg=\"[runner/exclusion_paths] skipped 3 paths\"\nlevel=warning msg=\"[runner/exclusion_rules] Skipped 0 issues by rules\"\n",
			Parsed: server.LintJSONResult{Issues: []server.LintIssue{}},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)
	assert.NotContains(t, stderr.String(), "runner/exclusion_paths")
	assert.NotContains(t, stderr.String(), "runner/exclusion_rules")
	assert.Contains(t, stdout.String(), "Auto-fix applied")
}

func TestIntercept_RawMode_OutputsRawJSON(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	rawJSON := `{"Issues":[{"FromLinter":"errcheck","Text":"unchecked error","Pos":{"Filename":"main.go","Line":10}}]}`
	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Stdout: rawJSON,
			Parsed: server.LintJSONResult{Issues: []server.LintIssue{
				{FromLinter: "errcheck", Text: "unchecked error"},
			}},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"--raw", "./..."}, &stdout, &stderr)
	require.NoError(t, err)
	assert.JSONEq(t, rawJSON, stdout.String())
	assert.NotContains(t, stdout.String(), "<summary>")
	assert.NotContains(t, stdout.String(), "<guidance>")
}

func TestIntercept_RawMode_BinaryNotFound(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{NotPath: true}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"--raw", "./..."}, &stdout, &stderr)
	require.Error(t, err)
	assert.Contains(t, stderr.String(), "binary not found")
	assert.Empty(t, stdout.String())
}

func TestIntercept_RawFlag_AfterPath(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	rawJSON := `{"Issues":[]}`
	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Stdout: rawJSON,
			Parsed: server.LintJSONResult{Issues: []server.LintIssue{}},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./...", "--raw"}, &stdout, &stderr)
	require.NoError(t, err)
	assert.JSONEq(t, rawJSON, stdout.String())
}

func TestIntercept_StderrPassthrough(t *testing.T) {
	original := lintRunFunc
	defer func() { lintRunFunc = original }()

	lintRunFunc = func(_ context.Context, _ string, _ time.Duration) server.LintRunResult {
		return server.LintRunResult{
			Stderr: "some warning from golangci-lint\n",
			Parsed: server.LintJSONResult{Issues: []server.LintIssue{}},
		}
	}

	var stdout, stderr bytes.Buffer
	err := RunIntercept(testInterceptFS(), []string{"./..."}, &stdout, &stderr)
	require.NoError(t, err)
	assert.Contains(t, stderr.String(), "some warning from golangci-lint")
	assert.Contains(t, stdout.String(), "Auto-fix applied")
}
