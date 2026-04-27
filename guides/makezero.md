# makezero

<instructions>
Makezero detects slice declarations with a non-zero initial length passed to `make` where the intent is likely to `append` — the pre-allocated elements contain zero values that get mixed with appended data. Use a zero length with the desired capacity instead.
</instructions>

<examples>
## Good
```go
items := make([]string, 0, 10)
items = append(items, "hello")
```
</examples>

<patterns>
- Use `make([]T, 0, n)` when the slice will be filled with `append`
- Set initial length to 0 and use capacity-only allocation for append-only slices
- Change `make([]T, n, m)` to `make([]T, 0, m)` when only `append` is used
</patterns>

<related>
prealloc, govet, staticcheck
</related>
