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
- Use `Err` prefix for sentinel error variables (e.g., `ErrNotFound`) instead of `Error` suffix
- Use `ErrXxx` naming convention for exported error variables consistently
- Use standard naming conventions for local error variables
- Add `Error` suffix to custom error type names (e.g., `NotFoundError`)
- Use consistent naming between error variables and error types
</patterns>

<related>
error-return, error-strings, errorf, errname
