# noinlineerr

<instructions>
Noinlineerr detects functions that inline error creation and return in a single expression, making it harder to add context or debug. It flags patterns like `return fmt.Errorf(...)` when the error should be assigned to a named return variable for inspection.

Assign errors to a named return variable before returning.
</instructions>

<examples>
## Bad
```go
func load() error {
    data, err := os.ReadFile("cfg.toml")
    if err != nil {
        return errors.Wrap(err, "load config")
    }
    return nil
}
```

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
- Inline `return errors.Wrap(...)` in error paths
- Unnamed error returns preventing defers from inspecting the value
- Multiple return paths where a named return would simplify cleanup
</patterns>

<related>
errcheck, wrapcheck, govet
