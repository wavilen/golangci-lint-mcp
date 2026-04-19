# makezero

<instructions>
Makezero detects slice declarations with a non-zero initial length passed to `make` where the intent is likely to `append` — the pre-allocated elements contain zero values that get mixed with appended data. Use a zero length with the desired capacity instead.
</instructions>

<examples>
## Bad
```go
items := make([]string, 10)
items = append(items, "hello")
// items[0:10] are empty strings, "hello" is at index 10
```

## Good
```go
items := make([]string, 0, 10)
items = append(items, "hello")
```
</examples>

<patterns>
- `make([]T, n)` followed by `append` — zero-value elements pollute the slice
- Pre-allocated slices used only with `append`, never with index assignment
- `make([]T, n, m)` where length n > 0 but only `append` is used
</patterns>

<related>
prealloc, govet, staticcheck
