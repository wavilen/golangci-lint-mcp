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
- Remove duplicate string keys in map literals
- Remove duplicate constant expressions used as map keys
- Avoid using slices as map keys — use a string key or struct instead
- Remove copy-paste duplicate entries in map literals
</patterns>

<related>
dupArg, dupCase, dupSubExpr
</related>
