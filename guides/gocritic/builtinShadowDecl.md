# gocritic: builtinShadowDecl

<instructions>
Detects declarations (variables, types, or functions) that shadow built-in identifiers such as `len`, `cap`, `new`, `make`, `append`, `copy`, `delete`, `close`, `panic`, `recover`, `print`, `println`, `complex`, `real`, `imag`, `error`, `true`, `false`, or `nil`. Shadowing builtins makes code confusing and error-prone.

Rename the declaration to avoid shadowing the builtin. Use a more descriptive name that reflects the variable's purpose.
</instructions>

<examples>
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
- Rename variables that shadow `len()` builtin — use `length` or `count`
- Rename variables that shadow `new()` builtin — use `newVal` or `created`
- Rename variables that shadow the `error` type — use `err` or `appError`
- Avoid using builtin names (`close`, `copy`, `append`) as local variable names
</patterns>

<related>
gocritic/dupArg, gocritic/sloppyReassign
</related>
