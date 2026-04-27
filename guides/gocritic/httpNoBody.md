# gocritic: httpNoBody

<instructions>
Detects uses of `nil` as the body in HTTP requests or responses where `http.NoBody` should be used instead. `nil` body works for requests, but `http.NoBody` is the explicit, canonical way to indicate no body and is required in some contexts (e.g., `httptest` expectations).

Replace `nil` body arguments with `http.NoBody`.
</instructions>

<examples>
## Good
```go
req, err := http.NewRequest("GET", url, http.NoBody)
resp := &http.Response{StatusCode: 200, Body: http.NoBody}
```
</examples>

<patterns>
- Replace `nil` body with `http.NoBody` in `http.NewRequest(method, url, http.NoBody)`
- Replace `nil` body with `http.NoBody` in `http.Response{Body: http.NoBody}`
- Replace `nil` body with `http.NoBody` in `http.NewRequestWithContext(ctx, method, url, http.NoBody)`
</patterns>

<related>
gocritic/wrapperFunc
</related>
