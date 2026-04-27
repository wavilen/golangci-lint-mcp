# gocritic: sloppyReassign

<instructions>
Detects reassignment of variables that were just assigned, particularly in `if`-`else` chains or short variable declarations where the reassignment is redundant or indicates a logic error. This often happens when a variable is assigned inside both branches of an if-else with the same expression.

Assign the variable once before the conditional, or fix the branch that should have a different value.
</instructions>

<examples>
## Good
```go
port := 443
if !useTLS {
    port = 80
}
```
</examples>

<patterns>
- Extract identical assignments from `if`/`else` branches to a single assignment before the branch
- Remove reassignment immediately after declaration with the same value
- Remove redundant reassignment in switch cases
- Remove conditional that always assigns the same result — assign once unconditionally
</patterns>

<related>
gocritic/sloppyLen, gocritic/sloppyTypeAssert, gocritic/dupBranchBody
</related>
