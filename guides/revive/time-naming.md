# revive: time-naming

<instructions>
Enforces consistent naming for time-related variables. Variables of type `time.Time` should use a `Time` suffix or use the full word "time" (e.g., `startTime`, `createdAt`) rather than abbreviations like `ts`, `t`, or `tm` at package scope. Clear naming prevents confusion between timestamps and durations.

Use descriptive names with a `Time` suffix or time-related prefix (e.g., `startTime`, `deadlineTime`, `lastModified`).
</instructions>

<examples>
## Bad
```go
var ts time.Time
var tm = time.Now()
```

## Good
```go
var startTime time.Time
var now = time.Now()
```
</examples>

<patterns>
- Short cryptic names like `ts`, `tm`, `t` for package-level time variables
- Time variables named after the unit (e.g., `seconds` for a `time.Time`)
- Inconsistent naming: some with `Time` suffix, some without
- Duration variables named like timestamps (e.g., `timeout` for a `time.Duration`)
- Variables named `date` when they hold a `time.Time` (misleading)
</patterns>

<related>
time-date, time-equal, var-naming
