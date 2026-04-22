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
- Replace copy-paste errors in long if-else-if chains by updating the branch that should differ
- Combine status or state conditions with identical logic using `||` in the condition
- Combine error classification branches that give multiple error types the same treatment
- Combine conditions in merged feature branches that were not combined
- Combine multi-condition checks into a single condition using logical OR
</patterns>

<related>
identical-branches, identical-ifelseif-conditions, identical-switch-branches
