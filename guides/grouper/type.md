# grouper: type

<instructions>
Detects multiple consecutive `type` declarations that can be grouped into a single `type` block. Scattered type definitions make it harder to see related types at a glance. Group them into `type ( ... )` for better organization, especially for small supporting types.
</instructions>

<examples>
## Bad
```go
type Config struct {
    Host string
    Port int
}
type Option func(*Config)
type OptionList []Option
```

## Good
```go
type (
    Config     struct{ Host string; Port int }
    Option     func(*Config)
    OptionList []Option
)
```
</examples>

<patterns>
- Multiple `type` declarations in sequence — group into `type ( ... )`
- Related small types declared separately — group for clarity
- Type aliases or interfaces defined individually — group when adjacent
</patterns>

<related>
const, var, import
