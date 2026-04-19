# gocritic: redundantSprint

<instructions>
Detects `fmt.Sprintf` calls where the format string is a single `%v` or `%s` with one argument, or where `fmt.Sprint` is called on a single string argument. These are redundant — the argument can be used directly or converted with a simpler method.

Use the value directly, or call `string()` for simple conversions instead of formatting.
</instructions>

<examples>
## Bad
```go
msg := fmt.Sprintf("%v", err)
name := fmt.Sprint(s)
```

## Good
```go
msg := err.Error()
name := s
```
</examples>

<patterns>
- `fmt.Sprintf("%v", x)` → `fmt.Sprint(x)` or direct use
- `fmt.Sprint(s)` where `s` is already a string → `s`
- `fmt.Sprintf("%s", str)` → `str`
</patterns>

<related>
stringConcatSimplify, wrapperFunc
</related>
