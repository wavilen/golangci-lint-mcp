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
- Remove `//go:build linux` from files already named `_linux.go` — the suffix already implies the constraint
- Remove build tags that are always true (e.g., `//go:build true`)
- Combine duplicate build constraints into a single correct tag
- Remove obsolete `// +build` tags when `//go:build` is already present
- Eliminate build tags that are subsets of the file suffix convention
</patterns>

<related>
redundant-import-alias, redundant-test-main-exit
