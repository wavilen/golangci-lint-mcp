# gocritic: initClause

<instructions>
Detects `if` statements where a simple assignment in the init clause could be moved before the `if` for clarity. While Go allows `if init; condition`, overusing init clauses for non-trivial logic reduces readability.

Move the init statement before the `if` when the assignment is reused later or when it improves clarity.
</instructions>

<examples>
## Bad
```go
if err := doSomething(); err != nil {
	return err
}
if result, err := compute(); err == nil {
	use(result)
}
```

## Good
```go
err := doSomething()
if err != nil {
	return err
}

result, err := compute()
if err == nil {
	use(result)
}
```
</examples>

<patterns>
- Move init-clause variables to outer scope when used after the `if` block
- Simplify complex init clauses with multiple assignments — split into separate statements
- Move init-clause logic to broader scope when error handling needs the variable later
</patterns>

<related>
unnecessaryBlock, nestingReduce
</related>
