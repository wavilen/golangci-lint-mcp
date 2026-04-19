# revive: unnecessary-format

<instructions>
Detects unnecessary use of `fmt.Sprintf` or other formatting functions where a simpler alternative exists. Using `fmt.Sprintf("%s", s)` instead of just `s`, or `fmt.Sprintf(msg)` with no format verbs, adds unnecessary overhead and reduces readability.

Replace `fmt.Sprintf` with the simpler alternative: use string concatenation, direct value, or `fmt.Sprint` when no format verbs are needed.
</instructions>

<examples>
## Bad
```go
msg := fmt.Sprintf("%s", name)
msg := fmt.Sprintf("hello") // no verbs
```

## Good
```go
msg := name
msg := "hello"
```
</examples>

<patterns>
- `fmt.Sprintf` with a single `%s` verb and one string argument
- `fmt.Sprintf` with no format verbs at all (just a plain string)
- `fmt.Sprintf("%v", x)` when `fmt.Sprint(x)` would suffice
- Using `fmt.Sprintf` for simple string concatenation instead of `+`
- Format calls where the result is identical to one of the arguments
</patterns>

<related>
unnecessary-if, unnecessary-stmt, use-fmt-print
