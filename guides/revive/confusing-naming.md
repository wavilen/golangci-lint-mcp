# revive: confusing-naming

<instructions>
Detects methods and fields that differ only by letter casing (e.g., `doThing` vs `dothing`, `ID` vs `id`). Such names are easily confused by developers and may cause subtle bugs when the wrong one is used.

Rename one of the conflicting identifiers to be clearly distinct. Follow Go naming conventions: use camelCase, avoid unnecessary abbreviations, and ensure names are visually distinguishable.
</instructions>

<examples>
## Good
```go
type Config struct {
    HTTPTimeout      int
    ConnectionTimeout int // clearly different name
}

func (c *Client) readData()    {}
func (c *Client) fetchData()   {} // distinct name
```
</examples>

<patterns>
- Rename struct fields that differ only in casing to clearly distinct names
- Avoid methods with names that differ only by case from other methods on the same type
- Ensure package-level identifiers are visually distinguishable from cased variants
- Use JSON tag names with Go field names to prevent casing collisions
- Rename case-conflicting identifiers produced by code generation from external schemas
</patterns>

<related>
revive/confusing-results, revive/exported
</related>
