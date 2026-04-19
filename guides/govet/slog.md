# govet: slog

<instructions>
Reports invalid arguments to `slog.Info`, `slog.Warn`, `slog.Error`, `slog.LogAttrs`, and related functions. Common issues include missing key-value pairs (odd argument count), non-string keys, or invalid attribute types.

Always pass key-value pairs: `slog.Info("msg", "key", value, "key2", value2)`.
</instructions>

<examples>
## Bad
```go
slog.Info("user logged in", userID, "admin") // first arg not a string key
```

## Good
```go
slog.Info("user logged in", "user_id", userID, "role", "admin")
```
</examples>

<patterns>
- Odd number of arguments to slog functions (missing value)
- Non-string keys in key-value pairs
- Passing raw values without key strings
- Using `slog.With` with mismatched pairs
</patterns>

<related>
printf, stdmethods
</related>
