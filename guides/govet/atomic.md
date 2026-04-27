# govet: atomic

<instructions>
Detects incorrect usage of `sync/atomic` operations. Common mistakes include mixing atomic and non-atomic access to the same variable, passing a value instead of a pointer to atomic functions, or swapping `*addr` for `atomic.Load*`.

Use `atomic.AddInt32`, `atomic.LoadInt32`, `atomic.StoreInt32`, etc. consistently — always pass a pointer and never mix with direct field access.
</instructions>

<examples>
## Good
```go
var counter int64
atomic.AddInt64(&counter, 1)  // consistent atomic access
val := atomic.LoadInt64(&counter)
```
</examples>

<patterns>
- Use only `atomic.Load`/`atomic.Store` for fields accessed concurrently — never mix with direct reads/writes
- Pass pointers (not values) to all `atomic` functions (`atomic.AddInt32(&x, 1)`, not `atomic.AddInt32(x, 1)`)
- Use `atomic.LoadPtr(&addr)` instead of dereferencing `*addr` for atomic values
- Use `atomic.Store` when reassigning atomic pointers — never plain assignment
</patterns>

<related>
govet/copylocks, govet/defers, gocritic/badLock
</related>
