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
- Use `strconv.Itoa` or `fmt.Sprintf` to convert integers to strings — never `string(intVar)`
- Replace `string(int(...))` with `strconv.Itoa` for display output
- Use `strconv` functions instead of `string()` for int-to-string conversion
</patterns>

<related>
printf, shift
</related>
