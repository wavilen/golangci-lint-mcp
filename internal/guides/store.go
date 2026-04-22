package guides

import (
	"fmt"
	"io/fs"
	"strings"
	"sync"
)

// Store holds an in-memory index of all guides loaded from an [fs.FS].
type Store struct {
	mu        sync.RWMutex
	guides    map[string]*Guide // key: "linter" or "linter/rule"
	linterSet map[string]bool   // all known linter names (for suggestions)
}

// NewStore loads all guides from the provided [fs.FS] filesystem and builds the lookup index.
// The filesystem should contain a "guides" directory with .md files.
func NewStore(fsys fs.FS) (*Store, error) {
	store := &Store{
		mu:        sync.RWMutex{},
		guides:    make(map[string]*Guide),
		linterSet: make(map[string]bool),
	}
	err := store.loadAll(fsys)
	if err != nil {
		return nil, fmt.Errorf("loading guides: %w", err)
	}
	return store, nil
}

// Lookup finds a guide by linter name and optional rule.
// Returns (nil, false) if not found.
func (store *Store) Lookup(linter, rule string) (*Guide, bool) {
	key := linter
	if rule != "" {
		key = linter + "/" + rule
	}
	store.mu.RLock()
	defer store.mu.RUnlock()
	guide, ok := store.guides[key]
	if !ok {
		return nil, false
	}
	return guide, true
}

// ListRules returns all available rule IDs for a compound linter.
// Returns empty slice for simple linters or unknown linters.
func (store *Store) ListRules(linter string) []string {
	store.mu.RLock()
	defer store.mu.RUnlock()
	var rules []string
	prefix := linter + "/"
	for key := range store.guides {
		if trimmed, ok := strings.CutPrefix(key, prefix); ok {
			rules = append(rules, trimmed)
		}
	}
	return rules
}

// LinterNames returns all known linter names.
func (store *Store) LinterNames() []string {
	store.mu.RLock()
	defer store.mu.RUnlock()
	names := make([]string, 0, len(store.linterSet))
	for name := range store.linterSet {
		names = append(names, name)
	}
	return names
}

// Suggest returns the closest matching linter name for a misspelled input.
// Uses Levenshtein distance to find the best match (D-03).
func (store *Store) Suggest(input string) string {
	store.mu.RLock()
	defer store.mu.RUnlock()
	input = strings.ToLower(input)
	best := ""
	bestDist := len(input) + 1
	for name := range store.linterSet {
		d := levenshtein(input, strings.ToLower(name))
		if d < bestDist {
			bestDist = d
			best = name
		}
	}
	// Only suggest if distance is reasonable (<= half the input length + 1)
	if bestDist <= len(input)/2+1 {
		return best
	}
	return ""
}

// loadAll walks the filesystem and parses all .md guide files under "guides/".
func (store *Store) loadAll(fsys fs.FS) error {
	walkErr := fs.WalkDir(fsys, "guides", func(path string, d fs.DirEntry, dirErr error) error {
		if dirErr != nil {
			return dirErr
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") || strings.HasPrefix(d.Name(), "_") {
			return nil
		}
		content, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		relPath := strings.TrimPrefix(path, "guides/")
		guide, err := Parse(relPath, content)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}
		store.guides[guide.Key()] = guide
		store.linterSet[guide.Linter] = true
		return nil
	})
	if walkErr != nil {
		return fmt.Errorf("walking guides directory: %w", walkErr)
	}
	return nil
}

// levenshtein computes the edit distance between two strings using Wagner-Fischer algorithm.
func levenshtein(strA, strB string) int {
	lenA, lenB := len(strA), len(strB)
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	prev := make([]int, lenB+1)
	curr := make([]int, lenB+1)

	for col := 0; col <= lenB; col++ {
		prev[col] = col
	}

	for row := 1; row <= lenA; row++ {
		curr[0] = row
		for col := 1; col <= lenB; col++ {
			cost := 1
			if strA[row-1] == strB[col-1] {
				cost = 0
			}
			curr[col] = min(
				prev[col]+1,
				curr[col-1]+1,
				prev[col-1]+cost,
			)
		}
		prev, curr = curr, prev
	}

	return prev[lenB]
}
