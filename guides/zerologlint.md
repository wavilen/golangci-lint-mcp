# zerologlint

<instructions>
Zerologlint detects incorrect usage of the zerolog structured logging library. It finds chained log calls that silently drop events because `Send()` or `Msg()` is never called, resulting in log lines that are never written.

Always end zerolog chains with `.Send()` or `.Msg("message")` to emit the log event.
</instructions>

<examples>
## Good
```go
log.Info().Str("key", "value").Int("count", 42).Send()
```
</examples>

<patterns>
- End every zerolog chain with `.Send()` or `.Msg("message")` to emit the log event
- Ensure conditional logging branches always call `.Send()` or `.Msg()`
- Replace `.Msg("")` with `.Send()` for event emission without a message
- Complete chains split across multiple lines with a terminal `.Send()` or `.Msg()`
</patterns>

<related>
loggercheck
</related>
