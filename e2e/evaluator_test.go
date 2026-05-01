//go:build e2e

package e2e

import (
	"context"
	"fmt"
	"strings"
)

func CountLintIssues(ctx context.Context, workDir string) (int, error) {
	_, output, err := containerExec(ctx, []string{
		"sh", "-c",
		fmt.Sprintf("cd %s && golangci-lint run --output.json.path stdout ./...", workDir),
	})
	if err != nil {
		return 0, fmt.Errorf("golangci-lint exec: %w", err)
	}
	return countJSONIssues(output), nil
}

func Evaluate(ctx context.Context, workDir string) (*EvalResult, error) {
	result := &EvalResult{}

	lintCode, lintOutput, err := containerExec(ctx, []string{
		"sh", "-c",
		fmt.Sprintf("cd %s && golangci-lint run --output.json.path stdout ./...", workDir),
	})
	if err != nil {
		return result, fmt.Errorf("lint exec: %w", err)
	}
	result.AfterIssues = countJSONIssues(lintOutput)
	result.LintClean = lintCode == 0 && result.AfterIssues == 0

	buildCode, _, err := containerExec(ctx, []string{
		"sh", "-c",
		fmt.Sprintf("cd %s && go build ./...", workDir),
	})
	if err != nil {
		return result, fmt.Errorf("build exec: %w", err)
	}
	result.BuildsClean = buildCode == 0

	_, nolintOutput, err := containerExec(ctx, []string{
		"sh", "-c",
		fmt.Sprintf("cd %s && grep -rc '//nolint' . --include='*.go' || true", workDir),
	})
	if err == nil {
		result.NolintCount = strings.Count(nolintOutput, "//nolint")
	}

	result.Pass = result.LintClean && result.BuildsClean
	return result, nil
}

func CheckConfigTamper(ctx context.Context, workDir string, fixtureDir string) (bool, string) {
	_, original, _ := containerExec(ctx, []string{
		"cat", "/workspace/testdata/" + fixtureDir + "/.golangci.yml",
	})
	_, current, _ := containerExec(ctx, []string{
		"cat", workDir + "/.golangci.yml",
	})
	if original == current {
		return false, ""
	}
	diff := computeDiff(original, current)
	return true, diff
}

func computeDiff(original, current string) string {
	origLines := strings.Split(original, "\n")
	curLines := strings.Split(current, "\n")
	var sb strings.Builder
	maxLen := max(len(curLines), len(origLines))
	for i := range maxLen {
		var oLine, cLine string
		if i < len(origLines) {
			oLine = origLines[i]
		}
		if i < len(curLines) {
			cLine = curLines[i]
		}
		if oLine != cLine {
			fmt.Fprintf(&sb, "-%d: %s\n+%d: %s\n", i+1, oLine, i+1, cLine)
		}
	}
	return sb.String()
}
