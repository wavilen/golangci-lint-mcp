# revive: add-constant

<instructions>
Detects magic numbers and string literals that should be extracted into named constants. Numeric literals embedded directly in expressions make code harder to understand and maintain — a named constant documents intent and provides a single point of change.

Extract repeated or meaningful literals into `const` declarations at the package or function level.
</instructions>

<examples>
## Bad
```go
if timeout > 30 {
    return fmt.Errorf("exceeded %d seconds", 30)
}
if status == 404 {
    handleNotFound()
}
```

## Good
```go
const (
    defaultTimeoutSec = 30
    httpStatusNotFound = 404
)

if timeout > defaultTimeoutSec {
    return fmt.Errorf("exceeded %d seconds", defaultTimeoutSec)
}
if status == httpStatusNotFound {
    handleNotFound()
}
```
</examples>

<patterns>
- Extract numeric literals used in comparisons or calculations into named constants
- Replace hardcoded HTTP status codes, port numbers, or timeout values with named constants
- Define constants for repeated string literals that represent fixed values
- Rename array indices or slice bounds with unexplained numeric values
- Move configuration-like values scattered through business logic into `const` declarations
</patterns>

<related>
mnd, goconst
