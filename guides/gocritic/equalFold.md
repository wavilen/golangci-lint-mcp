# gocritic: equalFold

<instructions>
Detects case-insensitive string comparisons using `strings.ToLower(a) == strings.ToLower(b)` (or `ToUpper`). This allocates two new strings unnecessarily. Use `strings.EqualFold(a, b)` instead — it performs case-insensitive comparison without allocating and handles Unicode correctly.

Replace any `strings.ToLower(x) == strings.ToLower(y)` or `ToUpper` variant with `strings.EqualFold(x, y)`.
</instructions>

<examples>
## Good
```go
if strings.EqualFold(s1, s2) {
    slog.Info("equal")
}
```
</examples>

<patterns>
- Use `strings.EqualFold` for case-insensitive HTTP header or MIME type comparisons
- Use `strings.EqualFold` for case-insensitive user input matching like "yes"/"NO"
- Replace `strings.ToLower(a) == strings.ToLower(b)` with `strings.EqualFold(a, b)`
- Use `strings.EqualFold` for case-insensitive enum or config value matching
</patterns>

<related>
gocritic/stringXbytes, gocritic/preferStringWriter
</related>
