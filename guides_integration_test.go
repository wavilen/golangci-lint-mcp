package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"
)

// TestSimpleLinterIntegration verifies the store returns useful guidance
// for every simple linter relevant to testdata/messy_example.go.
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
			g, ok := store.Lookup(linter, "")
			require.True(t, ok, "linter %q should have a guide", linter)

			assert.NotEmpty(t, g.Instructions,
				"guide %q must have instructions", linter)
			assert.NotEmpty(t, g.Examples,
				"guide %q must have examples", linter)
			assert.True(t, len(g.RawBody) > 100,
				"guide %q must have substantial content beyond headers (got %d bytes)", linter, len(g.RawBody))
		})
	}
}

// TestCrossReferenceValidity checks that all <related> cross-references
// in simple linter guides point to existing linters.
func TestCrossReferenceValidity(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	// Build a set of all known linter names.
	nameSet := make(map[string]bool)
	for _, name := range store.LinterNames() {
		nameSet[name] = true
	}

	// Iterate all simple linters (Lookup succeeds, no sub-rules).
	for _, name := range store.LinterNames() {
		g, ok := store.Lookup(name, "")
		if !ok || len(store.ListRules(name)) > 0 {
			continue
		}

		t.Run(name, func(t *testing.T) {
			related := extractRelatedTag(g.RawBody)
			if related == "" {
				return // no related tag is acceptable
			}

			refs := strings.Split(related, ",")
			for _, ref := range refs {
				ref = strings.TrimSpace(ref)
				if ref == "" {
					continue
				}

				// Check if reference contains a compound linter rule (e.g., "gocritic/appendAssign").
				if strings.Contains(ref, "/") {
					parts := strings.SplitN(ref, "/", 2)
					linterPart := strings.TrimSpace(parts[0])
					rulePart := strings.TrimSpace(parts[1])

					assert.True(t, nameSet[linterPart],
						"guide %q references unknown linter %q in <related>", name, linterPart)
					if nameSet[linterPart] {
						rules := store.ListRules(linterPart)
						found := false
						for _, r := range rules {
							if r == rulePart {
								found = true
								break
							}
						}
						assert.True(t, found,
							"guide %q references unknown rule %q for linter %q in <related>", name, rulePart, linterPart)
					}
				} else {
					assert.True(t, nameSet[ref],
						"guide %q references unknown linter %q in <related>", name, ref)
				}
			}
		})
	}
}

// actionVerbs lists verbs that indicate actionable guidance in guide instructions.
var actionVerbs = []string{
	"use", "replace", "add", "remove", "run", "avoid",
	"refactor", "simplify", "extract", "move", "rename",
	"pass", "handle", "break", "fix", "wrap", "defer",
	"check", "close", "reorder", "redesign", "group",
	"configure", "assign",
}

// TestContentQualitySpotCheck verifies that sampled simple linter guides
// have actionable instructions and Bad/Good code examples.
func TestContentQualitySpotCheck(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	sampleLinters := []string{
		"errcheck", "gofmt", "nakedret", "nestif", "funlen",
		"gocyclo", "dupl", "mnd", "misspell", "goconst",
		"ineffassign", "varnamelen", "wrapcheck", "gofumpt", "bodyclose",
	}

	for _, linter := range sampleLinters {
		t.Run(linter, func(t *testing.T) {
			g, ok := store.Lookup(linter, "")
			require.True(t, ok, "linter %q should have a guide", linter)

			// Check Instructions contains at least one action verb.
			lowerInstructions := strings.ToLower(g.Instructions)
			hasAction := false
			for _, verb := range actionVerbs {
				if strings.Contains(lowerInstructions, verb) {
					hasAction = true
					break
				}
			}
			assert.True(t, hasAction,
				"guide %q instructions should contain at least one action verb", linter)

			// Check Examples contains Bad and Good section markers.
			assert.Contains(t, g.Examples, "Bad",
				"guide %q examples should have a Bad section", linter)
			assert.Contains(t, g.Examples, "Good",
				"guide %q examples should have a Good section", linter)

			// Check overall content is meaningful.
			allContent := g.Instructions + " " + g.Examples + " " + g.Patterns
			wordCount := len(strings.Fields(allContent))
			assert.True(t, wordCount > 20,
				"guide %q must have meaningful content (got %d words), not a stub", linter, wordCount)
		})
	}
}

// extractRelatedTag extracts content between <related>...</related> tags.
func extractRelatedTag(content string) string {
	open := "<related>"
	close := "</related>"

	startIdx := strings.Index(content, open)
	if startIdx == -1 {
		return ""
	}
	startIdx += len(open)

	endIdx := strings.Index(content[startIdx:], close)
	if endIdx == -1 {
		return ""
	}

	return strings.TrimSpace(content[startIdx : startIdx+endIdx])
}

// TestCompoundLinterIntegration verifies all 10 compound linters have the
// correct number of indexed rules and each guide has meaningful content.
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
				g, ok := store.Lookup(linter, rule)
				require.True(t, ok, "%s/%s should be found via Lookup", linter, rule)
				assert.NotEmpty(t, g.Instructions,
					"%s/%s must have instructions", linter, rule)
				assert.NotEmpty(t, g.Examples,
					"%s/%s must have examples", linter, rule)
				assert.True(t, len(g.RawBody) > 50,
					"%s/%s must have substantial content (got %d bytes)", linter, rule, len(g.RawBody))
			}
		})
	}
}

// TestCompoundCrossReferenceValidity checks that all <related> cross-references
// in compound linter guides (526 total) point to existing linters and rules.
func TestCompoundCrossReferenceValidity(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	// Build a set of all known linter names.
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
			rules := store.ListRules(linter)
			for _, rule := range rules {
				t.Run(rule, func(t *testing.T) {
					g, ok := store.Lookup(linter, rule)
					require.True(t, ok, "guide %s/%s should be found", linter, rule)

					related := extractRelatedTag(g.RawBody)
					if related == "" {
						return // no related tag is acceptable
					}

					refs := strings.Split(related, ",")
					for _, ref := range refs {
						ref = strings.TrimSpace(ref)
						if ref == "" {
							continue
						}

						if strings.Contains(ref, "/") {
							// Cross-linter reference: "gocritic/appendAssign"
							parts := strings.SplitN(ref, "/", 2)
							linterPart := strings.TrimSpace(parts[0])
							rulePart := strings.TrimSpace(parts[1])

							assert.True(t, nameSet[linterPart],
								"guide %s/%s references unknown linter %q in <related>", linter, rule, linterPart)
							if nameSet[linterPart] {
								rules := store.ListRules(linterPart)
								found := false
								for _, r := range rules {
									if r == rulePart {
										found = true
										break
									}
								}
								assert.True(t, found,
									"guide %s/%s references unknown rule %q for linter %q in <related>",
									linter, rule, rulePart, linterPart)
							}
						} else if nameSet[ref] {
							// Cross-linter bare reference: "errcheck" (linter name)
						} else {
							// Within-linter reference: bare checker name like "weakCond" in gocritic
							// Verify it exists as a rule of the current linter
							rules := store.ListRules(linter)
							found := false
							for _, r := range rules {
								if r == ref {
									found = true
									break
								}
							}
							assert.True(t, found,
								"guide %s/%s references unknown %q in <related> (not a linter name and not a rule of %s)",
								linter, rule, ref, linter)
						}
					}
				})
			}
		})
	}
}

// TestCompoundContentQualitySpotCheck verifies ~33 sampled compound guides have
// actionable instructions, Bad/Good examples, and substantial content.
func TestCompoundContentQualitySpotCheck(t *testing.T) {
	store, err := guides.NewStore(guideFS)
	require.NoError(t, err)

	sampleRules := []struct {
		linter, rule string
	}{
		// gocritic (5)
		{"gocritic", "appendAssign"}, {"gocritic", "badCond"},
		{"gocritic", "dupImport"}, {"gocritic", "elseif"}, {"gocritic", "truncateCmp"},
		// staticcheck (5)
		{"staticcheck", "SA1019"}, {"staticcheck", "SA4004"},
		{"staticcheck", "S1003"}, {"staticcheck", "QF1001"}, {"staticcheck", "ST1005"},
		// revive (5)
		{"revive", "add-constant"}, {"revive", "cognitive-complexity"},
		{"revive", "exported"}, {"revive", "unreachable-code"}, {"revive", "unused-parameter"},
		// gosec (5)
		{"gosec", "G101"}, {"gosec", "G201"}, {"gosec", "G301"}, {"gosec", "G401"}, {"gosec", "G501"},
		// govet (5)
		{"govet", "atomic"}, {"govet", "copylocks"},
		{"govet", "loopclosure"}, {"govet", "printf"}, {"govet", "structtag"},
		// modernize (2)
		{"modernize", "errorf"}, {"modernize", "simplifyrange"},
		// testifylint (2)
		{"testifylint", "bool-compare"}, {"testifylint", "formatter"},
		// ginkgolinter (2)
		{"ginkgolinter", "async-assertion"}, {"ginkgolinter", "focus-container"},
		// errorlint (1)
		{"errorlint", "asserts"},
		// grouper (1)
		{"grouper", "const"},
	}

	for _, sample := range sampleRules {
		t.Run(sample.linter+"/"+sample.rule, func(t *testing.T) {
			g, ok := store.Lookup(sample.linter, sample.rule)
			require.True(t, ok, "%s/%s should have a guide", sample.linter, sample.rule)

			// Check Instructions contains at least one action verb.
			lowerInstructions := strings.ToLower(g.Instructions)
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

			// Check Examples contains Bad and Good section markers.
			assert.Contains(t, g.Examples, "Bad",
				"%s/%s examples should have a Bad section", sample.linter, sample.rule)
			assert.Contains(t, g.Examples, "Good",
				"%s/%s examples should have a Good section", sample.linter, sample.rule)

			// Check overall content is meaningful.
			allContent := g.Instructions + " " + g.Examples + " " + g.Patterns
			wordCount := len(strings.Fields(allContent))
			assert.True(t, wordCount > 20,
				"%s/%s must have meaningful content (got %d words), not a stub",
				sample.linter, sample.rule, wordCount)
		})
	}
}
