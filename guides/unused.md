# unused

<instructions>
Unused detects declared variables, constants, functions, and types that are never referenced anywhere in the codebase. Dead code increases maintenance burden and can indicate incomplete implementations or leftover refactoring artifacts.

Remove the unused declaration, or if intentionally kept for future use, add a comment explaining why.
</instructions>

<examples>
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
- Remove constants and variables left unused after refactoring
- Delete helper functions that were superseded by newer implementations
- Remove private struct types that are defined but never used
- Remove import aliases that are never referenced in code
</patterns>

<related>
ineffassign, unparam, govet
</related>
