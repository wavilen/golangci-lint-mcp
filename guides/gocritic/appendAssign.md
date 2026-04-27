# gocritic: appendAssign

<instructions>
Detects `append` calls where the result is assigned to a different variable than the original slice. This is almost always a mistake — `append` may or may not allocate a new backing array, so reassigning to a different variable leads to inconsistent state.

Always assign the result of `append` back to the same variable: `s = append(s, x)`.
</instructions>

<examples>
## Good
```go
items := make([]string, 0, 3)
items = append(items, "a", "b", "c")
```
</examples>

<patterns>
- Assign the result of `append` back to the same slice variable: `s = append(s, x)`
- Avoid creating new slice variables from `append` — the original and new may share backing arrays
- Replace pattern `more := append(items, ...)` with `items = append(items, ...)` for consistent state
</patterns>

<related>
gocritic/appendCombine, gocritic/rangeExprCopy
</related>
