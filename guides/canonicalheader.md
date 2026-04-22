# canonicalheader

<instructions>
Canonicalheader detects uses of non-canonical HTTP header keys. The Go `http.Header` type canonicalizes keys (e.g., `content-type` → `Content-Type`), so using non-canonical forms in Get/Set/Del calls works but is inconsistent and can mask bugs.

Use the canonical header key form (Title-Case with hyphens) when accessing headers.
</instructions>

<examples>
## Bad
```go
ct := resp.Header.Get("content-type")
resp.Header.Set("x-request-id", id)
```

## Good
```go
ct := resp.Header.Get("Content-Type")
resp.Header.Set("X-Request-Id", id)
```
</examples>

<patterns>
- Use canonical header keys with `http.Header.Get` — Go normalizes via `textproto.CanonicalMIMEHeaderKey`
- Replace all-caps keys like `"CONTENT-TYPE"` with canonical form `"Content-Type"`
- Set response headers using canonical key forms to ensure consistent access
- Use canonical key form consistently when accessing custom headers across all handlers
</patterns>

<related>
bodyclose, noctx, nosprintfhostport
