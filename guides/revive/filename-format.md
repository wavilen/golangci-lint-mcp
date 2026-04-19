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
- Files using PascalCase or camelCase instead of snake_case
- Files containing spaces or hyphens in the name
- Files with inconsistent naming patterns within a package
- Test files not following the `*_test.go` convention
- Platform-specific files with wrong naming (e.g., `linux_amd64.go`)
</patterns>

<related>
file-header, file-length-limit
