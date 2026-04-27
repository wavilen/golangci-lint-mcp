# funlen

<instructions>
Funlen checks that functions don't exceed a maximum number of lines (default 60) or statements (default 40). Long functions are hard to understand, test, and maintain.

Break long functions into smaller, named functions that each do one thing.
</instructions>

<examples>
## Good
```go
func ProcessOrder(o Order) error {
    if err := validateOrder(o); err != nil {
        return err
    }
    total := calculateTotal(o)
    return saveOrder(o, total)
}
```
</examples>

<patterns>
- Decompose god functions handling validation, business logic, and persistence into separate named functions
- Split request handlers that do routing, auth, and response formatting into focused handlers
- Extract setup functions with sequential configuration blocks into smaller, composable setup steps
</patterns>

<related>
cyclop, gocyclo, gocognit, nestif, revive/function-length
</related>
