package server

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

const (
	largeOutputIssueThreshold       = 30
	largeOutputPackageThreshold     = 3
	subagentRecommendationThreshold = 30
)

type packageEntry struct {
	path  string
	count int
}

// extractPackagesFromIssues groups issues by directory of their filename,
// returning entries sorted by count descending, then path ascending for ties.
func extractPackagesFromIssues(issues []lintIssue) []packageEntry {
	counts := make(map[string]int)
	for _, issue := range issues {
		dir := filepath.Dir(issue.Pos.Filename)
		if dir == "." || dir == "" {
			dir = "."
		}
		counts[dir]++
	}

	entries := make([]packageEntry, 0, len(counts))
	for path, count := range counts {
		entries = append(entries, packageEntry{path: path, count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].path < entries[j].path
	})

	return entries
}

// buildPackageBreakdown formats package entries as "path: N issues" lines
// followed by a total line, sorted by count descending.
func buildPackageBreakdown(packages []packageEntry) string {
	if len(packages) == 0 {
		return ""
	}

	lines := make([]string, 0, len(packages)+1)
	total := 0
	for _, pkg := range packages {
		lines = append(lines, fmt.Sprintf("%s: %d issues", pkg.path, pkg.count))
		total += pkg.count
	}
	lines = append(lines, fmt.Sprintf("TOTAL: %d issues across %d packages", total, len(packages)))
	return strings.Join(lines, "\n")
}

func buildStrategyInstructions(strategyName string, packages []packageEntry, totalIssues int) string {
	switch strategyName {
	case "subagent-per-package":
		var pkgList strings.Builder
		for _, pkg := range packages {
			if pkg.count > 0 {
				fmt.Fprintf(&pkgList, "\n   %s (%d issues)", pkg.path, pkg.count)
			}
		}
		promptTmpl := "Use golangci_lint_run to find and fix all lint issues in " +
			pkgList.String() + ". Verify with golangci_lint_run after fixing."
		return "\n\n<strategy_instructions>\n" +
			"IMPORTANT: You MUST use subagents. " +
			"Do NOT attempt to fix all issues yourself in a single pass.\n" +
			"Single-agent mode will exceed context limits at this scale.\n\n" +
			"REQUIRED steps:\n" +
			"1. For EACH package listed below, spawn a separate subagent:\n" +
			"   task(description=\"Fix golangci-lint in {package_path}\", " +
			"prompt=\"" + promptTmpl + "\")\n" +
			"2. After ALL subagents complete, run " +
			"golangci_lint_run(path=\"./...\") to verify zero issues.\n\n" +
			"Packages requiring subagents:" + pkgList.String() + "\n" +
			"</strategy_instructions>"

	case "subagent-per-file":
		header := fmt.Sprintf(
			"IMPORTANT: You MUST use subagents. "+
				"Do NOT attempt to fix all %d issues yourself in a single pass.\n",
			totalIssues)
		return "\n\n<strategy_instructions>\n" +
			header +
			"Single-agent mode will exceed context limits at this scale.\n\n" +
			"REQUIRED steps:\n" +
			"1. Use glob(\"**/*.go\") to identify all Go files.\n" +
			"2. For EACH file with issues, spawn a separate subagent:\n" +
			"   task(description=\"Fix golangci-lint in {filename}\", " +
			"prompt=\"Use golangci_lint_run to find and fix all lint issues. " +
			"Verify with golangci_lint_run after fixing.\")\n" +
			"3. After ALL subagents complete, run " +
			"golangci_lint_run(path=\"./...\") to verify zero issues.\n" +
			"</strategy_instructions>"

	default:
		return ""
	}
}

// recommendStrategy returns a strategy name and reason based on issue and package counts.
func recommendStrategy(totalIssues, totalPackages int) (string, string) {
	if totalPackages > largeOutputPackageThreshold {
		return "subagent-per-package",
			fmt.Sprintf("%d issues across %d packages — use subagent-per-package strategy",
				totalIssues, totalPackages)
	}
	if totalIssues > subagentRecommendationThreshold {
		return "subagent-per-file",
			fmt.Sprintf(">%d issues in %d packages — use subagent-per-file strategy",
				totalIssues, totalPackages)
	}
	return "single-agent",
		fmt.Sprintf("≤%d issues across %d packages — single-agent flow",
			totalIssues, totalPackages)
}

// isLargeOutput returns true when the output exceeds the large-output thresholds.
func isLargeOutput(totalIssues, totalPackages int) bool {
	return totalIssues > largeOutputIssueThreshold || totalPackages > largeOutputPackageThreshold
}
