# noctx

<instructions>
Noctx detects HTTP requests made without a context.Context. Without context, requests cannot be cancelled, timed out, or carry trace information, which leads to hung connections and poor observability.

Use `http.NewRequestWithContext` instead of `http.NewRequest`, passing a context from the caller.
</instructions>

<examples>
## Bad
```go
req, err := http.NewRequest("GET", "https://example.com", nil)
resp, err := http.DefaultClient.Do(req)
```

## Good
```go
req, err := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)
resp, err := http.DefaultClient.Do(req)
```
</examples>

<patterns>
- `http.NewRequest` calls without context propagation
- `http.Get`/`http.Post` shorthand functions (no context support)
- Custom transport wrappers that drop context
- Long-running requests that cannot be cancelled on shutdown
</patterns>

<related>
contextcheck, bodyclose, errcheck
</related>
