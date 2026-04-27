# gocritic: indexAlloc

<instructions>
Detects `strings.Index` or `strings.IndexByte` used to check substring/byte presence where `strings.Contains` or `bytes.Contains` would be clearer and avoids the integer return value. While the performance difference is negligible, `Contains` expresses intent more directly.

Replace `strings.Index(s, sub) >= 0` with `strings.Contains(s, sub)`, and `strings.IndexByte(s, b) >= 0` with `strings.Contains(s, string(b))` or `bytes.IndexByte` equivalents.
</instructions>

<examples>
## Good
```go
if strings.Contains(line, "error") {
    log.Println("error found")
}
```
</examples>

<patterns>
- Replace `strings.Index(s, sub) != -1` with `strings.Contains(s, sub)`
- Replace `bytes.Index(b, sub) >= 0` with `bytes.Contains(b, sub)`
- Replace `strings.IndexByte(s, ',') != -1` with `strings.ContainsByte(s, ',')`
- Use `strings.Contains` for all substring presence checks
</patterns>

<related>
gocritic/equalFold, gocritic/preferDecodeRune, gocritic/stringXbytes
</related>
