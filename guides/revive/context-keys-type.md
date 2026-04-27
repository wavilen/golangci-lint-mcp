# revive: context-keys-type

<instructions>
Detects `context.WithValue` calls that use built-in types (string, int, etc.) as keys. Built-in types can collide across packages, causing one package to accidentally overwrite another's context values. Define a custom key type to prevent collisions.

Create an unexported custom type for the context key and use it consistently.
</instructions>

<examples>
## Good
```go
type contextKey string

const requestIDKey contextKey = "requestID"

ctx = context.WithValue(ctx, requestIDKey, "abc-123")
```
</examples>

<patterns>
- Define a custom type for context keys instead of using string literals directly
- Define an unexported key type instead of using `int` or built-in types
- Avoid cross-package key collisions by wrapping each key in a package-specific type
- Use typed context keys in tests instead of plain string keys for convenience
- Ensure type safety when storing and retrieving values from context
</patterns>

<related>
revive/context-as-argument, staticcheck/SA1029
</related>
