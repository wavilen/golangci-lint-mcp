# unused

<instructions>
Unused detects declared variables, constants, functions, and types that are never referenced anywhere in the codebase. Dead code increases maintenance burden and can indicate incomplete implementations or leftover refactoring artifacts.

Remove the unused declaration, or if intentionally kept for future use, add a comment explaining why.
</instructions>

<examples>
## Bad
```go
const defaultTimeout = 30 // never referenced

func helper() string { // never called
    return "help"
}
```

## Good
```go
const defaultTimeout = 30

func process() error {
    ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
    defer cancel()
    return doWork(ctx)
}
```
</examples>

<patterns>
- Unused constants and variables after refactoring
- Helper functions that were superseded but not removed
- Private struct types defined but never instantiated
- Import aliases that are never referenced
</patterns>

<related>
ineffassign, unparam, govet
