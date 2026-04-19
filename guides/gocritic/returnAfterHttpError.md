# gocritic: returnAfterHttpError

<instructions>
Detects missing `return` after writing an HTTP error response. After calling `http.Error(w, msg, code)`, the handler should return immediately. Continuing execution after sending an error response leads to duplicate writes to the response writer, which causes a superfluous `http: superfluous response.WriteHeader call` warning or corrupted responses.

Add `return` immediately after `http.Error()` calls to stop handler execution.
</instructions>

<examples>
## Bad
```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        // execution continues — may write another response
    }
    w.Write([]byte("ok"))
}
```

## Good
```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("ok"))
}
```
</examples>

<patterns>
- `http.Error()` without subsequent `return`
- Error logging after response already written
- Continuing processing after writing error status code
- Multiple `w.WriteHeader()` calls in same handler
</patterns>

<related>
exitAfterDefer, badCall
</related>
