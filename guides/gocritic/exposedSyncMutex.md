# gocritic: exposedSyncMutex

<instructions>
Detects exported struct fields of type `sync.Mutex`, `sync.RWMutex`, or `sync.WaitGroup` in public types. Exposing these synchronization primitives breaks encapsulation — external callers can lock, unlock, or modify them arbitrarily.

Make sync fields unexported and provide methods that encapsulate the locking behavior.
</instructions>

<examples>
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
- Unexport `sync.Mutex` fields — rename `Mu sync.Mutex` to `mu sync.Mutex`
- Unexport `sync.RWMutex` fields — rename `Lock sync.RWMutex` to `lock sync.RWMutex`
- Unexport `sync.WaitGroup` fields — rename `Wg sync.WaitGroup` to `wg sync.WaitGroup`
- Unexport pointer-to-mutex fields — rename `Mu *sync.Mutex` to `mu *sync.Mutex`
</patterns>

<related>
gocritic/captLocal, gocritic/underef
</related>
