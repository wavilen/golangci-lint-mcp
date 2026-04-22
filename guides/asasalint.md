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
- Match format verbs to argument types: use `%d` for integers, `%s` for strings
- Add type assertions or conversions before passing `any` values to `fmt` functions
- Check aliased types against format verbs — a type alias does not change the underlying format
</patterns>

<related>
errcheck, govet, gosmopolitan