package e2e_test

import (
	"testing"
)

func TestCountJSONIssues(t *testing.T) {
	t.Run("counts single issue", func(t *testing.T) {
		output := `{"FromLinter":"errcheck","Pos":{"Filename":"main.go","Line":10}}`
		got := countJSONIssues(output)
		if got != 1 {
			t.Fatalf("expected 1 issue, got %d", got)
		}
	})

	t.Run("counts multiple issues", func(t *testing.T) {
		output := `{"FromLinter":"errcheck"}
{"FromLinter":"staticcheck","Pos":{"Line":5}}
{"FromLinter":"govet"}
`
		got := countJSONIssues(output)
		if got != 3 {
			t.Fatalf("expected 3 issues, got %d", got)
		}
	})

	t.Run("returns zero for empty string", func(t *testing.T) {
		got := countJSONIssues("")
		if got != 0 {
			t.Fatalf("expected 0 issues, got %d", got)
		}
	})

	t.Run("returns zero for non-JSON input without FromLinter", func(t *testing.T) {
		got := countJSONIssues("some random text without linter info")
		if got != 0 {
			t.Fatalf("expected 0 issues, got %d", got)
		}
	})

	t.Run("returns zero for valid JSON without FromLinter", func(t *testing.T) {
		output := `{"Type":"success","Message":"no issues found"}`
		got := countJSONIssues(output)
		if got != 0 {
			t.Fatalf("expected 0 issues, got %d", got)
		}
	})

	t.Run("counts across multiple lines with mixed content", func(t *testing.T) {
		output := `{"FromLinter":"errcheck","Pos":{"Filename":"a.go"}}
running...
{"FromLinter":"staticcheck","Pos":{"Filename":"b.go"}}
done
`
		got := countJSONIssues(output)
		if got != 2 {
			t.Fatalf("expected 2 issues, got %d", got)
		}
	})
}
