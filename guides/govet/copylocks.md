# govet: copylocks

<instructions>
Detects copying of values that contain mutexes or other lock types (`sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Cond`, `sync.Once`). Copying a lock after first use leads to race conditions because both copies share internal state.

Use pointers to structs containing locks, or avoid copying them after initialization.
</instructions>

<examples>
## Bad
```go
type Server struct {
    mu    sync.Mutex
    conns map[string]net.Conn
}

func (s Server) Clone() Server {
    return s // copies the mutex — dangerous
}
```

## Good
```go
type Server struct {
    mu    sync.Mutex
    conns map[string]net.Conn
}

func (s *Server) Clone() *Server {
    return &Server{conns: s.conns} // new mutex, shared state via pointer
}
```
</examples>

<patterns>
- Use pointer receivers for structs containing `sync.Mutex` or `sync.RWMutex` fields
- Return pointers to lock-containing structs instead of copying by value
- Use pointers when assigning lock-containing structs to new variables
- Use pointer range (`for i := range &slice`) or index-based loop to avoid copying mutex structs
</patterns>

<related>
atomic, waitgroup
</related>
