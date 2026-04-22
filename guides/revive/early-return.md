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
- Flatten nested if-else logic with early returns — handle error/edge cases first, then proceed with main logic
- Guard against errors first with early return instead of wrapping the entire function body in `if err == nil`
- Reduce nesting by converting multiple if-else levels to sequential early returns
- Convert "success path" code indented 3+ levels into a flat sequence with guard clauses
- Extract validation checks as early returns before the main business logic block
</patterns>

<related>
indent-error-flow, if-return, cognitive-complexity
