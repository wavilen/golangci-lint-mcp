# gocritic: badSorting

<instructions>
Detects calls to `sort.Slice` or similar where the comparison uses `<=` or `>=` instead of strict `<`. The `sort.Slice` comparison function must return whether element `i` should sort before element `j`, and must use strict less-than to avoid non-deterministic ordering and potential infinite loops.

Use strict `<` in sort comparison functions instead of `<=`.
</instructions>

<examples>
## Good
```go
sort.Slice(items, func(i, j int) bool {
    return items[i].Score < items[j].Score
})
```
</examples>

<patterns>
- Use strict `<` in `sort.Slice` comparators — never `<=`
- Use strict `<` in `sort.SliceStable` comparators — never `>=`
- Replace `cmp.Compare` used as a boolean with a direct comparison
- Return `false` for equal elements in sort comparators to avoid instability
</patterns>

<related>
gocritic/sortSlice, gocritic/sloppyLen
</related>
