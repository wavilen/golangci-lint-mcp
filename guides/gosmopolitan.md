# gosmopolitan

<instructions>
Gosmopolitan detects code that assumes a single locale or encoding. It flags `time.Local` usage, locale-specific string operations, and non-UTC time zones that may behave differently across environments.

Use `time.UTC` explicitly. For user-facing times, accept a `time.Location` parameter rather than relying on server locale. Use `strings.EqualFold` instead of case conversions that depend on locale.
</instructions>

<examples>
## Good
```go
func schedule(loc *time.Location) time.Time {
    return time.Now().In(loc).Truncate(24 * time.Hour)
}
```
</examples>

<patterns>
- Use explicit `time.UTC` or pass `*time.Location` instead of relying on `time.Local`
- Avoid `strings.ToLower`/`ToUpper` on user-facing text without locale consideration
- Load timezone names from configuration rather than hardcoding them
- Use locale-aware formatting for dates presented to users
</patterns>

<related>
predeclared, asciicheck, godot
</related>
