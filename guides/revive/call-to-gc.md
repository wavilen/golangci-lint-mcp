# revive: call-to-gc

<instructions>
Detects explicit calls to `runtime.GC()`. The Go garbage collector is concurrent and self-tuning — manually triggering it is almost always unnecessary and can hurt performance by interrupting normal GC pacing. Explicit GC calls usually indicate a misunderstanding of how Go manages memory.

Remove the `runtime.GC()` call. If you need to force a GC cycle for benchmarking, use `testing.B.ReportAllocs` or `runtime/debug.FreeOSMemory` in test code only.
</instructions>

<examples>
## Bad
```go
func processBatch(items []Item) {
    for _, item := range items {
        handle(item)
        runtime.GC() // forces full GC on every iteration
    }
}
```

## Good
```go
func processBatch(items []Item) {
    for _, item := range items {
        handle(item)
    }
    // Let the GC run naturally
}
```
</examples>

<patterns>
- Remove `runtime.GC()` calls from loops and let the garbage collector run naturally
- Eliminate periodic GC calls in long-running goroutines
- Replace manual GC triggers in finalizers with normal GC pacing
- Use `testing.B.ReportAllocs` instead of forcing GC in benchmarks
- Remove `runtime.GC()` calls that attempt to free memory before returning large data structures — let the GC handle it
</patterns>

<related>
deep-exit
