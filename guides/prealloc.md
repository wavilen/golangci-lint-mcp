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
- Preallocate slices with `make([]T, 0, n)` when the final size is knowable
- Preallocate maps with `make(map[K]V, n)` when populating from a known-size loop
- Use `make([]T, 0, len(src))` when building a slice from another collection
- Avoid `append` chains in loops without pre-allocation — they trigger repeated growth
</patterns>

<related>
makezero, govet, staticcheck
