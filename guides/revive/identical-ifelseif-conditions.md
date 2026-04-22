# revive: identical-ifelseif-conditions

<instructions>
Detects else-if chains where two branches check the same condition. The second branch is unreachable dead code — execution will always match the first branch.

Remove the duplicate condition. If the branches were meant to check different things, fix the condition.
</instructions>

<examples>
## Bad
```go
if x > 100 {
    handleLarge(x)
} else if x > 100 { // unreachable
    handleExtraLarge(x)
}
```

## Good
```go
if x > 100 {
    handleLarge(x)
} else if x > 50 {
    handleMedium(x)
}
```
</examples>

<patterns>
- Remove duplicate conditions in else-if chains caused by copy-paste errors
- Replace conditions that were not updated during refactoring, making branches unreachable
- Simplify long if-else chains by removing unreachable duplicate conditions
- Remove redundant feature flag checks that appear multiple times in the same chain
- Simplify variable-based conditions that evaluate identically due to intermediate mutations
</patterns>

<related>
identical-ifelseif-branches, identical-switch-conditions
