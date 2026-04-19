# revive: unchecked-type-assertion

<instructions>
Detects type assertions without the comma-ok safety check. Writing `v.(string)` panics at runtime if the underlying type doesn't match. Always use the two-value form `v, ok := x.(string)` to handle the mismatch gracefully.

Add the comma-ok check and handle the `false` case with an error return, default value, or log message. For type switches, ensure a `default` case.
</instructions>

<examples>
## Bad
```go
name := val.(string)
handler := v.(http.Handler)
```

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
- Direct type assertions without comma-ok on `interface{}`
- Assertions on JSON unmarshaled values (which are `interface{}`)
- Type assertions on context values without checking
- Slice element assertions after `reflect` or `interface{}` conversion
- Assertions on interface types from external packages
</patterns>

<related>
unhandled-error, unreachable-code, forcetypeassert
