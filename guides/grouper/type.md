# grouper: type

<instructions>
Detects multiple consecutive `type` declarations that can be grouped into a single `type` block. Scattered type definitions make it harder to see related types at a glance. Group them into `type ( ... )` for better organization, especially for small supporting types.
</instructions>

<examples>
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
- Group sequential `type` declarations into a single `type ( ... )` block
- Group related small types declared separately for clarity
- Group adjacent type aliases or interfaces defined individually
</patterns>

<related>
grouper/const, grouper/var, grouper/import
</related>
