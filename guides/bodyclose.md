# bodyclose

<instructions>
Bodyclose detects HTTP response bodies that are not closed. Unclosed bodies leak file descriptors and TCP connections, eventually exhausting connection pools and causing failures under load.

Always defer `resp.Body.Close()` immediately after checking the response error.
</instructions>

<examples>
## Bad
```go
resp, err := http.Get("https://example.com/api")
if err != nil {
    return err
}
data, _ := io.ReadAll(resp.Body)
// resp.Body never closed — FD leak
```

## Good
```go
resp, err := http.Get("https://example.com/api")
if err != nil {
    return err
}
defer resp.Body.Close()
data, _ := io.ReadAll(resp.Body)
```
</examples>

<patterns>
- Always `defer resp.Body.Close()` immediately after `http.Client.Do()` requests
- Ensure body close runs before any early return — place `defer` before conditional logic
- Close the response body in all code paths, not just error branches
- Add `defer resp.Body.Close()` after every `http.Get`/`http.Post` call
</patterns>

<related>
noctx, sqlclosecheck, rowserrcheck
