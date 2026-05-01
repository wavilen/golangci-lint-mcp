// Package linttypes provides shared types and utilities for golangci-lint output parsing
// used by both the MCP server and the crosscheck command.
package linttypes

import "strings"

// LintIssue represents a single golangci-lint diagnostic.
type LintIssue struct {
	FromLinter string `json:"FromLinter"`
	Text       string `json:"Text"`
	Pos        struct {
		Filename string `json:"Filename"`
		Line     int    `json:"Line"`
		Column   int    `json:"Column"`
	} `json:"Pos"`
}

// LintJSONResult represents the top-level golangci-lint JSON output.
type LintJSONResult struct {
	Issues []LintIssue `json:"Issues"`
}

// diagnostic is an unexported key type for issue deduplication.
type diagnostic struct {
	linter string
	rule   string
}

// ExtractRule extracts the rule name from a lint issue text.
// Issue text format is "ruleName: description" — this returns the part before ": ".
func ExtractRule(text string) string {
	before, _, found := strings.Cut(text, ": ")
	if !found {
		return ""
	}
	return strings.TrimSpace(before)
}

// DeduplicateIssues removes duplicate (linter, rule) pairs from the issue slice.
func DeduplicateIssues(issues []LintIssue) []LintIssue {
	seen := make(map[diagnostic]bool)
	unique := make([]LintIssue, 0, len(issues))
	for _, issue := range issues {
		rule := ExtractRule(issue.Text)
		key := diagnostic{linter: issue.FromLinter, rule: rule}
		if !seen[key] {
			seen[key] = true
			unique = append(unique, issue)
		}
	}
	return unique
}
