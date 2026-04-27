# intrange

<instructions>
Intrange detects `for` loops with integer indices that can be simplified using Go 1.22+ range-over-integer syntax. The `for i := range n` form is clearer and eliminates the common `i++` pattern.

Replace `for i := 0; i < n; i++` with `for i := range n` when the index starts at 0 and increments by 1.
</instructions>

<examples>
## Good
```go
for i := range items {
    process(items[i])
}
```
</examples>

<patterns>
- Replace `for i := 0; i < n; i++` with `for i := range n` when the index starts at 0 and steps by 1
- Use `for i := range n` for simple integer iteration (Go 1.22+)
- Convert `for i := 0; i < len(slice); i++` to `for i := range slice` for collection iteration
- Prefer range-over-int over three-clause loops that only increment by 1
</patterns>

<related>
prealloc, perfsprint, copyloopvar
</related>
