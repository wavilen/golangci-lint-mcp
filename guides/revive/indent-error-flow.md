# revive: indent-error-flow

<instructions>
Enforces reducing indentation in error handling paths. Instead of nesting the happy path inside an `if err == nil` block, handle the error first with an early return and keep the main logic at a lower indentation level. This mirrors the early-return pattern but focuses specifically on error flows.

Invert the error check: handle the error case first with a return, then continue with the success path unindented.
</instructions>

<examples>
## Bad
```go
func process(data []byte) error {
    result, err := parse(data)
    if err == nil {
        if err := validate(result); err == nil {
            return save(result)
        } else {
            return errors.Wrap(err, "validate")
        }
    }
    return errors.Wrap(err, "parse")
}
```

## Good
```go
func process(data []byte) error {
    result, err := parse(data)
    if err != nil {
        return errors.Wrap(err, "parse")
    }
    if err := validate(result); err != nil {
        return errors.Wrap(err, "validate")
    }
    return save(result)
}
```
</examples>

<patterns>
- Error checks wrapping the main logic in nested `if err == nil` blocks
- Happy path code indented multiple levels due to error checks
- Functions where the success path is harder to read than the error path
- `if err != nil` followed by an else block containing the main logic
- Deeply nested error handling that should use guard clauses
</patterns>

<related>
early-return, if-return, error-return
