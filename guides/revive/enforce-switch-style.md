# revive: enforce-switch-style

<instructions>
Enforces using switch statements instead of if-else chains when comparing a single variable against multiple values. Switch statements are more readable and easier to extend than long if-else chains doing equality checks.

Replace if-else chains that compare one variable against different values with a switch statement.
</instructions>

<examples>
## Bad
```go
func httpStatus(code int) string {
    if code == 200 {
        return "OK"
    } else if code == 404 {
        return "Not Found"
    } else if code == 500 {
        return "Internal Server Error"
    }
    return "Unknown"
}
```

## Good
```go
func httpStatus(code int) string {
    switch code {
    case 200:
        return "OK"
    case 404:
        return "Not Found"
    case 500:
        return "Internal Server Error"
    default:
        return "Unknown"
    }
}
```
</examples>

<patterns>
- Convert if-else chains comparing one variable against constants to switch statements
- Replace multiple if statements that could be switch cases with a switch
- Use switch for status code or enum value handling instead of if-else chains
- Switch to switch for command or event dispatch instead of if-else
- Convert state machine transitions from if-else chains to switch statements
</patterns>

<related>
enforce-map-style, enforce-slice-style, identical-switch-branches
