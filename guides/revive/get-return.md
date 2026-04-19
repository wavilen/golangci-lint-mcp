# revive: get-return

<instructions>
Enforces that getter methods return the expected type and don't perform unexpected side effects. A method named `GetX` should return `X` (or `X, error`) and not modify state, return a boolean status, or perform I/O. Getters should be simple field accessors.

Rename the method if it does more than getting a value. If it performs computation, use a verb like `Calculate`, `Fetch`, or `Load` instead of `Get`.
</instructions>

<examples>
## Bad
```go
func (c *Cache) GetSize() error {
    c.computeSize() // modifies state — not a simple getter
    return c.lastError
}
```

## Good
```go
func (c *Cache) Size() int {
    return c.size
}

func (c *Cache) RecalculateSize() error {
    c.computeSize()
    return c.lastError
}
```
</examples>

<patterns>
- Getter methods that perform side effects like logging or metrics
- `Get` methods that return an error instead of the value
- Methods named `GetX` that fetch from a database or network
- Getters that mutate internal state (lazy initialization)
- Methods named `Get` that return a boolean status instead of the value
</patterns>

<related>
modifies-value-receiver, exported
