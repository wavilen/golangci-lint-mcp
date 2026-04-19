# gocritic: unnecessaryBlock

<instructions>
Detects unnecessary block statements — braces `{ }` that create a new scope but are not associated with a control flow statement (`if`, `for`, `switch`). Standalone blocks are occasionally useful for scoping variables, but more often they indicate a mistake or leftover code.

Remove the unnecessary block, or if intentional (for variable scoping), add a comment explaining why.
</instructions>

<examples>
## Bad
```go
func process() {
	{
		x := compute()
		use(x)
	}
	return
}
```

## Good
```go
func process() {
	x := compute()
	use(x)
	return
}
```
</examples>

<patterns>
- Standalone `{ }` block inside a function with no control flow
- Block used only for scoping that could be inlined
- Accidental extra braces from copy-paste or refactoring
</patterns>

<related>
nestingReduce, initClause
</related>
