# revive: time-naming

<instructions>
Enforces consistent naming for time-related variables. Variables of type `time.Time` should use a `Time` suffix or use the full word "time" (e.g., `startTime`, `createdAt`) rather than abbreviations like `ts`, `t`, or `tm` at package scope. Clear naming prevents confusion between timestamps and durations.

Use descriptive names with a `Time` suffix or time-related prefix (e.g., `startTime`, `deadlineTime`, `lastModified`).
</instructions>

<examples>
## Good
```go
var startTime time.Time
var now = time.Now()
```
</examples>

<patterns>
- Use descriptive names with `Time` suffix for package-level `time.Time` variables instead of `ts`, `tm`, `t`
- Rename time variables after what they represent (e.g., `startTime`) rather than the unit (e.g., `seconds`)
- Use consistent `Time` suffix or time-related naming across all time variables in a package
- Use distinct names for duration variables vs timestamp variables (e.g., `timeoutDuration`)
- Use `time.Time`-specific names instead of `date` when the variable holds a `time.Time`
</patterns>

<related>
revive/time-date, revive/time-equal, revive/var-naming
</related>
