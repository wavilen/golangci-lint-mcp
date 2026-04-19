package server

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

const gosecAutofixTimeout = 60 * time.Second

func makeGosecAutofixHandler(opts Options) func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path, err := req.RequireString("path")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("missing required parameter 'path': %v", err)), nil
		}
		path = strings.TrimSpace(path)
		if path == "" {
			return mcp.NewToolResultError("parameter 'path' must not be empty"), nil
		}

		args := []string{"-ai-api-provider=" + opts.GosecAIProvider, "-ai-api-key=" + opts.GosecAIKey}
		if opts.GosecAIBaseURL != "" {
			args = append(args, "-ai-base-url="+opts.GosecAIBaseURL)
		}
		if opts.GosecAISkipSSL {
			args = append(args, "-ai-skip-ssl")
		}
		args = append(args, path)

		ctx, cancel := context.WithTimeout(ctx, gosecAutofixTimeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, "gosec", args...)
		out, runErr := cmd.CombinedOutput()

		output := string(out)
		// Belt-and-suspenders: strip the API key from any output
		if opts.GosecAIKey != "" {
			output = strings.ReplaceAll(output, opts.GosecAIKey, "***")
		}

		if runErr != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return mcp.NewToolResultError(
					"gosec AI autofix timed out after " + gosecAutofixTimeout.String() +
						". Fall back to the guide's <instructions> and <examples> sections."), nil
			}
			return mcp.NewToolResultError(
				"gosec AI autofix failed: " + sanitizeOutput(output, runErr.Error())), nil
		}

		return mcp.NewToolResultText(output), nil
	}
}

func sanitizeOutput(output string, errMsg string) string {
	return output + ": " + errMsg
}
