# revive: redundant-build-tag

<instructions>
Detects redundant or obsolete build constraint tags in Go source files. Build tags that duplicate the file location (e.g., a `_linux.go` suffix with a `//go:build linux` tag) or tags that are always true are unnecessary noise.

Remove the redundant build tag, or consolidate multiple tags into a single correct constraint.
</instructions>

<examples>
## Bad
```go
//go:build linux

// file: foo_linux.go (suffix already implies linux)
package pkg
```

## Good
```go
// file: foo_linux.go
package pkg
```
</examples>

<patterns>
- `//go:build linux` in a file already named `_linux.go`
- Build tags that are always true (e.g., `//go:build true`)
- Duplicate build constraints that say the same thing twice
- Obsolete `// +build` tags when `//go:build` is already present
- Build tags that are subsets of the file suffix convention
</patterns>

<related>
redundant-import-alias, redundant-test-main-exit
