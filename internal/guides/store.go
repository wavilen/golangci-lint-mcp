package guides

import (
	"fmt"
	"io/fs"
	"strings"
	"sync"
)

// Store holds an in-memory index of all guides loaded from an fs.FS.
type Store struct {
	mu        sync.RWMutex
	guides    map[string]*Guide // key: "linter" or "linter/rule"
	linterSet map[string]bool   // all known linter names (for suggestions)
}

// NewStore loads all guides from the provided filesystem and builds the lookup index.
// The filesystem should contain a "guides" directory with .md files.
func NewStore(fsys fs.FS) (*Store, error) {
	s := &Store{
		guides:    make(map[string]*Guide),
		linterSet: make(map[string]bool),
	}
	if err := s.loadAll(fsys); err != nil {
		return nil, fmt.Errorf("loading guides: %w", err)
	}
	return s, nil
}

// Lookup finds a guide by linter name and optional rule.
// Returns (nil, false) if not found.
func (s *Store) Lookup(linter, rule string) (*Guide, bool) {
	key := linter
	if rule != "" {
		key = linter + "/" + rule
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	g, ok := s.guides[key]
	if !ok {
		return nil, false
	}
	return g, true
}

// ListRules returns all available rule IDs for a compound linter.
// Returns empty slice for simple linters or unknown linters.
func (s *Store) ListRules(linter string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var rules []string
	prefix := linter + "/"
	for key := range s.guides {
		if strings.HasPrefix(key, prefix) {
			rules = append(rules, strings.TrimPrefix(key, prefix))
		}
	}
	return rules
}

// LinterNames returns all known linter names.
func (s *Store) LinterNames() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	names := make([]string, 0, len(s.linterSet))
	for name := range s.linterSet {
		names = append(names, name)
	}
	return names
}

// Suggest returns the closest matching linter name for a misspelled input.
// Uses Levenshtein distance to find the best match (D-03).
func (s *Store) Suggest(input string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	input = strings.ToLower(input)
	best := ""
	bestDist := len(input) + 1 // max possible distance + 1
	for name := range s.linterSet {
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
func (s *Store) loadAll(fsys fs.FS) error {
	return fs.WalkDir(fsys, "guides", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") || strings.HasPrefix(d.Name(), "_") {
			return nil
		}
		content, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		// path is "guides/errcheck.md" or "guides/gocritic/badcall.md"
		relPath := strings.TrimPrefix(path, "guides/")
		guide, err := Parse(relPath, content)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}
		s.guides[guide.Key()] = guide
		s.linterSet[guide.Linter] = true
		return nil
	})
}

// levenshtein computes the edit distance between two strings using Wagner-Fischer algorithm.
func levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	// Use two rows instead of full matrix for space efficiency.
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)

	for j := 0; j <= lb; j++ {
		prev[j] = j
	}

	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			curr[j] = min(
				prev[j]+1,      // deletion
				curr[j-1]+1,    // insertion
				prev[j-1]+cost, // substitution
			)
		}
		prev, curr = curr, prev
	}

	return prev[lb]
}
