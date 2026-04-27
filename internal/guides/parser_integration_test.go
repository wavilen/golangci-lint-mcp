package guides

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getProjectRoot(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	require.True(t, ok)
	// filename: .../golangci-lint-mcp/internal/guides/parser_integration_test.go
	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func TestRuntimeRelatedCoverage(t *testing.T) {
	root := getProjectRoot(t)
	fsys := os.DirFS(root)
	store, err := NewStore(fsys)
	require.NoError(t, err, "failed to create store from real guides at %s", root)

	var totalGuides int
	var relatedInBody int
	var relatedPopulated int
	var mismatches []string

	for _, linter := range store.LinterNames() {
		rules := store.ListRules(linter)
		if len(rules) == 0 {
			guide, found := store.Lookup(linter, "")
			if !found {
				continue
			}
			totalGuides++
			checkGuideRelated(t, guide, &relatedInBody, &relatedPopulated, &mismatches)
		} else {
			for _, rule := range rules {
				guide, found := store.Lookup(linter, rule)
				if !found {
					continue
				}
				totalGuides++
				checkGuideRelated(t, guide, &relatedInBody, &relatedPopulated, &mismatches)
			}
		}
	}

	t.Logf("Total guides: %d, Guides with <related> in body: %d, Related populated: %d",
		totalGuides, relatedInBody, relatedPopulated)

	if len(mismatches) > 0 {
		t.Logf("Guides with <related> in body but nil Related field (%d):\n%s",
			len(mismatches), strings.Join(mismatches, "\n"))
	}

	// Informational per D-05: log but do not fail
	assert.Equal(t, totalGuides, relatedPopulated,
		"Expected ALL guides to have Related populated after normalization")
	assert.Equal(t, relatedInBody, relatedPopulated,
		"Every guide with <related> in body should have Related field populated")
}

func checkGuideRelated(t *testing.T, guide *Guide, relatedInBody, relatedPopulated *int, mismatches *[]string) {
	t.Helper()
	hasTagInBody := strings.Contains(guide.RawBody, "<related>")
	if hasTagInBody {
		*relatedInBody++
	}
	if guide.Related != nil {
		*relatedPopulated++
	}
	if hasTagInBody && guide.Related == nil {
		*mismatches = append(*mismatches, guide.Key())
	}
}
