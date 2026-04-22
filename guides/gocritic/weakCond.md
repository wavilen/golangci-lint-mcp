# gocritic: weakCond

<instructions>
Detects weak or insufficient conditional checks that may not fully guard against the intended issue. Common examples include checking only one field of a struct when multiple fields should be validated, or using `!= nil` checks that don't prevent all error paths.

Strengthen the condition to cover all necessary cases, or add separate validation for each field.
</instructions>

<examples>
## Bad
```go
if err != nil || resp == nil { // weak — resp.StatusCode not checked
    return errors.New("request failed")
}
```

## Good
```go
if err != nil {
    return errors.Wrap(err, "request failed")
}
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("unexpected status: %d", resp.StatusCode)
}
```
</examples>

<patterns>
- Validate response status after checking `err != nil` — both must be verified
- Validate the pointed-to data after checking for non-nil pointer — not just the pointer
- Validate all struct fields in the conditional — avoid partial validation
- Guard nil map access with proper nil check — not just `len(s) > 0`
</patterns>

<related>
badCond, dupSubExpr, offBy1
</related>
