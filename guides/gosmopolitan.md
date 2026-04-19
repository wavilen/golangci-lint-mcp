# gosmopolitan

<instructions>
Gosmopolitan detects code that assumes a single locale or encoding. It flags `time.Local` usage, locale-specific string operations, and non-UTC time zones that may behave differently across environments.

Use `time.UTC` explicitly. For user-facing times, accept a `time.Location` parameter rather than relying on server locale. Use `strings.EqualFold` instead of case conversions that depend on locale.
</instructions>

<examples>
## Bad
```go
func schedule() time.Time {
    return time.Now().Truncate(24 * time.Hour)
}
```

## Good
```go
func schedule(loc *time.Location) time.Time {
    return time.Now().In(loc).Truncate(24 * time.Hour)
}
```
</examples>

<patterns>
- `time.Local` used implicitly via `time.Now()` without timezone awareness
- `strings.ToLower`/`ToUpper` on user-facing text without locale consideration
- Hardcoded timezone names instead of loading from configuration
- Date formatting with locale-sensitive format strings
</patterns>

<related>
predeclared, asciicheck, godot
</related>
