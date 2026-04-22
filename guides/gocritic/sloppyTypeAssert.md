# gocritic: sloppyTypeAssert

<instructions>
Detects type assertions without the comma-ok safety check. A plain `v := i.(T)` panics if `i` does not hold type `T`. The safe form `v, ok := i.(T)` returns a boolean indicating success.

Use the comma-ok form for type assertions unless a panic on wrong type is truly desired (rare). For type switches, the assertion is safe and this check does not apply.
</instructions>

<examples>
## Bad
```go
func process(v interface{}) {
    s := v.(string) // panics if v is not a string
    slog.Info("value", "s", s)
}
```

## Good
```go
func process(v interface{}) {
    s, ok := v.(string)
    if !ok {
        return fmt.Errorf("expected string, got %T", v)
    }
    slog.Info("value", "s", s)
}
```
</examples>

<patterns>
- Use comma-ok form `v, ok := i.(Type)` for type assertions in general code paths
- Guard type assertions on `interface{}` from external sources with comma-ok
- Guard unmarshaled JSON `interface{}` values with comma-ok before asserting type
- Guard `context.Value()` results with comma-ok before type assertion
</patterns>

<related>
sloppyReassign, sloppyLen, badCond
</related>
