# revive: range

<instructions>
Simplifies verbose range loops where the index or value is unnecessary. Writing `for i := range s` instead of `for i := 0; i < len(s); i++` is more idiomatic and less error-prone. Similarly, use `for range s` when neither the index nor value is needed.

Replace manual index loops with `for i := range` or `for _, v := range` forms. Use `for range` when the body doesn't reference index or value.
</instructions>

<examples>
## Bad
```go
for i := 0; i < len(items); i++ {
    total += items[i]
}
```

## Good
```go
for _, item := range items {
    total += item
}
```
</examples>

<patterns>
- Manual `for i := 0; i < len(s); i++` loops that could be range loops
- Range loops that use only the index to access elements
- Loops where neither index nor value is used (e.g., counting)
- Index-based iteration over maps (which is not even possible in Go)
- C-style for loops over slices or maps
</patterns>

<related>
range-val-address, range-val-in-closure
