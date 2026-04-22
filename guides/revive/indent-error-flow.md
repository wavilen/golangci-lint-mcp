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
- Handle errors first with early return instead of wrapping main logic in `if err == nil` blocks
- Reduce indentation by returning on error before the happy path code
- Convert error checks so the success path is more readable than the error path
- Remove else blocks after `if err != nil` — the happy path continues below
- Replace deeply nested error handling with sequential guard clauses
</patterns>

<related>
early-return, if-return, error-return
