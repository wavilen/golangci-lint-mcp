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
- Replace sequential `if _, ok := x.(T); ok` with a type switch
- Replace multiple comma-ok type assertions on the same expression with a type switch
- Replace type dispatch spread across `if` blocks with a single `switch x.(type)`
</patterns>

<related>
typeSwitchVar, ifElseChain, sloppyTypeAssert
</related>
