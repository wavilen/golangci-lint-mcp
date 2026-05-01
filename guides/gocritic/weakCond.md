# gocritic: weakCond

<instructions>
Detects conditional checks that don't fully guard against the intended issue — validating only one struct field, checking error without checking status, or using conditions with logical gaps. Incomplete guards let invalid data propagate to downstream code, causing failures far from the root cause. Strengthen the condition to cover all necessary cases, or add separate validation for each field.
</instructions>

<examples>
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
- Check all relevant struct fields in guard conditions — e.g., validate both `resp.StatusCode` and `err` from `http.Do`, not just one
- Guard nil map access with proper nil check — not just `len(s) > 0`
</patterns>

<related>
gocritic/badCond, gocritic/dupSubExpr, gocritic/offBy1
</related>
