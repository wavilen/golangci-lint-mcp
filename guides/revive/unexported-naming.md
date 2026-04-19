# revive: unexported-naming

<instructions>
Flags unexported names in package-level declarations that could cause confusion. This rule checks for naming inconsistencies where package-level types, constants, or variables use names that suggest they should be exported (e.g., acronyms in unexpected casing), or where internal-only names leak into documentation.

Follow Go naming conventions: use PascalCase for exported names, camelCase for unexported. Ensure the naming makes the visibility intent clear.
</instructions>

<examples>
## Bad
```go
type httpHandler struct{} // unexported but looks like it wraps http
const max_retries = 3     // snake_case in Go
```

## Good
```go
type httpHandler struct{} // unexported, fine if truly internal
const maxRetries = 3      // camelCase for unexported
```
</examples>

<patterns>
- Snake_case constant or variable names at package level
- Unexported types with names that suggest public API usage
- Inconsistent casing in unexported names (e.g., `httpHandler` vs `HTTPHandler`)
- Package-level names using ALL_CAPS as if they were C constants
- Mixed naming styles within the same package
</patterns>

<related>
unexported-return, var-naming, receiver-naming
