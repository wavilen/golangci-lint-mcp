# gocritic: appendCombine

<instructions>
Detects consecutive `append` calls to the same slice that can be combined into a single `append` with multiple arguments. Each separate `append` call incurs a potential reallocation, so combining them reduces allocations and improves readability.

Merge multiple `append` calls on the same variable into one call with variadic arguments.
</instructions>

<examples>
## Good
```go
xs = append(xs, 1, 2, 3)
```
</examples>

<patterns>
- Combine consecutive `append` calls into a single variadic `append(xs, 1, 2, 3)`
- Replace loop-based appending with `append(s, slice...)` for bulk additions
- Combine compile-time-known append chains into a single literal or variadic call
- Replace repeated single-element appends in constructors with one combined call
</patterns>

<related>
gocritic/appendAssign, gocritic/rangeExprCopy
</related>
