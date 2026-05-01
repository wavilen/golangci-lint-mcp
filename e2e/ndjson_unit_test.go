package e2e_test

import (
	"strings"
	"testing"
)

func TestParseNDJSON(t *testing.T) {
	t.Run("parses known tool_use and tool_result events", func(t *testing.T) {
		input := `{"type":"tool_use","data":{"name":"golangci_lint_guide"}}
{"type":"tool_result","data":{"content":"fix applied"}}
{"type":"input_json","data":{"prompt":"hello"}}
`
		result, err := ParseNDJSON(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, len(result.Events), 3, "events")
		assertInt(t, result.ToolCalls, 3, "tool calls")
		assertInt(t, len(result.RawLines), 3, "raw lines")
	})

	t.Run("skips non-JSON lines and empty lines", func(t *testing.T) {
		input := `{"type":"tool_use","data":{}}

not json at all
{"type":"assistant","data":{"text":"thinking"}}
another bad line
`
		result, err := ParseNDJSON(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, len(result.Events), 2, "events")
		assertInt(t, result.ToolCalls, 1, "tool calls")
		assertInt(t, len(result.RawLines), 4, "raw lines")
	})

	t.Run("handles empty input", func(t *testing.T) {
		result, err := ParseNDJSON(strings.NewReader(""))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, len(result.Events), 0, "events")
		assertInt(t, result.ToolCalls, 0, "tool calls")
		assertInt(t, len(result.RawLines), 0, "raw lines")
	})

	t.Run("counts only tool_use_tool_result_input_json events", func(t *testing.T) {
		input := `{"type":"tool_use","data":{}}
{"type":"assistant","data":{}}
{"type":"tool_result","data":{}}
{"type":"input_json","data":{}}
{"type":"assistant","data":{}}
`
		result, err := ParseNDJSON(strings.NewReader(input))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.ToolCalls, 3, "tool calls")
		assertInt(t, len(result.Events), 5, "events")
	})
}

func assertInt(t *testing.T, got, want int, label string) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %d %s, got %d", want, label, got)
	}
}
