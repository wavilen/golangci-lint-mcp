# gocritic: exposedSyncMutex

<instructions>
Detects exported struct fields of type `sync.Mutex`, `sync.RWMutex`, or `sync.WaitGroup` in public types. Exposing these synchronization primitives breaks encapsulation — external callers can lock, unlock, or modify them arbitrarily.

Make sync fields unexported and provide methods that encapsulate the locking behavior.
</instructions>

<examples>
## Bad
```go
type Server struct {
	Mu    sync.Mutex
	Count int
}
```

## Good
```go
type Server struct {
	mu    sync.Mutex
	Count int
}

func (s *Server) Increment() {
	s.mu.Lock()
	s.Count++
	s.mu.Unlock()
}
```
</examples>

<patterns>
- Exported `sync.Mutex` field: `Mu sync.Mutex`
- Exported `sync.RWMutex` field: `Lock sync.RWMutex`
- Exported `sync.WaitGroup` field: `Wg sync.WaitGroup`
- Pointer to mutex exported: `Mu *sync.Mutex`
</patterns>

<related>
captLocal, underef
</related>
