# gocritic: ptrToRefParam

<instructions>
Detects function parameters of type `*T` (pointer) that should be `T` (value) when the pointer is never modified through the parameter. If the function only reads the value, passing by value is safer and clearer.

Use value receivers and parameters unless the function needs to modify the caller's copy.
</instructions>

<examples>
## Bad
```go
func greet(name *string) {
	slog.Info("greeting", "name", *name)
}
```

## Good
```go
func greet(name string) {
	slog.Info("greeting", "name", name)
}
```
</examples>

<patterns>
- Replace pointer parameters used only for reading with value parameters
- Replace value receivers on non-mutating methods with pointer receivers only when needed
- Replace `*T` parameters with `T` when the function never writes through the pointer
</patterns>

<related>
underef, exposedSyncMutex
</related>
