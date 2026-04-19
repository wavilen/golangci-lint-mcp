# funcorder

<instructions>
Funcorder checks that exported functions on a type are defined before unexported ones, and that methods are grouped by receiver. This improves scanability and makes public API surfaces easier to find.

Reorder methods so exported methods come first, then unexported. Group all methods for the same receiver together in the file.
</instructions>

<examples>
## Bad
```go
func (s *Server) start() error { /* ... */ }
func (s *Server) Start() error { /* ... */ }
func (s *Server) Stop()  error { /* ... */ }
```

## Good
```go
func (s *Server) Start() error { /* ... */ }
func (s *Server) Stop()  error { /* ... */ }
func (s *Server) start() error { /* ... */ }
```
</examples>

<patterns>
- Unexported methods appearing before exported ones on same type
- Methods for the same receiver scattered across different files
- Constructor functions placed far from their type
- Interface satisfaction methods not grouped with other methods
</patterns>

<related>
decorder, gochecknoinits, gochecknoglobals
</related>
