# exhaustive

<instructions>
Checks that switch statements on enum types cover every defined member. Unhandled enum values silently fall through to `default`, causing runtime errors or undefined behavior that won't surface until production. Add a case for every enum member, or use an explicit `default` with a panic for truly unexpected values.
</instructions>

<examples>
## Good
```go
func handle(s Status) {
    switch s {
    case StatusNew:
        slog.Info("new")
    case StatusPending:
        slog.Info("pending")
    case StatusDone:
        slog.Info("done")
    default:
        panic(fmt.Sprintf("unhandled status: %d", s))
    }
}
```
</examples>

<patterns>
- Audit all `switch` statements on `iota`-based enum types when adding a new constant — add a case for each new value
- Include all interface implementations in `switch t.(type)` statements
- Run `exhaustive` after modifying `const` blocks to catch missing switch cases
</patterns>

<related>
exhaustruct, gochecksumtype, govet
</related>
