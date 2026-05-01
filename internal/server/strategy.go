package server

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/linttypes"
)

const (
	largeOutputPackageThreshold     = 3
	subagentRecommendationThreshold = 30
	strategySubagentPerPackage      = "subagent-per-package"
	strategySubagentPerFile         = "subagent-per-file"
	strategySingleAgent             = "single-agent"
)

type PackageEntry struct {
	Path  string
	Count int
}

// GuideRef represents a unique (linter, rule) pair for guide call generation.
type GuideRef struct {
	Linter string
	Rule   string
}

// StrategyResult holds the complete analysis of a lint run, produced by AnalyzeStrategy.
type StrategyResult struct {
	RawIssues      []LintIssue     // original issues before deduplication
	UniqueIssues   []LintIssue     // deduplicated issues
	Packages       []PackageEntry  // packages sorted by count desc
	TotalRawIssues int             // len(RawIssues) — used for strategy decisions per D-01
	StrategyName   string          // "single-agent", "subagent-per-package", or "subagent-per-file"
	StrategyReason string          // human-readable reason for the strategy choice
	SummaryOnly    bool            // true when unique count > IssueCountThreshold
	EscalatedPkgs  map[string]bool // package paths escalated to per-file (D-05)
	FileThreshold  int             // resolved per-package escalation threshold
	RawCounts      map[string]int  // raw issue counts per package path (for escalation display)
}

// extractGuideRefsByDir groups deduplicated GuideRefs by directory.
func extractGuideRefsByDir(issues []LintIssue) map[string][]GuideRef {
	result := make(map[string][]GuideRef)
	seen := make(map[string]map[GuideRef]bool) // dir -> set of seen refs

	for _, issue := range issues {
		dir := filepath.Dir(issue.Pos.Filename)
		if dir == "." || dir == "" {
			dir = "."
		}
		ref := GuideRef{
			Linter: issue.FromLinter,
			Rule:   linttypes.ExtractRule(issue.Text),
		}
		if seen[dir] == nil {
			seen[dir] = make(map[GuideRef]bool)
		}
		if !seen[dir][ref] {
			seen[dir][ref] = true
			result[dir] = append(result[dir], ref)
		}
	}

	// Sort each dir's refs by linter then rule
	for dir := range result {
		sort.Slice(result[dir], func(i, j int) bool {
			if result[dir][i].Linter != result[dir][j].Linter {
				return result[dir][i].Linter < result[dir][j].Linter
			}
			return result[dir][i].Rule < result[dir][j].Rule
		})
	}

	return result
}

// extractGuideRefsByFile groups deduplicated GuideRefs by filename.
func extractGuideRefsByFile(issues []LintIssue) map[string][]GuideRef {
	result := make(map[string][]GuideRef)
	seen := make(map[string]map[GuideRef]bool) // file -> set of seen refs

	for _, issue := range issues {
		filename := issue.Pos.Filename
		ref := GuideRef{
			Linter: issue.FromLinter,
			Rule:   linttypes.ExtractRule(issue.Text),
		}
		if seen[filename] == nil {
			seen[filename] = make(map[GuideRef]bool)
		}
		if !seen[filename][ref] {
			seen[filename][ref] = true
			result[filename] = append(result[filename], ref)
		}
	}

	// Sort each file's refs by linter then rule
	for filename := range result {
		sort.Slice(result[filename], func(i, j int) bool {
			if result[filename][i].Linter != result[filename][j].Linter {
				return result[filename][i].Linter < result[filename][j].Linter
			}
			return result[filename][i].Rule < result[filename][j].Rule
		})
	}

	return result
}

// formatBatchCall formats a slice of GuideRefs as a single batch golangci_lint_guide tool call.
// Returns empty string if refs is empty. Empty rule fields are omitted (D-03).
// If len(refs) exceeds defaultBatchMax, output is truncated with a remainder note (D-05).
func formatBatchCall(refs []GuideRef) string {
	if len(refs) == 0 {
		return ""
	}

	visible := refs
	remainder := 0
	if len(refs) > defaultBatchMax {
		visible = refs[:defaultBatchMax]
		remainder = len(refs) - defaultBatchMax
	}

	items := formatBatchQueryItems(visible)
	result := fmt.Sprintf("golangci_lint_guide(queries=[%s])", items)

	if remainder > 0 {
		result += fmt.Sprintf(
			"\n     Note: %d additional (linter, rule) pairs not shown. Use golangci_lint_parse for full diagnostics.",
			remainder,
		)
	}

	return result
}

// formatBatchQueryItems builds comma-separated Go-style struct strings for batch queries.
func formatBatchQueryItems(refs []GuideRef) string {
	parts := make([]string, 0, len(refs))
	for _, ref := range refs {
		if ref.Rule != "" {
			parts = append(parts, fmt.Sprintf(`{linter:%q, rule:%q}`, ref.Linter, ref.Rule))
		} else {
			parts = append(parts, fmt.Sprintf(`{linter:%q}`, ref.Linter))
		}
	}
	return strings.Join(parts, ", ")
}

// isSubagentStrategy returns true for "subagent-per-package" or "subagent-per-file".
func isSubagentStrategy(strategyName string) bool {
	return strategyName == strategySubagentPerPackage || strategyName == strategySubagentPerFile
}

// ExtractPackagesFromIssues groups issues by directory of their filename,
// returning entries sorted by count descending, then path ascending for ties.
func ExtractPackagesFromIssues(issues []LintIssue) []PackageEntry {
	counts := make(map[string]int)
	for _, issue := range issues {
		dir := filepath.Dir(issue.Pos.Filename)
		if dir == "." || dir == "" {
			dir = "."
		}
		counts[dir]++
	}

	entries := make([]PackageEntry, 0, len(counts))
	for path, count := range counts {
		entries = append(entries, PackageEntry{Path: path, Count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Path < entries[j].Path
	})

	return entries
}

// buildPackageBreakdown formats package entries as "path: N issues" lines
// followed by a total line, sorted by count descending.
func buildPackageBreakdown(packages []PackageEntry) string {
	if len(packages) == 0 {
		return ""
	}

	lines := make([]string, 0, len(packages)+1)
	total := 0
	for _, pkg := range packages {
		lines = append(lines, fmt.Sprintf("%s: %d issues", pkg.Path, pkg.Count))
		total += pkg.Count
	}
	lines = append(lines, fmt.Sprintf("TOTAL: %d issues across %d packages", total, len(packages)))
	return strings.Join(lines, "\n")
}

// RecommendStrategy returns a strategy name and reason based on issue and package counts.
func RecommendStrategy(totalIssues, totalPackages int) (string, string) {
	if totalPackages > largeOutputPackageThreshold {
		return strategySubagentPerPackage,
			fmt.Sprintf("%d issues across %d packages — use subagent-per-package strategy",
				totalIssues, totalPackages)
	}
	if totalIssues > subagentRecommendationThreshold {
		return strategySubagentPerFile,
			fmt.Sprintf(">%d issues in %d packages — use subagent-per-file strategy",
				totalIssues, totalPackages)
	}
	return strategySingleAgent,
		fmt.Sprintf("≤%d issues across %d packages — single-agent flow",
			totalIssues, totalPackages)
}

// fileThresholdFromEnv returns the per-package escalation threshold from
// GOLANGCI_LINT_FILE_THRESHOLD env var, defaulting to IssueCountThreshold (30).
// Invalid values (non-numeric, zero, negative) are silently ignored.
func fileThresholdFromEnv() int {
	val := os.Getenv("GOLANGCI_LINT_FILE_THRESHOLD")
	if val != "" {
		n, err := strconv.Atoi(val)
		if err == nil && n > 0 {
			return n
		}
	}
	return IssueCountThreshold
}

// AnalyzeStrategy performs the full analysis pipeline:
//
//	deduplicate → extract packages → recommend strategy → determine routing → per-package escalation.
//
// Per D-01: uses total raw issue count (len(rawIssues)) for RecommendStrategy, not unique count.
func AnalyzeStrategy(rawIssues []LintIssue) StrategyResult {
	unique := linttypes.DeduplicateIssues(rawIssues)
	packages := ExtractPackagesFromIssues(unique)
	totalRaw := len(rawIssues)

	// D-01: Always pass total raw count to RecommendStrategy
	strategyName, strategyReason := RecommendStrategy(totalRaw, len(packages))

	fileThreshold := fileThresholdFromEnv()
	summaryOnly := len(unique) > IssueCountThreshold

	// D-05: Per-package escalation for subagent-per-package strategy
	escalatedPkgs := make(map[string]bool)
	rawCounts := make(map[string]int)
	if strategyName == strategySubagentPerPackage {
		for _, issue := range rawIssues {
			dir := filepath.Dir(issue.Pos.Filename)
			if dir == "." || dir == "" {
				dir = "."
			}
			rawCounts[dir]++
		}
		if fileThreshold > 0 {
			for _, pkg := range packages {
				if rawCounts[pkg.Path] > fileThreshold {
					escalatedPkgs[pkg.Path] = true
				}
			}
		}
	}

	return StrategyResult{
		RawIssues:      rawIssues,
		UniqueIssues:   unique,
		Packages:       packages,
		TotalRawIssues: totalRaw,
		StrategyName:   strategyName,
		StrategyReason: strategyReason,
		SummaryOnly:    summaryOnly,
		EscalatedPkgs:  escalatedPkgs,
		FileThreshold:  fileThreshold,
		RawCounts:      rawCounts,
	}
}

// buildStrategyInstructions produces strategy instructions from a StrategyResult,
// supporting per-package escalation to per-file for high-issue packages (D-05, D-07).
func buildStrategyInstructions(result StrategyResult) string {
	switch result.StrategyName {
	case strategySubagentPerPackage:
		return buildSubagentPerPackageInstructions(result)
	case strategySubagentPerFile:
		return buildSubagentPerFileInstructions(result)
	default:
		return ""
	}
}

// formatTaskBlock produces a concrete task(description=..., prompt=...) block.
// Uses %q for description (simple text) and "%s" for prompt (preserves guide call formatting).
// Inner double quotes in prompt are escaped so the prompt="..." wrapper remains parseable.
func formatTaskBlock(description, prompt string) string {
	escaped := strings.ReplaceAll(prompt, `"`, `\"`)
	return fmt.Sprintf(`task(description=%q, prompt="%s")`, description, escaped)
}

// buildSubagentPerPackageInstructions generates concrete task() blocks for the
// subagent-per-package strategy, including per-package escalation to per-file
// for high-issue packages (D-05, D-06, D-07).
func buildSubagentPerPackageInstructions(result StrategyResult) string {
	var taskList strings.Builder
	dirRefs := extractGuideRefsByDir(result.UniqueIssues)
	fileRefs := extractGuideRefsByFile(result.UniqueIssues)

	for _, pkg := range result.Packages {
		if pkg.Count == 0 {
			continue
		}

		if result.EscalatedPkgs[pkg.Path] {
			buildEscalatedPackageEntry(result, pkg, fileRefs, &taskList)
		} else {
			buildStandardPackageEntry(result, pkg, dirRefs, &taskList)
		}
	}

	return "\n\n<strategy_instructions>\n" +
		"IMPORTANT: You MUST use subagents. " +
		"Do NOT attempt to fix all issues yourself in a single pass.\n" +
		"Each task() block below is ready to use — " +
		"spawn each as a separate subagent.\n\n" +
		taskList.String() + "\n\n" +
		"After ALL subagents complete, run " +
		"golangci_lint_run(path=\"./...\") to verify zero issues.\n" +
		"</strategy_instructions>"
}

// buildEscalatedPackageEntry produces one task() block per file in an escalated package (D-06).
// Escalated packages exceed the file threshold and get per-file task blocks.
func buildEscalatedPackageEntry(
	result StrategyResult,
	pkg PackageEntry,
	fileRefs map[string][]GuideRef,
	pkgList *strings.Builder,
) {
	filenames := sortedKeysContainingPath(fileRefs, pkg.Path)
	for _, filename := range filenames {
		refs := fileRefs[filename]
		desc := fmt.Sprintf("Fix golangci-lint in %s", filename)
		prompt := fmt.Sprintf("Fix all golangci-lint issues in %s. "+
			"REQUIRED: call %s to understand how to fix each issue type. "+
			"Verify with golangci_lint_run after fixing.",
			filename, formatBatchCall(refs))
		fmt.Fprintf(pkgList, "\n%s", formatTaskBlock(desc, prompt))
	}
}

// buildStandardPackageEntry produces one task() block per package (D-07).
func buildStandardPackageEntry(
	_ StrategyResult,
	pkg PackageEntry,
	dirRefs map[string][]GuideRef,
	pkgList *strings.Builder,
) {
	desc := fmt.Sprintf("Fix golangci-lint in %s", pkg.Path)
	promptCore := ""
	if refs, ok := dirRefs[pkg.Path]; ok && len(refs) > 0 {
		promptCore = fmt.Sprintf(
			" REQUIRED: call %s to understand how to fix each issue type.",
			formatBatchCall(refs))
	}
	prompt := fmt.Sprintf("Fix all golangci-lint issues in %s files.%s "+
		"Verify with golangci_lint_run after fixing.",
		pkg.Path, promptCore)
	fmt.Fprintf(pkgList, "\n%s", formatTaskBlock(desc, prompt))
}

// buildSubagentPerFileInstructions generates concrete task() blocks for the
// subagent-per-file strategy, one per file.
func buildSubagentPerFileInstructions(result StrategyResult) string {
	header := fmt.Sprintf(
		"IMPORTANT: You MUST use subagents. "+
			"Do NOT attempt to fix all %d issues yourself in a single pass.\n",
		result.TotalRawIssues)

	fileRefs := extractGuideRefsByFile(result.UniqueIssues)
	var taskList strings.Builder
	if len(fileRefs) > 0 {
		filenames := make([]string, 0, len(fileRefs))
		for filename := range fileRefs {
			filenames = append(filenames, filename)
		}
		sort.Strings(filenames)

		for _, filename := range filenames {
			refs := fileRefs[filename]
			desc := fmt.Sprintf("Fix golangci-lint in %s", filename)
			prompt := fmt.Sprintf("Fix all golangci-lint issues in %s. "+
				"REQUIRED: call %s to understand how to fix each issue type. "+
				"Verify with golangci_lint_run after fixing.",
				filename, formatBatchCall(refs))
			fmt.Fprintf(&taskList, "\n%s", formatTaskBlock(desc, prompt))
		}
	}

	return "\n\n<strategy_instructions>\n" +
		header +
		"Each task() block below is ready to use — " +
		"spawn each as a separate subagent.\n\n" +
		taskList.String() + "\n\n" +
		"After ALL subagents complete, run " +
		"golangci_lint_run(path=\"./...\") to verify zero issues.\n" +
		"</strategy_instructions>"
}

// sortedKeysContainingPath returns sorted filenames from fileRefs where
// the directory of the filename matches pkgPath.
func sortedKeysContainingPath(fileRefs map[string][]GuideRef, pkgPath string) []string {
	var filenames []string
	for filename := range fileRefs {
		if filepath.Dir(filename) == pkgPath {
			filenames = append(filenames, filename)
		}
	}
	sort.Strings(filenames)
	return filenames
}
