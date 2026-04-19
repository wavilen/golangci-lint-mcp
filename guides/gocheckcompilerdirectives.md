# gocheckcompilerdirectives

<instructions>
Gocheckcompilerdirectives validates that compiler directives (lines starting with `//`) are properly formatted. Misplaced or malformed directives like `//go:generate` or `//go:embed` that lack the exact spacing cause silent failures.

Ensure compiler directives use exactly `//go:` (two slashes, no space before `go:`) on the line immediately before the affected declaration. Do not place directives inside function bodies.
</instructions>

<examples>
## Bad
```go
// go:generate go run gen.go
type Config struct{}
```

## Good
```go
//go:generate go run gen.go
type Config struct{}
```
</examples>

<patterns>
- Space between `//` and `go:` in directives — `// go:generate` is a comment, not a directive
- `//go:embed` placed on wrong line or after the declaration
- Custom build tags with incorrect formatting
- `//go:build` constraints not paired with `// +build` for older Go versions
</patterns>

<related>
gomoddirectives, godot, goheader
</related>
