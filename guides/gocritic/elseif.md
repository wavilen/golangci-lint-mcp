# gocritic: elseif

<instructions>
Detects nested `if-else` chains that could be flattened into `else if` blocks. Deeply nested `if` statements inside `else` blocks reduce readability by adding unnecessary indentation.

Replace `else { if ... }` with `else if ...` to flatten the control flow.
</instructions>

<examples>
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
- Replace `else { if cond { ... } }` with `else if cond { ... }`
- Replace `else { if cond { ... } else { ... } }` with `else if cond { ... } else { ... }`
- Flatten multi-level nested if-else into a flat `else if` chain
</patterns>

<related>
gocritic/ifElseChain, gocritic/nestingReduce
</related>
