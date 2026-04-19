# gocritic: hugeParam

<instructions>
Detects function parameters that are large value types (typically structs ≥80 bytes). Passing large types by value causes expensive copies on every call. Use a pointer receiver or parameter instead to avoid the copy overhead.

Change the parameter type from `T` to `*T` when `T` is a large struct and the function does not need to modify the original.
</instructions>

<examples>
## Bad
```go
type BigConfig struct {
    Data [1024]byte
    Name string
    Tags []string
}

func Process(cfg BigConfig) error {
    // cfg is copied — 1024+ bytes per call
    return nil
}
```

## Good
```go
func Process(cfg *BigConfig) error {
    // cfg is a pointer — no copy overhead
    return nil
}
```
</examples>

<patterns>
- Large structs passed by value as function parameters or method receivers
- Config or option structs exceeding ~80 bytes passed by value
- Value receivers on large structs that only read data
- Functions accepting large arrays by value instead of pointer or slice
</patterns>

<related>
rangeValCopy, rangeExprCopy, unexportedCall
