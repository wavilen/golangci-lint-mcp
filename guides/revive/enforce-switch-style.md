# revive: enforce-switch-style

<instructions>
Enforces using switch statements instead of if-else chains when comparing a single variable against multiple values. Switch statements are more readable and easier to extend than long if-else chains doing equality checks.

Replace if-else chains that compare one variable against different values with a switch statement.
</instructions>

<examples>
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
revive/enforce-map-style, revive/enforce-slice-style, revive/identical-switch-branches
</related>
