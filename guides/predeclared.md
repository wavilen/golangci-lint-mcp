# predeclared

<instructions>
Predeclared detects when Go's predeclared identifiers (like `true`, `false`, `nil`, `iota`, `error`, `string`, `len`, `cap`, `copy`, etc.) are shadowed by local variable, type, or constant declarations. Shadowing built-in names causes confusion and subtle bugs.

Rename the shadowing identifier to something distinct from the predeclared name.
</instructions>

<examples>
## Good
```go
type appError struct {
    msg string
}
```
</examples>

<patterns>
- Rename types that shadow built-in `error` to a distinct name (e.g., `appError`)
- Rename variables that shadow built-in functions like `copy`
- Rename type aliases that shadow built-in `string`
- Rename local variables that shadow predeclared identifiers like `true`, `nil`, `len`
</patterns>

<related>
govet, revive, staticcheck
</related>
