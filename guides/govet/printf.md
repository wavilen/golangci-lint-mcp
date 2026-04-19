# govet: printf

<instructions>
Reports `fmt.Printf`/`Sprintf`/`Fprintf`/`Errorf` format string mismatches. This includes wrong format verbs for argument types (`%d` for a string), wrong argument count (too few or too many), and non-constant format strings when a constant is expected.

Fix the format verbs to match argument types and ensure the argument count matches the format specifiers.
</instructions>

<examples>
## Bad
```go
name := "Alice"
age := 30
fmt.Printf("Name: %d, Age: %s\n", name, age) // %d for string, %s for int
```

## Good
```go
name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age) // correct verbs
```
</examples>

<patterns>
- Wrong format verb for argument type (`%d` for string, `%s` for int)
- More arguments than format specifiers
- Fewer arguments than format specifiers
- `%` at end of format string (truncated verb)
</patterns>

<related>
slog, stringintconv
</related>
