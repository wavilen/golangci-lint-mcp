# revive: time-date

<instructions>
Detects calls to `time.Date` that could be replaced with `time.Now`. When all arguments to `time.Date` match the current time components, it's an overly verbose way to get the current time. More commonly, `time.Date` is used when `time.Now()` or `time.Parse` would be clearer.

Use `time.Now()` to get the current time. Use `time.Parse` with a layout string to construct a specific time from a string representation.
</instructions>

<examples>
## Bad
```go
now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
```

## Good
```go
now := time.Now().Truncate(24 * time.Hour)
```
</examples>

<patterns>
- Constructing the current date/time manually via `time.Date`
- Using `time.Date` with `time.Now()` components instead of calling `time.Now()` directly
- Complex date arithmetic that could use `time.Now().Add` or `Truncate`
- Building timestamps from individual fields parsed from the current time
- Using `time.Date` for date-only values when `time.Now().Date()` suffices
</patterns>

<related>
time-equal, time-naming
