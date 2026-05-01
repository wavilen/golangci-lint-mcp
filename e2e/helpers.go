package e2e

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"time"
)

type EvalResult struct {
	TestCase       string  `json:"TestCase"`
	Model          string  `json:"Model"`
	BeforeIssues   int     `json:"BeforeIssues"`
	AfterIssues    int     `json:"AfterIssues"`
	BuildsClean    bool    `json:"BuildsClean"`
	LintClean      bool    `json:"LintClean"`
	NolintCount    int     `json:"NolintCount"`
	IssueReduction float64 `json:"IssueReduction"`
	ToolCalls      int     `json:"ToolCalls"`
	Retries        int     `json:"Retries"`
	Pass           bool    `json:"Pass"`
	Error          string  `json:"Error"`
	ConfigModified bool    `json:"ConfigModified"`
	ConfigDiff     string  `json:"ConfigDiff,omitempty"`

	SubagentToolCalls int        `json:"SubagentToolCalls"`
	SubagentCount     int        `json:"SubagentCount"`
	TokenUsage        TokenUsage `json:"TokenUsage"`
	TotalCost         float64    `json:"TotalCost"`
}

type TokenUsage struct {
	Input      int `json:"Input"`
	Output     int `json:"Output"`
	Reasoning  int `json:"Reasoning"`
	CacheRead  int `json:"CacheRead"`
	CacheWrite int `json:"CacheWrite"`
}

func (tu TokenUsage) Total() int {
	return tu.Input + tu.Output + tu.Reasoning + tu.CacheRead + tu.CacheWrite
}

type OpenCodeEvent struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data,omitempty"`
}

type NDJSONResult struct {
	Events    []OpenCodeEvent
	ToolCalls int
	RawLines  []string
}

type ReportData struct {
	Timestamp    time.Time
	Results      []*EvalResult
	Models       []string
	Fixtures     []string
	PassCount    int
	TotalCount   int
	PassRate     float64
	AvgReduction float64
	Errors       []string
}

func countJSONIssues(output string) int {
	output = strings.TrimSpace(output)
	if output == "" {
		return 0
	}
	// Try v2 wrapped JSON first: {"Issues":[...],"Report":{...}}
	// Use json.Decoder so trailing text after the JSON blob is ignored.
	// Decode into raw map to check for "Issues" key presence (distinguishes
	// v2 wrapped format from NDJSON where first line has no Issues key).
	var raw map[string]json.RawMessage
	if json.NewDecoder(strings.NewReader(output)).Decode(&raw) == nil {
		if issuesRaw, ok := raw["Issues"]; ok {
			var issues []struct{}
			if json.Unmarshal(issuesRaw, &issues) == nil {
				return len(issues)
			}
		}
	}
	// Fallback: v1 NDJSON (one JSON per line)
	count := 0
	for line := range strings.SplitSeq(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var obj map[string]json.RawMessage
		if json.Unmarshal([]byte(line), &obj) == nil {
			if _, ok := obj["FromLinter"]; ok {
				count++
			}
		}
	}
	return count
}

func runResultName(testCase, model string, retry int) string {
	baseName := fmt.Sprintf("%s-%s",
		strings.ToLower(testCase),
		strings.ToLower(strings.ReplaceAll(model, ".", "-")),
	)
	if retry > 0 {
		return fmt.Sprintf("%s-r%d-result.json", baseName, retry)
	}
	return fmt.Sprintf("%s-result.json", baseName)
}

const (
	scannerBufSize  = 1024 * 1024
	scannerMaxRatio = 10
	percentage      = 100
)

func ParseNDJSON(reader io.Reader) (*NDJSONResult, error) {
	result := new(NDJSONResult)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, scannerBufSize), scannerMaxRatio*scannerBufSize)

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

		switch event.Type {
		case "tool_use", "tool_result", "input_json":
			result.ToolCalls++
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("scanning NDJSON: %w", err)
	}

	return result, nil
}

func GenerateReport(results []*EvalResult, outputPath string) error {
	modelsSeen := make(map[string]bool)
	fixturesSeen := make(map[string]bool)
	passCount := 0
	totalReduction := 0.0
	var errors []string

	for _, r := range results {
		modelsSeen[r.Model] = true
		fixturesSeen[r.TestCase] = true
		if r.Pass {
			passCount++
		}
		totalReduction += r.IssueReduction
		if r.Error != "" {
			errors = append(errors, fmt.Sprintf("%s/%s: %s", r.TestCase, r.Model, r.Error))
		}
	}

	models := make([]string, 0, len(modelsSeen))
	for m := range modelsSeen {
		models = append(models, m)
	}
	fixtures := make([]string, 0, len(fixturesSeen))
	for f := range fixturesSeen {
		fixtures = append(fixtures, f)
	}

	avgReduction := 0.0
	if len(results) > 0 {
		avgReduction = totalReduction / float64(len(results))
	}

	data := ReportData{
		Timestamp:    time.Now(),
		Results:      results,
		Models:       models,
		Fixtures:     fixtures,
		PassCount:    passCount,
		TotalCount:   len(results),
		PassRate:     0,
		AvgReduction: avgReduction,
		Errors:       errors,
	}
	if len(results) > 0 {
		data.PassRate = float64(passCount) / float64(len(results)) * percentage
	}

	funcMap := template.FuncMap{
		"formatTokens": FormatTokenCount,
	}
	tmpl, err := template.New("report.html.tpl").Funcs(funcMap).ParseFiles("report.html.tpl")
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating report: %w", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return fmt.Errorf("executing report template: %w", err)
	}

	return nil
}

// --- Export JSON parsing types (internal) ---

type sessionListEntry struct {
	ID string `json:"id"`
}

type sessionExport struct {
	Messages []exportMessage `json:"messages"`
}

type exportMessage struct {
	Info  exportMessageInfo `json:"info"`
	Parts []exportPart      `json:"parts"`
}

type exportMessageInfo struct {
	Tokens *exportTokens `json:"tokens,omitempty"`
	Cost   *float64      `json:"cost,omitempty"`
}

type exportTokens struct {
	Input     int          `json:"input"`
	Output    int          `json:"output"`
	Reasoning int          `json:"reasoning"`
	Cache     *exportCache `json:"cache,omitempty"`
}

type exportCache struct {
	Read  int `json:"read"`
	Write int `json:"write"`
}

type exportPart struct {
	Type  string       `json:"type"`
	Tool  string       `json:"tool,omitempty"`
	State *exportState `json:"state,omitempty"`
}

type exportState struct {
	Metadata *exportMeta `json:"metadata,omitempty"`
}

type exportMeta struct {
	SessionID string `json:"sessionId"`
}

// ParseExportResult holds the aggregated data extracted from an opencode session export.
type ParseExportResult struct {
	ToolCalls   int
	SubagentIDs []string
	TokenUsage  TokenUsage
	TotalCost   float64
}

// ParseSessionList parses opencode session list --format json output and returns the first session ID.
func ParseSessionList(output string) (string, error) {
	var entries []sessionListEntry
	if err := json.Unmarshal([]byte(output), &entries); err != nil {
		return "", fmt.Errorf("parsing session list: %w", err)
	}
	if len(entries) == 0 {
		return "", fmt.Errorf("session list is empty")
	}
	return entries[0].ID, nil
}

// ParseSessionExport parses a single opencode export <id> JSON output and extracts
// tool call counts, subagent session IDs, token usage, and total cost.
func ParseSessionExport(output string) (*ParseExportResult, error) {
	var exp sessionExport
	if err := json.Unmarshal([]byte(output), &exp); err != nil {
		return nil, fmt.Errorf("parsing session export: %w", err)
	}

	result := &ParseExportResult{}

	for _, msg := range exp.Messages {
		// Sum tokens
		if msg.Info.Tokens != nil {
			result.TokenUsage.Input += msg.Info.Tokens.Input
			result.TokenUsage.Output += msg.Info.Tokens.Output
			result.TokenUsage.Reasoning += msg.Info.Tokens.Reasoning
			if msg.Info.Tokens.Cache != nil {
				result.TokenUsage.CacheRead += msg.Info.Tokens.Cache.Read
				result.TokenUsage.CacheWrite += msg.Info.Tokens.Cache.Write
			}
		}

		// Sum cost
		if msg.Info.Cost != nil {
			result.TotalCost += *msg.Info.Cost
		}

		// Count tool calls and extract subagent IDs
		for _, part := range msg.Parts {
			if part.Type == "tool" {
				result.ToolCalls++
				if part.Tool == "task" && part.State != nil && part.State.Metadata != nil {
					if sid := part.State.Metadata.SessionID; sid != "" {
						result.SubagentIDs = append(result.SubagentIDs, sid)
					}
				}
			}
		}
	}

	return result, nil
}

// FormatTokenCount formats a token count with k/M suffix for human-readable display.
func FormatTokenCount(count int) string {
	if count >= 1_000_000 {
		s := fmt.Sprintf("%.1fM", float64(count)/1_000_000)
		if strings.HasSuffix(s, ".0M") {
			return fmt.Sprintf("%.0fM", float64(count)/1_000_000)
		}
		return s
	}
	if count >= 1_000 {
		s := fmt.Sprintf("%.1fk", float64(count)/1_000)
		if strings.HasSuffix(s, ".0k") {
			return fmt.Sprintf("%.0fk", float64(count)/1_000)
		}
		return s
	}
	return fmt.Sprintf("%d", count)
}
