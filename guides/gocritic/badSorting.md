# gocritic: badSorting

<instructions>
Detects calls to `sort.Slice` or similar where the comparison uses `<=` or `>=` instead of strict `<`. The `sort.Slice` comparison function must return whether element `i` should sort before element `j`, and must use strict less-than to avoid non-deterministic ordering and potential infinite loops.

Use strict `<` in sort comparison functions instead of `<=`.
</instructions>

<examples>
## Bad
```go
sort.Slice(items, func(i, j int) bool {
    return items[i].Score <= items[j].Score // should be <
})
```

## Good
```go
sort.Slice(items, func(i, j int) bool {
    return items[i].Score < items[j].Score
})
```
</examples>

<patterns>
- `sort.Slice` with `<=` in comparator
- `sort.SliceStable` with `>=` in comparator
- Using `cmp.Compare` incorrectly as a boolean
- Forgetting that equal elements should return `false`
</patterns>

<related>
sortSlice, sloppyLen
</related>
