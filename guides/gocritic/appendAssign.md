# gocritic: appendAssign

<instructions>
Detects `append` calls where the result is assigned to a different variable than the original slice. This is almost always a mistake — `append` may or may not allocate a new backing array, so reassigning to a different variable leads to inconsistent state.

Always assign the result of `append` back to the same variable: `s = append(s, x)`.
</instructions>

<examples>
## Bad
```go
items := []string{"a", "b"}
more := append(items, "c") // items and more share backing array
```

## Good
```go
items := []string{"a", "b"}
items = append(items, "c")
```
</examples>

<patterns>
- Assign the result of `append` back to the same slice variable: `s = append(s, x)`
- Avoid creating new slice variables from `append` — the original and new may share backing arrays
- Replace pattern `more := append(items, ...)` with `items = append(items, ...)` for consistent state
</patterns>

<related>
appendCombine, rangeExprCopy
</related>
