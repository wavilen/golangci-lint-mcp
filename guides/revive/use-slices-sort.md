# revive: use-slices-sort

<instructions>
Suggests using `slices.Sort` from the standard library (Go 1.21+) instead of `sort.Slice` or `sort.SliceStable`. The `slices.Sort` function is type-safe, avoids reflection, and produces cleaner code without the comparison function boilerplate.

Replace `sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })` with `slices.Sort(s)`. For custom ordering, use `slices.SortFunc`.
</instructions>

<examples>
## Bad
```go
sort.Slice(names, func(i, j int) bool {
    return names[i] < names[j]
})
```

## Good
```go
slices.Sort(names)
```
</examples>

<patterns>
- Replace `sort.Slice` with `slices.Sort` for simple less-than comparisons on slice elements
- Use `slices.SortStableFunc` instead of `sort.SliceStable` for type-safe stable sorting
- Replace `sort.Ints`, `sort.Strings`, and `sort.Float64s` with `slices.Sort`
- Use `slices.SortFunc` instead of custom `sort.Interface` implementations for simple orderings
- Replace `sort.Slice` with `slices.Sort` to avoid reflection-heavy comparison callbacks
</patterns>

<related>
use-any, use-errors-new, use-fmt-print
