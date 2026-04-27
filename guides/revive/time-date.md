# revive: time-date

<instructions>
Detects calls to `time.Date` that could be replaced with `time.Now`. When all arguments to `time.Date` match the current time components, it's an overly verbose way to get the current time. More commonly, `time.Date` is used when `time.Now()` or `time.Parse` would be clearer.

Use `time.Now()` to get the current time. Use `time.Parse` with a layout string to construct a specific time from a string representation.
</instructions>

<examples>
## Good
```go
now := time.Now().Truncate(24 * time.Hour)
```
</examples>

<patterns>
- Use `time.Now()` instead of constructing the current time manually via `time.Date`
- Replace `time.Date` with `time.Now()` components by calling `time.Now()` directly
- Simplify date arithmetic with `time.Now().Add` or `Truncate` instead of `time.Date`
- Use `time.Parse` to build timestamps from strings instead of assembling individual fields
- Use `time.Now().Date()` for date-only values instead of `time.Date`
</patterns>

<related>
revive/time-equal, revive/time-naming
</related>
