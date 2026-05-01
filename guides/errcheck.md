# errcheck

<instructions>
Detects error return values that are discarded or assigned but never checked. Unchecked errors cause silent data corruption and security vulnerabilities that are extremely difficult to trace in production. Check every error return with `if err != nil`; when intentionally ignoring, use `_ =` with a comment explaining why.
</instructions>

<examples>
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
err113, errname, wrapcheck, govet, nilerr, rowserrcheck
</related>
