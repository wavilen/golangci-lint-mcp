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
- `sort.Slice` with a simple less-than comparison on the slice elements
- `sort.SliceStable` where `slices.SortStableFunc` would be more type-safe
- `sort.Ints`, `sort.Strings`, `sort.Float64s` where `slices.Sort` handles all types
- Custom `sort.Interface` implementations for simple orderings
- `sort.Slice` with reflection-heavy comparison callbacks
</patterns>

<related>
use-any, use-errors-new, use-fmt-print
