# unqueryvet

<instructions>
Unqueryvet detects struct types that implement the `url.Queryer` or `url.ValuesQueryer` interface but return ignored error values, or types that implement the interface incorrectly. Misimplementations lead to silent failures when encoding structs as URL query parameters.

Ensure `Query()` methods return `(url.Values, error)` and always handle the error.
</instructions>

<examples>
## Bad
```go
func (f *Filter) Query() (url.Values, error) {
    v := url.Values{}
    v.Set("status", string(f.Status))
    return v, nil
    // error never checked by callers
}
```

## Good
```go
func (f *Filter) Query() (url.Values, error) {
    v := url.Values{}
    v.Set("status", string(f.Status))
    return v, nil
}

values, err := filter.Query()
if err != nil {
    return errors.Wrap(err, "building query")
}
```
</examples>

<patterns>
- `Query()` method returns values but callers ignore the error
- Incorrect `Query()` signature not matching `url.Queryer`
- Missing error return in `Query()` implementation
- Calling `Encode()` on values from `Query()` without checking error
</patterns>

<related>
errcheck, musttag
</related>
