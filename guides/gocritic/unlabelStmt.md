# gocritic: unlabelStmt

<instructions>
Detects labels on `for`, `switch`, or `select` statements that are never referenced by a `break`, `continue`, or `goto`. Unused labels add visual noise and suggest the control flow was once more complex than it is now.

Remove the unused label. Keep labels only when needed for `break` or `continue` targeting outer loops.
</instructions>

<examples>
## Bad
```go
loop:
	for i := 0; i < n; i++ {
		process(i)
	}
```

## Good
```go
for i := 0; i < n; i++ {
	process(i)
}
```
</examples>

<patterns>
- Labels on single loops with no nested break/continue targeting them
- Labels on `switch` statements (rarely needed)
- Leftover labels from refactored code
</patterns>

<related>
unnecessaryBlock, nestingReduce
</related>
