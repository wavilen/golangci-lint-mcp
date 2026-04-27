# gocritic: syncMapLoadAndDelete

<instructions>
Detects separate `sync.Map.Load` followed by `sync.Map.Delete` calls that should be combined into a single `sync.Map.LoadAndDelete` call. Using two separate operations is not atomic — the value may change between the load and delete. `LoadAndDelete` performs both operations atomically.

Replace the separate `Load` + `Delete` with a single `LoadAndDelete` call.
</instructions>

<examples>
## Good
```go
if v, loaded := m.LoadAndDelete(key); loaded {
    return v.(int)
}
```
</examples>

<patterns>
- Replace `sync.Map.Load` + `sync.Map.Delete` on the same key with `sync.Map.LoadAndDelete`
- Use `sync.Map.LoadAndDelete` instead of `LoadOrStore` when deletion is intended
- Replace separate check-and-delete patterns on `sync.Map` with atomic `LoadAndDelete`
- Use `LoadAndDelete` to avoid race conditions between read and remove operations
</patterns>

<related>
gocritic/badLock, gocritic/badSyncOnceFunc
</related>
