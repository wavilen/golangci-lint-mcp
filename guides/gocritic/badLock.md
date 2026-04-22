# gocritic: badLock

<instructions>
Detects incorrect mutex usage patterns, such as locking a mutex on one type but unlocking a different one, or calling `Lock`/`Unlock` on value copies instead of pointers. Since `sync.Mutex` and `sync.RWMutex` must not be copied after first use, operating on a copy silently disables synchronization.

Always lock and unlock the same mutex instance, and ensure mutexes are accessed via pointer receivers.
</instructions>

<examples>
## Bad
```go
func (s Server) Process() {
    s.mu.Lock()
    defer s.mu2.Unlock() // wrong mutex unlocked
    // ...
}
```

## Good
```go
func (s *Server) Process() {
    s.mu.Lock()
    defer s.mu.Unlock()
    // ...
}
```
</examples>

<patterns>
- Ensure `Lock` and `Unlock` target the same mutex instance
- Use pointer receivers on methods that access mutex fields — value receivers copy the mutex
- Pair every `mu.Lock()` with `defer mu.Unlock()` on the same mutex
- Avoid copying structs with embedded `sync.Mutex` — always use pointers
</patterns>

<related>
badSyncOnceFunc, unnecessaryDefer, deferInLoop
</related>
