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
- Parameters added for future use but never referenced
- Interface implementation parameters not needed by this specific method
- Callback functions where only some parameters are relevant
- Refactored functions where parameters were removed from the body but not the signature
- Constructor parameters stored elsewhere and no longer needed locally
</patterns>

<related>
unused-receiver, unnecessary-stmt, context-as-argument
