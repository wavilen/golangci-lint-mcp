# revive: identical-branches

<instructions>
Detects `if-else` or `switch` statements where two branches contain identical code. This is usually a copy-paste error — the developer intended different behavior for each branch but forgot to modify one.

Either remove the duplicate branch (if it is truly identical behavior) or fix the branch that should differ.
</instructions>

<examples>
## Bad
```go
if isPriority {
    process(item)
    notify(item)
} else {
    process(item)
    notify(item) // copy-paste — should this do something different?
}
```

## Good
```go
if isPriority {
    processPriority(item)
    notify(item)
} else {
    process(item)
}
```
</examples>

<patterns>
- Copy-paste if-else blocks where one branch was not updated
- Switch cases with identical implementations that should be merged
- Refactoring leftovers where one branch became identical to another
- Feature flag checks where both branches accidentally do the same thing
- Error handling with identical recovery logic in different branches
</patterns>

<related>
identical-ifelseif-branches, identical-switch-branches, constant-logical-expr
