# revive: unused-parameter

<instructions>
Detects function parameters that are never referenced in the function body. Unused parameters signal incomplete implementations, leftover refactoring, or interface implementations where the parameter isn't needed. They add noise to the function signature and may confuse callers.

Remove the parameter if the signature is under your control. For interface implementations, prefix the name with `_` (e.g., `_ context.Context`) to signal intentional non-use.
</instructions>

<examples>
## Bad
```go
func Process(ctx context.Context, data []byte, mode string) error {
    return handle(data) // ctx and mode unused
}
```

## Good
```go
func Process(_ context.Context, data []byte, _ string) error {
    return handle(data)
}
```
</examples>

<patterns>
- Remove parameters added for future use that are never referenced, or add `_` prefix
- Add `_` prefix to unused interface implementation parameters to signal intentional non-use
- Use `_` for callback function parameters where only some are relevant
- Remove unused parameters from refactored function signatures
- Remove constructor parameters stored elsewhere that are no longer needed locally
</patterns>

<related>
unused-receiver, unnecessary-stmt, context-as-argument
