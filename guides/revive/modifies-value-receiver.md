# revive: modifies-value-receiver

<instructions>
Detects methods with value receivers that modify the receiver. Since value receivers receive a copy, modifications are discarded when the method returns. This is almost always a bug — the developer intended to modify the original but used the wrong receiver type.

Change the receiver to a pointer (`*T`) if mutation is intended, or remove the mutation if the value receiver was intentional.
</instructions>

<examples>
## Good
```go
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++ // modifies the original
}
```
</examples>

<patterns>
- Switch to a pointer receiver `*T` when the method needs to modify fields visible to the caller
- Switch to pointer receiver when methods increment counters or append to slices
- Use pointer receiver for builder-pattern methods that must chain on the original instance
- Add the `*` to receiver declarations copied from pointer receiver methods
- Use pointer receiver consistently for types where any method modifies the receiver
</patterns>

<related>
revive/modifies-parameter, revive/get-return
</related>
