# gocritic: syncMapLoadAndDelete

<instructions>
Detects separate `sync.Map.Load` followed by `sync.Map.Delete` calls that should be combined into a single `sync.Map.LoadAndDelete` call. Using two separate operations is not atomic — the value may change between the load and delete. `LoadAndDelete` performs both operations atomically.

Replace the separate `Load` + `Delete` with a single `LoadAndDelete` call.
</instructions>

<examples>
## Bad
```go
if v, ok := m.Load(key); ok {
    m.Delete(key) // not atomic with Load
    return v.(int)
}
```

## Good
```go
if v, loaded := m.LoadAndDelete(key); loaded {
    return v.(int)
}
```
</examples>

<patterns>
- `sync.Map.Load` + `sync.Map.Delete` on same key
- `sync.Map.LoadOrStore` when `LoadAndDelete` is intended
- Separate check-and-delete patterns on `sync.Map`
- Race conditions between read and remove operations
</patterns>

<related>
badLock, badSyncOnceFunc
</related>
