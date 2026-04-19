# recvcheck

<instructions>
Recvcheck detects receiver type inconsistencies — when the same type has some methods with pointer receivers and others with value receivers. Mixing receiver types can cause confusion about whether an interface is satisfied and whether mutations are visible.

Use pointer receivers for all methods if any method needs a pointer receiver, or use value receivers consistently for immutable types.
</instructions>

<examples>
## Bad
```go
type Counter struct {
    count int
}

func (c Counter) Value() int {
    return c.count
}

func (c *Counter) Increment() {
    c.count++
}
```

## Good
```go
type Counter struct {
    count int
}

func (c *Counter) Value() int {
    return c.count
}

func (c *Counter) Increment() {
    c.count++
}
```
</examples>

<patterns>
- Structs with value-receiver getters and pointer-receiver mutators
- Types that implement interfaces with value receivers but have pointer-receiver methods
- Mixed receiver styles inherited from different authors or code generations
</patterns>

<related>
gocritic, revive, govet
</related>
