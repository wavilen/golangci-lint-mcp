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
- Calling `runtime.GC()` after processing each item in a loop
- Periodic GC calls in long-running goroutines
- Manual GC triggers in finalizers or cleanup functions
- Benchmark setup code forcing GC before measurements
- Attempting to free memory before returning large data structures
</patterns>

<related>
deep-exit
