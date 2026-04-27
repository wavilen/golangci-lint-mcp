# revive: datarace

<instructions>
Detects patterns that commonly cause data races in concurrent Go code. Data races occur when two goroutines access the same memory simultaneously, with at least one write, and no synchronization. This rule flags unsafe patterns like accessing shared variables without synchronization.

Protect shared state with `sync.Mutex`, `sync.RWMutex`, channels, or `sync/atomic` operations.
</instructions>

<examples>
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
- Guard shared variables accessed from goroutines with `sync.Mutex` or `sync/atomic`
- Use `sync.Map` or mutex-protected maps for concurrent map access from multiple goroutines
- Guard shared slice appends from concurrent goroutines with a mutex
- Call `wg.Add` before spawning a goroutine — never inside it
- Use loop variables by value when closing over them in concurrent goroutines
</patterns>

<related>
revive/atomic, copyloopvar
</related>
