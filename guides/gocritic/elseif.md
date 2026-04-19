# gocritic: elseif

<instructions>
Detects nested `if-else` chains that could be flattened into `else if` blocks. Deeply nested `if` statements inside `else` blocks reduce readability by adding unnecessary indentation.

Replace `else { if ... }` with `else if ...` to flatten the control flow.
</instructions>

<examples>
## Bad
```go
if x > 0 {
	positive()
} else {
	if x < 0 {
		negative()
	} else {
		zero()
	}
}
```

## Good
```go
if x > 0 {
	positive()
} else if x < 0 {
	negative()
} else {
	zero()
}
```
</examples>

<patterns>
- `else { if cond { ... } }` → `else if cond { ... }`
- `else { if cond { ... } else { ... } }` → `else if cond { ... } else { ... }`
- Multi-level nesting of if-else that should be a flat chain
</patterns>

<related>
ifElseChain, nestingReduce
</related>
