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
- Locking `mu1` but unlocking `mu2`
- Value receiver methods that copy the mutex
- `defer mu.Unlock()` paired with a different `mu.Lock()`
- Copying structs containing embedded mutexes
</patterns>

<related>
badSyncOnceFunc, unnecessaryDefer, deferInLoop
</related>
