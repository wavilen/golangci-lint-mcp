# revive: context-keys-type

<instructions>
Detects `context.WithValue` calls that use built-in types (string, int, etc.) as keys. Built-in types can collide across packages, causing one package to accidentally overwrite another's context values. Define a custom key type to prevent collisions.

Create an unexported custom type for the context key and use it consistently.
</instructions>

<examples>
## Bad
```go
ctx = context.WithValue(ctx, "requestID", "abc-123")
ctx = context.WithValue(ctx, 42, someData)
```

## Good
```go
type contextKey string

const requestIDKey contextKey = "requestID"

ctx = context.WithValue(ctx, requestIDKey, "abc-123")
```
</examples>

<patterns>
- Using string literals directly as context keys
- Using `int` or other built-in types as keys
- Keys defined in one package potentially colliding with keys in another
- Tests using `context.WithValue` with plain string keys for convenience
- Missing type safety when retrieving values from context
</patterns>

<related>
context-as-argument, staticcheck SA1029
