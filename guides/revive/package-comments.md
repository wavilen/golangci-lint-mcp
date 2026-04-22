# revive: package-comments

<instructions>
Requires every package to have a doc comment on the `package` declaration. Package-level comments explain the purpose of the package and appear in `go doc` output. The convention is to write a single sentence beginning with "Package {name}" on the line immediately preceding the `package` keyword.

Add a comment starting with `// Package {name} ...` directly above the package declaration in one file (typically `doc.go` or the main file).
</instructions>

<examples>
## Bad
```go
package user
```

## Good
```go
// Package user provides types and functions for managing user accounts
// in the system.
package user
```
</examples>

<patterns>
- Add a comment preceding the `package` declaration in every package
- Add a comment starting with "Package {name}" to every package
- Add doc comments to generated packages
- Document `main` packages with a comment explaining the command's purpose
- Replace TODO-only comments above package declarations with proper doc comments
</patterns>

<related>
package-naming, package-directory-mismatch, comment-spacings
