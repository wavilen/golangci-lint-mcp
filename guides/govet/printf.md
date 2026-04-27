# govet: printf

<instructions>
Reports `fmt.Printf`/`Sprintf`/`Fprintf`/`Errorf` format string mismatches. This includes wrong format verbs for argument types (`%d` for a string), wrong argument count (too few or too many), and non-constant format strings when a constant is expected.

Fix the format verbs to match argument types and ensure the argument count matches the format specifiers.
</instructions>

<examples>
## Good
```go
name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age) // correct verbs
```
</examples>

<patterns>
- Match format verbs to argument types (`%s` for string, `%d` for int, `%v` as fallback)
- Add format specifiers for every argument — remove extra arguments
- Add arguments for every format specifier — provide all values
- Complete truncated format verbs — never end a format string with `%`
</patterns>

<related>
govet/slog, govet/stringintconv
</related>
