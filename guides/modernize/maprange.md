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
- Use `maps.Keys(m)` instead of collecting map keys into a slice with a manual loop
- Use `maps.Values(m)` instead of collecting map values into a slice with a manual loop
- Use `maps.Keys` instead of `for k := range m { append(keys, k) }`
</patterns>

<related>
mapval, simplifyrange
