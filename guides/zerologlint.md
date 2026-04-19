# zerologlint

<instructions>
Zerologlint detects incorrect usage of the zerolog structured logging library. It finds chained log calls that silently drop events because `Send()` or `Msg()` is never called, resulting in log lines that are never written.

Always end zerolog chains with `.Send()` or `.Msg("message")` to emit the log event.
</instructions>

<examples>
## Bad
```go
log.Info().Str("key", "value").Int("count", 42)
```

## Good
```go
log.Info().Str("key", "value").Int("count", 42).Send()
```
</examples>

<patterns>
- Log chains ending with a field method instead of `Send()` or `Msg()`
- Conditional logging that builds the event but never sends
- Using `.Msg("")` instead of `.Send()` (works but semantically wrong)
- Chains split across multiple lines where the terminal call is missing
</patterns>

<related>
loggercheck, sloglint
</related>
