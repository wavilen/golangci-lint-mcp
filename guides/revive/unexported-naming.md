# revive: unexported-naming

<instructions>
Flags unexported names in package-level declarations that could cause confusion. This rule checks for naming inconsistencies where package-level types, constants, or variables use names that suggest they should be exported (e.g., acronyms in unexpected casing), or where internal-only names leak into documentation.

Follow Go naming conventions: use PascalCase for exported names, camelCase for unexported. Ensure the naming makes the visibility intent clear.
</instructions>

<examples>
## Good
```go
type httpHandler struct{} // unexported, fine if truly internal
const maxRetries = 3      // camelCase for unexported
```
</examples>

<patterns>
- Use camelCase for unexported constants and variables at package level instead of snake_case
- Rename unexported types that suggest public API usage to reflect their internal purpose
- Use consistent casing in unexported names — e.g., `httpHandler` consistently, not mixed with `HTTPHandler`
- Replace ALL_CAPS package-level names with Go-style camelCase for unexported constants
- Use naming style for unexported names within the same package
</patterns>

<related>
revive/unexported-return, revive/var-naming, revive/receiver-naming
</related>
