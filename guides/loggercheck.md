# loggercheck

<instructions>
Loggercheck detects incorrect use of structured logging libraries (slog, zap, zerolog, logr). It flags key-value pair mismatches, non-string keys, and odd argument counts that cause silent data loss or garbled output.

Ensure logging calls have even numbers of key-value arguments and use string keys.
</instructions>

<examples>
## Bad
```go
slog.Info("request completed",
    "method", req.Method,
    "status", resp.StatusCode,
    "duration", // missing value
)
```

## Good
```go
slog.Info("request completed",
    "method", req.Method,
    "status", resp.StatusCode,
    "duration", elapsed,
)
```
</examples>

<patterns>
- Odd number of key-value arguments in structured log calls
- Using non-string types as log field keys (e.g., int keys)
- Mixing positional and key-value arguments in slog/zap calls
- Passing structs directly instead of key-value pairs to structured loggers
</patterns>

<related>
sloglint, zerologlint, sloglint
