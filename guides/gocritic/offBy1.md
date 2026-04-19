# gocritic: offBy1

<instructions>
Detects common off-by-one errors in index-based operations, such as accessing `s[len(s)]` (out of bounds) or slicing `s[:len(s)-1]` when `s[:len(s)]` was intended. Also flags suspicious loop bounds and slice indices.

Use the correct index. Remember that Go slices are zero-indexed and the last valid index is `len(s)-1`. For slicing, `s[:n]` excludes index `n`.
</instructions>

<examples>
## Bad
```go
last := items[len(items)] // panic: index out of range
```

## Good
```go
last := items[len(items)-1]
```
</examples>

<patterns>
- `s[len(s)]` instead of `s[len(s)-1]`
- Loop bound `<= len(s)` instead of `< len(s)`
- Slice expression `s[:len(s)]` which is equivalent to `s`
- `s[0:len(s)-1]` when `s[:len(s)]` or `s` was intended
</patterns>

<related>
truncateCmp, badCond, sloppyLen
</related>
