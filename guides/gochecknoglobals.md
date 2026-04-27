# gochecknoglobals

<instructions>
Gochecknoglobals checks that no global variables exist in the package. Globals introduce hidden state, make testing difficult, and can cause data races in concurrent programs.

Move global state into structs passed explicitly as dependencies. Use function-scoped variables or dependency injection instead of package-level mutable state.
</instructions>

<examples>
## Good
```go
type Cache struct {
    data map[string]string
}

func (c *Cache) Get(key string) string {
    return c.data[key]
}
```
</examples>

<patterns>
- Move package-level maps and slices into struct fields passed via dependency injection
- Replace global config variables with explicit setup functions called from `main()`
- Use sync.Once inside a constructor instead of singleton package-level vars
- Pass counters and metrics through struct receivers rather than package variables
</patterns>

<related>
gochecknoinits, reassign, varnamelen
</related>
