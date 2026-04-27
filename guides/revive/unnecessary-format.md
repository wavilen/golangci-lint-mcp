# revive: unnecessary-format

<instructions>
Detects unnecessary use of `fmt.Sprintf` or other formatting functions where a simpler alternative exists. Using `fmt.Sprintf("%s", s)` instead of just `s`, or `fmt.Sprintf(msg)` with no format verbs, adds unnecessary overhead and reduces readability.

Replace `fmt.Sprintf` with the simpler alternative: use string concatenation, direct value, or `fmt.Sprint` when no format verbs are needed.
</instructions>

<examples>
## Good
```go
msg := name
msg := "hello"
```
</examples>

<patterns>
- Use the string directly instead of `fmt.Sprintf("%s", str)`
- Replace `fmt.Sprintf` with no format verbs with a plain string literal
- Use `fmt.Sprint(x)` instead of `fmt.Sprintf("%v", x)`
- Replace `fmt.Sprintf` for simple concatenation with the `+` operator
- Use the argument directly when the format call result is identical to one of the arguments
</patterns>

<related>
revive/unnecessary-if, revive/unnecessary-stmt, revive/use-fmt-print
</related>
