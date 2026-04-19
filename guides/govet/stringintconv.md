# govet: stringintconv

<instructions>
Reports explicit `string(x)` conversions from integer types. `string(65)` produces the Unicode character `"A"` (a rune), not the decimal string `"65"`. This is almost never the intended behavior.

Use `strconv.Itoa(x)` to get the decimal string representation, or `string(rune(x))` explicitly if a rune conversion is intended.
</instructions>

<examples>
## Bad
```go
port := 8080
addr := "localhost:" + string(port) // produces "localhost: лиц" (rune), not "localhost:8080"
```

## Good
```go
port := 8080
addr := "localhost:" + strconv.Itoa(port) // produces "localhost:8080"
```
</examples>

<patterns>
- `string(intVar)` producing a rune instead of decimal string
- `string(int(expr))` conversion for display purposes
- Converting integer constants to strings without `strconv`
</patterns>

<related>
printf, shift
</related>
