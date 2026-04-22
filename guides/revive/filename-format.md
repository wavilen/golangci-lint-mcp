# revive: filename-format

<instructions>
Enforces that source file names follow a consistent naming convention. Go convention uses `lowercase_with_underscores.go` (snake_case) for file names, with no spaces or special characters. Test files use `_test.go` suffix.

Rename files to follow the configured naming convention. Use `go fmt`-compatible names: lowercase, underscores for separators, `.go` extension.
</instructions>

<examples>
## Bad
```go
// userHandler.go
// User-Service.go
// user service.go
```

## Good
```go
// user_handler.go
// user_service.go
// user.go
```
</examples>

<patterns>
- Rename files using PascalCase or camelCase to snake_case
- Remove spaces or hyphens from file names — use underscores instead
- Use a consistent snake_case naming pattern across all files in a package
- Rename test files following the `*_test.go` convention
- Use platform-specific naming for build-constraint files (e.g., `_linux_amd64.go`)
</patterns>

<related>
file-header, file-length-limit
