# nilerr

<instructions>
Nilerr detects cases where an error is checked but the function returns a nil error instead of the actual error, or returns a non-nil error when the check succeeded. This is typically a copy-paste bug where the wrong variable is returned.

Return the checked error variable, not nil or a different error.
</instructions>

<examples>
## Bad
```go
func process() error {
    result, err := doWork()
    if err != nil {
        return nil // should return err
    }
    return result.Save()
}
```

## Good
```go
func process() error {
    result, err := doWork()
    if err != nil {
        return err
    }
    return result.Save()
}
```
</examples>

<patterns>
- `if err != nil { return nil }` — returns nil instead of the checked error
- `if err == nil { return err }` — returns nil error incorrectly
- Confusing error variable names leading to returning the wrong one
</patterns>

<related>
errcheck, govet, staticcheck
