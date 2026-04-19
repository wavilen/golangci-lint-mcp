# revive: identical-switch-branches

<instructions>
Detects switch statements where two or more cases contain identical code bodies. This is usually a copy-paste error, or the cases should be combined into a single case with comma-separated match expressions.

Combine the duplicate cases using the Go multi-value case syntax: `case "a", "b":`.
</instructions>

<examples>
## Bad
```go
switch role {
case "admin":
    grantFullAccess()
case "superuser":
    grantFullAccess() // identical — should be combined
case "user":
    grantLimitedAccess()
}
```

## Good
```go
switch role {
case "admin", "superuser":
    grantFullAccess()
case "user":
    grantLimitedAccess()
}
```
</examples>

<patterns>
- Switch cases with identical bodies that should use multi-value syntax
- Copy-paste errors where a case body was not updated for a new value
- Enum or status handling with identical behavior for different values
- Type switch cases with identical handling for multiple types
- Default and a specific case having identical implementations
</patterns>

<related>
identical-branches, identical-switch-conditions, identical-ifelseif-branches
