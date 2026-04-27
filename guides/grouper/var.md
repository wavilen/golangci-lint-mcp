# grouper: var

<instructions>
Detects multiple consecutive `var` declarations that can be grouped into a single `var` block. Scattered `var` statements for related variables are harder to scan. Group them using `var ( ... )` to show they belong together, especially for package-level configuration or state variables.
</instructions>

<examples>
## Good
```go
var (
    defaultTimeout = 30 * time.Second
    maxRetries     = 3
    enableCache    = true
)
```
</examples>

<patterns>
- Group sequential `var` declarations into a single `var ( ... )` block
- Group package-level configuration variables declared separately
- Group related state variables for readability
</patterns>

<related>
grouper/const, grouper/type, grouper/import
</related>
