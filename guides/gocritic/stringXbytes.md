# gocritic: stringXbytes

<instructions>
Detects unnecessary `string(b)` or `[]byte(s)` conversions in specific contexts where they can be avoided. Each conversion allocates a new copy of the data. In `fmt.Print` family calls and string concatenation with `+`, `[]byte` arguments are auto-converted — explicit conversion is redundant.

Remove the explicit `string()` or `[]byte()` conversion where the target accepts the original type directly.
</instructions>

<examples>
## Bad
```go
data := []byte("hello")
fmt.Println(string(data)) // unnecessary string() conversion
```

## Good
```go
data := []byte("hello")
fmt.Println(data) // fmt handles []byte natively
```
</examples>

<patterns>
- `fmt.Println(string(byteSlice))` — `fmt` formats `[]byte` as a string automatically
- `string(b) == "expected"` comparisons — use `bytes.Equal(b, []byte("expected"))`
- `[]byte(s)` then immediately passing to a function accepting `string`
- Redundant conversions in hot loops where each allocation adds up
</patterns>

<related>
equalFold, preferStringWriter, preferFprint
