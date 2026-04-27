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

		var result lintJSONResult
		unmarshalErr := json.Unmarshal([]byte(firstLine), &result)
		if unmarshalErr != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid JSON: %v", unmarshalErr)), nil
		}

		if len(result.Issues) == 0 {
			return mcp.NewToolResultText("No issues found."), nil
		}

		unique := deduplicateIssues(result.Issues)
		packages := extractPackagesFromIssues(unique)
		strategyName, strategyReason := recommendStrategy(len(unique), len(packages))
		linterBreakdown := buildLinterBreakdown(unique)
		packageBreakdown := buildPackageBreakdown(packages)

		var builder strings.Builder
		fmt.Fprintf(&builder, "## Summary\n\n")
		fmt.Fprintf(&builder, "- Total issues: %d\n", len(result.Issues))
		fmt.Fprintf(&builder, "- Unique diagnostics: %d\n", len(unique))
		fmt.Fprintf(&builder, "- Packages affected: %d\n", len(packages))
		fmt.Fprintf(&builder, "- Strategy: %s (%s)\n", strategyName, strategyReason)

		if packageBreakdown != "" {
			fmt.Fprintf(&builder, "\n## Package Breakdown\n\n%s\n", packageBreakdown)
		}

		fmt.Fprintf(&builder, "\n## Linter Breakdown\n\n%s\n", linterBreakdown)

		return mcp.NewToolResultText(builder.String()), nil
	}
}
