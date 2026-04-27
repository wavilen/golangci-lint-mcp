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

// recommendStrategy returns a strategy name and reason based on issue and package counts.
func recommendStrategy(totalIssues, totalPackages int) (string, string) {
	if totalIssues > subagentRecommendationThreshold || totalPackages > largeOutputPackageThreshold {
		return "subagent-per-package",
			fmt.Sprintf(">%d issues across %d packages — use subagent-per-package strategy",
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
