# revive: string-of-int

<instructions>
Detects `string(intValue)` conversions which produce a Unicode character rather than the decimal representation of the number. For example, `string(65)` gives `"A"` (the rune), not `"65"`. This is almost always a mistake — use `strconv.Itoa` or `fmt.Sprintf` for numeric-to-string conversion.

Replace `string(n)` with `strconv.Itoa(n)` or `fmt.Sprintf("%d", n)` to get the decimal string representation.
</instructions>

<examples>
## Bad
```go
port := 8080
addr := "localhost:" + string(port) // gets rune, not "8080"
```

## Good
```go
port := 8080
addr := "localhost:" + strconv.Itoa(port)
```
</examples>

<patterns>
- Use `strconv.Itoa` or `fmt.Sprintf("%d", n)` to convert int to string for display
- Replace `string()` on numeric types from external APIs with `strconv.Itoa`
- Convert byte counts or sizes using `strconv.Itoa` or `fmt.Sprintf`
- Use `strconv.Itoa(count)` instead of `string(count)` for log messages
- Use `strconv.Itoa` for numeric formatting instead of `string()` on rune values
</patterns>

<related>
use-fmt-print, unnecessary-format, gocritic/stringX
