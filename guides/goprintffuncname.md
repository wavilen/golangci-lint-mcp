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
- Rename custom logging wrappers to include the `f` suffix when using format strings
- Add `f` suffix to `log*` functions that call `fmt.Sprintf` internally
- Rename test helper functions wrapping `t.Logf` with the `f` suffix
- Rename variadic string functions that use format strings to end with `f`
</patterns>

<related>
forbidigo, errcheck, unparam
</related>
