# modernize: slicesort

<instructions>
Detects `sort.Slice` and `sort.SliceStable` calls that can be replaced with `slices.Sort` from the standard library (Go 1.21+). The `slices.Sort` function is type-safe, avoids reflection, and produces cleaner code. When the slice elements implement `cmp.Ordered`, `slices.Sort` is a direct drop-in replacement.
</instructions>

<examples>
## Bad
```go
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})
```

## Good
```go
slices.Sort(nums)
```
</examples>

<patterns>
- `sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })` on ordered types — use `slices.Sort`
- `sort.SliceStable(s, ...)` with simple less function — use `slices.SortStableFunc`
- Sorting with `sort.Float64s`, `sort.Ints`, `sort.Strings` — use `slices.Sort`
</patterns>

<related>
sortfunc, sliceclear
