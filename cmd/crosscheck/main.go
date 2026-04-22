package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type extraction struct {
	guidePath    string
	relativePath string
	code         string
}

type violation struct {
	Guide         string `json:"guide"`
	ExtractedFile string `json:"extracted_file"`
	Linter        string `json:"linter"`
	Line          int    `json:"line"`
	Message       string `json:"message"`
}

type report struct {
	TotalGuides     int         `json:"total_guides"`
	Passing         int         `json:"passing"`
	Failing         int         `json:"failing"`
	TotalViolations int         `json:"total_violations"`
	Violations      []violation `json:"violations"`
}

type golangciLintIssue struct {
	Pos struct {
		Filename string `json:"Filename"`
		Line     int    `json:"Line"`
		Column   int    `json:"Column"`
	} `json:"Pos"`
	Text       string `json:"Text"`
	FromLinter string `json:"FromLinter"`
}

type golangciLintOutput struct {
	Issues []golangciLintIssue `json:"Issues"`
}

var declarationKeywords = []string{
	"func ", "type ", "var ", "const ", "import ", "var(", "const(", "import(",
}

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get working directory: %v", err)
	}

	extractions, skipped := extractGoodBlocks(filepath.Join(projectRoot, "guides"))
	relSkipped := make([]string, len(skipped))
	for idx, skipPath := range skipped {
		rel, relErr := filepath.Rel(projectRoot, skipPath)
		if relErr != nil {
			rel = skipPath
		}
		relSkipped[idx] = rel
	}
	log.Printf("Extracted %d/%d guides (%d skipped: %s)",
		len(extractions), len(extractions)+len(skipped), len(skipped), strings.Join(relSkipped, ", "))

	if len(extractions) == 0 {
		log.Fatal("no Good code blocks found — check guides/ directory")
	}

	tmpDir := filepath.Join(projectRoot, "tmp", "crosscheck")
	writeErr := writeExtractions(tmpDir, extractions)
	if writeErr != nil {
		log.Fatalf("writing extractions: %v", writeErr)
	}

	modErr := writeGoMod(tmpDir)
	if modErr != nil {
		log.Fatalf("writing go.mod: %v", modErr)
	}

	runGoimports(tmpDir, extractions)

	issues := runGolangciLint(projectRoot, tmpDir, extractions)

	rpt := buildReport(extractions, issues)

	reportPath := filepath.Join(tmpDir, "violations.json")
	reportErr := writeJSONReport(reportPath, rpt)
	if reportErr != nil {
		log.Fatalf("writing report: %v", reportErr)
	}

	printSummary(rpt)
}

func extractGoodBlocks(guidesDir string) ([]extraction, []string) {
	examplesRe := regexp.MustCompile(`(?s)<examples>(.*?)</examples>`)
	goodRe := regexp.MustCompile(`(?s)## Good\s*\n` + "`" + "`" + "`go\n(.*?)" + "`" + "`" + "`")

	var extractions []extraction
	var skipped []string

	walkErr := filepath.WalkDir(guidesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("warning: cannot access %s: %v", path, err)
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		data, readErr := os.ReadFile(path)
		if readErr != nil {
			log.Printf("warning: cannot read %s: %v", path, readErr)
			return nil
		}
		content := string(data)

		examplesMatch := examplesRe.FindStringSubmatch(content)
		if examplesMatch == nil {
			skipped = append(skipped, path)
			return nil
		}
		examplesContent := examplesMatch[1]

		goodMatch := goodRe.FindStringSubmatch(examplesContent)
		if goodMatch == nil {
			skipped = append(skipped, path)
			return nil
		}
		code := goodMatch[1]

		rel, relErr := filepath.Rel(guidesDir, path)
		if relErr != nil {
			log.Printf("warning: cannot compute relative path for %s: %v", path, relErr)
			skipped = append(skipped, path)
			return nil
		}
		rel = strings.TrimSuffix(rel, ".md")
		rel = filepath.ToSlash(rel)

		extractions = append(extractions, extraction{
			guidePath:    path,
			relativePath: rel,
			code:         code,
		})
		return nil
	})
	if walkErr != nil {
		log.Printf("warning: error walking guides directory: %v", walkErr)
	}

	return extractions, skipped
}

func stripPackageDecl(code string) string {
	lines := strings.Split(code, "\n")
	var result []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(strings.TrimSpace(line), "//") {
			result = append(result, line)
			continue
		}
		if strings.HasPrefix(strings.TrimSpace(line), "package ") {
			continue
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

func needsFuncWrap(code string) bool {
	for _, line := range strings.Split(code, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") {
			continue
		}
		for _, keyword := range declarationKeywords {
			if strings.HasPrefix(line, keyword) {
				return false
			}
		}
		return true
	}
	return false
}

func wrapCodeForExtraction(code string) string {
	if !needsFuncWrap(code) {
		return code
	}

	lines := strings.Split(code, "\n")
	wrapEnd := len(lines)
	for idx, line := range lines {
		stripped := strings.TrimSpace(line)
		if stripped == "" || strings.HasPrefix(stripped, "//") || strings.HasPrefix(stripped, "/*") {
			continue
		}
		isDecl := false
		for _, keyword := range declarationKeywords {
			if strings.HasPrefix(stripped, keyword) {
				isDecl = true
				break
			}
		}
		if isDecl && idx > 0 {
			wrapEnd = idx
			for wrapEnd > 0 && strings.TrimSpace(lines[wrapEnd-1]) == "" {
				wrapEnd--
			}
			break
		}
	}

	prefix := strings.Join(lines[:wrapEnd], "\n")
	suffix := strings.Join(lines[wrapEnd:], "\n")

	wrapped := "func _() {\n" + prefix + "\n}"
	if suffix != "" {
		wrapped += "\n" + suffix
	}
	return wrapped
}

func writeExtractions(tmpDir string, extractions []extraction) error {
	_, statErr := os.Stat(tmpDir)
	if statErr == nil {
		removeErr := os.RemoveAll(tmpDir)
		if removeErr != nil {
			return fmt.Errorf("clearing tmp/crosscheck: %w", removeErr)
		}
	}

	for _, ext := range extractions {
		outDir := filepath.Join(tmpDir, ext.relativePath)
		mkdirErr := os.MkdirAll(outDir, 0o750)
		if mkdirErr != nil {
			return fmt.Errorf("creating directory %s: %w", outDir, mkdirErr)
		}

		code := stripPackageDecl(ext.code)
		code = wrapCodeForExtraction(code)
		var builder strings.Builder
		builder.WriteString("package lintcheck\n\n")
		guideComment := "guides/" + ext.relativePath + ".md"
		builder.WriteString("// Extracted from: " + guideComment + "\n\n")
		builder.WriteString(code)

		outFile := filepath.Join(outDir, "main.go")
		writeErr := os.WriteFile(outFile, []byte(builder.String()), 0o600)
		if writeErr != nil {
			return fmt.Errorf("writing %s: %w", outFile, writeErr)
		}
	}

	return nil
}

func writeGoMod(tmpDir string) error {
	goMod := "module github.com/wavilen/golangcilint-mcp/tmp/crosscheck\n\ngo 1.23\n"
	writeErr := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(goMod), 0o600)
	if writeErr != nil {
		return fmt.Errorf("writing go.mod: %w", writeErr)
	}
	return nil
}

func runGoimports(tmpDir string, extractions []extraction) {
	_, lookErr := exec.LookPath("goimports")
	if lookErr != nil {
		log.Println("warning: goimports not found — skipping import formatting")
		return
	}

	for _, ext := range extractions {
		file := filepath.Join(tmpDir, ext.relativePath, "main.go")
		cmd := exec.CommandContext(context.Background(), "goimports", "-w", file)
		out, cmdErr := cmd.CombinedOutput()
		if cmdErr != nil {
			log.Printf("warning: goimports failed on %s: %v\n%s", file, cmdErr, out)
		}
	}
	log.Println("Ran goimports on all extracted files")
}

func getExcludedLinters() map[string]bool {
	return map[string]bool{
		"typecheck":   true,
		"unused":      true,
		"whitespace":  true,
		"mnd":         true,
		"ineffassign": true,
		"forbidigo":   true,
		"noctx":       true,
		"sloglint":    true,
		"golines":     true,
		"gosec":       true,
	}
}

func runGolangciLint(projectRoot, tmpDir string, extractions []extraction) []golangciLintIssue {
	configPath := filepath.Join(projectRoot, "golden-config", ".golangci.yml")

	var allIssues []golangciLintIssue

	for _, ext := range extractions {
		pkgDir := filepath.Join(tmpDir, ext.relativePath)
		args := []string{
			"run",
			"--config", configPath,
			"--output.json.path", "stdout",
			".",
		}

		cmd := exec.CommandContext(context.Background(), "golangci-lint", args...)
		cmd.Dir = pkgDir

		out, cmdErr := cmd.Output()

		if len(out) == 0 && cmdErr != nil {
			log.Printf("warning: golangci-lint failed for %s: %v", ext.relativePath, cmdErr)
			continue
		}

		var output golangciLintOutput
		unmarshalErr := json.Unmarshal(out, &output)
		if unmarshalErr != nil {
			for _, line := range bytes.Split(out, []byte("\n")) {
				line = bytes.TrimSpace(line)
				if len(line) == 0 {
					continue
				}
				var obj golangciLintOutput
				lineErr := json.Unmarshal(line, &obj)
				if lineErr == nil && len(obj.Issues) > 0 {
					output.Issues = append(output.Issues, obj.Issues...)
				}
			}
		}
		allIssues = append(allIssues, output.Issues...)
	}

	excludedLinters := getExcludedLinters()
	var filtered []golangciLintIssue
	for _, issue := range allIssues {
		if !excludedLinters[issue.FromLinter] {
			filtered = append(filtered, issue)
		}
	}

	log.Printf("golangci-lint checked %d packages, found %d issues (%d after filtering excluded linters)",
		len(extractions), len(allIssues), len(filtered))
	return filtered
}

func buildReport(extractions []extraction, issues []golangciLintIssue) report {
	fileToGuide := make(map[string]string)
	for _, ext := range extractions {
		guide := "guides/" + ext.relativePath + ".md"
		fileToGuide[filepath.Join("tmp", "crosscheck", ext.relativePath, "main.go")] = guide
		fileToGuide[filepath.Join("..", "tmp", "crosscheck", ext.relativePath, "main.go")] = guide
		fileToGuide[filepath.Join(ext.relativePath, "main.go")] = guide
	}

	violations := make([]violation, 0, len(issues))
	failingGuides := make(map[string]bool)

	for _, issue := range issues {
		guide := fileToGuide[issue.Pos.Filename]
		if guide == "" {
			guide = filenameToGuide(issue.Pos.Filename)
		}

		violations = append(violations, violation{
			Guide:         guide,
			ExtractedFile: issue.Pos.Filename,
			Linter:        issue.FromLinter,
			Line:          issue.Pos.Line,
			Message:       issue.Text,
		})
		failingGuides[guide] = true
	}

	totalGuides := len(extractions)
	failingCount := len(failingGuides)
	passingCount := totalGuides - failingCount

	return report{
		TotalGuides:     totalGuides,
		Passing:         passingCount,
		Failing:         failingCount,
		TotalViolations: len(violations),
		Violations:      violations,
	}
}

func filenameToGuide(filename string) string {
	rel := filename
	for strings.HasPrefix(rel, "../") || strings.HasPrefix(rel, "./") {
		rel = strings.TrimPrefix(rel, "../")
		rel = strings.TrimPrefix(rel, "./")
	}
	rel = strings.TrimPrefix(rel, "tmp/crosscheck/")
	rel = strings.TrimSuffix(rel, "/main.go")
	return "guides/" + rel + ".md"
}

func writeJSONReport(path string, rpt report) error {
	data, marshalErr := json.MarshalIndent(rpt, "", "  ")
	if marshalErr != nil {
		return fmt.Errorf("marshaling report: %w", marshalErr)
	}
	writeErr := os.WriteFile(path, data, 0o600)
	if writeErr != nil {
		return fmt.Errorf("writing report file: %w", writeErr)
	}
	return nil
}

func printSummary(rpt report) {
	fmt.Fprintf(os.Stdout, "\nCrosscheck Results: %d/%d guides passing (%d failing, %d total violations)\n\n",
		rpt.Passing, rpt.TotalGuides, rpt.Failing, rpt.TotalViolations)

	if rpt.TotalViolations == 0 {
		fmt.Fprintln(os.Stdout, "All guides pass — no violations found!")
		fmt.Fprintln(os.Stdout, "\nFull report: tmp/crosscheck/violations.json")
		return
	}

	linterCounts := make(map[string]int)
	linterGuides := make(map[string]map[string]bool)
	for _, violation := range rpt.Violations {
		linterCounts[violation.Linter]++
		if linterGuides[violation.Linter] == nil {
			linterGuides[violation.Linter] = make(map[string]bool)
		}
		linterGuides[violation.Linter][violation.Guide] = true
	}

	linters := make([]string, 0, len(linterCounts))
	for name := range linterCounts {
		linters = append(linters, name)
	}
	sort.Slice(linters, func(left, right int) bool {
		return linterCounts[linters[left]] > linterCounts[linters[right]]
	})

	fmt.Fprintln(os.Stdout, "Violations by linter:")
	for _, name := range linters {
		fmt.Fprintf(
			os.Stdout,
			"  %s: %d violations across %d guides\n",
			name,
			linterCounts[name],
			len(linterGuides[name]),
		)
	}

	guideViolations := make(map[string][]violation)
	for _, violation := range rpt.Violations {
		guideViolations[violation.Guide] = append(guideViolations[violation.Guide], violation)
	}

	guides := make([]string, 0, len(guideViolations))
	for guide := range guideViolations {
		guides = append(guides, guide)
	}
	sort.Slice(guides, func(left, right int) bool {
		return len(guideViolations[guides[left]]) > len(guideViolations[guides[right]])
	})

	fmt.Fprintln(os.Stdout, "\nFailing guides:")
	for _, guide := range guides {
		violations := guideViolations[guide]
		linterSet := make(map[string]bool)
		for _, violation := range violations {
			linterSet[violation.Linter] = true
		}
		linterNames := make([]string, 0, len(linterSet))
		for name := range linterSet {
			linterNames = append(linterNames, name)
		}
		sort.Strings(linterNames)
		fmt.Fprintf(os.Stdout, "  %s — %d violations (%s)\n", guide, len(violations), strings.Join(linterNames, ", "))
	}

	fmt.Fprintln(os.Stdout, "\nFull report: tmp/crosscheck/violations.json")
}
