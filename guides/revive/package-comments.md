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
- Package declarations without any preceding comment
- Packages with comments that don't start with "Package {name}"
- Generated packages missing doc comments
- Main packages that skip documentation
- Packages with only TODO comments above the declaration
</patterns>

<related>
package-naming, package-directory-mismatch, comment-spacings
