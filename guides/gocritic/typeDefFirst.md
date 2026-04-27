# gocritic: typeDefFirst

<instructions>
Detects type definitions in `switch` `case` clauses where the new type name is defined before its first use. This checker enforces defining types at the top level or before the function body rather than inline in control flow.

Move the type definition out of the `case` clause to the package level or before the function.
</instructions>

<examples>
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
- Define named types for anonymous structs used in `case` clauses
- Replace `case struct{ ... }` with a named type defined before the switch
</patterns>

<related>
gocritic/typeSwitchVar, gocritic/typeUnparen
</related>
