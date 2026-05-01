# godot

<instructions>
Checks that doc comments end with a period as required by Go documentation conventions. godoc renders comments as documentation pages on pkg.go.dev and IDE tooltips — missing punctuation produces fragmented, hard-to-read sentences. Add a period at the end of each comment sentence; for multi-sentence comments, ensure every sentence is punctuated.
</instructions>

<examples>
## Good
```go
// Parse reads the input and returns a structured result.
func Parse(input string) (*Result, error) {
```
</examples>

<patterns>
- Add a period to the end of `//` doc comments on exported functions, types, and constants — required for `godoc` sentence parsing
- Ensure every sentence in multi-line doc comments ends with `.` — `godoc` treats periods as sentence boundaries for formatting
- End package-level `// Package foo ...` comments with proper sentence punctuation
- Annotate TODO/FIXME with trailing periods unless excluded by configuration
</patterns>

<related>
dupword, godoclint, godox
</related>
