# modernize: maprange

<instructions>
Detects range over a map where `maps.Keys` or `maps.Values` (Go 1.21+) would be more expressive. When you only need the keys or values of a map as a slice, use the `maps` package functions instead of collecting them manually in a loop.
</instructions>

<examples>
## Bad
```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
```

## Good
```go
keys := maps.Keys(m)
```
</examples>

<patterns>
- Collecting map keys into a slice in a loop — use `maps.Keys(m)`
- Collecting map values into a slice — use `maps.Values(m)`
- `for k := range m` followed by `append(keys, k)` — use `maps.Keys`
</patterns>

<related>
mapval, simplifyrange
