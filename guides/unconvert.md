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
- Remove `int(x)` when `x` is already `int`
- Remove `string(b)` when `b` is already `string`
- Remove `[]byte(s)` when converting between identical types
- Remove redundant type conversions in return statements
</patterns>

<related>
gosimple, staticcheck, govet
