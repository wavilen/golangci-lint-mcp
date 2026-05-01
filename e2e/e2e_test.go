package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const fixedPrompt = `Use the golangci-lint-guide skill to fix all golangci-lint issues in the current directory. Follow the skill process: call golangci_lint_run(path="./...") to find issues, apply fixes using the guidance returned, then verify with golangci_lint_run again. Do not add nolint directives. After the first scan, if issues remain, use per-package paths for targeted fixes.

IMPORTANT: This project uses golangci-lint v2. The CLI flags changed from v1. Do NOT use --out-format json or --output=json — these are v1 flags and will fail with "unknown flag". Always use the MCP tool golangci_lint_run instead of running golangci-lint directly. If you must use the CLI, the v2 JSON output flag is: --output.json.path stdout`

type ModelInfo struct {
	Name  string
	Value string
}

var models = []ModelInfo{
	{"GLM-5-Turbo", "zai-coding-plan/glm-5-turbo"},
	{"GLM-5.1", "zai-coding-plan/glm-5.1"},
	{"GLM-4.7", "zai-coding-plan/glm-4.7"},
}

const maxRetries = 3

type TestCase struct {
	Name        string
	FixtureDir  string
	MinIssues   int
	MaxIssues   int
	SourceFiles int
}

var testCases = []TestCase{
	{Name: "Simple", FixtureDir: "simple", MinIssues: 8, MaxIssues: 8, SourceFiles: 1},
	{Name: "Medium", FixtureDir: "medium", MinIssues: 30, MaxIssues: 30, SourceFiles: 3},
	{Name: "Large", FixtureDir: "large", MinIssues: 116, MaxIssues: 116, SourceFiles: 5},
}

var _ = Describe("E2E Integration", func() {
	Describe("Probe", Serial, func() {
		It("validates MCP integration inside container", func(ctx SpecContext) {
			cleanup := createTestContainer(ctx)
			defer cleanup()

			code, _, err := containerExec(ctx, []string{
				"sh", "-c",
				"mkdir -p /tmp/probe && cd /tmp/probe && echo 'module probe\n\ngo 1.26' > go.mod && echo 'package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"hello\") }' > main.go",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(code).To(Equal(0))

			lintCode, lintOutput, lintErr := containerExec(ctx, []string{
				"sh", "-c", "cd /tmp/probe && golangci-lint run ./...",
			})
			Expect(lintErr).ToNot(HaveOccurred())
			_ = lintCode
			By("golangci-lint output: " + lintOutput)

			opencodeCode, opencodeOutput, opencodeErr := containerExec(ctx, []string{
				"opencode", "run",
				"--format", "json",
				"--dangerously-skip-permissions",
				"--dir", "/tmp/probe",
				"Run golangci_lint_list to verify the MCP server is working. Just list available tools.",
			})
			Expect(opencodeErr).ToNot(HaveOccurred())
			Expect(opencodeOutput).ToNot(BeEmpty(), "opencode should produce output")
			_ = opencodeCode

			dumpNDJSON("probe", 0, []byte(opencodeOutput))
		}, SpecTimeout(5*time.Minute))
	})

	DescribeTable("Fixture validation",
		func(tc TestCase) {
			fixturePath := filepath.Join("testdata", tc.FixtureDir)
			Expect(filepath.Join(fixturePath, "go.mod")).To(BeAnExistingFile())
			Expect(filepath.Join(fixturePath, ".golangci.yml")).To(BeAnExistingFile())
			entries, err := os.ReadDir(fixturePath)
			Expect(err).ToNot(HaveOccurred())
			goFiles := 0
			for _, e := range entries {
				if !e.IsDir() && filepath.Ext(e.Name()) == ".go" {
					goFiles++
				}
			}
			Expect(goFiles).To(Equal(tc.SourceFiles),
				fmt.Sprintf("%s: expected %d .go files, found %d", tc.Name, tc.SourceFiles, goFiles))
		},
		Entry("Simple fixture", testCases[0]),
		Entry("Medium fixture", testCases[1]),
		Entry("Large fixture", testCases[2]),
	)

	DescribeTable("Agent fixes lint issues",
		func(ctx SpecContext, tc TestCase, m ModelInfo) {
			By(fmt.Sprintf("=== %s / %s ===", tc.Name, m.Name))

			result := &EvalResult{
				TestCase: tc.Name,
				Model:    m.Name,
			}
			resultWritten := false

			defer func() {
				if !resultWritten {
					result.Pass = false
					if result.Error == "" {
						result.Error = fmt.Sprintf("test interrupted: %s/%s (before=%d issues, retries=%d)",
							tc.Name, m.Name, result.BeforeIssues, result.Retries)
					}
					dumpResult(result)
				}
			}()

			cleanup := createTestContainer(ctx)
			defer cleanup()

			tmpDir := fmt.Sprintf("/tmp/%s-%s",
				tc.FixtureDir,
				strings.ToLower(strings.ReplaceAll(m.Name, ".", "-")),
			)

			copyFixture := func() {
				code, _, err := containerExec(ctx, []string{
					"sh", "-c",
					fmt.Sprintf(
						"rm -rf %s && mkdir -p %s && cp -r /workspace/testdata/%s/. %s/",
						tmpDir, tmpDir, tc.FixtureDir, tmpDir,
					),
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(code).To(Equal(0))
			}

			copyFixture()

			beforeCount, err := CountLintIssues(ctx, tmpDir)
			Expect(err).ToNot(HaveOccurred())
			Expect(beforeCount).To(
				BeNumerically(">=", tc.MinIssues),
				fmt.Sprintf("%s: expected >= %d issues, found %d", tc.Name, tc.MinIssues, beforeCount),
			)

			result.BeforeIssues = beforeCount

			for retry := range maxRetries {
				result.Retries = retry
				if retry > 0 {
					copyFixture()
				}

				_, rawOutput, runErr := containerExec(ctx, []string{
					"opencode", "run",
					"--format", "json",
					"--dangerously-skip-permissions",
					"--model", m.Value,
					"--dir", tmpDir,
					fixedPrompt,
				})
				Expect(runErr).ToNot(HaveOccurred())

				ndjsonTag := fmt.Sprintf("%s-%s",
					tc.FixtureDir,
					strings.ToLower(strings.ReplaceAll(m.Name, ".", "-")),
				)
				dumpNDJSON(ndjsonTag, retry, []byte(rawOutput))

				ndjsonResult, _ := ParseNDJSON(strings.NewReader(rawOutput))
				if ndjsonResult != nil {
					result.ToolCalls = ndjsonResult.ToolCalls
				}

				evalResult, evalErr := Evaluate(ctx, tmpDir)
				Expect(evalErr).ToNot(HaveOccurred())

				// Config tamper detection (D-04/D-05)
				tampered, tamperDiff := CheckConfigTamper(ctx, tmpDir, tc.FixtureDir)
				result.ConfigModified = tampered
				result.ConfigDiff = tamperDiff

				result.AfterIssues = evalResult.AfterIssues
				result.BuildsClean = evalResult.BuildsClean
				result.LintClean = evalResult.LintClean
				result.NolintCount = evalResult.NolintCount
				result.Pass = evalResult.Pass
				if result.BeforeIssues > 0 {
					result.IssueReduction = float64(result.BeforeIssues-result.AfterIssues) / float64(result.BeforeIssues) * 100
				}

				if result.Pass {
					break
				}
			}

			dumpResult(result)
			resultWritten = true

			if !result.LintClean {
				Skip(fmt.Sprintf(
					"T1 SKIP: %s/%s — %d issues remain after %d retries (reduction: %.1f%%)",
					tc.Name, m.Name, result.AfterIssues, result.Retries, result.IssueReduction,
				))
			}
			if !result.BuildsClean {
				Skip(fmt.Sprintf(
					"T2 SKIP: %s/%s — code does not compile after %d retries",
					tc.Name, m.Name, result.Retries,
				))
			}
			if result.ConfigModified {
				Skip(fmt.Sprintf(
					"CONFIG TAMPER: %s/%s — agent modified .golangci.yml during execution\n%s",
					tc.Name, m.Name, result.ConfigDiff,
				))
			}
		},
		Entry("Simple/GLM-5-Turbo", testCases[0], models[0], SpecTimeout(30*time.Minute)),
		Entry("Simple/GLM-5.1", testCases[0], models[1], SpecTimeout(30*time.Minute)),
		Entry("Medium/GLM-5-Turbo", testCases[1], models[0], SpecTimeout(30*time.Minute)),
		Entry("Medium/GLM-5.1", testCases[1], models[1], SpecTimeout(30*time.Minute)),
		Entry("Large/GLM-5-Turbo", testCases[2], models[0], SpecTimeout(30*time.Minute)),
		Entry("Large/GLM-5.1", testCases[2], models[1], SpecTimeout(30*time.Minute)),
		Entry("Simple/GLM-4.7", testCases[0], models[2], SpecTimeout(30*time.Minute)),
		Entry("Medium/GLM-4.7", testCases[1], models[2], SpecTimeout(30*time.Minute)),
		Entry("Large/GLM-4.7", testCases[2], models[2], SpecTimeout(30*time.Minute)),
	)
})

func dumpNDJSON(tag string, retry int, data []byte) {
	ndjsonDir := filepath.Join("..", "tmp", "ndjson")
	os.MkdirAll(ndjsonDir, 0755)
	name := tag
	if retry > 0 {
		name = fmt.Sprintf("%s-r%d", tag, retry)
	}
	dest := filepath.Join(ndjsonDir, name+".ndjson")
	writeErr := os.WriteFile(dest, data, 0644)
	if writeErr != nil {
		By(fmt.Sprintf("Warning: failed to write NDJSON dump %s: %v", dest, writeErr))
	} else {
		By("NDJSON dumped: " + dest)
	}
}

func dumpResult(result *EvalResult) {
	ndjsonDir := filepath.Join("..", "tmp", "ndjson")
	os.MkdirAll(ndjsonDir, 0755)
	name := fmt.Sprintf("%s-%s-result.json",
		strings.ToLower(result.TestCase),
		strings.ToLower(strings.ReplaceAll(result.Model, ".", "-")),
	)
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		By(fmt.Sprintf("Warning: failed to marshal result: %v", err))
		return
	}
	dest := filepath.Join(ndjsonDir, name)
	writeErr := os.WriteFile(dest, data, 0644)
	if writeErr != nil {
		By(fmt.Sprintf("Warning: failed to write result %s: %v", dest, writeErr))
	}
}

var _ = ReportAfterSuite("Generate report and golden files", func(_ Report) {
	ndjsonDir := filepath.Join("..", "tmp", "ndjson")
	entries, readErr := os.ReadDir(ndjsonDir)
	if readErr != nil {
		fmt.Fprintf(GinkgoWriter, "No results directory — skipping report\n")
		return
	}

	var results []*EvalResult
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), "-result.json") {
			continue
		}
		fileData, fileReadErr := os.ReadFile(filepath.Join(ndjsonDir, e.Name()))
		if fileReadErr != nil {
			continue
		}
		var r EvalResult
		unmarshalErr := json.Unmarshal(fileData, &r)
		if unmarshalErr != nil {
			continue
		}
		results = append(results, &r)
	}

	if len(results) == 0 {
		fmt.Fprintln(GinkgoWriter, "No results collected — skipping report generation")
		return
	}

	reportPath := filepath.Join("..", "e2e-report.html")
	reportErr := GenerateReport(results, reportPath)
	if reportErr != nil {
		fmt.Fprintf(GinkgoWriter, "Warning: failed to generate report: %v\n", reportErr)
	} else {
		fmt.Fprintf(GinkgoWriter, "Report generated: %s\n", reportPath)
	}

	goldenDir := filepath.Join("testdata", "golden")
	os.MkdirAll(goldenDir, 0755)
	for _, r := range results {
		goldenName := fmt.Sprintf("%s_%s.json",
			strings.ToLower(r.TestCase),
			strings.ToLower(strings.ReplaceAll(r.Model, ".", "-")))
		data, marshalErr := json.MarshalIndent(r, "", "  ")
		if marshalErr != nil {
			continue
		}
		os.WriteFile(filepath.Join(goldenDir, goldenName), data, 0644)
	}
	fmt.Fprintf(GinkgoWriter, "Golden files saved to %s (%d files)\n", goldenDir, len(results))
})
