# gocritic: httpNoBody

<instructions>
Detects uses of `nil` as the body in HTTP requests or responses where `http.NoBody` should be used instead. `nil` body works for requests, but `http.NoBody` is the explicit, canonical way to indicate no body and is required in some contexts (e.g., `httptest` expectations).

Replace `nil` body arguments with `http.NoBody`.
</instructions>

<examples>
## Bad
```go
req, err := http.NewRequest("GET", url, nil)
resp := &http.Response{StatusCode: 200, Body: nil}
```

## Good
```go
req, err := http.NewRequest("GET", url, http.NoBody)
resp := &http.Response{StatusCode: 200, Body: http.NoBody}
```
</examples>

<patterns>
- `http.NewRequest(method, url, nil)` → use `http.NoBody`
- `http.Response{Body: nil}` → use `http.NoBody`
- `http.NewRequestWithContext(ctx, method, url, nil)` → use `http.NoBody`
</patterns>

<related>
wrapperFunc
</related>
