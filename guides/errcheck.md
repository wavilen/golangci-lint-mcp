# errcheck

<instructions>
Errcheck detects unchecked error returns in Go code. When a function returns an error and you don't handle it, failures can go silently unnoticed, leading to bugs that are hard to trace.

Always check error return values. If intentionally ignoring, use `_ =` with a comment explaining why.
</instructions>

<examples>
## Bad
```go
file, _ := os.Open("config.yaml")
data, _ := io.ReadAll(file)
```

## Good
```go
file, err := os.Open("config.yaml")
if err != nil {
    return errors.Wrap(err, "opening config")
}
data, err := io.ReadAll(file)
if err != nil {
    return errors.Wrap(err, "reading config")
}
```
</examples>

<patterns>
- Always check errors from `os.Open`, `os.Create`, `io.ReadAll` for file operations
- Use `v, ok := m[key]` two-value form for safe map access
- Use `v, ok := val.(T)` comma-ok pattern for type assertions
- Use `v, ok := <-ch` to detect closed channels on receive
</patterns>

<related>
err113, errname, wrapcheck, govet
</related>
