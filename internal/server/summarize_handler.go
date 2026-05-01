package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

func makeSummarizeHandler(
	_ *guides.Store,
) func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output, err := req.RequireString("output")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'output': %v", err)), nil
		}
		output = strings.TrimSpace(output)
		if output == "" {
			return mcp.NewToolResultError("parameter 'output' must not be empty"), nil
		}

		firstLine := output
		if before, _, found := strings.Cut(output, "\n"); found {
			firstLine = before
		}

		var result LintJSONResult
		unmarshalErr := json.Unmarshal([]byte(firstLine), &result)
		if unmarshalErr != nil || len(result.Issues) == 0 {
			ndjsonResult := parseNDJSON(output)
			if len(ndjsonResult.Issues) > 0 {
				result = ndjsonResult
			} else if unmarshalErr != nil {
				return mcp.NewToolResultError(
					fmt.Sprintf(
						"invalid JSON: %v. Try golangci_lint_parse for full guidance, "+
							"or golangci_lint_run to re-run and parse automatically.",
						unmarshalErr,
					)), nil
			}
		}

		if len(result.Issues) == 0 {
			return mcp.NewToolResultText("No issues found."), nil
		}

		// Unified pipeline: analyze → build response (D-03)
		// includeGuidance=false: summarize never shows guidance (D-09)
		strategyResult := AnalyzeStrategy(result.Issues)
		return mcp.NewToolResultText(
			BuildResponse(strategyResult, "", nil, Options{}, false, false)), nil
	}
}
