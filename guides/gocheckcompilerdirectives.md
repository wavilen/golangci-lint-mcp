# gocheckcompilerdirectives

<instructions>
Gocheckcompilerdirectives validates that compiler directives (lines starting with `//`) are properly formatted. Misplaced or malformed directives like `//go:generate` or `//go:embed` that lack the exact spacing cause silent failures.

Ensure compiler directives use exactly `//go:` (two slashes, no space before `go:`) on the line immediately before the affected declaration. Do not place directives inside function bodies.
</instructions>

<examples>
## Good
```go
//go:generate go run gen.go
type Config struct{}
```
</examples>

<patterns>
- Remove space between `//` and `go:` — use `//go:generate` not `// go:generate`
- Place `//go:embed` on the line immediately before the declaration it annotates
- Format custom build tags without spaces: `//go:build linux,amd64`
- Pair `//go:build` with `// +build` for compatibility with Go versions before 1.17
</patterns>

<related>
gomoddirectives, godot, goheader
</related>
