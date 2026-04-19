# gocritic: typeAssertChain

<instructions>
Detects repeated type assertions on the same expression across multiple `if` or `switch` branches. A chain of `if _, ok := x.(T1); ok` statements should be rewritten as a type switch for clarity and efficiency.

Use a `switch x := x.(type)` statement to handle multiple possible types in one construct.
</instructions>

<examples>
## Bad
```go
if v, ok := val.(int); ok {
	useInt(v)
} else if v, ok := val.(string); ok {
	useStr(v)
} else if v, ok := val.(bool); ok {
	useBool(v)
}
```

## Good
```go
switch v := val.(type) {
case int:
	useInt(v)
case string:
	useStr(v)
case bool:
	useBool(v)
}
```
</examples>

<patterns>
- Sequential `if _, ok := x.(T); ok` patterns
- Multiple comma-ok type assertions on the same expression
- Type dispatch spread across multiple `if` blocks
</patterns>

<related>
typeSwitchVar, ifElseChain, sloppyTypeAssert
</related>
