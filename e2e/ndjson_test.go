package e2e_test

import (
	"bufio"
	"encoding/json"
	"io"
)

// OpenCodeEvent represents a single NDJSON event from opencode --format json.
type OpenCodeEvent struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data,omitempty"`
}

// NDJSONResult holds parsed results from opencode output.
type NDJSONResult struct {
	Events    []OpenCodeEvent
	ToolCalls int      // count of tool_use/tool_result events
	RawLines  []string // raw lines for debug
}

// ParseNDJSON reads NDJSON output from opencode run --format json.
// Non-JSON lines are silently skipped.
func ParseNDJSON(reader io.Reader) (*NDJSONResult, error) {
	result := &NDJSONResult{}
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 1024*1024), 10*1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		result.RawLines = append(result.RawLines, line)

		var event OpenCodeEvent
		err := json.Unmarshal([]byte(line), &event)
		if err != nil {
			continue
		}
		result.Events = append(result.Events, event)

		// Count tool-related events (schema discovered empirically from probe test)
		switch event.Type {
		case "tool_use", "tool_result", "input_json":
			result.ToolCalls++
		}
	}

	return result, scanner.Err()
}
