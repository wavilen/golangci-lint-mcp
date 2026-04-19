# gocritic: builtinShadowDecl

<instructions>
Detects declarations (variables, types, or functions) that shadow built-in identifiers such as `len`, `cap`, `new`, `make`, `append`, `copy`, `delete`, `close`, `panic`, `recover`, `print`, `println`, `complex`, `real`, `imag`, `error`, `true`, `false`, or `nil`. Shadowing builtins makes code confusing and error-prone.

Rename the declaration to avoid shadowing the builtin. Use a more descriptive name that reflects the variable's purpose.
</instructions>

<examples>
## Bad
```go
func process(items []string) {
    len := 0 // shadows builtin len
    for range items {
        len++
    }
    slog.Info("value", "count", len)
}
```

## Good
```go
func process(items []string) {
    count := 0
    for range items {
        count++
    }
    slog.Info("value", "count", count)
}
```
</examples>

<patterns>
- `len := 0` shadowing the `len()` builtin
- `new := ...` shadowing the `new()` builtin
- `error := ...` shadowing the `error` type
- `close := ...` shadowing the `close()` builtin
- Local variable names matching builtin function names
</patterns>

<related>
dupArg, sloppyReassign
</related>
