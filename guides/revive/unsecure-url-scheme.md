# revive: unsecure-url-scheme

<instructions>
Detects HTTP (non-HTTPS) URL literals in code. Using plain HTTP for API calls, redirects, or resource fetching exposes data to interception and man-in-the-middle attacks. Modern services should use HTTPS exclusively.

Replace `http://` URLs with `https://` equivalents. If the service genuinely requires HTTP (e.g., localhost during development), add a `//nolint` comment with justification.
</instructions>

<examples>
## Good
```go
const apiURL = "https://api.example.com/v1/users"
resp, err := http.Get("https://metadata.internal/role")
```
</examples>

<patterns>
- Replace hardcoded `http://` URLs for external services with `https://`
- Use HTTPS in configuration defaults and constants instead of HTTP
- Replace HTTP URLs in test fixtures for production-like services with HTTPS
- Replace `http://` in redirect logic with `https://`
- Use HTTPS in documentation examples to prevent insecure copy-paste patterns
</patterns>

<related>
revive/unhandled-error, revive/datarace
</related>
