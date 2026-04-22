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
- Flatten `if`-inside-`if`-inside-`for` — use guard clauses to reduce nesting
- Move the happy path to the top — use early returns for error conditions
- Replace nested error handling with early `return err` to flatten the function body
- Extract guard clauses to reduce nesting levels
</patterns>

<related>
elseif, ifElseChain, unnecessaryBlock
</related>
