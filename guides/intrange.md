# intrange

<instructions>
Intrange detects `for` loops with integer indices that can be simplified using Go 1.22+ range-over-integer syntax. The `for i := range n` form is clearer and eliminates the common `i++` pattern.

Replace `for i := 0; i < n; i++` with `for i := range n` when the index starts at 0 and increments by 1.
</instructions>

<examples>
## Bad
```go
for i := 0; i < len(items); i++ {
    process(items[i])
}
```

## Good
```go
for i := range items {
    process(items[i])
}
```
</examples>

<patterns>
- Classic three-clause for loops where index starts at 0 and steps by 1
- `for i := 0; i < n; i++` patterns that can use range-over-int
- Iterating with an index variable that only increments by 1
- Using `for i := 0; i < len(slice); i++` instead of range
</patterns>

<related>
prealloc, perfsprint, copyloopvar
