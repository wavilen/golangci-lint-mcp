# gocritic: equalFold

<instructions>
Detects case-insensitive string comparisons using `strings.ToLower(a) == strings.ToLower(b)` (or `ToUpper`). This allocates two new strings unnecessarily. Use `strings.EqualFold(a, b)` instead — it performs case-insensitive comparison without allocating and handles Unicode correctly.

Replace any `strings.ToLower(x) == strings.ToLower(y)` or `ToUpper` variant with `strings.EqualFold(x, y)`.
</instructions>

<examples>
## Bad
```go
if strings.ToLower(s1) == strings.ToLower(s2) {
    slog.Info("equal")
}
```

## Good
```go
if strings.EqualFold(s1, s2) {
    slog.Info("equal")
}
```
</examples>

<patterns>
- Case-insensitive header or MIME type comparisons in HTTP handlers
- Comparing user input case-insensitively (e.g., "yes"/"NO" flags)
- String equality checks after `strings.ToLower` or `strings.ToUpper` on both operands
- Case-insensitive enum or config value matching
</patterns>

<related>
stringXbytes, preferStringWriter, unconvenName
