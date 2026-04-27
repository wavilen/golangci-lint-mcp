# modernize: mapval

<instructions>
Detects double map lookup patterns where the same key is accessed twice: once to check existence and again to use the value. This is wasteful when the value can be captured in the initial lookup. Use the value returned from the existence check directly instead of performing a second lookup.
</instructions>

<examples>
## Good
```go
if v, ok := cache[key]; ok {
    process(v) // use value from the same lookup
}
```
</examples>

<patterns>
- Use the value from the initial lookup instead of `_, ok := m[k]` followed by `m[k]`
- Use comma-ok idiom `v, ok := m[k]` instead of `if m[k] != ""` followed by `use(m[k])`
- Use the value returned from the existence check instead of double lookups
</patterns>

<related>
modernize/maprange, modernize/errorf
</related>
