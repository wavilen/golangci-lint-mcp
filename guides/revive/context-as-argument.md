# revive: context-as-argument

<instructions>
Enforces that `context.Context` is always the first parameter of a function and has type `context.Context` (not a pointer or wrapper). Go convention established by the standard library is that context flows as the first argument, making it easy to thread cancellation and deadlines through call chains.

Move the `context.Context` parameter to the first position. Do not wrap it in another type.
</instructions>

<examples>
## Bad
```go
func Process(data []byte, ctx context.Context) error {
    return nil
}

func Handle(*CustomContext) error { // custom wrapper
    return nil
}
```

## Good
```go
func Process(ctx context.Context, data []byte) error {
    return nil
}

func Handle(ctx context.Context) error {
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
context-keys-type, inamedparam
