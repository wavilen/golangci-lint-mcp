# noinlineerr

<instructions>
Noinlineerr detects functions that inline error creation and return in a single expression, making it harder to add context or debug. It flags patterns like `return fmt.Errorf(...)` when the error should be assigned to a named return variable for inspection.

Assign errors to a named return variable before returning.
</instructions>

<examples>
## Good
```go
func load() (err error) {
    data, err := os.ReadFile("cfg.toml")
    if err != nil {
        return errors.Wrap(err, "load config")
    }
    return nil
}
```
</examples>

<patterns>
- Use named return values and assign errors in conditional blocks rather than inline `return errors.Wrap(...)`
- Declare named error returns when defers need to inspect or modify the error value
- Use named returns to simplify cleanup across multiple return paths
</patterns>

<related>
errcheck, wrapcheck, govet
</related>
