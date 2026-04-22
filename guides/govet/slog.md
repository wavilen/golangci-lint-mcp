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
- Provide key-value pairs in slog functions — always an even number of arguments
- Use string keys in all slog key-value pairs — wrap values with `"key", value`
- Pair every value with a string key in slog calls
- Ensure `slog.With` receives matching key-value pairs — no dangling keys
</patterns>

<related>
printf, stdmethods
</related>
