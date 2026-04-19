# gocritic: flagDeref

<instructions>
Detects immediate dereferencing of `flag.*` pointer results. When you call `flag.String`, `flag.Int`, etc., they return a pointer. Dereferencing the pointer immediately after declaration (e.g., in a package-level `var`) reads the value before `flag.Parse` has been called, so you get the zero value instead of the actual flag value.

Use the pointer directly and dereference only after `flag.Parse()`, or call `flag.Parse()` in `init()` before dereferencing.
</instructions>

<examples>
## Bad
```go
var port = *flag.Int("port", 8080, "listen port")
// port is always 0 — flag.Parse hasn't been called yet
```

## Good
```go
var port = flag.Int("port", 8080, "listen port")

func main() {
    flag.Parse()
    slog.Info("listening", "port", *port)
}
```
</examples>

<patterns>
- Package-level `var x = *flag.String(...)` dereferences before parse
- Dereferencing flag pointers in `var` blocks
- `*flag.Bool(...)` in global variable declarations
- Using flag values in `init()` functions
</patterns>

<related>
flagName, badCall
</related>
