# revive: atomic

<instructions>
Detects misuse of `sync/atomic` operations on shared variables. Common mistakes include using plain assignment to update a value that is read with `atomic.Load`, or mixing atomic and non-atomic accesses to the same variable. This causes data races.

Ensure every read and write to a shared variable uses the corresponding `atomic` function consistently.
</instructions>

<examples>
## Bad
```go
var counter int64

// Non-atomic write — races with atomic reads
counter = 42

// Atomic read — sees torn value
val := atomic.LoadInt64(&counter)
```

## Good
```go
var counter int64

atomic.StoreInt64(&counter, 42)
val := atomic.LoadInt64(&counter)
```
</examples>

<patterns>
- Assigning directly to a variable that is read atomically elsewhere
- Mixing mutex-protected and atomic access to the same field
- Using `atomic.Add` but reading with a plain access instead of `atomic.Load`
- Forgetting that pointer atomic operations require `unsafe.Pointer` casting
- Swapping values without using `atomic.Swap` or `atomic.CompareAndSwap`
</patterns>

<related>
datarace
