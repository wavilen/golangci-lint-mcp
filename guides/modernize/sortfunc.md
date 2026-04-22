# modernize: sortfunc

<instructions>
Detects `sort.Slice` or `sort.SliceStable` calls with comparison functions that can be replaced with `slices.SortFunc` from Go 1.21+. The new API uses a two-argument comparator returning `int` via `cmp.Compare` or `cmp.Or`, which is more readable and type-safe than index-based `func(i, j int) bool` callbacks.
</instructions>

<examples>
## Bad
```go
sort.Slice(users, func(i, j int) bool {
    return users[i].Name < users[j].Name
})
```

## Good
```go
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Name, b.Name)
})
```
</examples>

<patterns>
- Use `slices.SortFunc` instead of `sort.Slice(x, func(i, j) bool { return x[i].Field < x[j].Field })`
- Use `slices.SortStableFunc` instead of `sort.SliceStable` with custom comparators
- Use `cmp.Or` chain in `SortFunc` instead of multi-field `sort.Slice`
</patterns>

<related>
slicesort, sliceclear
