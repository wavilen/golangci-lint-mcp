package e2e

import (
	"regexp"
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

func TestRetryResultNaming(t *testing.T) {
	t.Run("attempt 0 has no retry suffix", func(t *testing.T) {
		name := runResultName("Simple", "GLM-5-Turbo", 0)
		assertString(t, name, "simple-glm-5-turbo-result.json")
	})
	t.Run("retry 1 includes -r1 suffix", func(t *testing.T) {
		name := runResultName("Simple", "GLM-5-Turbo", 1)
		assertString(t, name, "simple-glm-5-turbo-r1-result.json")
	})
	t.Run("retry 2 includes -r2 suffix", func(t *testing.T) {
		name := runResultName("Simple", "GLM-5-Turbo", 2)
		assertString(t, name, "simple-glm-5-turbo-r2-result.json")
	})
	t.Run("model with dots gets dashes", func(t *testing.T) {
		name := runResultName("Medium", "GLM-4.7", 0)
		assertString(t, name, "medium-glm-4-7-result.json")
	})
}

func TestRetryResultFilterRegex(t *testing.T) {
	pat := regexp.MustCompile(`-r\d+-result\.json$`)
	t.Run("final result passes filter", func(t *testing.T) {
		if pat.MatchString("simple-glm-5-turbo-result.json") {
			t.Fatal("should not match final result")
		}
	})
	t.Run("retry results are filtered", func(t *testing.T) {
		if !pat.MatchString("simple-glm-5-turbo-r1-result.json") {
			t.Fatal("should match retry 1 result")
		}
		if !pat.MatchString("simple-glm-5-turbo-r2-result.json") {
			t.Fatal("should match retry 2 result")
		}
	})
}

func TestParseSessionList(t *testing.T) {
	t.Run("returns first session ID from valid list", func(t *testing.T) {
		input := `[{"id":"ses_abc123","title":"test","updated":1,"created":1,"projectId":"p","directory":"/tmp"}]`
		id, err := ParseSessionList(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertString(t, id, "ses_abc123")
	})

	t.Run("returns error for empty array", func(t *testing.T) {
		_, err := ParseSessionList("[]")
		if err == nil {
			t.Fatal("expected error for empty array")
		}
	})

	t.Run("returns error for invalid JSON", func(t *testing.T) {
		_, err := ParseSessionList("not json")
		if err == nil {
			t.Fatal("expected error for invalid JSON")
		}
	})

	t.Run("returns error for null", func(t *testing.T) {
		_, err := ParseSessionList("null")
		if err == nil {
			t.Fatal("expected error for null")
		}
	})
}

func TestParseSessionExport(t *testing.T) {
	t.Run("counts tool calls and extracts subagent IDs", func(t *testing.T) {
		input := `{
			"messages": [
				{
					"info": {"role": "assistant", "tokens": {"input": 100, "output": 50, "reasoning": 10, "cache": {"read": 5, "write": 3}}, "cost": 0.002},
					"parts": [
						{"type": "tool", "tool": "golangci_lint_guide"},
						{"type": "tool", "tool": "task", "state": {"status": "completed", "metadata": {"sessionId": "ses_sub1"}}},
						{"type": "tool", "tool": "bash", "state": {"status": "completed"}}
					]
				}
			]
		}`
		result, err := ParseSessionExport(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.ToolCalls, 3, "tool calls")
		if len(result.SubagentIDs) != 1 {
			t.Fatalf("expected 1 subagent ID, got %d", len(result.SubagentIDs))
		}
		assertString(t, result.SubagentIDs[0], "ses_sub1")
		assertInt(t, result.TokenUsage.Input, 100, "input tokens")
		assertInt(t, result.TokenUsage.Output, 50, "output tokens")
		assertInt(t, result.TokenUsage.Reasoning, 10, "reasoning tokens")
		assertInt(t, result.TokenUsage.CacheRead, 5, "cache read")
		assertInt(t, result.TokenUsage.CacheWrite, 3, "cache write")
		assertFloat64(t, result.TotalCost, 0.002, "total cost")
	})

	t.Run("handles export with zero messages", func(t *testing.T) {
		input := `{"messages":[]}`
		result, err := ParseSessionExport(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.ToolCalls, 0, "tool calls")
		if len(result.SubagentIDs) != 0 {
			t.Fatalf("expected 0 subagent IDs, got %d", len(result.SubagentIDs))
		}
		assertInt(t, result.TokenUsage.Total(), 0, "total tokens")
		assertFloat64(t, result.TotalCost, 0, "total cost")
	})

	t.Run("handles tool parts without metadata", func(t *testing.T) {
		input := `{
			"messages": [
				{
					"info": {"role": "assistant"},
					"parts": [
						{"type": "tool", "tool": "bash"},
						{"type": "tool", "tool": "read"}
					]
				}
			]
		}`
		result, err := ParseSessionExport(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.ToolCalls, 2, "tool calls")
		if len(result.SubagentIDs) != 0 {
			t.Fatalf("expected 0 subagent IDs, got %d", len(result.SubagentIDs))
		}
	})

	t.Run("sums tokens across multiple messages", func(t *testing.T) {
		input := `{
			"messages": [
				{
					"info": {"role": "assistant", "tokens": {"input": 100, "output": 50, "reasoning": 10, "cache": {"read": 5, "write": 3}}},
					"parts": []
				},
				{
					"info": {"role": "assistant", "tokens": {"input": 200, "output": 75, "reasoning": 20, "cache": {"read": 10, "write": 7}}},
					"parts": []
				}
			]
		}`
		result, err := ParseSessionExport(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.TokenUsage.Input, 300, "input tokens")
		assertInt(t, result.TokenUsage.Output, 125, "output tokens")
		assertInt(t, result.TokenUsage.Reasoning, 30, "reasoning tokens")
		assertInt(t, result.TokenUsage.CacheRead, 15, "cache read")
		assertInt(t, result.TokenUsage.CacheWrite, 10, "cache write")
	})

	t.Run("handles nil tokens and cost gracefully", func(t *testing.T) {
		input := `{
			"messages": [
				{
					"info": {"role": "assistant"},
					"parts": [{"type": "tool", "tool": "bash"}]
				}
			]
		}`
		result, err := ParseSessionExport(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertInt(t, result.TokenUsage.Total(), 0, "total tokens")
		assertFloat64(t, result.TotalCost, 0, "total cost")
	})
}

func TestFormatTokenCount(t *testing.T) {
	cases := []struct {
		input  int
		expect string
	}{
		{500, "500"},
		{999, "999"},
		{1000, "1k"},
		{1500, "1.5k"},
		{18700, "18.7k"},
		{100000, "100k"},
		{1500000, "1.5M"},
	}
	for _, tc := range cases {
		got := FormatTokenCount(tc.input)
		assertString(t, got, tc.expect)
	}
}

func TestTokenUsageTotal(t *testing.T) {
	t.Run("sums all fields", func(t *testing.T) {
		tu := TokenUsage{Input: 100, Output: 50, Reasoning: 10, CacheRead: 30, CacheWrite: 5}
		assertInt(t, tu.Total(), 195, "total tokens")
	})
	t.Run("zero value returns zero", func(t *testing.T) {
		var tu TokenUsage
		assertInt(t, tu.Total(), 0, "total tokens")
	})
}

func assertInt(t *testing.T, got, want int, label string) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %d %s, got %d", want, label, got)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func assertFloat64(t *testing.T, got, want float64, label string) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %f %s, got %f", want, label, got)
	}
}
