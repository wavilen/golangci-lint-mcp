# noinlineerr

<instructions>
Detects functions that create and return errors in a single expression like `return fmt.Errorf(...)`. Inline returns prevent debuggers from inspecting the error, `defer` from modifying it, and logs from capturing the failure point. Assign errors to a named return variable before returning.
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
