# gocritic: indexAlloc

<instructions>
Detects `strings.Index` or `strings.IndexByte` used to check substring/byte presence where `strings.Contains` or `bytes.Contains` would be clearer and avoids the integer return value. While the performance difference is negligible, `Contains` expresses intent more directly.

Replace `strings.Index(s, sub) >= 0` with `strings.Contains(s, sub)`, and `strings.IndexByte(s, b) >= 0` with `strings.Contains(s, string(b))` or `bytes.IndexByte` equivalents.
</instructions>

<examples>
## Bad
```go
if strings.Index(line, "error") != -1 {
    log.Println("error found")
}
```

## Good
```go
if strings.Contains(line, "error") {
    log.Println("error found")
}
```
</examples>

<patterns>
- Checking substring presence with `strings.Index(...) != -1` or `>= 0`
- Using `bytes.Index` for presence checks instead of `bytes.Contains`
- `strings.IndexByte(s, ',') != -1` for single-byte presence tests
- Legacy code predating `strings.Contains` (added in Go 1.1)
</patterns>

<related>
equalFold, preferDecodeRune, stringXbytes
