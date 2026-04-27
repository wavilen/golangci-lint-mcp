# revive: context-as-argument

<instructions>
Enforces that `context.Context` is always the first parameter of a function and has type `context.Context` (not a pointer or wrapper). Go convention established by the standard library is that context flows as the first argument, making it easy to thread cancellation and deadlines through call chains.

Move the `context.Context` parameter to the first position. Do not wrap it in another type.
</instructions>

<examples>
## Good
```go
func Process(_ context.Context, _ []byte) error {
    return nil
}

func Handle(_ context.Context) error {
    return nil
}
```
</examples>

<patterns>
- Move `context.Context` to the first parameter position in function signatures
- Use `context.Context` directly instead of custom wrapper types
- Ensure methods place context as the first parameter after the receiver
- Pass `context.Context` through the call chain rather than storing it in a struct field
- Pass context as the first argument in HTTP handlers instead of storing it in struct fields
</patterns>

<related>
revive/context-keys-type, inamedparam
</related>
