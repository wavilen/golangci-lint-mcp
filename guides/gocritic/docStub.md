# gocritic: docStub

<instructions>
Detects documentation comments that are stub placeholders — they merely restate the function or type name without adding useful information. Examples include `// Foo ...`, `// Foo is a Foo`, or `//nolint` used as the sole doc comment.

Write documentation that describes the purpose, behavior, or usage of the symbol. If nothing meaningful can be said, omit the comment entirely.
</instructions>

<examples>
## Good
```go
// Handler processes incoming HTTP requests and routes them
// to the appropriate service endpoint.
func Handler(_ http.ResponseWriter, _ *http.Request) {}

// Config holds application-wide configuration values
// loaded from environment variables and config files.
type Config struct{}
```
</examples>

<patterns>
- Replace `// Name ...` placeholder ellipsis with actual documentation
- Replace `// Name is a Name` tautologies with meaningful descriptions
- Add text to empty doc comments — `// ` with no content provides no value
- Add proper doc comments instead of `//nolint` as the only comment on exported symbols
</patterns>

<related>
gocritic/commentFormatting, gocritic/captLocal
</related>
