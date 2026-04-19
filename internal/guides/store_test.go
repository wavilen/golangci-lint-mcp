package guides

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeTestFS() fstest.MapFS {
	return fstest.MapFS{
		"guides/testlinter.md": &fstest.MapFile{
			Data: []byte("# testlinter\n\n<instructions>Test instructions for simple linter</instructions>\n\n<examples>Some example code</examples>\n"),
		},
		"guides/testcompound/rule1.md": &fstest.MapFile{
			Data: []byte("# testcompound: rule1\n\n<instructions>Compound rule one</instructions>\n\n<examples>Code here</examples>\n"),
		},
		"guides/testcompound/rule2.md": &fstest.MapFile{
			Data: []byte("# testcompound: rule2\n\n<instructions>Compound rule two</instructions>\n"),
		},
		"guides/another.md": &fstest.MapFile{
			Data: []byte("# another\n\n<instructions>Another linter guide</instructions>\n\n<patterns>Pattern info</patterns>\n"),
		},
	}
}

func TestNewStore(t *testing.T) {
	s, err := NewStore(makeTestFS())
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestLookup(t *testing.T) {
	s, err := NewStore(makeTestFS())
	require.NoError(t, err)

	t.Run("simple linter lookup", func(t *testing.T) {
		g, ok := s.Lookup("testlinter", "")
		require.True(t, ok)
		require.NotNil(t, g)
		assert.Equal(t, "testlinter", g.Linter)
		assert.Equal(t, "", g.Rule)
		assert.Contains(t, g.Instructions, "Test instructions for simple linter")
	})

	t.Run("compound linter rule lookup", func(t *testing.T) {
		g, ok := s.Lookup("testcompound", "rule1")
		require.True(t, ok)
		require.NotNil(t, g)
		assert.Equal(t, "testcompound", g.Linter)
		assert.Equal(t, "rule1", g.Rule)
		assert.Contains(t, g.Instructions, "Compound rule one")
	})

	t.Run("nonexistent linter returns nil false", func(t *testing.T) {
		g, ok := s.Lookup("nonexistent", "")
		assert.Nil(t, g)
		assert.False(t, ok)
	})

	t.Run("nonexistent rule returns nil false", func(t *testing.T) {
		g, ok := s.Lookup("testcompound", "nonexistent")
		assert.Nil(t, g)
		assert.False(t, ok)
	})
}

func TestSuggest(t *testing.T) {
	s, err := NewStore(makeTestFS())
	require.NoError(t, err)

	t.Run("fuzzy match for misspelling", func(t *testing.T) {
		suggestion := s.Suggest("testlintr")
		assert.Equal(t, "testlinter", suggestion)
	})

	t.Run("fuzzy match for partial name", func(t *testing.T) {
		suggestion := s.Suggest("testlnter")
		assert.Equal(t, "testlinter", suggestion)
	})

	t.Run("no match for completely different name", func(t *testing.T) {
		suggestion := s.Suggest("zzzzzzzz")
		assert.Equal(t, "", suggestion)
	})
}

func TestListRules(t *testing.T) {
	s, err := NewStore(makeTestFS())
	require.NoError(t, err)

	t.Run("compound linter returns rules", func(t *testing.T) {
		rules := s.ListRules("testcompound")
		assert.Contains(t, rules, "rule1")
		assert.Contains(t, rules, "rule2")
		assert.Len(t, rules, 2)
	})

	t.Run("simple linter returns empty slice", func(t *testing.T) {
		rules := s.ListRules("testlinter")
		assert.Empty(t, rules)
	})

	t.Run("unknown linter returns empty slice", func(t *testing.T) {
		rules := s.ListRules("nonexistent")
		assert.Empty(t, rules)
	})
}

func TestLinterNames(t *testing.T) {
	s, err := NewStore(makeTestFS())
	require.NoError(t, err)

	names := s.LinterNames()
	assert.Contains(t, names, "testlinter")
	assert.Contains(t, names, "testcompound")
	assert.Contains(t, names, "another")
	assert.Len(t, names, 3)
}
