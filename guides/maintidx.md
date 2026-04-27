# maintidx

<instructions>
Maintidx computes the Maintainability Index — a composite of cyclomatic complexity, Halstead volume, and lines of code. Low scores indicate expensive-to-maintain code. Default threshold is 20.

Improve scores by extracting helpers, simplifying conditionals, and replacing magic numbers with named constants.
</instructions>

<examples>
## Good
```go
func calc(p Price, qty int, discount, taxRate float64) float64 {
    subtotal := p.Base * float64(qty)
    subtotal = applyDiscount(subtotal, discount)
    subtotal = applyTax(subtotal, taxRate)
    subtotal = applyBulkDiscounts(subtotal)
    return subtotal
}
```
</examples>

<patterns>
- Extract business rules, calculations, and discount logic into separate functions
- Simplify data-processing functions by extracting each conditional modifier into its own function
- Reduce complexity by replacing inline operator chains with named helper functions
</patterns>

<related>
gocyclo, gocognit, cyclop, funlen
</related>
