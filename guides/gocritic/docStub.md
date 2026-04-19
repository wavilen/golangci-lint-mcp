# gocritic: docStub

<instructions>
Detects documentation comments that are stub placeholders — they merely restate the function or type name without adding useful information. Examples include `// Foo ...`, `// Foo is a Foo`, or `//nolint` used as the sole doc comment.

Write documentation that describes the purpose, behavior, or usage of the symbol. If nothing meaningful can be said, omit the comment entirely.
</instructions>

<examples>
## Bad
```go
// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {}

// Config is a config.
type Config struct{}
```

## Good
```go
// Handler processes incoming HTTP requests and routes them
// to the appropriate service endpoint.
func Handler(w http.ResponseWriter, r *http.Request) {}

// Config holds application-wide configuration values
// loaded from environment variables and config files.
type Config struct{}
```
</examples>

<patterns>
- `// Name ...` — placeholder ellipsis pattern
- `// Name is a Name` — tautological description
- Empty doc comment: `// ` with no text
- `//nolint` as the only doc comment on an exported symbol
</patterns>

<related>
commentFormatting, captLocal
</related>
