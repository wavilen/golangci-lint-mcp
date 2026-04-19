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
		"guides/errcheck.md": &fstest.MapFile{
			Data: []byte("# errcheck\n\n<instructions>Errcheck detects unchecked errors</instructions>"),
		},
		"guides/gosec/G101.md": &fstest.MapFile{
			Data: []byte("# G101\n\n<instructions>Detects hardcoded credentials</instructions>"),
		},
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
	srv := NewServer(store, Options{GosecAI: true, GosecAIKey: "test-key"})
	assert.NotNil(t, srv)
}

func TestNewServer_WithGosecAI(t *testing.T) {
	store := setupStore(t)
	srv := NewServer(store, Options{
		GosecAI:         true,
		GosecAIKey:      "test-key",
		GosecAIProvider: "test-provider",
	})
	assert.NotNil(t, srv)
}

func TestGosecAIConfigured_True(t *testing.T) {
	opts := Options{GosecAI: true, GosecAIKey: "some-key"}
	assert.True(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_NoFlag(t *testing.T) {
	opts := Options{GosecAIKey: "some-key"}
	assert.False(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_NoKey(t *testing.T) {
	opts := Options{GosecAI: true}
	assert.False(t, opts.GosecAIConfigured())
}

func TestGosecAIConfigured_False_Neither(t *testing.T) {
	opts := Options{}
	assert.False(t, opts.GosecAIConfigured())
}
