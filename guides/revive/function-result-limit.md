# revive: function-result-limit

<instructions>
Enforces a maximum number of return values per function (default is typically 3). Functions returning many values are hard to use correctly — the caller must remember the position and type of each value. This often indicates the function should return a struct instead.

Group related return values into a result struct. Keep the function focused on returning its primary result and an error.
</instructions>

<examples>
## Bad
```go
func ParseHeader(raw string) (proto string, version float64, headers map[string]string, body string, err error) {
    // 5 return values — easy to mix up at call site
    return proto, ver, hdrs, body, nil
}
```

## Good
```go
type HeaderResult struct {
    Proto   string
    Version float64
    Headers map[string]string
    Body    string
}

func ParseHeader(raw string) (HeaderResult, error) {
    return HeaderResult{Proto: proto, Version: ver, Headers: hdrs, Body: body}, nil
}
```
</examples>

<patterns>
- Functions returning many primitive values of the same type
- API functions returning status, data, metadata, and error separately
- Functions with named returns used as output parameters instead of a struct
- Gradual accumulation of return values as the function evolves
- Functions returning multiple error types alongside results
</patterns>

<related>
argument-limit, function-length, confusing-results
