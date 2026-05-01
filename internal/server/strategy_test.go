package server

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractPackagesFromIssues(t *testing.T) {
	issues := []LintIssue{
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/handler.go", Line: 1, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/middleware.go", Line: 2, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/db/connection.go", Line: 3, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 4, Column: 1}},
	}

	packages := ExtractPackagesFromIssues(issues)

	assert.Len(t, packages, 3)

	// pkg/auth has 2 issues, should be first
	assert.Equal(t, "pkg/auth", packages[0].Path)
	assert.Equal(t, 2, packages[0].Count)

	// root ("." for main.go) has 1 issue — "." sorts before "pkg/" alphabetically
	assert.Equal(t, ".", packages[1].Path)
	assert.Equal(t, 1, packages[1].Count)

	// pkg/db has 1 issue
	assert.Equal(t, "pkg/db", packages[2].Path)
	assert.Equal(t, 1, packages[2].Count)
}

func TestExtractPackagesFromIssues_Empty(t *testing.T) {
	packages := ExtractPackagesFromIssues(nil)
	assert.Empty(t, packages)
}

func TestBuildPackageBreakdown(t *testing.T) {
	packages := []PackageEntry{
		{Path: "pkg/auth", Count: 15},
		{Path: "pkg/db", Count: 8},
		{Path: "pkg/util", Count: 3},
	}

	result := buildPackageBreakdown(packages)

	assert.Contains(t, result, "pkg/auth: 15 issues")
	assert.Contains(t, result, "pkg/db: 8 issues")
	assert.Contains(t, result, "pkg/util: 3 issues")
	assert.Contains(t, result, "TOTAL: 26 issues across 3 packages")
}

func TestBuildPackageBreakdown_Empty(t *testing.T) {
	result := buildPackageBreakdown(nil)
	assert.Empty(t, result)
}

func TestRecommendStrategy_SingleAgent(t *testing.T) {
	name, reason := RecommendStrategy(20, 2)
	assert.Equal(t, "single-agent", name)
	assert.Contains(t, reason, "≤20")
	assert.Contains(t, reason, "2 packages")
	assert.Contains(t, reason, "single-agent flow")
}

func TestRecommendStrategy_SingleAgent_EdgeCase(t *testing.T) {
	name, _ := RecommendStrategy(30, 3)
	assert.Equal(t, "single-agent", name)
}

func TestRecommendStrategy_SubagentPerFile(t *testing.T) {
	name, reason := RecommendStrategy(45, 2)
	assert.Equal(t, "subagent-per-file", name)
	assert.Contains(t, reason, ">45")
	assert.Contains(t, reason, "2 packages")
	assert.Contains(t, reason, "subagent-per-file strategy")
}

func TestRecommendStrategy_SubagentPerPackage(t *testing.T) {
	name, reason := RecommendStrategy(5, 5)
	assert.Equal(t, "subagent-per-package", name)
	assert.Contains(t, reason, "5 issues")
	assert.Contains(t, reason, "5 packages")
}

func TestRecommendStrategy_SubagentPerPackageManyIssues(t *testing.T) {
	name, reason := RecommendStrategy(100, 6)
	assert.Equal(t, "subagent-per-package", name)
	assert.Contains(t, reason, "100 issues")
	assert.Contains(t, reason, "6 packages")
}

func TestBuildPackageBreakdown_Ordering(t *testing.T) {
	// Input is already sorted by count desc, path asc (as ExtractPackagesFromIssues returns)
	packages := []PackageEntry{
		{Path: "pkg/c", Count: 10},
		{Path: "pkg/a", Count: 5},
		{Path: "pkg/b", Count: 5},
	}

	result := buildPackageBreakdown(packages)
	lines := strings.Split(result, "\n")

	// First line should be pkg/c (highest count)
	assert.Contains(t, lines[0], "pkg/c: 10 issues")
	// Next two should be sorted by path (a before b) — already sorted by caller
	assert.Contains(t, lines[1], "pkg/a: 5 issues")
	assert.Contains(t, lines[2], "pkg/b: 5 issues")
}

// --- Task 1 TDD Tests (RED phase) ---

func TestExtractGuideRefsByDir(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/handler.go", Line: 1, Column: 1}},
		{FromLinter: "staticcheck", Text: "SA1000: invalid regex", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/handler.go", Line: 2, Column: 1}},
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/middleware.go", Line: 3, Column: 1}},
	}

	result := extractGuideRefsByDir(issues)

	// pkg/auth should have 2 unique guide refs: errcheck and staticcheck/SA1000
	authRefs, ok := result["pkg/auth"]
	require.True(t, ok, "expected pkg/auth key in result")
	require.Len(t, authRefs, 2, "expected 2 unique guide refs for pkg/auth")

	// Verify the refs are sorted by linter then rule
	assert.Equal(t, "errcheck", authRefs[0].Linter)
	assert.Empty(t, authRefs[0].Rule)
	assert.Equal(t, "staticcheck", authRefs[1].Linter)
	assert.Equal(t, "SA1000", authRefs[1].Rule)
}

func TestExtractGuideRefsByDir_Deduplication(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/a.go", Line: 1, Column: 1}},
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/b.go", Line: 2, Column: 1}},
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/c.go", Line: 3, Column: 1}},
	}

	result := extractGuideRefsByDir(issues)

	authRefs := result["pkg/auth"]
	require.Len(t, authRefs, 1, "multiple errcheck issues in same dir should produce single guide ref")
	assert.Equal(t, "errcheck", authRefs[0].Linter)
}

func TestExtractGuideRefsByFile(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 1, Column: 1}},
		{FromLinter: "govet", Text: "Printf argument mismatch", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 2, Column: 1}},
	}

	result := extractGuideRefsByFile(issues)

	mainRefs, ok := result["main.go"]
	require.True(t, ok, "expected main.go key in result")
	require.Len(t, mainRefs, 2, "expected 2 unique guide refs for main.go")

	// Sorted by linter then rule
	assert.Equal(t, "errcheck", mainRefs[0].Linter)
	assert.Empty(t, mainRefs[0].Rule)
	assert.Equal(t, "govet", mainRefs[1].Linter)
	assert.Empty(t, mainRefs[1].Rule)
}

func TestExtractGuideRefsByFile_Deduplication(t *testing.T) {
	issues := []LintIssue{
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 1, Column: 1}},
		{FromLinter: "errcheck", Text: "Error return value not checked", Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 2, Column: 1}},
	}

	result := extractGuideRefsByFile(issues)

	mainRefs := result["main.go"]
	require.Len(t, mainRefs, 1, "multiple errcheck issues in same file should produce single guide ref")
}

func TestFormatGuideCall(t *testing.T) {
	tests := []struct {
		name     string
		ref      GuideRef
		expected string
	}{
		{
			name:     "linter only, no rule",
			ref:      GuideRef{Linter: "errcheck", Rule: ""},
			expected: `golangci_lint_guide(linter="errcheck")`,
		},
		{
			name:     "linter with rule",
			ref:      GuideRef{Linter: "staticcheck", Rule: "SA1000"},
			expected: `golangci_lint_guide(linter="staticcheck", rule="SA1000")`,
		},
		{
			name:     "gosec with rule",
			ref:      GuideRef{Linter: "gosec", Rule: "G101"},
			expected: `golangci_lint_guide(linter="gosec", rule="G101")`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, formatGuideCall(tt.ref))
		})
	}
}

func TestIsSubagentStrategy(t *testing.T) {
	assert.True(t, isSubagentStrategy("subagent-per-package"))
	assert.True(t, isSubagentStrategy("subagent-per-file"))
	assert.False(t, isSubagentStrategy("single-agent"))
	assert.False(t, isSubagentStrategy(""))
}

// --- Plan 111 Task 1: AnalyzeStrategy pipeline + fileThresholdFromEnv ---

// makeIssue creates a LintIssue with minimal boilerplate.
func makeIssue(linter, text, filename string) LintIssue {
	return LintIssue{
		FromLinter: linter,
		Text:       text,
		Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: filename, Line: 1, Column: 1},
	}
}

// makeIssues creates n identical issues in the given directory (file: dir/file.go).
func makeIssues(n int, linter, text, dir string) []LintIssue {
	issues := make([]LintIssue, n)
	for i := range issues {
		issues[i] = makeIssue(linter, text, dir+"/file.go")
	}
	return issues
}

func TestAnalyzeStrategy_SingleAgent_NoEscalation(t *testing.T) {
	// 20 raw issues across 2 packages → single-agent, no escalation
	issues := append(
		makeIssues(10, "errcheck", "error not checked", "pkg/a"),
		makeIssues(10, "errcheck", "error not checked", "pkg/b")...,
	)

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "single-agent", result.StrategyName)
	assert.Equal(t, 20, result.TotalRawIssues)
	assert.Empty(t, result.EscalatedPkgs)
	assert.Equal(t, IssueCountThreshold, result.FileThreshold)
}

func TestAnalyzeStrategy_PerPackage_NoEscalation(t *testing.T) {
	// 5 raw issues across 5 packages → per-package, no escalation (no package exceeds threshold)
	// Each issue needs unique (linter, rule) to survive dedup and produce 5 packages
	var issues []LintIssue
	for i := range 5 {
		dir := fmt.Sprintf("pkg/%d", i)
		issues = append(issues, makeIssue("errcheck", fmt.Sprintf("SA%d: error", i), dir+"/file.go"))
	}

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "subagent-per-package", result.StrategyName)
	assert.Equal(t, 5, result.TotalRawIssues)
	assert.Len(t, result.Packages, 5)
	assert.Empty(t, result.EscalatedPkgs)
}

func TestAnalyzeStrategy_PerFileStrategy_NoEscalation(t *testing.T) {
	// 45 raw issues across 2 packages → per-file strategy, EscalatedPkgs empty (per-file applies globally)
	issues := append(
		makeIssues(25, "errcheck", "error not checked", "pkg/a"),
		makeIssues(20, "errcheck", "error not checked", "pkg/b")...,
	)

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "subagent-per-file", result.StrategyName)
	assert.Equal(t, 45, result.TotalRawIssues)
	assert.Empty(t, result.EscalatedPkgs, "per-file applies globally, no per-package escalation")
}

func TestAnalyzeStrategy_PerPackage_WithEscalation(t *testing.T) {
	// 100+ raw issues across 6 packages where 2 packages have 35+ issues each → per-package with 2 escalated
	var issues []LintIssue
	// pkg/big1: 35 issues (escalated) — unique linter so it survives dedup
	issues = append(issues, makeIssues(35, "errcheck", "SA1001: error", "pkg/big1")...)
	// pkg/big2: 35 issues (escalated) — different unique linter+rule
	issues = append(issues, makeIssues(35, "govet", "composites", "pkg/big2")...)
	// 4 remaining packages with 8 issues each (not escalated)
	for i := range 4 {
		dir := fmt.Sprintf("pkg/small%d", i)
		issues = append(issues, makeIssues(8, fmt.Sprintf("linter%d", i), "error", dir+"/file.go")...)
	}

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "subagent-per-package", result.StrategyName)
	assert.Equal(t, 102, result.TotalRawIssues)
	assert.Len(t, result.Packages, 6)
	assert.Len(t, result.EscalatedPkgs, 2)
	assert.True(t, result.EscalatedPkgs["pkg/big1"])
	assert.True(t, result.EscalatedPkgs["pkg/big2"])
}

func TestAnalyzeStrategy_EscalationThreshold_Exactly30(t *testing.T) {
	// One package with exactly 30 issues → no escalation (must exceed, not meet threshold)
	issues := makeIssues(30, "errcheck", "SA1001: error", "pkg/exact")
	// Add 4 more packages with unique linterrs to trigger per-package strategy
	for i := range 4 {
		dir := fmt.Sprintf("pkg/other%d", i)
		issues = append(issues, makeIssue(fmt.Sprintf("linter%d", i), "error", dir+"/file.go"))
	}

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "subagent-per-package", result.StrategyName)
	assert.Empty(t, result.EscalatedPkgs, "exactly 30 issues should NOT trigger escalation")
}

func TestAnalyzeStrategy_EscalationThreshold_31Issues(t *testing.T) {
	// One package with 31 issues → escalation triggered
	issues := makeIssues(31, "errcheck", "SA1001: error", "pkg/over")
	for i := range 4 {
		dir := fmt.Sprintf("pkg/other%d", i)
		issues = append(issues, makeIssue(fmt.Sprintf("linter%d", i), "error", dir+"/file.go"))
	}

	result := AnalyzeStrategy(issues)

	assert.Equal(t, "subagent-per-package", result.StrategyName)
	assert.Len(t, result.EscalatedPkgs, 1)
	assert.True(t, result.EscalatedPkgs["pkg/over"])
}

func TestAnalyzeStrategy_EmptyInput(t *testing.T) {
	result := AnalyzeStrategy(nil)

	assert.Equal(t, "single-agent", result.StrategyName)
	assert.Equal(t, 0, result.TotalRawIssues)
	assert.Empty(t, result.Packages)
	assert.Empty(t, result.EscalatedPkgs)
	assert.Empty(t, result.RawIssues)
	assert.Empty(t, result.UniqueIssues)
}

func TestAnalyzeStrategy_UsesRawCount(t *testing.T) {
	// D-01 fix: RecommendStrategy must be called with len(rawIssues), not len(unique)
	// Create duplicate issues (same linter + text + file) so dedup reduces count
	// 35 raw issues → per-file. If deduped to <30 → single-agent. Test that we get per-file.
	issues := makeIssues(35, "errcheck", "error not checked", "pkg/a")

	result := AnalyzeStrategy(issues)

	// 35 raw issues → subagent-per-file (even though unique might be 1)
	assert.Equal(t, "subagent-per-file", result.StrategyName,
		"D-01: must use raw count (35) not unique count for RecommendStrategy")
	assert.Equal(t, 35, result.TotalRawIssues)
}

func TestFileThresholdFromEnv_Default(t *testing.T) {
	// No env var → returns IssueCountThreshold (30)
	result := fileThresholdFromEnv()
	assert.Equal(t, IssueCountThreshold, result)
}

func TestFileThresholdFromEnv_ValidValue(t *testing.T) {
	t.Setenv("GOLANGCI_LINT_FILE_THRESHOLD", "50")

	result := fileThresholdFromEnv()
	assert.Equal(t, 50, result)
}

func TestFileThresholdFromEnv_InvalidValue(t *testing.T) {
	t.Setenv("GOLANGCI_LINT_FILE_THRESHOLD", "abc")

	result := fileThresholdFromEnv()
	assert.Equal(t, IssueCountThreshold, result, "invalid value should return default")
}

func TestFileThresholdFromEnv_ZeroValue(t *testing.T) {
	t.Setenv("GOLANGCI_LINT_FILE_THRESHOLD", "0")

	result := fileThresholdFromEnv()
	assert.Equal(t, IssueCountThreshold, result, "zero should return default")
}

func TestFileThresholdFromEnv_NegativeValue(t *testing.T) {
	t.Setenv("GOLANGCI_LINT_FILE_THRESHOLD", "-5")

	result := fileThresholdFromEnv()
	assert.Equal(t, IssueCountThreshold, result, "negative should return default")
}

// --- Plan 111 Task 2: buildStrategyInstructions (formerly V2) ---

func TestBuildStrategyInstructions_FromResult_SingleAgent(t *testing.T) {
	result := StrategyResult{StrategyName: "single-agent"}
	assert.Empty(t, buildStrategyInstructions(result))
}

func TestBuildStrategyInstructions_FromResult_PerFile(t *testing.T) {
	issues := []LintIssue{
		makeIssue("errcheck", "error not checked", "main.go"),
		makeIssue("govet", "argument mismatch", "pkg/auth/handler.go"),
	}
	result := StrategyResult{
		StrategyName:   "subagent-per-file",
		TotalRawIssues: 2,
		UniqueIssues:   issues,
	}

	out := buildStrategyInstructions(result)

	assert.Contains(t, out, "<strategy_instructions>")
	assert.Contains(t, out, "</strategy_instructions>")
	assert.Contains(t, out, "main.go — call")
	assert.Contains(t, out, "pkg/auth/handler.go — call")
	assert.Contains(t, out, `golangci_lint_guide(linter="errcheck")`)
	assert.Contains(t, out, `golangci_lint_guide(linter="govet")`)
}

func TestBuildStrategyInstructions_FromResult_PerPackage_NoEscalation(t *testing.T) {
	issues := []LintIssue{
		makeIssue("errcheck", "error not checked", "pkg/auth/handler.go"),
		makeIssue("govet", "argument mismatch", "pkg/db/conn.go"),
	}
	packages := []PackageEntry{
		{Path: "pkg/auth", Count: 1},
		{Path: "pkg/db", Count: 1},
	}
	result := StrategyResult{
		StrategyName:  "subagent-per-package",
		Packages:      packages,
		UniqueIssues:  issues,
		EscalatedPkgs: map[string]bool{},
		FileThreshold: IssueCountThreshold,
	}

	out := buildStrategyInstructions(result)

	assert.Contains(t, out, "<strategy_instructions>")
	assert.Contains(t, out, "</strategy_instructions>")
	assert.Contains(t, out, "pkg/auth")
	assert.Contains(t, out, "pkg/db")
	assert.Contains(t, out, `golangci_lint_guide(linter="errcheck")`)
	assert.Contains(t, out, `golangci_lint_guide(linter="govet")`)
}

func TestBuildStrategyInstructions_FromResult_PerPackage_WithEscalation(t *testing.T) {
	// 3 packages: pkg/big escalated, pkg/small1 and pkg/small2 not escalated
	issues := []LintIssue{
		makeIssue("errcheck", "SA1001: error", "pkg/big/a.go"),
		makeIssue("errcheck", "SA1001: error", "pkg/big/b.go"),
		makeIssue("govet", "argument mismatch", "pkg/big/c.go"),
		makeIssue("staticcheck", "SA5001: error", "pkg/small1/a.go"),
		makeIssue("gosec", "G101: hardcoded", "pkg/small2/a.go"),
	}
	packages := []PackageEntry{
		{Path: "pkg/big", Count: 3},
		{Path: "pkg/small1", Count: 1},
		{Path: "pkg/small2", Count: 1},
	}
	result := StrategyResult{
		StrategyName:  "subagent-per-package",
		Packages:      packages,
		UniqueIssues:  issues,
		EscalatedPkgs: map[string]bool{"pkg/big": true},
		FileThreshold: IssueCountThreshold,
	}

	out := buildStrategyInstructions(result)

	assert.Contains(t, out, "<strategy_instructions>")
	assert.Contains(t, out, "</strategy_instructions>")

	// Escalated package gets [PER-FILE] label and individual files
	assert.Contains(t, out, "[PER-FILE] pkg/big", "escalated package should have PER-FILE label")
	assert.Contains(t, out, "a.go", "escalated package should show individual files")
	assert.Contains(t, out, "b.go", "escalated package should show individual files")
	assert.Contains(t, out, "c.go", "escalated package should show individual files")
	assert.Contains(t, out, "escalated", "escalated packages should mention escalation")

	// Non-escalated packages get [PER-PACKAGE] label
	assert.Contains(t, out, "[PER-PACKAGE] pkg/small1", "non-escalated package should have PER-PACKAGE label")
	assert.Contains(t, out, "[PER-PACKAGE] pkg/small2", "non-escalated package should have PER-PACKAGE label")

	// Non-escalated packages show package-level guide calls
	assert.Contains(t, out, `golangci_lint_guide(linter="staticcheck", rule="SA5001")`)
	assert.Contains(t, out, `golangci_lint_guide(linter="gosec", rule="G101")`)
}

func TestSortedKeysContainingPath(t *testing.T) {
	fileRefs := map[string][]GuideRef{
		"pkg/a/one.go":   {{Linter: "errcheck"}},
		"pkg/a/two.go":   {{Linter: "govet"}},
		"pkg/b/three.go": {{Linter: "staticcheck"}},
	}

	result := sortedKeysContainingPath(fileRefs, "pkg/a")

	assert.Equal(t, []string{"pkg/a/one.go", "pkg/a/two.go"}, result)
}

func TestSortedKeysContainingPath_Empty(t *testing.T) {
	fileRefs := map[string][]GuideRef{
		"pkg/b/one.go": {{Linter: "errcheck"}},
	}

	result := sortedKeysContainingPath(fileRefs, "pkg/a")

	assert.Empty(t, result)
}
