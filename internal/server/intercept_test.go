package server

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

// setupInterceptTestStore creates a minimal guide store for intercept tests.
func setupInterceptTestStore(t *testing.T) *guides.Store {
	t.Helper()
	testFS := fstest.MapFS{
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
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)
	return store
}

func TestIntercept_ValidateRunPath_ValidPath(t *testing.T) {
	cleaned, err := ValidateRunPath("./...")
	require.NoError(t, err)
	assert.Equal(t, "./...", cleaned)
}

func TestIntercept_ValidateRunPath_PackagePath(t *testing.T) {
	cleaned, err := ValidateRunPath("./pkg/auth/...")
	require.NoError(t, err)
	assert.Equal(t, "pkg/auth/...", cleaned)
}

func TestIntercept_ValidateRunPath_Empty(t *testing.T) {
	_, err := ValidateRunPath("")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must not be empty")
}

func TestIntercept_ValidateRunPath_Absolute(t *testing.T) {
	_, err := ValidateRunPath("/usr/local/src")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "absolute path")
}

func TestIntercept_ValidateRunPath_Traversal(t *testing.T) {
	_, err := ValidateRunPath("../../../etc/passwd")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "traverse")
}

func TestIntercept_DeduplicateIssues(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "Error return value not checked"},
		{FromLinter: "errcheck", Text: "Error return value not checked"},
		{FromLinter: "govet", Text: "Printf issue"},
	}
	unique := DeduplicateIssues(issues)
	assert.Len(t, unique, 2)
}

func TestIntercept_ExtractRule(t *testing.T) {
	assert.Equal(t, "G101", ExtractRule("G101: hardcoded credentials"))
	assert.Empty(t, ExtractRule("no rule here"))
}

func TestIntercept_BuildResponse_PerPackage(t *testing.T) {
	store := setupInterceptTestStore(t)

	issues := []LintIssue{
		{
			FromLinter: "errcheck",
			Text:       "Error return value not checked",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "main.go", Line: 1, Column: 1},
		},
	}

	strategyResult := AnalyzeStrategy(issues)
	response := BuildResponse(strategyResult, "", store, Options{}, true, true)

	assert.Contains(t, response, "<summary>")
	assert.Contains(t, response, "<guidance>")
	assert.Contains(t, response, "errcheck")
}

func TestIntercept_BuildResponse_SummaryOnly(t *testing.T) {
	// 35 unique issues → SummaryOnly
	issues := make([]LintIssue, 0, 35)
	for idx := range 35 {
		issues = append(issues, LintIssue{
			FromLinter: "linter_" + string(rune('A'+idx%26)) + string(rune('0'+idx/26)),
			Text:       "R001: some issue",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/a/foo.go", Line: idx, Column: 1},
		})
	}

	strategyResult := AnalyzeStrategy(issues)
	response := BuildResponse(strategyResult, "./...", nil, Options{}, true, true)

	assert.Contains(t, response, "<summary>")
	assert.Contains(t, response, "Total issues:")
	assert.NotContains(t, response, "<guidance>")
}

func TestIntercept_RecommendStrategy(t *testing.T) {
	name, _ := RecommendStrategy(10, 2)
	assert.Equal(t, "single-agent", name)

	name, _ = RecommendStrategy(50, 2)
	assert.Equal(t, "subagent-per-file", name)

	name, _ = RecommendStrategy(10, 5)
	assert.Equal(t, "subagent-per-package", name)
}

func TestIntercept_BuildPanicResponse(t *testing.T) {
	stderr := "panic: runtime error: index out of range\n\ngoroutine 1 [running]:\n" +
		"github.com/golangci/golangci-lint/pkg/golinters/exhaustruct.analyze(0x0)"
	resp := BuildPanicResponse(stderr)
	assert.Contains(t, resp, "<summary>")
	assert.Contains(t, resp, "golangci-lint crashed")
}

func TestIntercept_IssueCountThreshold(t *testing.T) {
	assert.Equal(t, 30, IssueCountThreshold)
}

func TestIntercept_ParsePartialOutput(t *testing.T) {
	stdout := `{"Issues":[{"FromLinter":"errcheck","Text":"Error return value not checked","Pos":{"Filename":"main.go","Line":1,"Column":1}}],"Report":{}}`
	issues := ParsePartialOutput(stdout)
	assert.Len(t, issues, 1)
	assert.Equal(t, "errcheck", issues[0].FromLinter)
}

func TestIntercept_FullFlow_SummaryOnly(t *testing.T) {
	// Create 35 unique issues with distinct linter names → each gets a unique
	// diagnostic key after dedup, so unique count > IssueCountThreshold.
	issues := make([]LintIssue, 0, 35)
	for idx := range 35 {
		issues = append(issues, LintIssue{
			FromLinter: "linter_" + string(rune('A'+idx%26)) + string(rune('0'+idx/26)),
			Text:       "R001: some issue",
			Pos: struct {
				Filename string `json:"Filename"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			}{Filename: "pkg/a/foo.go", Line: idx, Column: 1},
		})
	}
	unique := DeduplicateIssues(issues)
	assert.Greater(t, len(unique), IssueCountThreshold,
		"expected >%d unique issues, got %d", IssueCountThreshold, len(unique))

	strategyResult := AnalyzeStrategy(issues)
	response := BuildResponse(strategyResult, "./...", nil, Options{}, true, true)

	assert.Contains(t, response, "<summary>")
	assert.NotContains(t, response, "<guidance>")
}
