package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

func TestSimpleLinterIntegration(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	messyLinters := []string{
		"errcheck", "nakedret", "nestif", "mnd", "varnamelen",
		"funlen", "gocyclo", "goconst", "ineffassign", "dupl",
		"misspell", "gofmt",
	}

	for _, linter := range messyLinters {
		t.Run(linter, func(t *testing.T) {
			guide, ok := store.Lookup(linter, "")
			require.True(t, ok, "linter %q should have a guide", linter)

			assert.NotEmpty(t, guide.Instructions,
				"guide %q must have instructions", linter)
			assert.NotEmpty(t, guide.Examples,
				"guide %q must have examples", linter)
			assert.Greater(t, len(guide.RawBody), 100,
				"guide %q must have substantial content beyond headers (got %d bytes)", linter, len(guide.RawBody))
		})
	}
}

func TestCrossReferenceValidity(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	nameSet := make(map[string]bool)
	for _, name := range store.LinterNames() {
		nameSet[name] = true
	}

	for _, name := range store.LinterNames() {
		guide, ok := store.Lookup(name, "")
		if !ok || len(store.ListRules(name)) > 0 {
			continue
		}

		t.Run(name, func(t *testing.T) {
			related := extractRelatedTag(guide.RawBody)
			if related == "" {
				return
			}

			refs := strings.Split(related, ",")
			for _, ref := range refs {
				assertRefValid(t, name, ref, nameSet, store)
			}
		})
	}
}

func TestContentQualitySpotCheck(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	sampleLinters := []string{
		"errcheck", "gofmt", "nakedret", "nestif", "funlen",
		"gocyclo", "dupl", "mnd", "misspell", "goconst",
		"ineffassign", "varnamelen", "wrapcheck", "gofumpt", "bodyclose",
	}

	actionVerbs := []string{
		"use", "replace", "add", "remove", "run", "avoid",
		"refactor", "simplify", "extract", "move", "rename",
		"pass", "handle", "break", "fix", "wrap", "defer",
		"check", "close", "reorder", "redesign", "group",
		"configure", "assign",
	}

	for _, linter := range sampleLinters {
		t.Run(linter, func(t *testing.T) {
			guide, ok := store.Lookup(linter, "")
			require.True(t, ok, "linter %q should have a guide", linter)

			lowerInstructions := strings.ToLower(guide.Instructions)
			hasAction := false
			for _, verb := range actionVerbs {
				if strings.Contains(lowerInstructions, verb) {
					hasAction = true
					break
				}
			}
			assert.True(t, hasAction,
				"guide %q instructions should contain at least one action verb", linter)

			assert.Contains(t, guide.Examples, "Good",
				"guide %q examples should have a Good section", linter)

			allContent := guide.Instructions + " " + guide.Examples + " " + guide.Patterns
			wordCount := len(strings.Fields(allContent))
			assert.Greater(t, wordCount, 20,
				"guide %q must have meaningful content (got %d words), not a stub", linter, wordCount)
		})
	}
}

func assertRefValid(t *testing.T, name string, ref string, nameSet map[string]bool, store *guides.Store) {
	t.Helper()
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return
	}

	if strings.Contains(ref, "/") {
		parts := strings.SplitN(ref, "/", 2)
		linterPart := strings.TrimSpace(parts[0])
		rulePart := strings.TrimSpace(parts[1])

		assert.True(t, nameSet[linterPart],
			"guide %q references unknown linter %q in <related>", name, linterPart)
		if nameSet[linterPart] {
			rules := store.ListRules(linterPart)
			assert.True(t,
				slices.Contains(rules, rulePart),
				"guide %q references unknown rule %q for linter %q in <related>",
				name, rulePart, linterPart)
		}
	} else {
		assert.True(t, nameSet[ref],
			"guide %q references unknown linter %q in <related>", name, ref)
	}
}

func assertCompoundRefValid(
	t *testing.T,
	linter string,
	rule string,
	ref string,
	nameSet map[string]bool,
	store *guides.Store,
) {
	t.Helper()
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return
	}

	switch {
	case strings.Contains(ref, "/"):
		parts := strings.SplitN(ref, "/", 2)
		linterPart := strings.TrimSpace(parts[0])
		rulePart := strings.TrimSpace(parts[1])

		assert.True(t, nameSet[linterPart],
			"guide %s/%s references unknown linter %q in <related>", linter, rule, linterPart)
		if nameSet[linterPart] {
			crossRules := store.ListRules(linterPart)
			assert.True(t, slices.Contains(crossRules, rulePart),
				"guide %s/%s references unknown rule %q for linter %q in <related>",
				linter, rule, rulePart, linterPart)
		}
	case nameSet[ref]:
		// Cross-linter bare reference.
	default:
		withinRules := store.ListRules(linter)
		assert.True(
			t,
			slices.Contains(withinRules, ref),
			"guide %s/%s references unknown %q in <related> (not a linter name and not a rule of %s)",
			linter,
			rule,
			ref,
			linter,
		)
	}
}

func extractRelatedTag(content string) string {
	open := "<related>"
	closeTag := "</related>"

	startIdx := strings.Index(content, open)
	if startIdx == -1 {
		return ""
	}
	startIdx += len(open)

	endIdx := strings.Index(content[startIdx:], closeTag)
	if endIdx == -1 {
		return ""
	}

	return strings.TrimSpace(content[startIdx : startIdx+endIdx])
}

func TestCompoundLinterIntegration(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	compoundLinters := map[string]int{
		"gocritic":     108,
		"staticcheck":  172,
		"revive":       101,
		"gosec":        61,
		"govet":        35,
		"modernize":    10,
		"testifylint":  20,
		"ginkgolinter": 12,
		"errorlint":    3,
		"grouper":      4,
	}

	for linter, expectedCount := range compoundLinters {
		t.Run(linter, func(t *testing.T) {
			rules := store.ListRules(linter)
			assert.NotEmpty(t, rules, "%s should have indexed rules", linter)
			assert.Len(t, rules, expectedCount,
				"%s should have exactly %d rules, got %d", linter, expectedCount, len(rules))

			if len(rules) > 0 {
				rule := rules[0]
				guide, ok := store.Lookup(linter, rule)
				require.True(t, ok, "%s/%s should be found via Lookup", linter, rule)
				assert.NotEmpty(t, guide.Instructions,
					"%s/%s must have instructions", linter, rule)
				assert.NotEmpty(t, guide.Examples,
					"%s/%s must have examples", linter, rule)
				assert.Greater(t, len(guide.RawBody), 50,
					"%s/%s must have substantial content (got %d bytes)", linter, rule, len(guide.RawBody))
			}
		})
	}
}

func TestCompoundCrossReferenceValidity(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	nameSet := make(map[string]bool)
	for _, name := range store.LinterNames() {
		nameSet[name] = true
	}

	compoundLinters := []string{
		"gocritic", "staticcheck", "revive", "gosec", "govet",
		"modernize", "testifylint", "ginkgolinter", "errorlint", "grouper",
	}

	for _, linter := range compoundLinters {
		t.Run(linter, func(t *testing.T) {
			linterRules := store.ListRules(linter)
			for _, rule := range linterRules {
				t.Run(rule, func(t *testing.T) {
					guide, ok := store.Lookup(linter, rule)
					require.True(t, ok, "guide %s/%s should be found", linter, rule)

					related := extractRelatedTag(guide.RawBody)
					if related == "" {
						return
					}

					refs := strings.Split(related, ",")
					for _, ref := range refs {
						assertCompoundRefValid(t, linter, rule, ref, nameSet, store)
					}
				})
			}
		})
	}
}

func TestCompoundContentQualitySpotCheck(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	sampleRules := []struct {
		linter, rule string
	}{
		{"gocritic", "appendAssign"}, {"gocritic", "badCond"},
		{"gocritic", "dupImport"}, {"gocritic", "elseif"}, {"gocritic", "truncateCmp"},
		{"staticcheck", "SA1019"}, {"staticcheck", "SA4004"},
		{"staticcheck", "S1003"}, {"staticcheck", "QF1001"}, {"staticcheck", "ST1005"},
		{"revive", "add-constant"}, {"revive", "cognitive-complexity"},
		{"revive", "exported"}, {"revive", "unreachable-code"}, {"revive", "unused-parameter"},
		{"gosec", "G101"}, {"gosec", "G201"}, {"gosec", "G301"}, {"gosec", "G401"}, {"gosec", "G501"},
		{"govet", "atomic"}, {"govet", "copylocks"},
		{"govet", "loopclosure"}, {"govet", "printf"}, {"govet", "structtag"},
		{"modernize", "errorf"}, {"modernize", "simplifyrange"},
		{"testifylint", "bool-compare"}, {"testifylint", "formatter"},
		{"ginkgolinter", "async-assertion"}, {"ginkgolinter", "focus-container"},
		{"errorlint", "asserts"},
		{"grouper", "const"},
	}

	actionVerbs := []string{
		"use", "replace", "add", "remove", "run", "avoid",
		"refactor", "simplify", "extract", "move", "rename",
		"pass", "handle", "break", "fix", "wrap", "defer",
		"check", "close", "reorder", "redesign", "group",
		"configure", "assign",
	}

	for _, sample := range sampleRules {
		t.Run(sample.linter+"/"+sample.rule, func(t *testing.T) {
			guide, ok := store.Lookup(sample.linter, sample.rule)
			require.True(t, ok, "%s/%s should have a guide", sample.linter, sample.rule)

			lowerInstructions := strings.ToLower(guide.Instructions)
			hasAction := false
			for _, verb := range actionVerbs {
				if strings.Contains(lowerInstructions, verb) {
					hasAction = true
					break
				}
			}
			assert.True(t, hasAction,
				"%s/%s instructions should contain at least one action verb",
				sample.linter, sample.rule)

			assert.Contains(t, guide.Examples, "Good",
				"%s/%s examples should have a Good section", sample.linter, sample.rule)

			allContent := guide.Instructions + " " + guide.Examples + " " + guide.Patterns
			wordCount := len(strings.Fields(allContent))
			assert.Greater(t, wordCount, 20,
				"%s/%s must have meaningful content (got %d words), not a stub",
				sample.linter, sample.rule, wordCount)
		})
	}
}
