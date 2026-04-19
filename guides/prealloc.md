# prealloc

<instructions>
Prealloc detects slice and map declarations that are populated with a known number of elements via `append` or assignment in a loop, without pre-allocating the underlying array. Pre-allocating avoids repeated allocations and copies as the slice grows.

Use `make([]T, 0, n)` or `make(map[K]V, n)` with the known capacity.
</instructions>

<examples>
## Bad
```go
var results []string
for _, item := range items {
    results = append(results, item.Name)
}
```

## Good
```go
results := make([]string, 0, len(items))
for _, item := range items {
    results = append(results, item.Name)
}
```
</examples>

<patterns>
- `var s []T` followed by `append` in a loop where the length is knowable
- Map declared without size hint then populated in a loop
- Slice built from another collection without pre-allocation
- `append` chains in loops that trigger repeated growth allocations
</patterns>

<related>
makezero, govet, staticcheck
