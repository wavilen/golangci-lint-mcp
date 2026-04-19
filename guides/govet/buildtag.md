# govet: buildtag

<instructions>
Reports invalid or misplaced `//go:build` and `// +build` directives. Build constraints must appear before the package clause at the top of the file, and the syntax must follow the build constraint format (e.g., `//go:build linux && amd64`).

Fix the constraint syntax and ensure the directive is at the very top of the file, before `package`.
</instructions>

<examples>
## Bad
```go
package main // build tag after package clause

//go:build linux
```

## Good
```go
//go:build linux

package main
```
</examples>

<patterns>
- Build tag placed after the package declaration
- Invalid boolean expression in build tag (`||` without parentheses)
- Malformed `// +build` syntax
- Missing blank line between `//go:build` and `// +build` lines
</patterns>

<related>
directive, stdversion
</related>
