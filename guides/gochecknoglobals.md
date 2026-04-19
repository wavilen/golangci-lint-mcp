# gochecknoglobals

<instructions>
Gochecknoglobals checks that no global variables exist in the package. Globals introduce hidden state, make testing difficult, and can cause data races in concurrent programs.

Move global state into structs passed explicitly as dependencies. Use function-scoped variables or dependency injection instead of package-level mutable state.
</instructions>

<examples>
## Bad
```go
var cache = map[string]string{}

func Get(key string) string {
    return cache[key]
}
```

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
- Package-level maps or slices used as registries
- Global configuration variables mutated at init time
- Singleton patterns using package-level vars
- Counters or metrics stored as package variables
</patterns>

<related>
gochecknoinits, reassign, varnamelen
</related>
