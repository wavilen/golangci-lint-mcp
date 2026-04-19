package guides

// Guide represents a parsed linter guide with XML-tagged sections.
type Guide struct {
	Linter       string // linter name (e.g., "errcheck", "gocritic")
	Rule         string // rule ID for compound linters (e.g., "badcall"), empty for simple linters
	RawBody      string // complete raw markdown content
	Instructions string // content from <instructions> tag
	Examples     string // content from <examples> tag
	Patterns     string // content from <patterns> tag
}

// Key returns the lookup key: "linter" for simple, "linter/rule" for compound.
func (g *Guide) Key() string {
	if g.Rule != "" {
		return g.Linter + "/" + g.Rule
	}
	return g.Linter
}

// IsCompound returns true if this guide is for a specific compound linter rule.
func (g *Guide) IsCompound() bool {
	return g.Rule != ""
}
