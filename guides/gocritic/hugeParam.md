# gocritic: hugeParam

<instructions>
Detects function parameters that are large value types (typically structs ≥80 bytes). Passing large types by value causes expensive copies on every call. Use a pointer receiver or parameter instead to avoid the copy overhead.

Change the parameter type from `T` to `*T` when `T` is a large struct and the function does not need to modify the original.
</instructions>

<examples>
## Good
```go
func Process(cfg *BigConfig) error {
    // cfg is a pointer — no copy overhead
    return nil
}
```
</examples>

<patterns>
- Pass large structs by pointer instead of value as function parameters or method receivers
- Pass config or option structs exceeding ~80 bytes by pointer
- Use pointer receivers on large structs that only read data
- Replace large array value parameters with pointer or slice parameters
</patterns>

<related>
gocritic/rangeValCopy, gocritic/rangeExprCopy, prealloc
</related>
