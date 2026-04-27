# loggercheck

<instructions>
Loggercheck detects incorrect use of structured logging libraries (slog, zap, zerolog, logr). It flags key-value pair mismatches, non-string keys, and odd argument counts that cause silent data loss or garbled output.

Ensure logging calls have even numbers of key-value arguments and use string keys.
</instructions>

<examples>
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
- Ensure an even number of key-value arguments in structured log calls
- Use string keys only for log field names — never int or other types
- Separate positional arguments from key-value pairs in slog/zap calls
- Decompose structs into key-value pairs when passing to structured loggers
</patterns>

<related>
sloglint, zerologlint
</related>
