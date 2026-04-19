# gocritic: dynamicFmtString

<instructions>
Detects `fmt.Sprintf` or similar formatting calls where the format string is a dynamic variable rather than a string literal. When the format string is not a compile-time constant, the formatter cannot verify format verbs match the arguments, and user-controlled format strings can cause panics or information leaks.

Use a literal format string whenever possible. If the string must be dynamic, use `fmt.Fprint` or `fmt.Sprint` without format verbs, or validate the format string.
</instructions>

<examples>
## Bad
```go
msg := fmt.Sprintf(userInput, data) // format string is dynamic
```

## Good
```go
msg := fmt.Sprintf("%s", userInput) // literal format string
// or simply:
msg := userInput
```
</examples>

<patterns>
- Format string from user input or external data
- Variable format string in `fmt.Printf` family
- Template strings constructed at runtime
- Config-driven format strings passed to `fmt.Sprintf`
</patterns>

<related>
sprintfQuotedString, badCall
</related>
