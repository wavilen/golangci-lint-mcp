# unconvert

<instructions>
Unconvert detects unnecessary type conversions where the source and target types are identical. Redundant conversions add noise and can obscure the intent of the code.

Remove the unnecessary conversion.
</instructions>

<examples>
## Bad
```go
func add(a, b int) int {
    return int(a) + int(b)
}
```

## Good
```go
func add(a, b int) int {
    return a + b
}
```
</examples>

<patterns>
- `int(x)` where `x` is already `int`
- `string(b)` where `b` is already `string`
- `[]byte(s)` when converting from `[]byte` to `[]byte`
- Converting between identical types in return statements
</patterns>

<related>
gosimple, staticcheck, govet
