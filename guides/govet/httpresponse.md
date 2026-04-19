# govet: httpresponse

<instructions>
Detects HTTP response bodies that are not closed. After a successful `http.Get`, `http.Post`, or `http.Client.Do` call, the response body must be closed to release the underlying connection. Failing to close it leaks connections and can exhaust file descriptors.

Always `defer resp.Body.Close()` immediately after checking for errors.
</instructions>

<examples>
## Bad
```go
resp, err := http.Get(url)
if err != nil {
    return err
}
// resp.Body never closed — connection leak
data, _ := io.ReadAll(resp.Body)
return nil
```

## Good
```go
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
data, _ := io.ReadAll(resp.Body)
return nil
```
</examples>

<patterns>
- `http.Get`/`http.Post`/`http.Do` without `defer resp.Body.Close()`
- Closing body before reading it
- Returning without closing the response body on error paths
</patterns>

<related>
hostport, defers
</related>
