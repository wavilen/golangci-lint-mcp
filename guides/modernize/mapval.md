# modernize: mapval

<instructions>
Detects double map lookup patterns where the same key is accessed twice: once to check existence and again to use the value. This is wasteful when the value can be captured in the initial lookup. Use the value returned from the existence check directly instead of performing a second lookup.
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
if v, ok := cache[key]; ok {
    process(v) // use value from the same lookup
}
```
</examples>

<patterns>
- `_, ok := m[k]` followed by `m[k]` — capture the value in the initial lookup
- `if m[k] != ""` followed by `use(m[k])` — use comma-ok idiom instead
- Checking existence then immediately retrieving — always use the returned value
</patterns>

<related>
maprange, errorf
