# gocritic: sloppyReassign

<instructions>
Detects reassignment of variables that were just assigned, particularly in `if`-`else` chains or short variable declarations where the reassignment is redundant or indicates a logic error. This often happens when a variable is assigned inside both branches of an if-else with the same expression.

Assign the variable once before the conditional, or fix the branch that should have a different value.
</instructions>

<examples>
## Bad
```go
var port int
if useTLS {
    port = 443
} else {
    port = 443 // same value in both branches
}
```

## Good
```go
port := 443
if !useTLS {
    port = 80
}
```
</examples>

<patterns>
- Same value assigned in both `if` and `else` branches
- Variable reassigned immediately after declaration with same value
- Redundant reassignment in switch cases
- Conditional that always assigns the same result
</patterns>

<related>
sloppyLen, sloppyTypeAssert, dupBranchBody
</related>
