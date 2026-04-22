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
- Extract validation, business logic, and I/O into separate helper functions
- Move each substantial switch case into its own function
- Separate handler functions into parsing, validation, and response building helpers
- Decompose functions that grew over time by extracting edge cases into focused functions
- Extract inline state setup from test helper functions into dedicated setup functions
</patterns>

<related>
cyclomatic, cognitive-complexity, function-result-limit, funlen
