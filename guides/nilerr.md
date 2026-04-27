# nilerr

<instructions>
Nilerr detects cases where an error is checked but the function returns a nil error instead of the actual error, or returns a non-nil error when the check succeeded. This is typically a copy-paste bug where the wrong variable is returned.

Return the checked error variable, not nil or a different error.
</instructions>

<examples>
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
- Return the checked error in `if err != nil` branches instead of returning `nil`
- Return `nil` explicitly in `if err == nil` branches when no error occurred
- Rename error variables clearly to avoid returning the wrong error in conditional branches
</patterns>

<related>
errcheck, govet, staticcheck, nilnesserr
</related>
