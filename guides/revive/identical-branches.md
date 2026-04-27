# revive: identical-branches

<instructions>
Detects `if-else` or `switch` statements where two branches contain identical code. This is usually a copy-paste error — the developer intended different behavior for each branch but forgot to modify one.

Either remove the duplicate branch (if it is truly identical behavior) or fix the branch that should differ.
</instructions>

<examples>
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
- Replace copy-paste if-else blocks where one branch was not updated with its intended logic
- Combine switch cases with identical implementations using multi-value case syntax
- Remove refactoring leftovers where one branch became identical to another
- Replace feature flag checks where both branches accidentally do the same thing
- Combine identical error recovery logic in different branches into a shared handler
</patterns>

<related>
revive/identical-ifelseif-branches, revive/identical-switch-branches, revive/constant-logical-expr
</related>
