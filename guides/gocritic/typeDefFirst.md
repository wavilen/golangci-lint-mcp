# gocritic: typeDefFirst

<instructions>
Detects type definitions in `switch` `case` clauses where the new type name is defined before its first use. This checker enforces defining types at the top level or before the function body rather than inline in control flow.

Move the type definition out of the `case` clause to the package level or before the function.
</instructions>

<examples>
## Bad
```go
switch v := x.(type) {
case struct{ Name string }:
	slog.Info("name", "value", v.Name)
}
```

## Good
```go
type namedItem struct{ Name string }

switch v := x.(type) {
case namedItem:
	slog.Info("name", "value", v.Name)
}
```
</examples>

<patterns>
- Anonymous struct types in `case` clauses
- Inline type definitions in `switch` arms
- `case struct{ ... }` instead of named type
</patterns>

<related>
typeSwitchVar, typeUnparen
</related>
