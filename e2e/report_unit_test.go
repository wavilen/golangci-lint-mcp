package e2e_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateReport(t *testing.T) {
	tmpDir := t.TempDir()

	t.Run("generates HTML report from passing results", func(t *testing.T) {
		outputPath := filepath.Join(tmpDir, "report.html")
		results := []*EvalResult{
			{TestCase: "fixture-a", Model: "model-x", BeforeIssues: 10, AfterIssues: 0, BuildsClean: true, LintClean: true, IssueReduction: 100.0, ToolCalls: 5, Pass: true},
			{TestCase: "fixture-b", Model: "model-y", BeforeIssues: 8, AfterIssues: 2, BuildsClean: true, LintClean: false, IssueReduction: 75.0, ToolCalls: 3, Pass: false},
		}
		html := generateAndRead(t, results, outputPath)
		assertContains(t, html, "Integration Test Report", "title")
		assertContains(t, html, "fixture-a", "fixture-a")
		assertContains(t, html, "model-x", "model-x")
		assertContains(t, html, "model-y", "model-y")
		assertContains(t, html, "PASS", "PASS status")
	})

	t.Run("handles empty results", func(t *testing.T) {
		outputPath := filepath.Join(tmpDir, "empty-report.html")
		html := generateAndRead(t, []*EvalResult{}, outputPath)
		assertContains(t, html, "Integration Test Report", "title")
	})

	t.Run("includes errors in report", func(t *testing.T) {
		outputPath := filepath.Join(tmpDir, "error-report.html")
		results := []*EvalResult{{TestCase: "fixture-err", Model: "model-z", Pass: false, Error: "container crashed"}}
		html := generateAndRead(t, results, outputPath)
		assertContains(t, html, "container crashed", "error message")
		assertContains(t, html, "ERROR", "ERROR status")
	})

	t.Run("computes pass rate from multiple results", func(t *testing.T) {
		outputPath := filepath.Join(tmpDir, "rate-report.html")
		results := []*EvalResult{
			{TestCase: "a", Model: "m1", Pass: true, IssueReduction: 90.0},
			{TestCase: "b", Model: "m1", Pass: true, IssueReduction: 80.0},
			{TestCase: "c", Model: "m1", Pass: false, IssueReduction: 50.0},
		}
		html := generateAndRead(t, results, outputPath)
		assertContains(t, html, "67%", "pass rate")
		assertContains(t, html, "2/3 tests", "test count")
	})
}

func generateAndRead(t *testing.T, results []*EvalResult, outputPath string) string {
	t.Helper()
	err := GenerateReport(results, outputPath)
	if err != nil {
		t.Fatalf("GenerateReport failed: %v", err)
	}
	content, readErr := os.ReadFile(outputPath)
	if readErr != nil {
		t.Fatalf("reading report: %v", readErr)
	}
	return string(content)
}

func assertContains(t *testing.T, haystack, needle, label string) {
	t.Helper()
	if !strings.Contains(haystack, needle) {
		t.Fatalf("report missing %s: %s", label, needle)
	}
}
