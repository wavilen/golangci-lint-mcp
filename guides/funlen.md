# funlen

<instructions>
Funlen checks that functions don't exceed a maximum number of lines (default 60) or statements (default 40). Long functions are hard to understand, test, and maintain.

Break long functions into smaller, named functions that each do one thing.
</instructions>

<examples>
## Bad
```go
func ProcessOrder(o Order) error {
    if o.ID == "" { return errors.New("missing id") }
    if o.Customer == "" { return errors.New("missing customer") }
    if len(o.Items) == 0 { return errors.New("no items") }
    for _, item := range o.Items {
        if item.Quantity <= 0 { return errors.New("invalid qty") }
    }
    var total float64
    for _, item := range o.Items {
        total += float64(item.Quantity) * item.Price
    }
    total *= 1.08 // tax
    return saveOrder(o, total)
}
```

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
- God functions handling validation, business logic, and persistence in one body
- Request handlers that do routing, auth, and response formatting
- Setup functions with sequential configuration blocks
</patterns>

<related>
cyclop, gocyclo, gocognit, nestif
</related>
