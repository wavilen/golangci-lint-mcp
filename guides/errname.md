# errname

<instructions>
Errname checks that error type names and sentinel error variable names follow the Go convention: error types should end in `Error` and sentinel error variables should start with `Err` or `err`. This improves readability and consistency across codebases.

Rename error types and variables to follow the standard naming convention.
</instructions>

<examples>
## Bad
```go
type NotFound struct {
    Msg string
}

var ErrorNotFound = errors.New("not found")
```

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
