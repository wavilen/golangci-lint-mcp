# unqueryvet

<instructions>
Unqueryvet detects struct types that implement the `url.Queryer` or `url.ValuesQueryer` interface but return ignored error values, or types that implement the interface incorrectly. Misimplementations lead to silent failures when encoding structs as URL query parameters.

Ensure `Query()` methods return `(url.Values, error)` and always handle the error.
</instructions>

<examples>
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
- Handle the error return value from `Query()` method calls
- Match the `url.Queryer` interface signature exactly: `Query() (url.Values, error)`
- Return an error from `Query()` implementations instead of only returning values
- Check the error before calling `Encode()` on values returned by `Query()`
</patterns>

<related>
errcheck, musttag
</related>
