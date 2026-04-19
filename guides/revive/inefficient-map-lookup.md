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
- Checking `m[key]` existence then immediately accessing `m[key]` again
- Using `_, ok := m[k]` for existence check followed by `m[k]` for value
- Double map access in cache hit/miss patterns
- Set-membership checks using `delete` followed by re-check
- Pattern matching on map keys where value is also needed
</patterns>

<related>
datarace
