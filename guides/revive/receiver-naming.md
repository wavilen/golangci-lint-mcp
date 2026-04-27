# revive: receiver-naming

<instructions>
Enforces consistent and idiomatic receiver names in method declarations. The receiver name should be a short abbreviation (1-2 characters) of the type name, consistent across all methods of that type. Avoid generic names like `this` or `self` — Go uses concise, type-derived names.

Use the first letter or a meaningful short abbreviation of the type name. Apply the same receiver name to every method on the type.
</instructions>

<examples>
## Good
```go
func (s *Server) Start() error { ... }
func (s *Server) Stop() error  { ... }
func (c *Client) Send() error  { ... }
```
</examples>

<patterns>
- Replace `this` or `self` receiver names with a short type-derived abbreviation (e.g., `s` for `Server`)
- Use the same receiver name across all methods of a given type
- Simplify overly long receiver names to 1-2 characters derived from the type name
- Use a receiver letter that matches the type name (e.g., `s` for `Server`, not `x`)
- Use consistent value or pointer receivers across all methods of the same type
</patterns>

<related>
revive/unexported-naming, revive/unexported-return, revive/var-naming
</related>
