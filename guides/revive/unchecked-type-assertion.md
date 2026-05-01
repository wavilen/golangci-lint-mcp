# revive: unchecked-type-assertion

<instructions>
Detects single-value type assertions `v.(T)` that panic at runtime when the underlying type doesn't match. A failed assertion crashes the goroutine with an unrecoverable `panic`, bypassing any error handling. Use the two-value form `v, ok := x.(T)` and handle the `false` case with an error return or default value.
</instructions>

<examples>
## Good
```go
name, ok := val.(string)
if !ok {
    return fmt.Errorf("expected string, got %T", val)
}

handler, ok := v.(http.Handler)
if !ok {
    return errors.New("not an http.Handler")
}
```
</examples>

<patterns>
- Use the comma-ok form `v, ok := x.(string)` for type assertions on `interface{}`
- Check the ok value for type assertions on JSON unmarshaled values (which are `interface{}`)
- Validate type assertions on context values with the comma-ok pattern
- Use comma-ok for slice element assertions after `reflect` or `interface{}` conversion
- Handle the false case for type assertions on interface types from external packages
</patterns>

<related>
revive/unhandled-error, revive/unreachable-code, forcetypeassert
</related>
