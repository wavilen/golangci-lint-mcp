# revive: function-length

<instructions>
Enforces a maximum number of lines or statements per function body. Long functions are hard to understand, test, and maintain. They often indicate the function has too many responsibilities and should be decomposed.

Extract logical sections into helper functions with descriptive names. Each function should do one thing well.
</instructions>

<examples>
## Bad
```go
func ProcessOrder(o Order) error {
    // 80 lines of validation, calculation, persistence,
    // notification, and logging mixed together...
    return nil
}
```

## Good
```go
func ProcessOrder(o Order) error {
    if err := validateOrder(o); err != nil {
        return errors.Wrap(err, "validate")
    }
    total := calculateTotal(o)
    if err := persistOrder(o, total); err != nil {
        return errors.Wrap(err, "persist")
    }
    return notifyCustomer(o)
}
```
</examples>

<patterns>
- Functions mixing validation, business logic, and I/O
- Long switch-case functions where each case is substantial
- Handler functions with inline parsing, validation, and response building
- Functions that grew over time as edge cases were added
- Test helper functions that set up complex state inline
</patterns>

<related>
cyclomatic, cognitive-complexity, function-result-limit, funlen
