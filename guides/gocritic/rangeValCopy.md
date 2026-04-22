# gocritic: rangeValCopy

<instructions>
Detects `for _, v := range` loops where `v` is a large value type. Each iteration copies the current element into `v`, which is expensive for large structs. Use an index loop or take a pointer to avoid the per-element copy.

Change `for _, v := range slice` to `for i := range slice` and access `slice[i]` directly, or iterate over a slice of pointers.
</instructions>

<examples>
## Bad
```go
type Record struct {
    Data [1024]byte
    Name string
}

records := []Record{{}, {}}
for _, r := range records {
    // each Record (~1KB) is copied per iteration
    _ = r.Name
}
```

## Good
```go
for i := range records {
    r := &records[i]
    _ = r.Name
}
```
</examples>

<patterns>
- Use `for i := range slice` with index access instead of `for _, v := range` for large structs
- Use pointer elements in the slice instead of copying large value-type elements
- Replace `for _, v := range` with index access when only a few fields are needed from large structs
- Use `for i := range` when the value variable size exceeds ~80 bytes
</patterns>

<related>
rangeExprCopy, hugeParam, indexAlloc
