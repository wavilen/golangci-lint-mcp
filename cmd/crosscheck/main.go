package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
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

type excludeConfig struct {
	Excludes map[string][]string `json:"excludes"`
}

type fixSuggestion struct {
	Guide        string `json:"guide"`
	Linter       string `json:"linter"`
	Message      string `json:"message"`
	SuggestedFix string `json:"suggested_fix"`
}

type fixReport struct {
	Trivial []fixSuggestion `json:"trivial"`
	Complex []fixSuggestion `json:"complex"`
}

type badValidation struct {
	Guide        string `json:"guide"`
	TargetLinter string `json:"target_linter"`
	Triggered    bool   `json:"triggered"`
	Issues       int    `json:"issues"`
	Details      string `json:"details,omitempty"`
}

var declarationKeywords = []string{
	"func ", "type ", "var ", "const ", "import ", "var(", "const(", "import(",
}

const linterGovet = "govet"

const minPathParts = 2

func main() {
	excludeConfigPath := flag.String(
		"exclude-config",
		"cmd/crosscheck/excludes.json",
		"path to per-guide exclude config JSON",
	)
	fixMode := flag.Bool("fix", false, "categorize violations into trivial/complex and output fixes.json")
	validateBad := flag.Bool("validate-bad", false, "validate Bad code blocks trigger their target linter")
	flag.Parse()

	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get working directory: %v", err)
	}

	// --validate-bad mode: separate pipeline for Bad examples
	if *validateBad {
		runBadValidation(projectRoot)
		return
	}

	perGuideExcludes := loadExcludeConfig(*excludeConfigPath)

	extractions, skipped := extractBlocks(filepath.Join(projectRoot, "guides"), "Good")
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

	issues := runGolangciLint(projectRoot, tmpDir, extractions, perGuideExcludes)

	rpt := buildReport(extractions, issues)

	reportPath := filepath.Join(tmpDir, "violations.json")
	reportErr := writeJSONReport(reportPath, rpt)
	if reportErr != nil {
		log.Fatalf("writing report: %v", reportErr)
	}

	printSummary(rpt)

	if *fixMode {
		fixRpt := categorizeViolations(rpt)
		fixPath := filepath.Join(tmpDir, "fixes.json")
		fixErr := writeFixReport(fixPath, fixRpt)
		if fixErr != nil {
			log.Fatalf("writing fix report: %v", fixErr)
		}
		fmt.Fprintf(os.Stdout, "\n%d trivial violations (auto-fixable), %d complex violations (need review)\n",
			len(fixRpt.Trivial), len(fixRpt.Complex))
		fmt.Fprintln(os.Stdout, "Fix report: "+fixPath)
	}
}

func extractBlocks(guidesDir, sectionHeader string) ([]extraction, []string) {
	examplesRe := regexp.MustCompile(`(?s)<examples>(.*?)</examples>`)
	sectionRe := regexp.MustCompile(fmt.Sprintf(`(?s)## %s\s*\n`+"`"+"`"+"`"+`go\n(.*?)`+"`"+"`"+"`", sectionHeader))

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

		sectionMatch := sectionRe.FindStringSubmatch(examplesContent)
		if sectionMatch == nil {
			skipped = append(skipped, path)
			return nil
		}
		code := sectionMatch[1]

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

func runGolangciLint(
	projectRoot, tmpDir string,
	extractions []extraction,
	perGuideExcludes map[string][]string,
) []golangciLintIssue {
	configPath := filepath.Join(projectRoot, "golden-config", ".golangci.yml")

	var allIssues []golangciLintIssue

	for _, ext := range extractions {
		issues := lintSinglePackage(configPath, tmpDir, ext)
		allIssues = append(allIssues, issues...)
	}

	filtered := filterIssues(allIssues, perGuideExcludes)

	log.Printf("golangci-lint checked %d packages, found %d issues (%d after filtering excluded linters)",
		len(extractions), len(allIssues), len(filtered))
	return filtered
}

func lintSinglePackage(configPath, tmpDir string, ext extraction) []golangciLintIssue {
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
		return nil
	}

	output := parseLintJSON(out)
	return output.Issues
}

func parseLintJSON(out []byte) golangciLintOutput {
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
	return output
}

func filterIssues(allIssues []golangciLintIssue, perGuideExcludes map[string][]string) []golangciLintIssue {
	excludedLinters := getExcludedLinters()
	var filtered []golangciLintIssue
	for _, issue := range allIssues {
		if excludedLinters[issue.FromLinter] {
			continue
		}
		guidePath := filenameToGuide(issue.Pos.Filename)
		if guideExcludes, ok := perGuideExcludes[guidePath]; ok {
			if slices.Contains(guideExcludes, issue.FromLinter) {
				continue
			}
		}
		filtered = append(filtered, issue)
	}
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

func loadExcludeConfig(path string) map[string][]string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("warning: cannot read exclude config %s: %v — using empty excludes", path, err)
		return make(map[string][]string)
	}
	var cfg excludeConfig
	unmarshalErr := json.Unmarshal(data, &cfg)
	if unmarshalErr != nil {
		log.Printf("warning: cannot parse exclude config %s: %v — using empty excludes", path, unmarshalErr)
		return make(map[string][]string)
	}
	if cfg.Excludes == nil {
		return make(map[string][]string)
	}
	return cfg.Excludes
}

func categorizeViolations(rpt report) fixReport {
	trivialLinters := map[string]bool{
		"errcheck":    true,
		"intrange":    true,
		"prealloc":    true,
		"perfsprint":  true,
		"godot":       true,
		"modernize":   true,
		"exhaustruct": true,
		"varnamelen":  true,
	}

	complexLinters := map[string]bool{
		"wrapcheck":   true,
		"noinlineerr": true,
		"godox":       true,
	}

	var fixRpt fixReport
	for _, violation := range rpt.Violations {
		suggestion := fixSuggestion{
			Guide:        violation.Guide,
			Linter:       violation.Linter,
			Message:      violation.Message,
			SuggestedFix: suggestFix(violation.Linter, violation.Message),
		}

		switch {
		case trivialLinters[violation.Linter]:
			fixRpt.Trivial = append(fixRpt.Trivial, suggestion)
		case complexLinters[violation.Linter]:
			fixRpt.Complex = append(fixRpt.Complex, suggestion)
		case violation.Linter == "revive" && strings.Contains(violation.Message, "unused-parameter"):
			fixRpt.Trivial = append(fixRpt.Trivial, suggestion)
		case violation.Linter == linterGovet && (strings.Contains(violation.Message, "shadow") || strings.Contains(violation.Message, "unusedresult")):
			fixRpt.Complex = append(fixRpt.Complex, suggestion)
		default:
			fixRpt.Complex = append(fixRpt.Complex, suggestion)
		}
	}
	return fixRpt
}

func suggestFix(linter, message string) string {
	switch {
	case linter == "errcheck":
		return "Check returned error or explicitly discard with _"
	case linter == "intrange":
		return "Replace for loop with integer range (for i := range n)"
	case linter == "prealloc":
		return "Preallocate slice with make() using known length"
	case linter == "perfsprint":
		return "Replace fmt.Sprintf with strconv or direct string concatenation"
	case linter == "godot":
		return "Add period at end of comment"
	case linter == "modernize":
		return "Use modern Go idiom as suggested by linter message"
	case linter == "exhaustruct":
		return "Initialize all struct fields or add to exhaustruct exclude pattern"
	case linter == "varnamelen":
		return "Use longer variable name matching its scope"
	case linter == "wrapcheck":
		return "Wrap error returned from external package with fmt.Errorf(\"...: %w\", err)"
	case linter == "noinlineerr":
		return "Extract error check: assign err first, then check in separate if statement"
	case linter == "godox":
		return "Convert TODO/FIXME to tracked issue or add nolint directive"
	case linter == "revive" && strings.Contains(message, "unused-parameter"):
		return "Rename unused parameter to '_'"
	case linter == linterGovet && strings.Contains(message, "shadow"):
		return "Rename variable to avoid shadowing outer scope variable"
	case linter == linterGovet && strings.Contains(message, "unusedresult"):
		return "Use or explicitly discard the result of the function call"
	default:
		return "Review and fix per linter suggestion"
	}
}

func writeFixReport(path string, fr fixReport) error {
	data, marshalErr := json.MarshalIndent(fr, "", "  ")
	if marshalErr != nil {
		return fmt.Errorf("marshaling fix report: %w", marshalErr)
	}
	writeErr := os.WriteFile(path, data, 0o600)
	if writeErr != nil {
		return fmt.Errorf("writing fix report file: %w", writeErr)
	}
	return nil
}

func targetLinterFromPath(relativePath string) string {
	parts := strings.Split(relativePath, "/")
	if len(parts) >= minPathParts {
		// Subdirectory guide: e.g., "gosec/G502" → "gosec", "staticcheck/SA1006" → "staticcheck"
		return parts[0]
	}
	// Root-level guide: e.g., "errcheck" → "errcheck"
	return parts[0]
}

// fixUnusedVars attempts to fix "declared and not used" compile errors in
// Bad extractions by adding `_ = varName` assignments. Bad code snippets
// often declare variables that aren't referenced within the snippet.
func fixUnusedVars(tmpDir string, extractions []extraction) {
	unusedRe := regexp.MustCompile(`declared and not used: (\w+)`)

	for _, ext := range extractions {
		file := filepath.Join(tmpDir, ext.relativePath, "main.go")

		// Try up to 3 rounds of fixing (each fix may expose new unused vars)
		for range 3 {
			cmd := exec.CommandContext(context.Background(), "go", "build", ".")
			cmd.Dir = filepath.Join(tmpDir, ext.relativePath)
			out, _ := cmd.CombinedOutput()

			matches := unusedRe.FindAllStringSubmatch(string(out), -1)
			if len(matches) == 0 {
				break
			}

			data, readErr := os.ReadFile(file)
			if readErr != nil {
				break
			}
			content := string(data)

			// Find the last closing brace (end of func _() {} wrapper)
			lastBrace := strings.LastIndex(content, "}")
			if lastBrace == -1 {
				break
			}

			// Build _ = statements for all unused variables
			var suppressions strings.Builder
			suppressions.WriteString("\t// suppress unused var errors for snippet extraction\n")
			for _, match := range matches {
				suppressions.WriteString("\t_ = " + match[1] + "\n")
			}

			// Insert before the last closing brace
			newContent := content[:lastBrace] + suppressions.String() + content[lastBrace:]
			writeErr := os.WriteFile(
				file, []byte(newContent), 0o600,
			)
			if writeErr != nil {
				log.Printf("warning: cannot fix unused vars in %s: %v", file, writeErr)
				break
			}
		}
	}
}

func runBadValidation(projectRoot string) {
	guidesDir := filepath.Join(projectRoot, "guides")
	extractions, skipped := extractBlocks(guidesDir, "Bad")
	log.Printf("Bad block extraction: %d with Bad blocks, %d skipped",
		len(extractions), len(skipped))

	if len(extractions) == 0 {
		fmt.Fprintln(os.Stdout, "No Bad code blocks found — validation skipped")
		return
	}

	tmpDir := filepath.Join(projectRoot, "tmp", "crosscheck-bad")
	prepareBadTmpDir(tmpDir)
	writeBadExtractions(tmpDir, extractions)

	badConfigPath := filepath.Join(projectRoot, "cmd", "crosscheck", "bad-validation-config.yml")
	validations := validateBadExtractions(tmpDir, badConfigPath, extractions)

	writeBadValidationReport(tmpDir, validations)
	printBadValidationSummary(validations, tmpDir)
}

func prepareBadTmpDir(tmpDir string) {
	_, statErr := os.Stat(tmpDir)
	if statErr == nil {
		removeErr := os.RemoveAll(tmpDir)
		if removeErr != nil {
			log.Fatalf("clearing tmp/crosscheck-bad: %v", removeErr)
		}
	}
	mkdirErr := os.MkdirAll(tmpDir, 0o750)
	if mkdirErr != nil {
		log.Fatalf("creating tmp/crosscheck-bad: %v", mkdirErr)
	}
}

func writeBadExtractions(tmpDir string, extractions []extraction) {
	writeErr := writeExtractions(tmpDir, extractions)
	if writeErr != nil {
		log.Fatalf("writing Bad extractions: %v", writeErr)
	}

	modErr := writeGoMod(tmpDir)
	if modErr != nil {
		log.Fatalf("writing go.mod: %v", modErr)
	}

	runGoimports(tmpDir, extractions)
	fixUnusedVars(tmpDir, extractions)
}

func validateBadExtractions(tmpDir, badConfigPath string, extractions []extraction) []badValidation {
	validations := make([]badValidation, 0, len(extractions))

	for _, ext := range extractions {
		guidePath := "guides/" + ext.relativePath + ".md"
		targetLinter := targetLinterFromPath(ext.relativePath)
		validation := lintBadPackage(tmpDir, badConfigPath, ext, guidePath, targetLinter)
		validations = append(validations, validation)
		log.Printf("Validated %s (target: %s) — triggered: %v, issues: %d",
			guidePath, targetLinter, validation.Triggered, validation.Issues)
	}

	return validations
}

func lintBadPackage(tmpDir, badConfigPath string, ext extraction, guidePath, targetLinter string) badValidation {
	pkgDir := filepath.Join(tmpDir, ext.relativePath)

	args := []string{
		"run",
		"--config", badConfigPath,
		"--enable-only", targetLinter,
		"--output.json.path", "stdout",
		".",
	}

	cmd := exec.CommandContext(context.Background(), "golangci-lint", args...)
	cmd.Dir = pkgDir

	out, cmdErr := cmd.Output()

	var issues int
	var triggered bool
	var details string

	if len(out) > 0 {
		output := parseLintJSON(out)
		for _, issue := range output.Issues {
			if issue.FromLinter == targetLinter {
				issues++
			}
		}
		triggered = issues > 0
	}

	if cmdErr != nil && len(out) == 0 {
		details = fmt.Sprintf("golangci-lint error: %v", cmdErr)
	}

	return badValidation{
		Guide:        guidePath,
		TargetLinter: targetLinter,
		Triggered:    triggered,
		Issues:       issues,
		Details:      details,
	}
}

func writeBadValidationReport(tmpDir string, validations []badValidation) {
	reportPath := filepath.Join(tmpDir, "bad-validation.json")
	data, marshalErr := json.MarshalIndent(validations, "", "  ")
	if marshalErr != nil {
		log.Fatalf("marshaling bad validation report: %v", marshalErr)
	}
	writeErr := os.WriteFile(reportPath, data, 0o600)
	if writeErr != nil {
		log.Fatalf("writing bad validation report: %v", writeErr)
	}
}

func printBadValidationSummary(validations []badValidation, tmpDir string) {
	reportPath := filepath.Join(tmpDir, "bad-validation.json")

	passing := 0
	for _, validation := range validations {
		if validation.Triggered {
			passing++
		}
	}
	fmt.Fprintf(os.Stdout, "\nBad Example Validation: %d/%d Bad examples trigger their target linter\n",
		passing, len(validations))

	failing := len(validations) - passing
	if failing > 0 {
		fmt.Fprintln(os.Stdout, "\nFailing Bad examples (do NOT trigger target linter):")
		for _, validation := range validations {
			if !validation.Triggered {
				detail := ""
				if validation.Details != "" {
					detail = fmt.Sprintf(" (%s)", validation.Details)
				}
				fmt.Fprintf(os.Stdout, "  %s — target: %s, issues: %d%s\n",
					validation.Guide, validation.TargetLinter, validation.Issues, detail)
			}
		}
	}

	fmt.Fprintln(os.Stdout, "\nFull report: "+reportPath)
}
