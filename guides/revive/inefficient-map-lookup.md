# revive: inefficient-map-lookup

<instructions>
Detects patterns where a map lookup is performed twice — once to check existence and again to retrieve the value. This is inefficient; Go's comma-ok idiom (`val, ok := m[key]`) does both in a single access.

Replace the double lookup with a single comma-ok map access.
</instructions>

<examples>
## Bad
```go
if _, ok := cache[key]; ok {
    result := cache[key] // second lookup
    process(result)
}
```

## Good
```go
if result, ok := cache[key]; ok {
    process(result)
}
```
</examples>

<patterns>
- Use comma-ok `val, ok := m[key]` to check existence and retrieve value in one access
- Replace `_, ok := m[k]` followed by `m[k]` with a single comma-ok lookup
- Combine double map access in cache hit/miss patterns into one lookup
- Use the ok value directly instead of checking membership with `delete` followed by re-check
- Use comma-ok to get the value alongside the existence check when pattern matching on map keys
</patterns>

<related>
datarace
