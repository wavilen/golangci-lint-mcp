# gocritic: sortSlice

<instructions>
Detects incorrect usage of `sort.Slice` where the comparison function captures the wrong variable or compares indices instead of values. Also flags when `sort.Slice` is used on a nil slice or when a more specific sort function (like `sort.Ints`, `sort.Strings`) would be clearer.

Use the correct slice elements in the comparison function, and prefer specialized sort functions for simple types.
</instructions>

<examples>
## Bad
```go
sort.Slice(names, func(i, j int) bool {
    return i < j // comparing indices, not values
})
```

## Good
```go
sort.Slice(names, func(i, j int) bool {
    return names[i] < names[j]
})
```
</examples>

<patterns>
- Use `slice[i] < slice[j]` not `i < j` in `sort.Slice` comparators
- Use indexed elements `slice[i]` in the comparator — avoid capturing outer variables
- Replace `sort.Slice` on `[]string` with `sort.Strings(slice)`
- Replace `sort.Slice` on `[]int` with `sort.Ints(slice)`
</patterns>

<related>
badSorting, sloppyLen, offBy1
</related>
