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
- Replace manual `for i := 0; i < len(s); i++` loops with `for i := range s`
- Use `for _, v := range` to access elements directly instead of indexing with `s[i]`
- Simplify loops to `for range s` when neither index nor value is needed
- Use `for k, v := range m` for map iteration instead of index-based approaches
- Convert C-style for loops over slices to range loops
</patterns>

<related>
range-val-address, range-val-in-closure
