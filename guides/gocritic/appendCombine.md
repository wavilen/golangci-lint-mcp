# gocritic: appendCombine

<instructions>
Detects consecutive `append` calls to the same slice that can be combined into a single `append` with multiple arguments. Each separate `append` call incurs a potential reallocation, so combining them reduces allocations and improves readability.

Merge multiple `append` calls on the same variable into one call with variadic arguments.
</instructions>

<examples>
## Bad
```go
xs = append(xs, 1)
xs = append(xs, 2)
xs = append(xs, 3)
```

## Good
```go
xs = append(xs, 1, 2, 3)
```
</examples>

<patterns>
- Multiple consecutive `append` calls appending individual values to the same slice
- Loops that could use variadic `append` or `append(s, slice...)` instead of repeated appends
- Append chains in initialization code where all values are known at compile time
- Building slices in constructors with repeated single-element appends
</patterns>

<related>
appendAssign, rangedExpr, makeLen
