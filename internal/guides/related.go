package guides

import (
	"strings"
	"unicode"
)

// wordSet splits a string into words, lowercases them, and filters out
// single-character tokens. Returns a set of unique words.
func wordSet(s string) map[string]bool {
	result := make(map[string]bool)
	var buf strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			if buf.Len() > 1 {
				result[buf.String()] = true
			}
			buf.Reset()
		}
	}
	if buf.Len() > 1 {
		result[buf.String()] = true
	}
	return result
}

// KeywordOverlapScore returns the number of words shared between text and bullet.
func KeywordOverlapScore(text, bullet string) int {
	textWords := wordSet(text)
	bulletWords := wordSet(bullet)
	count := 0
	for w := range textWords {
		if bulletWords[w] {
			count++
		}
	}
	return count
}

// extractPatternBullets splits patterns content on newlines and extracts
// lines starting with "- ". Returns nil if no bullets found.
func extractPatternBullets(patterns string) []string {
	if patterns == "" {
		return nil
	}
	var bullets []string
	for _, line := range strings.Split(patterns, "\n") {
		trimmed := strings.TrimSpace(line)
		if after, ok := strings.CutPrefix(trimmed, "- "); ok {
			bullets = append(bullets, after)
		}
	}
	return bullets
}

// BestPatternBullet selects the pattern bullet with the highest keyword overlap
// score against keywordSource. Falls back to the first bullet when all scores
// are zero. Returns empty string if no bullets found.
func BestPatternBullet(patterns string, keywordSource string) string {
	bullets := extractPatternBullets(patterns)
	if len(bullets) == 0 {
		return ""
	}

	bestIdx := 0
	bestScore := 0
	for i, b := range bullets {
		score := KeywordOverlapScore(keywordSource, b)
		if score > bestScore {
			bestScore = score
			bestIdx = i
		}
	}
	return bullets[bestIdx]
}
