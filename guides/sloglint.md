# sloglint

<instructions>
Sloglint enforces consistent usage of the `log/slog` structured logging package. It detects mixed logging styles (key-value pairs vs attributes), missing or duplicate keys, and ensures consistent argument passing.

Use a single style consistently: either key-value pairs (`slog.Info("msg", "key", val)`) or typed attributes (`slog.Info("msg", slog.String("key", val))`).
</instructions>

<examples>
## Good
```go
slog.Info("request failed",
    slog.Int("status", resp.StatusCode),
    slog.String("path", req.URL.Path),
)
```
</examples>

<patterns>
- Use either key-value pairs or slog attributes consistently within a single call
- Eliminate duplicate keys in a single log call
- Wrap values in type-safe `slog.String`/`slog.Int` attributes instead of passing raw values
- Adopt a single key naming convention (snake_case or camelCase) across all log calls
</patterns>

<related>
loggercheck
</related>
