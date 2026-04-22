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
- Place build tags before the `package` declaration at the top of the file
- Use valid boolean syntax in build constraints — wrap `||` in parentheses
- Fix malformed `// +build` syntax to match the constraint format
- Add a blank line between `//go:build` and `// +build` directives when both are present
</patterns>

<related>
directive, stdversion
</related>
