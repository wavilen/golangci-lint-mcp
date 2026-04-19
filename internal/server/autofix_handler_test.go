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

func TestGosecAutofixHandler_MissingPath(t *testing.T) {
	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "test-key",
	})

	result, err := handler(context.Background(), mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "missing required parameter 'path'")
}

func TestGosecAutofixHandler_EmptyPath(t *testing.T) {
	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "test-key",
	})

	result, err := handler(context.Background(), mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{"path": "   "},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "must not be empty")
}

func TestGosecAutofixHandler_GosecNotFound(t *testing.T) {
	// Use a PATH that doesn't contain gosec
	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "test-key",
	})

	// With gosec not in PATH, this should return an error
	// (unless gosec happens to be installed on this system, in which case skip)
	if _, err := exec.LookPath("gosec"); err == nil {
		t.Skip("gosec is installed; skipping gosec-not-found test")
	}

	result, err := handler(context.Background(), mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{"path": "./..."},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "gosec AI autofix failed")
}

func TestGosecAutofixHandler_OutputSanitized(t *testing.T) {
	// Create a temp script that echoes the API key
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\necho 'Results using key=SECRET-KEY-12345'"), 0755)
	require.NoError(t, err)

	// Override PATH to use our fake gosec
	origPath := os.Getenv("PATH")
	t.Cleanup(func() { _ = os.Setenv("PATH", origPath) })
	require.NoError(t, os.Setenv("PATH", tmpDir+":"+origPath))

	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "SECRET-KEY-12345",
	})

	result, err := handler(context.Background(), mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{"path": "./..."},
		},
	})
	require.NoError(t, err)
	require.False(t, result.IsError, "expected success result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.NotContains(t, text, "SECRET-KEY-12345", "API key should be stripped from output")
	assert.Contains(t, text, "***", "API key should be replaced with ***")
}

func TestGosecAutofixHandler_Timeout(t *testing.T) {
	// Create a temp script that blocks — use exec so the process can be killed directly
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\nexec sleep 300"), 0755)
	require.NoError(t, err)

	origPath := os.Getenv("PATH")
	t.Cleanup(func() { _ = os.Setenv("PATH", origPath) })
	require.NoError(t, os.Setenv("PATH", tmpDir+":"+origPath))

	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "test-key",
	})

	// Use a context with a very short external timeout — the handler's internal
	// timeout (60s) will be reduced to min(external, 60s) = external.
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	result, err := handler(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{"path": "./..."},
		},
	})
	require.NoError(t, err)
	require.True(t, result.IsError, "expected error result due to timeout")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "timed out")
	assert.Contains(t, text, "<instructions>")
}

func TestGosecAutofixHandler_Success(t *testing.T) {
	// Create a temp script that simulates successful gosec output
	tmpDir := t.TempDir()
	script := filepath.Join(tmpDir, "gosec")
	err := os.WriteFile(script, []byte("#!/bin/sh\necho 'AI autofix suggestions generated successfully'"), 0755)
	require.NoError(t, err)

	origPath := os.Getenv("PATH")
	t.Cleanup(func() { _ = os.Setenv("PATH", origPath) })
	require.NoError(t, os.Setenv("PATH", tmpDir+":"+origPath))

	handler := makeGosecAutofixHandler(Options{
		GosecAIProvider: "test-provider",
		GosecAIKey:      "test-key",
	})

	result, err := handler(context.Background(), mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      "gosec_ai_autofix",
			Arguments: map[string]any{"path": "./..."},
		},
	})
	require.NoError(t, err)
	require.False(t, result.IsError, "expected success result")
	text := result.Content[0].(mcp.TextContent).Text
	assert.Contains(t, text, "AI autofix suggestions generated successfully")
}
