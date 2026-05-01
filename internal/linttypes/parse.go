package linttypes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseNDJSON attempts to parse newline-delimited individual LintIssue objects.
// Non-JSON lines are skipped. Returns collected issues in a LintJSONResult.
func ParseNDJSON(input string) LintJSONResult {
	var issues []LintIssue
	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var issue LintIssue
		if json.Unmarshal([]byte(line), &issue) == nil && issue.FromLinter != "" {
			issues = append(issues, issue)
		}
	}
	return LintJSONResult{Issues: issues}
}

// ParseLintOutput parses golangci-lint JSON output, trying single-object JSON
// first, then falling back to NDJSON line-by-line parsing.
// Returns the parsed result or the initial JSON unmarshal error.
func ParseLintOutput(output string) (LintJSONResult, error) {
	firstLine := output
	if before, _, found := strings.Cut(output, "\n"); found {
		firstLine = before
	}
	var result LintJSONResult
	err := json.Unmarshal([]byte(firstLine), &result)
	if err != nil || len(result.Issues) == 0 {
		ndjsonResult := ParseNDJSON(output)
		if len(ndjsonResult.Issues) > 0 {
			return ndjsonResult, nil
		}
		if err != nil {
			return LintJSONResult{}, fmt.Errorf("parse lint output: %w", err)
		}
	}
	return result, nil
}
