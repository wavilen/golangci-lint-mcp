# govet: stdversion

<instructions>
Reports `//go:build` constraints that reference Go versions beyond the maximum version the toolchain recognizes. Using `//go:build go1.999` or an unrecognized future version means the constraint never matches or behaves unexpectedly.

Use correct, recognized Go version numbers in build constraints.
</instructions>

<examples>
## Good
```go
//go:build go1.22

package pkg // recognized Go version
```
</examples>

<patterns>
- Use Go versions supported by your toolchain in build constraints
- Fix typos in Go version numbers (e.g., `go1.2` → `go1.20`)
- Use only released Go version numbers in build constraints — remove speculative versions
</patterns>

<related>
govet/buildtag, govet/directive
</related>
