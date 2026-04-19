# goprintffuncname

<instructions>
Goprintffuncname checks that `Printf`-style functions are named consistently. Go convention requires `print`-like functions to be named `Printf`, `Sprintf`, `Fprintf`, or `Logf` so vet and tooling can verify format strings.

Rename logging or formatting functions to follow the `*f` suffix convention when they accept a format string. If not a format function, use a different name to avoid confusion.
</instructions>

<examples>
## Bad
```go
func logMessage(format string, args ...any) {
    log.Printf(format, args...)
}
```

## Good
```go
func logMessagef(format string, args ...any) {
    log.Printf(format, args...)
}
```
</examples>

<patterns>
- Custom logging wrappers with format strings but no `f` suffix
- Functions named `log*` that call `fmt.Sprintf` internally
- Helper functions wrapping `t.Logf` in tests
- Variadic string functions that look like format functions
</patterns>

<related>
forbidigo, errcheck, unparam
</related>
