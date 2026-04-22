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
- Group many same-type return values into a result struct
- Wrap status, data, and metadata into a single response struct instead of returning them separately
- Replace named returns used as output parameters with a result struct
- Combine gradually accumulated return values into a struct as the function evolves
- Combine multiple error and result returns into a structured result type
</patterns>

<related>
argument-limit, function-length, confusing-results
