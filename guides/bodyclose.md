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
- Forgetting `defer resp.Body.Close()` after http.Client requests
- Early returns before the body close statement
- Only closing the body in error branches, not the success path
- Using http.Get/Post without a deferred close on the response body
</patterns>

<related>
noctx, sqlclosecheck, rowserrcheck
