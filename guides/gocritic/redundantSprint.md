# gocritic: redundantSprint

<instructions>
Detects `fmt.Sprintf` calls where the format string is a single `%v` or `%s` with one argument, or where `fmt.Sprint` is called on a single string argument. These are redundant — the argument can be used directly or converted with a simpler method.

Use the value directly, or call `string()` for simple conversions instead of formatting.
</instructions>

<examples>
## Good
```go
msg := err.Error()
name := s
```
</examples>

<patterns>
- Replace `fmt.Sprintf("%v", x)` with `fmt.Sprint(x)` or use `x` directly
- Replace `fmt.Sprint(s)` with `s` when `s` is already a string
- Replace `fmt.Sprintf("%s", str)` with `str` directly
</patterns>

<related>
gocritic/stringConcatSimplify, gocritic/wrapperFunc
</related>
