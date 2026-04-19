# grouper: var

<instructions>
Detects multiple consecutive `var` declarations that can be grouped into a single `var` block. Scattered `var` statements for related variables are harder to scan. Group them using `var ( ... )` to show they belong together, especially for package-level configuration or state variables.
</instructions>

<examples>
## Bad
```go
var defaultTimeout = 30 * time.Second
var maxRetries = 3
var enableCache = true
```

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
- Multiple `var` declarations in sequence — group into `var ( ... )`
- Package-level configuration variables declared separately — group them
- Related state variables — group for readability
</patterns>

<related>
const, type, import
