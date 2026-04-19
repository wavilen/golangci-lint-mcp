# asasalint

<instructions>
Asasalint detects passing `any` values to `fmt.Sprintf`-like functions where the argument should match a specific format verb. This typically happens when `fmt.Sprintf("%d", anyValue)` is called with a value of the wrong type. Fix by ensuring the argument type matches the format verb, or by converting the value to the expected type.
</instructions>

<examples>
## Bad
```go
var val any = "hello"
fmt.Sprintf("%d", val)
```

## Good
```go
var val any = "hello"
fmt.Sprintf("%s", val)
```
</examples>

<patterns>
- Format verb mismatch: `%d` with string, `%s` with int
- Passing `interface{}` or `any` to fmt functions without type assertion
- Aliased types used with wrong format verb
</patterns>

<related>
errcheck, govet, gosmopolitan