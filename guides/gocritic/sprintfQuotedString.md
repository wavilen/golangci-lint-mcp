# gocritic: sprintfQuotedString

<instructions>
Detects `fmt.Sprintf("'%s'", s)` or `fmt.Sprintf("\"%s\"", s)` patterns that quote strings manually. This is fragile and produces incorrect output when the string itself contains quotes. Use `%q` format verb instead, which handles proper quoting and escaping automatically.

Replace manual quoting with `%q` in format strings.
</instructions>

<examples>
## Bad
```go
msg := fmt.Sprintf("'%s' not found", name)
```

## Good
```go
msg := fmt.Sprintf("%s not found", name)
// or if you need quotes:
msg := fmt.Sprintf("%q not found", name)
```
</examples>

<patterns>
- `fmt.Sprintf("'%s'", s)` for single-quote wrapping
- `fmt.Sprintf("\"%s\"", s)` for double-quote wrapping
- `"` + s + `"` string concatenation for quoting
- Manual escape sequences instead of `%q`
</patterns>

<related>
dynamicFmtString, badCall
</related>
