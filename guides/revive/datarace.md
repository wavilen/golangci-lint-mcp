# revive: datarace

<instructions>
Detects patterns that commonly cause data races in concurrent Go code. Data races occur when two goroutines access the same memory simultaneously, with at least one write, and no synchronization. This rule flags unsafe patterns like accessing shared variables without synchronization.

Protect shared state with `sync.Mutex`, `sync.RWMutex`, channels, or `sync/atomic` operations.
</instructions>

<examples>
## Bad
```go
var counter int

func inc() {
    go func() {
        counter++ // data race: unsynchronized write
    }()
}
```

## Good
```go
var counter int64

func inc() {
    go func() {
        atomic.AddInt64(&counter, 1)
    }()
}
```
</examples>

<patterns>
- Goroutines accessing shared variables without locks or atomics
- Non-thread-safe map access from multiple goroutines
- Shared slices appended to from concurrent goroutines
- WaitGroup `Add` called inside the spawned goroutine instead of before
- Closing over loop variables that are modified concurrently
</patterns>

<related>
atomic, copyloopvar
