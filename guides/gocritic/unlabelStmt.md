# gocritic: unlabelStmt

<instructions>
Detects labels on `for`, `switch`, or `select` statements that are never referenced by a `break`, `continue`, or `goto`. Unused labels add visual noise and suggest the control flow was once more complex than it is now.

Remove the unused label. Keep labels only when needed for `break` or `continue` targeting outer loops.
</instructions>

<examples>
## Good
```go
for i := 0; i < n; i++ {
	process(i)
}
```
</examples>

<patterns>
- Remove labels on single loops with no nested break/continue targeting them
- Remove labels on `switch` statements — rarely needed
- Remove leftover labels from refactored code
</patterns>

<related>
gocritic/unnecessaryBlock, gocritic/nestingReduce
</related>
