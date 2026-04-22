package server

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testAutofixOptions(key string) Options {
	return Options{
		GosecAI:         false,
		GosecAIProvider: "test-provider",
		GosecAIKey:      key,
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
}

func testCallToolRequest(args map[string]any) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "",
			Params: mcp.RequestParams{Meta: nil},
		},
		Header: nil,
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: args,
			Meta:      nil,
			Task:      nil,
		},
	}
}

func TestGosecAutofixHandler_MissingPath(t *testing.T) {
	handler := makeGosecAutofixHandler(testAutofixOptions("test-key"))

	result, err := handler(context.Background(), testCallToolRequest(map[string]any{}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "missing required parameter 'path'")
}

func TestGosecAutofixHandler_EmptyPath(t *testing.T) {
	handler := makeGosecAutofixHandler(testAutofixOptions("test-key"))

	result, err := handler(context.Background(), testCallToolRequest(map[string]any{"path": "   "}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "must not be empty")
}

func TestGosecAutofixHandler_GosecNotFound(t *testing.T) {
	handler := makeGosecAutofixHandler(testAutofixOptions("test-key"))

	_, err := exec.LookPath("gosec")
	if err == nil {
		t.Skip("gosec is installed; skipping gosec-not-found test")
	}

	result, err := handler(context.Background(), testCallToolRequest(map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gosec AI autofix failed")
}

func TestGosecAutofixHandler_OutputSanitized(t *testing.T) {
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\necho 'Results using key=SECRET-KEY-12345'"), 0o750)
	require.NoError(t, err)

	origPath := os.Getenv("PATH")
	t.Setenv("PATH", tmpDir+":"+origPath)

	handler := makeGosecAutofixHandler(testAutofixOptions("SECRET-KEY-12345"))

	result, err := handler(context.Background(), testCallToolRequest(map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.False(t, result.IsError, "expected success result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.NotContains(t, text, "SECRET-KEY-12345", "API key should be stripped from output")
	assert.Contains(t, text, "***", "API key should be replaced with ***")
}

func TestGosecAutofixHandler_Timeout(t *testing.T) {
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\nexec sleep 300"), 0o750)
	require.NoError(t, err)

	origPath := os.Getenv("PATH")
	t.Setenv("PATH", tmpDir+":"+origPath)

	handler := makeGosecAutofixHandler(testAutofixOptions("test-key"))

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	result, err := handler(ctx, testCallToolRequest(map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result due to timeout")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "timed out")
	assert.Contains(t, text, "<instructions>")
}

func TestGosecAutofixHandler_Success(t *testing.T) {
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\necho 'AI autofix suggestions generated successfully'"), 0o750)
	require.NoError(t, err)

	origPath := os.Getenv("PATH")
	t.Setenv("PATH", tmpDir+":"+origPath)

	handler := makeGosecAutofixHandler(testAutofixOptions("test-key"))

	result, err := handler(context.Background(), testCallToolRequest(map[string]any{"path": "./..."}))
	require.NoError(t, err)
	require.False(t, result.IsError, "expected success result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "AI autofix suggestions generated successfully")
}
