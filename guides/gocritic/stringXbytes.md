# gocritic: stringXbytes

<instructions>
Detects unnecessary `string(b)` or `[]byte(s)` conversions in specific contexts where they can be avoided. Each conversion allocates a new copy of the data. In `fmt.Print` family calls and string concatenation with `+`, `[]byte` arguments are auto-converted — explicit conversion is redundant.

Remove the explicit `string()` or `[]byte()` conversion where the target accepts the original type directly.
</instructions>

<examples>
## Good
```go
data := []byte("hello")
fmt.Println(data) // fmt handles []byte natively
```
</examples>

<patterns>
- Remove `string(byteSlice)` conversion when passing to `fmt.Println` — `fmt` formats `[]byte` as string
- Replace `string(b) == "expected"` with `bytes.Equal(b, []byte("expected"))` to avoid allocation
- Remove `[]byte(s)` conversion when the target function accepts `string`
- Eliminate redundant `string`↔`[]byte` conversions in hot loops — use consistent types
</patterns>

<related>
gocritic/equalFold, gocritic/preferStringWriter, gocritic/preferFprint
</related>
