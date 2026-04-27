# modernize: slicesort

<instructions>
Detects `sort.Slice` and `sort.SliceStable` calls that can be replaced with `slices.Sort` from the standard library (Go 1.21+). The `slices.Sort` function is type-safe, avoids reflection, and produces cleaner code. When the slice elements implement `cmp.Ordered`, `slices.Sort` is a direct drop-in replacement.
</instructions>

<examples>
## Good
```go
slices.Sort(nums)
```
</examples>

<patterns>
- Use `slices.Sort` instead of `sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })` on ordered types
- Use `slices.SortStableFunc` instead of `sort.SliceStable(s, ...)` with simple less functions
- Use `slices.Sort` instead of `sort.Float64s`, `sort.Ints`, `sort.Strings`
</patterns>

<related>
modernize/sortfunc, modernize/sliceclear
</related>
