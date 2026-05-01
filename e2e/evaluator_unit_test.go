package e2e

import (
	"testing"
)

func TestCountJSONIssues(t *testing.T) {
	t.Run("v2 wrapped JSON with issues", func(t *testing.T) {
		output := `{"Issues":[{"FromLinter":"errcheck","Text":"unhandled error"},{"FromLinter":"govet","Text":"printf"}],"Report":{"Warnings":[]}}`
		got := countJSONIssues(output)
		if got != 2 {
			t.Fatalf("expected %d issues, got %d", 2, got)
		}
	})

	t.Run("v2 wrapped JSON with empty issues", func(t *testing.T) {
		output := `{"Issues":[],"Report":{"Warnings":[]}}`
		got := countJSONIssues(output)
		if got != 0 {
			t.Fatalf("expected %d issues, got %d", 0, got)
		}
	})

	t.Run("v2 wrapped JSON with text after blob", func(t *testing.T) {
		output := "{\"Issues\":[{\"FromLinter\":\"errcheck\"}],\"Report\":{}}\n2 issues:\n* errcheck: 1\n"
		got := countJSONIssues(output)
		if got != 1 {
			t.Fatalf("expected %d issues, got %d", 1, got)
		}
	})

	t.Run("NDJSON with multiple issues", func(t *testing.T) {
		output := "{\"FromLinter\":\"errcheck\"}\n{\"FromLinter\":\"staticcheck\",\"Pos\":{\"Line\":5}}\n{\"FromLinter\":\"govet\"}\n"
		got := countJSONIssues(output)
		if got != 3 {
			t.Fatalf("expected %d issues, got %d", 3, got)
		}
	})

	t.Run("NDJSON with mixed content", func(t *testing.T) {
		output := "{\"FromLinter\":\"errcheck\",\"Pos\":{\"Filename\":\"a.go\"}}\nrunning...\n{\"FromLinter\":\"staticcheck\",\"Pos\":{\"Filename\":\"b.go\"}}\ndone\n"
		got := countJSONIssues(output)
		if got != 2 {
			t.Fatalf("expected %d issues, got %d", 2, got)
		}
	})

	t.Run("NDJSON line without FromLinter not counted", func(t *testing.T) {
		output := "{\"FromLinter\":\"errcheck\"}\n{\"Text\":\"some other thing\"}\n{\"FromLinter\":\"govet\"}\n"
		got := countJSONIssues(output)
		if got != 2 {
			t.Fatalf("expected %d issues, got %d", 2, got)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		got := countJSONIssues("")
		if got != 0 {
			t.Fatalf("expected %d issues, got %d", 0, got)
		}
	})

	t.Run("non-JSON text", func(t *testing.T) {
		got := countJSONIssues("some random text without linter info")
		if got != 0 {
			t.Fatalf("expected %d issues, got %d", 0, got)
		}
	})

	t.Run("valid JSON without Issues or FromLinter", func(t *testing.T) {
		output := `{"Type":"success","Message":"no issues found"}`
		got := countJSONIssues(output)
		if got != 0 {
			t.Fatalf("expected %d issues, got %d", 0, got)
		}
	})
}
