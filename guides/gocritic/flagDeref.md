# gocritic: flagDeref

<instructions>
Detects immediate dereferencing of `flag.*` pointer results. When you call `flag.String`, `flag.Int`, etc., they return a pointer. Dereferencing the pointer immediately after declaration (e.g., in a package-level `var`) reads the value before `flag.Parse` has been called, so you get the zero value instead of the actual flag value.

Use the pointer directly and dereference only after `flag.Parse()`, or call `flag.Parse()` in `init()` before dereferencing.
</instructions>

<examples>
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
- Dereference flags after `flag.Parse()` — avoid `var x = *flag.String(...)` at package level
- Move flag dereferences from `var` blocks to `func main()` after `flag.Parse()`
- Replace `*flag.Bool(...)` in global declarations with pointer access after parse
- Use flag pointer values only after `flag.Parse()` — never in `init()` functions
</patterns>

<related>
gocritic/flagName, gocritic/badCall
</related>
