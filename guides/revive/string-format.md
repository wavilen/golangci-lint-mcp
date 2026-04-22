# revive: string-format

<instructions>
Detects format string mismatches in `fmt.Sprintf`, `fmt.Printf`, and similar functions. This includes wrong format verbs for the argument type (e.g., `%d` for a string), more verbs than arguments, or unused arguments. These errors produce incorrect output at runtime.

Match the format verb to the argument type: `%s` for strings, `%d` for integers, `%v` for any type, `%w` for error wrapping. Ensure every verb has a corresponding argument.
</instructions>

<examples>
## Bad
```go
name := "Alice"
fmt.Sprintf("Hello %d", name) // %d for string
fmt.Sprintf("Hello %s %s", name) // 2 verbs, 1 arg
```

## Good
```go
name := "Alice"
fmt.Sprintf("Hello %s", name)
fmt.Sprintf("Hello %s, you are #%d", name, 1)
```
</examples>

<patterns>
- Ensure format verbs to argument types — use `%s` for strings, `%d` for integers
- Ensure every format verb has a corresponding argument — remove extra verbs
- Provide arguments for every format verb — remove unused arguments
- Use type-specific verbs like `%d` instead of `%v` to catch type mismatches at review time
- Use width/precision modifiers with the argument type
</patterns>

<related>
errorf, unnecessary-format, use-fmt-print
