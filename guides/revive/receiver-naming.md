# revive: receiver-naming

<instructions>
Enforces consistent and idiomatic receiver names in method declarations. The receiver name should be a short abbreviation (1-2 characters) of the type name, consistent across all methods of that type. Avoid generic names like `this` or `self` — Go uses concise, type-derived names.

Use the first letter or a meaningful short abbreviation of the type name. Apply the same receiver name to every method on the type.
</instructions>

<examples>
## Bad
```go
func (this *Server) Start() error { ... }
func (s *Server) Stop() error  { ... } // inconsistent with 'this' above
func (self *Client) Send() error { ... }
```

## Good
```go
func (s *Server) Start() error { ... }
func (s *Server) Stop() error  { ... }
func (c *Client) Send() error  { ... }
```
</examples>

<patterns>
- Receiver named `this` or `self` (non-idiomatic Go)
- Inconsistent receiver names across methods of the same type
- Overly long receiver names (e.g., `server` instead of `s`)
- Single-letter receiver names that don't match the type (e.g., `(x *Server)`)
- Mixed value and pointer receivers on the same type
</patterns>

<related>
unexported-naming, unexported-return, var-naming
