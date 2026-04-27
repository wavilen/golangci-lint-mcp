package guides

import (
	"fmt"
	"strings"
)

// Parse extracts XML-tagged sections from a markdown guide file.
// filename is a relative path like "errcheck.md" or "gocritic/badcall.md".
// content is the raw markdown bytes.
func Parse(filename string, content []byte) (*Guide, error) {
	raw := string(content)

	// Derive linter and rule from filename path.
	// "errcheck.md" → linter="errcheck", rule=""
	// "gocritic/badcall.md" → linter="gocritic", rule="badcall"
	linter, rule := parseFilename(filename)

	guide := &Guide{
		Linter:       linter,
		Rule:         rule,
		RawBody:      raw,
		Instructions: "",
		Examples:     "",
		Patterns:     "",
		Related:      nil,
	}

	recognizedTags := []string{"instructions", "examples", "patterns", "related"}

	// Extract each recognized tag's content.
	for _, tag := range recognizedTags {
		inner := extractTag(raw, tag)
		switch tag {
		case "instructions":
			guide.Instructions = inner
		case "examples":
			guide.Examples = inner
		case "patterns":
			guide.Patterns = inner
		case "related":
			guide.Related = parseRelatedRefs(inner)
		}
	}

	// Check that at least one primary tag (not related) has content.
	if guide.Instructions == "" && guide.Examples == "" && guide.Patterns == "" {
		return nil, fmt.Errorf(
			"guide %q must contain at least one recognized XML tag (<instructions>, <examples>, or <patterns>)",
			filename,
		)
	}

	return guide, nil
}

// parseFilename extracts linter name and optional rule from a relative file path.
func parseFilename(filename string) (string, string) {
	const pathParts = 2

	base := strings.TrimSuffix(filename, ".md")

	parts := strings.SplitN(base, "/", pathParts)
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
	closeTag := "</" + tag + ">"

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

// parseRelatedRefs splits a comma-separated related tag into trimmed references.
// Returns nil if no valid refs are found (empty input, whitespace only).
func parseRelatedRefs(raw string) []string {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	refs := make([]string, 0, len(parts))
	for _, p := range parts {
		ref := strings.TrimSpace(p)
		if ref != "" {
			refs = append(refs, ref)
		}
	}
	if len(refs) == 0 {
		return nil
	}
	return refs
}
