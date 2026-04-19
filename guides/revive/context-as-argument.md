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
- `context.Context` appearing as the second or later parameter
- Custom context wrapper types used instead of `context.Context`
- Methods with context as a non-first parameter
- Functions that take a context but pass it deeper without using it
- HTTP handlers storing context in a struct field instead of passing it
</patterns>

<related>
context-keys-type, inamedparam
