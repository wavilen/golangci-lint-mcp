# exhaustive

<instructions>
Exhaustive checks that switch statements on enum types cover all possible cases. Missing cases mean unhandled enum values silently fall through to the default, potentially causing runtime errors.

Add a case for every enum member, or use a `default` case explicitly if that is the intended behavior.
</instructions>

<examples>
## Bad
```go
type Status int

const (
    StatusNew Status = iota
    StatusPending
    StatusDone
)

func handle(s Status) {
    switch s {
    case StatusNew:
        slog.Info("new")
    case StatusDone:
        slog.Info("done")
    }
}
```

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
- Switch on enum types missing newly added enum members
- Type switch statements missing interface implementations
- Package-level enum switches scattered across the codebase
</patterns>

<related>
exhaustruct, gochecksumtype, govet
</related>
