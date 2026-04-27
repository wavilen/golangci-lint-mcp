# revive: identical-switch-branches

<instructions>
Detects switch statements where two or more cases contain identical code bodies. This is usually a copy-paste error, or the cases should be combined into a single case with comma-separated match expressions.

Combine the duplicate cases using the Go multi-value case syntax: `case "a", "b":`.
</instructions>

<examples>
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
- Combine switch cases with identical bodies using Go's multi-value case syntax `case "a", "b":`
- Replace copy-paste errors where a case body was not updated for a new value
- Combine enum or status cases with identical behavior into a single multi-value case
- Combine type switch cases with identical handling for multiple types
- Remove default cases that duplicate a specific case's implementation
</patterns>

<related>
revive/identical-branches, revive/identical-switch-conditions, revive/identical-ifelseif-branches
</related>
