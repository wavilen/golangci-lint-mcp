# gocritic: mapKey

<instructions>
Detects invalid or problematic map key types, particularly maps with keys that have underlying `[]byte`, `sync.Mutex`, or other non-comparable types. Also flags map literals with duplicate keys where the later value silently overwrites the earlier one.

Ensure map key types are comparable and remove duplicate keys from map literals.
</instructions>

<examples>
## Bad
```go
_ = map[string]int{
    "a": 1,
    "b": 2,
    "a": 3, // duplicate key — overwrites value 1
}
```

## Good
```go
_ = map[string]int{
    "a": 3,
    "b": 2,
}
```
</examples>

<patterns>
- Duplicate string keys in map literals
- Duplicate constant expressions as map keys
- Using slices as map keys (compile error, but gocritic may catch related patterns)
- Map literals where copy-paste introduced duplicate entries
</patterns>

<related>
dupArg, dupCase, dupSubExpr
</related>
