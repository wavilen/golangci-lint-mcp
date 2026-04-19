# revive: early-return

<instructions>
Detects code that could use early returns (guard clauses) to reduce nesting. Deep nesting makes code harder to read because the reader must track multiple indentation levels. Guard clauses handle error or edge cases first, then continue with the main logic.

Invert the condition and return immediately. Flatten the remaining code by one or more indentation levels.
</instructions>

<examples>
## Bad
```go
func process(data []byte) error {
    if len(data) > 0 {
        if isValid(data) {
            result := transform(data)
            return save(result)
        } else {
            return errors.New("invalid data")
        }
    } else {
        return errors.New("empty data")
    }
}
```

## Good
```go
func process(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    if !isValid(data) {
        return errors.New("invalid data")
    }
    result := transform(data)
    return save(result)
}
```
</examples>

<patterns>
- Functions where the main logic is deeply nested inside if-else blocks
- Error handling wrapping the entire function body
- Multiple levels of if-else that could be flattened with early returns
- Functions with a "success path" indented 3+ levels
- Validation checks followed by business logic all in one nested block
</patterns>

<related>
indent-error-flow, if-return, cognitive-complexity
