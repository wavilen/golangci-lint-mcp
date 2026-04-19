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
- `i.(Type)` without comma-ok in general code paths
- Type assertions on `interface{}` values from external sources
- Unmarshaled JSON decoded as `interface{}` with direct assertions
- Assertion on values from `context.Value()` without safety check
</patterns>

<related>
sloppyReassign, sloppyLen, badCond
</related>
