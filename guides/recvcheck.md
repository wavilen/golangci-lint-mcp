# recvcheck

<instructions>
Recvcheck detects receiver type inconsistencies — when the same type has some methods with pointer receivers and others with value receivers. Mixing receiver types can cause confusion about whether an interface is satisfied and whether mutations are visible.

Use pointer receivers for all methods if any method needs a pointer receiver, or use value receivers consistently for immutable types.
</instructions>

<examples>
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
- Use pointer receivers consistently when any method mutates state
- Align all methods on a type to use the same receiver kind (pointer or value)
- Unify mixed receiver styles inherited from different authors into a consistent pattern
</patterns>

<related>
gocritic, revive, govet
</related>
