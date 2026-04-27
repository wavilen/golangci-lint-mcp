# funcorder

<instructions>
Funcorder checks that exported functions on a type are defined before unexported ones, and that methods are grouped by receiver. This improves scanability and makes public API surfaces easier to find.

Reorder methods so exported methods come first, then unexported. Group all methods for the same receiver together in the file.
</instructions>

<examples>
## Good
```go
func (s *Server) Start() error { /* ... */ }
func (s *Server) Stop()  error { /* ... */ }
func (s *Server) start() error { /* ... */ }
```
</examples>

<patterns>
- Reorder exported methods before unexported methods on the same type
- Group methods for the same receiver into a single file
- Move constructor functions next to their type definition
- Group interface satisfaction methods with other methods on the same receiver
</patterns>

<related>
decorder, gochecknoinits, gochecknoglobals
</related>
