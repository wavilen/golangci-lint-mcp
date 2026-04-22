# revive: error-return

<instructions>
Enforces that errors are returned as the last return value in function signatures. Go convention strongly prefers error as the final return value. Placing error in any other position makes it easy to forget error handling at call sites.

Reorder return values to place the error last: `(resultType, error)`.
</instructions>

<examples>
## Bad
```go
func Lookup() (error, string) {
    // error first is unusual and confuses callers
    return nil, "found"
}

func Parse() (int, error, bool) {
    // error in the middle is even worse
    return 42, nil, true
}
```

## Good
```go
func Lookup() (string, error) {
    return "found", nil
}

func Parse() (int, bool, error) {
    return 42, true, nil
}
```
</examples>

<patterns>
- Move the error return to the last position when functions return error first
- Reorder return values so error is the last value in multi-return functions
- Ensure constructor functions return error as the last value
- Use Go convention by placing error last — avoid patterns from other languages
- Replace interface methods where implementations place error in a non-last position
</patterns>

<related>
error-naming, error-strings, errorf
