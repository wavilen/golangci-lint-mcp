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
- Iterating over slices of large structs with `for _, v := range` and only reading `v`
- Loop variables that are structs with large embedded arrays or many fields
- Processing batches of large value-type elements where only a few fields are accessed
- Any range loop where the value variable size exceeds ~80 bytes
</patterns>

<related>
rangeExprCopy, hugeParam, indexAlloc
