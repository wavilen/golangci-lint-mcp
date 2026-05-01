# errname

<instructions>
Detects error types that don't end in `Error` and sentinel error variables that don't start with `Err` or `err`. Inconsistent naming breaks the Go convention readers expect, making it harder to distinguish error types from regular types in code review. Rename error types to end in `Error` and sentinel variables to start with `Err`.
</instructions>

<examples>
## Good
```go
type NotFoundError struct {
    Msg string
}

var ErrNotFound = errors.New("not found")
```
</examples>

<patterns>
- Add `Error` suffix to error type names (e.g., `NotFound` → `NotFoundError`)
- Rename sentinel error variables to start with `Err` (e.g., `ErrorNotFound` → `ErrNotFound`)
- Use consistent `Err` prefix for sentinel errors and `Error` suffix for error types across the package
</patterns>

<related>
errcheck, revive, govet
</related>
