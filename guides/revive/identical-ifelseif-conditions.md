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
- Copy-paste errors duplicating the condition in else-if
- Refactoring that changed one branch but not the condition
- Long if-else chains where duplicate conditions are hard to spot
- Feature flags checked multiple times in the same chain
- Variable mutation between checks making conditions appear different but evaluate the same
</patterns>

<related>
identical-ifelseif-branches, identical-switch-conditions
