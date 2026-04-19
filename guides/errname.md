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
- Error types named without the `Error` suffix (e.g., `type NotFound`)
- Sentinel errors not prefixed with `Err` (e.g., `ErrorNotFound`)
- Inconsistent error naming within a package
</patterns>

<related>
errcheck, revive, govet
