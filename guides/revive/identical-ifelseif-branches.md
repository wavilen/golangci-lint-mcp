# revive: identical-ifelseif-branches

<instructions>
Detects else-if chains where two branches contain identical code bodies. Like the identical-branches rule but specifically for multi-branch if-else-if chains. This is typically a copy-paste error.

Merge the duplicate branches using `||` in the condition, or fix the branch that should differ.
</instructions>

<examples>
## Bad
```go
if status == "active" {
    enableFeature()
    logAction("enabled")
} else if status == "pending" {
    enableFeature()
    logAction("enabled") // identical to "active" branch
} else {
    disableFeature()
}
```

## Good
```go
if status == "active" || status == "pending" {
    enableFeature()
    logAction("enabled")
} else {
    disableFeature()
}
```
</examples>

<patterns>
- Copy-paste errors in long if-else-if chains
- Status or state handling with identical logic for different states
- Error classification where multiple error types get the same treatment
- Feature branches merged but conditions not consolidated
- Multi-condition checks that should be combined with logical OR
</patterns>

<related>
identical-branches, identical-ifelseif-conditions, identical-switch-branches
