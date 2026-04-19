package main

import (
	"errors"
	"testing"
	"testing/fstest"

	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun_Success(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test instructions</instructions>"),
		},
	}

	env := func(key string) string { return "" }

	err := run(testFS, false, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.NoError(t, err)
}

func TestRun_GuideLoadError(t *testing.T) {
	emptyFS := fstest.MapFS{}

	env := func(key string) string { return "" }

	err := run(emptyFS, false, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error loading guides")
}

func TestRun_ServeError(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test instructions</instructions>"),
		},
	}

	env := func(key string) string { return "" }

	err := run(testFS, false, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return errors.New("transport failure")
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "server error")
}

func TestRun_GosecAIWithKey(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test</instructions>"),
		},
	}

	env := func(key string) string {
		if key == "GOSEC_AI_API_KEY" {
			return "test-key"
		}
		return ""
	}

	err := run(testFS, true, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.NoError(t, err)
}

func TestRun_GosecAIWithoutKey(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test</instructions>"),
		},
	}

	env := func(key string) string { return "" }

	err := run(testFS, true, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.NoError(t, err)
}

func TestRun_SkipSSL(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test</instructions>"),
		},
	}

	env := func(key string) string {
		if key == "GOSEC_AI_SKIP_SSL" {
			return "true"
		}
		if key == "GOSEC_AI_API_KEY" {
			return "key"
		}
		return ""
	}

	err := run(testFS, true, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.NoError(t, err)
}

func TestRun_EnvVarsPropagated(t *testing.T) {
	testFS := fstest.MapFS{
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Test</instructions>"),
		},
	}

	env := func(key string) string {
		switch key {
		case "GOSEC_AI_API_PROVIDER":
			return "gemini"
		case "GOSEC_AI_API_KEY":
			return "mykey"
		case "GOSEC_AI_BASE_URL":
			return "https://custom.api.com"
		default:
			return ""
		}
	}

	err := run(testFS, true, env, func(srv *mcpserver.MCPServer, opts ...mcpserver.StdioOption) error {
		return nil
	})
	require.NoError(t, err)
}
