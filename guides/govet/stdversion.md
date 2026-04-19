# govet: stdversion

<instructions>
Reports `//go:build` constraints that reference Go versions beyond the maximum version the toolchain recognizes. Using `//go:build go1.999` or an unrecognized future version means the constraint never matches or behaves unexpectedly.

Use correct, recognized Go version numbers in build constraints.
</instructions>

<examples>
## Bad
```go
//go:build go1.999

package pkg // version 1.999 does not exist
```

## Good
```go
//go:build go1.22

package pkg // recognized Go version
```
</examples>

<patterns>
- Build constraint using a Go version beyond the toolchain's knowledge
- Typos in version numbers (e.g., `go1.2` instead of `go1.20`)
- Future version numbers that don't exist yet
</patterns>

<related>
buildtag, directive
</related>
