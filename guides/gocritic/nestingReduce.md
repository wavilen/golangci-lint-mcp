# gocritic: nestingReduce

<instructions>
Detects deeply nested `if`/`for`/`select` blocks that can be flattened by using early returns, `continue`, or `break`. Deep nesting makes code harder to understand and test. The checker suggests refactoring to reduce nesting depth.

Invert conditions and return early to reduce the nesting level of the main logic path.
</instructions>

<examples>
## Bad
```go
func process(data []byte) error {
	if data != nil {
		if len(data) > 0 {
			if isValid(data) {
				return handle(data)
			} else {
				return errors.New("invalid")
			}
		}
	}
	return nil
}
```

## Good
```go
func process(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}
	if !isValid(data) {
		return errors.New("invalid")
	}
	return handle(data)
}
```
</examples>

<patterns>
- `if` inside `if` inside `for` — three or more levels deep
- Happy path buried inside nested blocks
- Error handling nested instead of early return
- Guard clauses that could flatten the function body
</patterns>

<related>
elseif, ifElseChain, unnecessaryBlock
</related>
