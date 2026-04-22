package server

import (
	"testing"
	"testing/fstest"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupStore(t *testing.T) *guides.Store {
	t.Helper()
	testFS := fstest.MapFS{
		"guides/errcheck.md": testMapFile(
			"# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>",
		),
		"guides/gosec/G101.md": testMapFile("# G101\n\n<instructions>Detects hardcoded credentials</instructions>"),
	}
	store, err := guides.NewStore(testFS)
	require.NoError(t, err)
	return store
}

func TestNewServer_NoOptions(t *testing.T) {
	store := setupStore(t)
	srv := NewServer(store)
	assert.NotNil(t, srv)
}

func TestNewServer_WithOptions(t *testing.T) {
	store := setupStore(t)
	srv := NewServer(store, Options{
		GosecAI:         true,
		GosecAIProvider: "",
		GosecAIKey:      "test-key",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	})
	assert.NotNil(t, srv)
}

func TestNewServer_WithGosecAI(t *testing.T) {
	store := setupStore(t)
	srv := NewServer(store, Options{
		GosecAI:         true,
		GosecAIKey:      "test-key",
		GosecAIProvider: "test-provider",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	})
	assert.NotNil(t, srv)
}

func TestGosecAIConfigured_True(t *testing.T) {
	opts := Options{
		GosecAI:         true,
		GosecAIProvider: "",
		GosecAIKey:      "some-key",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
	assert.True(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_NoFlag(t *testing.T) {
	opts := Options{
		GosecAI:         false,
		GosecAIProvider: "",
		GosecAIKey:      "some-key",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
	assert.False(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_NoKey(t *testing.T) {
	opts := Options{
		GosecAI:         true,
		GosecAIProvider: "",
		GosecAIKey:      "",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
	assert.False(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_Neither(t *testing.T) {
	opts := Options{
		GosecAI:         false,
		GosecAIProvider: "",
		GosecAIKey:      "",
		GosecAIBaseURL:  "",
		GosecAISkipSSL:  false,
	}
	assert.False(t, opts.GosecAIConfigured())
}
