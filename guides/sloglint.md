# sloglint

<instructions>
Sloglint enforces consistent usage of the `log/slog` structured logging package. It detects mixed logging styles (key-value pairs vs attributes), missing or duplicate keys, and ensures consistent argument passing.

Use a single style consistently: either key-value pairs (`slog.Info("msg", "key", val)`) or typed attributes (`slog.Info("msg", slog.String("key", val))`).
</instructions>

<examples>
## Bad
```go
slog.Info("request failed", "status", resp.StatusCode, slog.String("path", req.URL.Path))
```

## Good
```go
slog.Info("request failed",
    slog.Int("status", resp.StatusCode),
    slog.String("path", req.URL.Path),
)
```
</examples>

<patterns>
- Mixing key-value pairs and slog attributes in one call
- Duplicate keys in a single log call
- Passing raw values without type-safe wrappers
- Inconsistent key naming conventions (camelCase vs snake_case)
</patterns>

<related>
loggercheck, godot
</related>
