package guides

import (
	"fmt"
	"strings"
)

// recognizedTags lists the XML tags that the parser extracts from guide markdown.
var recognizedTags = []string{"instructions", "examples", "patterns"}

// Parse extracts XML-tagged sections from a markdown guide file.
// filename is a relative path like "errcheck.md" or "gocritic/badcall.md".
// content is the raw markdown bytes.
func Parse(filename string, content []byte) (*Guide, error) {
	raw := string(content)

	// Derive linter and rule from filename path.
	// "errcheck.md" → linter="errcheck", rule=""
	// "gocritic/badcall.md" → linter="gocritic", rule="badcall"
	linter, rule := parseFilename(filename)

	g := &Guide{
		Linter:  linter,
		Rule:    rule,
		RawBody: raw,
	}

	// Extract each recognized tag's content.
	found := false
	for _, tag := range recognizedTags {
		inner := extractTag(raw, tag)
		if inner != "" {
			found = true
		}
		switch tag {
		case "instructions":
			g.Instructions = inner
		case "examples":
			g.Examples = inner
		case "patterns":
			g.Patterns = inner
		}
	}

	if !found {
		return nil, fmt.Errorf("guide %q must contain at least one recognized XML tag (<instructions>, <examples>, or <patterns>)", filename)
	}

	return g, nil
}

// parseFilename extracts linter name and optional rule from a relative file path.
func parseFilename(filename string) (linter, rule string) {
	// Remove .md extension
	base := strings.TrimSuffix(filename, ".md")

	// Split on directory separator
	parts := strings.SplitN(base, "/", 2)
	if len(parts) == 1 {
		// Simple linter: "errcheck.md" → linter="errcheck"
		return parts[0], ""
	}
	// Compound linter: "gocritic/badcall.md" → linter="gocritic", rule="badcall"
	return parts[0], parts[1]
}

// extractTag returns the content between <tag>...</tag>, or empty string if not found.
func extractTag(content, tag string) string {
	open := "<" + tag + ">"
	close := "</" + tag + ">"

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
