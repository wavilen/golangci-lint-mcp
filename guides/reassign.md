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
- Convert package-level config variables to `const` if never meant to change
- Replace `var` with `const` for magic numbers that are known at compile time
- Guard global state from accidental mutation across packages with unexported fields
</patterns>

<related>
goconst, govet, staticcheck
