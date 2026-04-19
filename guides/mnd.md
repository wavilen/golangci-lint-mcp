# mnd

<instructions>
MND (Magic Number Detector) finds magic numbers in code — numeric literals used directly in expressions without explanation. These obscure intent and make future changes error-prone.

Extract the number into a named constant or variable. Use a descriptive name that explains what the value represents and why that specific number is chosen.
</instructions>

<examples>
## Bad
```go
func totalPrice(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Price * 1.08
    }
    return total
}
```

## Good
```go
const taxRate = 1.08

func totalPrice(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Price * taxRate
    }
    return total
}
```
</examples>

<patterns>
- Tax rates, multipliers, or conversion factors used inline
- Time durations as raw numbers without `time.Duration` constants
- Buffer sizes or limits hardcoded in function bodies
- Array indices or slice bounds using magic numbers
</patterns>

<related>
goconst, gochecknoglobals, varnamelen
</related>
