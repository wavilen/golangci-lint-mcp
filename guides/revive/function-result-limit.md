# revive: function-result-limit

<instructions>
Enforces a maximum number of return values per function (default is typically 3). Functions returning many values are hard to use correctly — the caller must remember the position and type of each value. This often indicates the function should return a struct instead.

Group related return values into a result struct. Keep the function focused on returning its primary result and an error.
</instructions>

<examples>
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
revive/argument-limit, revive/function-length, revive/confusing-results
</related>
