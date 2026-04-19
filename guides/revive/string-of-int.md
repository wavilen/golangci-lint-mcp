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
- Converting int to string for display or concatenation
- Using `string()` on numeric types from external APIs
- Converting byte counts or sizes to strings
- Building log messages with `string(count)` instead of `strconv.Itoa`
- Converting rune values when the intent is actually numeric formatting
</patterns>

<related>
use-fmt-print, unnecessary-format, gocritic/stringX
