# revive: modifies-value-receiver

<instructions>
Detects methods with value receivers that modify the receiver. Since value receivers receive a copy, modifications are discarded when the method returns. This is almost always a bug — the developer intended to modify the original but used the wrong receiver type.

Change the receiver to a pointer (`*T`) if mutation is intended, or remove the mutation if the value receiver was intentional.
</instructions>

<examples>
## Bad
```go
type Counter struct {
    count int
}

func (c Counter) Increment() {
    c.count++ // modifies a copy — original unchanged
}
```

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
- Value receiver methods that set fields, expecting the caller to see changes
- Methods that increment counters or append to slices on a value receiver
- Builder-pattern methods on value receivers that should chain on the original
- Copy-paste from pointer receiver methods where the `*` was dropped
- Methods that only sometimes modify — should consistently use pointer receiver
</patterns>

<related>
modifies-parameter, get-return
