# govet: atomic

<instructions>
Detects incorrect usage of `sync/atomic` operations. Common mistakes include mixing atomic and non-atomic access to the same variable, passing a value instead of a pointer to atomic functions, or swapping `*addr` for `atomic.Load*`.

Use `atomic.AddInt32`, `atomic.LoadInt32`, `atomic.StoreInt32`, etc. consistently — always pass a pointer and never mix with direct field access.
</instructions>

<examples>
## Bad
```go
var counter int64
counter++                    // non-atomic access
_ = atomic.AddInt64(&counter, 1) // mixing atomic and non-atomic
```

## Good
```go
var counter int64
atomic.AddInt64(&counter, 1)  // consistent atomic access
val := atomic.LoadInt64(&counter)
```
</examples>

<patterns>
- Direct field access mixed with atomic operations on the same variable
- Passing value instead of pointer to atomic functions
- Using `*addr` instead of `atomic.LoadPtr`
- Reassigning atomic pointer with non-atomic store
</patterns>

<related>
copylocks, defers
</related>
