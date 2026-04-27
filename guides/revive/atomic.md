# revive: atomic

<instructions>
Detects misuse of `sync/atomic` operations on shared variables. Common mistakes include using plain assignment to update a value that is read with `atomic.Load`, or mixing atomic and non-atomic accesses to the same variable. This causes data races.

Ensure every read and write to a shared variable uses the corresponding `atomic` function consistently.
</instructions>

<examples>
## Good
```go
var counter int64

atomic.StoreInt64(&counter, 42)
val := atomic.LoadInt64(&counter)
```
</examples>

<patterns>
- Use `atomic.Store` instead of plain assignment for variables read atomically elsewhere
- Use either mutex or atomic access consistently — never mix both on the same field
- Use `atomic.Add` writes with `atomic.Load` reads instead of plain access
- Use `unsafe.Pointer` casting for pointer atomic operations like `atomic.LoadPointer`
- Use `atomic.Swap` or `atomic.CompareAndSwap` for conditional value updates
</patterns>

<related>
revive/datarace
</related>
