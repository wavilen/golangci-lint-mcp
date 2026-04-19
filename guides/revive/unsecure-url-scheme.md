# revive: unsecure-url-scheme

<instructions>
Detects HTTP (non-HTTPS) URL literals in code. Using plain HTTP for API calls, redirects, or resource fetching exposes data to interception and man-in-the-middle attacks. Modern services should use HTTPS exclusively.

Replace `http://` URLs with `https://` equivalents. If the service genuinely requires HTTP (e.g., localhost during development), add a `//nolint` comment with justification.
</instructions>

<examples>
## Bad
```go
const apiURL = "http://api.example.com/v1/users"
resp, err := http.Get("http://metadata.internal/role")
```

## Good
```go
const apiURL = "https://api.example.com/v1/users"
resp, err := http.Get("https://metadata.internal/role")
```
</examples>

<patterns>
- Hardcoded `http://` URLs for external services or APIs
- HTTP URLs in configuration defaults or constants
- HTTP URLs in test fixtures meant for production-like services
- String literals containing `http://` in redirect logic
- HTTP URLs in documentation examples that get copy-pasted
</patterns>

<related>
unhandled-error, datarace
