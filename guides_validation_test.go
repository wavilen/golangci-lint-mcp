package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

func TestGuideStructure(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	names := store.LinterNames()
	require.NotEmpty(t, names, "expected at least one linter guide")

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			g, ok := store.Lookup(name, "")
			if !ok {
				rules := store.ListRules(name)
				if len(rules) > 0 {
					return
				}
				t.Fatalf("linter %q found in LinterNames but not in Lookup", name)
			}

			assert.NotEmpty(t, g.Instructions,
				"guide %q must have <instructions> tag", name)
			assert.NotEmpty(t, g.Examples,
				"guide %q must have <examples> tag", name)
			assert.True(t, g.Instructions != "" || g.Examples != "" || g.Patterns != "",
				"guide %q must have at least one recognized tag", name)
		})
	}
}

func TestWordLimits(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	for _, name := range store.LinterNames() {
		g, ok := store.Lookup(name, "")
		if !ok {
			continue
		}

		t.Run(name, func(t *testing.T) {
			var total int
			for _, s := range []string{g.Instructions, g.Examples, g.Patterns} {
				total += len(strings.Fields(s))
			}
			assert.LessOrEqual(t, total, 200,
				"simple linter guide %q must be ≤ 200 words (got %d)", name, total)
		})
	}
}

func TestCompoundWordLimits(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	compoundLinters := []string{"gocritic", "staticcheck", "revive", "gosec", "govet", "modernize", "testifylint", "ginkgolinter", "errorlint", "grouper"}
	for _, linter := range compoundLinters {
		rules := store.ListRules(linter)
		for _, rule := range rules {
			t.Run(linter+"/"+rule, func(t *testing.T) {
				g, ok := store.Lookup(linter, rule)
				require.True(t, ok, "rule %s/%s not found", linter, rule)

				var total int
				for _, s := range []string{g.Instructions, g.Examples, g.Patterns} {
					total += len(strings.Fields(s))
				}
				assert.LessOrEqual(t, total, 500,
					"compound linter guide %s/%s must be ≤ 500 words (got %d)", linter, rule, total)
			})
		}
	}
}

func TestGocriticCheckerCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("gocritic")
	assert.Len(t, rules, 108,
		"expected exactly 108 gocritic checker guides, got %d", len(rules))
}

func TestReviveRuleCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("revive")
	expectedCount := 101 // 59 A-M + 42 N-Z
	assert.Len(t, rules, expectedCount,
		"expected %d revive rule guides, got %d", expectedCount, len(rules))
}

func TestStaticcheckCheckCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("staticcheck")
	expectedCount := 172 // SA:97 + S:39 + ST:24 + QF:12
	assert.Len(t, rules, expectedCount,
		"expected %d staticcheck check guides, got %d", expectedCount, len(rules))
}

func TestGosecRuleCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("gosec")
	expectedCount := 61 // G1xx:24 + G2xx:4 + G3xx:7 + G4xx:8 + G5xx:7 + G6xx:2 + G7xx:9
	assert.Len(t, rules, expectedCount,
		"expected %d gosec rule guides, got %d", expectedCount, len(rules))
}

func TestGovetAnalyzerCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("govet")
	expectedCount := 35
	assert.Len(t, rules, expectedCount,
		"expected %d govet analyzer guides, got %d", expectedCount, len(rules))
}

func TestModernizeAnalyzerCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("modernize")
	expectedCount := 10
	assert.Len(t, rules, expectedCount,
		"expected %d modernize analyzer guides, got %d", expectedCount, len(rules))
}

func TestTestifylintCheckerCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("testifylint")
	expectedCount := 20
	assert.Len(t, rules, expectedCount,
		"expected %d testifylint checker guides, got %d", expectedCount, len(rules))
}

func TestGinkgolinterCheckCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("ginkgolinter")
	expectedCount := 12
	assert.Len(t, rules, expectedCount,
		"expected %d ginkgolinter check guides, got %d", expectedCount, len(rules))
}

func TestErrorlintCheckCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("errorlint")
	expectedCount := 3
	assert.Len(t, rules, expectedCount,
		"expected %d errorlint check guides, got %d", expectedCount, len(rules))
}

func TestGrouperGroupCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	rules := store.ListRules("grouper")
	expectedCount := 4
	assert.Len(t, rules, expectedCount,
		"expected %d grouper group guides, got %d", expectedCount, len(rules))
}

func TestTemplateNotLoaded(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	_, found := store.Lookup("_template", "")
	assert.False(t, found, "_template.md should not be loaded as a guide")
}

func TestSimpleLinterCount(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	count := 0
	for _, name := range store.LinterNames() {
		_, ok := store.Lookup(name, "")
		if ok && len(store.ListRules(name)) == 0 {
			count++
		}
	}

	expectedCount := 103 // 102 original simple guides + gofmt.md
	assert.Equal(t, expectedCount, count,
		"expected %d simple linter guides, got %d", expectedCount, count)
}
