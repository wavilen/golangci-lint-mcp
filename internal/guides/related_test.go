package guides

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]bool
	}{
		{
			name:     "simple words",
			input:    "Hello, World! 123",
			expected: map[string]bool{"hello": true, "world": true, "123": true},
		},
		{
			name:     "skips single character tokens",
			input:    "a I the big",
			expected: map[string]bool{"the": true, "big": true},
		},
		{
			name:     "empty string",
			input:    "",
			expected: map[string]bool{},
		},
		{
			name:     "only single chars",
			input:    "a b c",
			expected: map[string]bool{},
		},
		{
			name:     "mixed case normalized",
			input:    "Go GO go",
			expected: map[string]bool{"go": true},
		},
		{
			name:     "special characters stripped",
			input:    "err-check (error) return!",
			expected: map[string]bool{"err": true, "check": true, "error": true, "return": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordSet(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestKeywordOverlapScore(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		bullet   string
		minScore int // score should be >= minScore
		maxScore int // score should be <= maxScore
	}{
		{
			name:     "high overlap between similar text",
			text:     "error return value not checked",
			bullet:   "Always check error return values",
			minScore: 2,
			maxScore: 10,
		},
		{
			name:     "low overlap between unrelated text",
			text:     "sql injection format",
			bullet:   "Validate user-supplied command arguments",
			minScore: 0,
			maxScore: 0,
		},
		{
			name:     "exact match",
			text:     "check error",
			bullet:   "check error",
			minScore: 2,
			maxScore: 2,
		},
		{
			name:     "empty text",
			text:     "",
			bullet:   "check error returns",
			minScore: 0,
			maxScore: 0,
		},
		{
			name:     "empty bullet",
			text:     "check error returns",
			bullet:   "",
			minScore: 0,
			maxScore: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := KeywordOverlapScore(tt.text, tt.bullet)
			assert.GreaterOrEqual(t, score, tt.minScore, "score should be >= %d but got %d", tt.minScore, score)
			assert.LessOrEqual(t, score, tt.maxScore, "score should be <= %d but got %d", tt.maxScore, score)
		})
	}
}

func TestExtractPatternBullets(t *testing.T) {
	tests := []struct {
		name     string
		patterns string
		expected []string
	}{
		{
			name:     "extracts bullets",
			patterns: "- bullet one\n- bullet two\nother text",
			expected: []string{"bullet one", "bullet two"},
		},
		{
			name:     "no bullets returns nil",
			patterns: "just regular text\nno bullets here",
			expected: nil,
		},
		{
			name:     "empty string",
			patterns: "",
			expected: nil,
		},
		{
			name:     "mixed bullets and text",
			patterns: "Some intro text\n- first bullet\nmore text\n- second bullet",
			expected: []string{"first bullet", "second bullet"},
		},
		{
			name:     "single bullet",
			patterns: "- only one",
			expected: []string{"only one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPatternBullets(tt.patterns)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBestPatternBullet(t *testing.T) {
	tests := []struct {
		name         string
		patterns     string
		keywordSrc   string
		expected     string
		expectPrefix string // if non-empty, result must start with this
	}{
		{
			name:         "selects highest keyword overlap",
			patterns:     "- check error returns from file ops\n- validate user input\n- use context properly",
			keywordSrc:   "error handling for file operations",
			expected:     "check error returns from file ops",
			expectPrefix: "",
		},
		{
			name:         "falls back to first bullet when all scores zero",
			patterns:     "- bullet about foo\n- another about bar",
			keywordSrc:   "completely unrelated text about sql",
			expected:     "bullet about foo",
			expectPrefix: "",
		},
		{
			name:         "empty patterns returns empty string",
			patterns:     "",
			keywordSrc:   "anything",
			expected:     "",
			expectPrefix: "",
		},
		{
			name:         "no bullets returns empty string",
			patterns:     "no bullets here",
			keywordSrc:   "anything",
			expected:     "",
			expectPrefix: "",
		},
		{
			name:         "single bullet returned",
			patterns:     "- the only bullet",
			keywordSrc:   "anything",
			expected:     "the only bullet",
			expectPrefix: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BestPatternBullet(tt.patterns, tt.keywordSrc)
			if tt.expected != "" {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
