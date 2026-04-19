# revive: error-naming

<instructions>
Enforces Go naming conventions for error variables. Error variables should be named `err` or start with `Err` (e.g., `ErrNotFound`). Sentinel errors stored as package-level variables must follow the `ErrXxx` convention.

Rename error variables to follow Go conventions. Use `err` for local error variables and `ErrXxx` for exported sentinel errors.
</instructions>

<examples>
## Bad
```go
var NotFoundError = errors.New("not found")
var errorPermission = errors.New("permission denied")
```

## Good
```go
var ErrNotFound = errors.New("not found")
var ErrPermissionDenied = errors.New("permission denied")
```
</examples>

<patterns>
- Sentinel errors named with `Error` suffix instead of `Err` prefix
- Exported error variables not following `ErrXxx` naming convention
- Local error variables using non-standard names
- Error type names not ending in `Error` (custom error types)
- Inconsistent naming between error variables and error types
</patterns>

<related>
error-return, error-strings, errorf, errname
