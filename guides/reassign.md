# reassign

<instructions>
Reassign detects reassignment of package-level variables that are used as constants or configuration. Package-level variables initialized once and never meant to change should be `const` or protected from accidental reassignment.

Convert the variable to `const` if the value is known at compile time, or document why reassignment is intentional.
</instructions>

<examples>
## Bad
```go
var MaxRetries = 3

func init() {
    MaxRetries = 5 // accidental override
}
```

## Good
```go
const MaxRetries = 3

// Or if runtime configuration is needed:
var MaxRetries = 3 //nolint:reassign // configurable via config file
```
</examples>

<patterns>
- Package-level config variables reassigned in `init()` or tests
- Magic numbers stored in `var` that should be `const`
- Global state mutated unexpectedly across packages
</patterns>

<related>
goconst, govet, staticcheck
