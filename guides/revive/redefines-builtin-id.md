# revive: redefines-builtin-id

<instructions>
Detects shadowing of built-in identifiers such as `true`, `false`, `nil`, `append`, `len`, `cap`, `close`, `delete`, `copy`, `new`, `make`, `panic`, `recover`, `print`, `println`, `complex`, `real`, `imag`, and `error`. Redefining builtins makes code confusing because readers expect these names to have their standard meaning.

Rename the variable or type to avoid colliding with the built-in identifier. Use a more descriptive name.
</instructions>

<examples>
## Bad
```go
var true = false
var append = func(items ...int) []int { ... }
error := "something went wrong"
```

## Good
```go
var isValid = false
var addItems = func(items ...int) []int { ... }
errMsg := "something went wrong"
```
</examples>

<patterns>
- Variables named `true`, `false`, or `nil` in tests or generated code
- Function parameters shadowing builtins like `close` or `copy`
- Local variables named `error` instead of `err`
- Loop variables named `copy` or `append`
- Type aliases named after builtins in generated protobuf code
</patterns>

<related>
var-naming, confusing-naming
