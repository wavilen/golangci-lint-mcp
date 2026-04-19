# dogsled

<instructions>
Dogsled checks for excessive blank identifiers in assignments. Functions returning many values where most are discarded with `_` indicate unclear intent and make code harder to maintain.

Replace long blank identifier runs with a struct return type, an options pattern, or explicit named variables. If discarding is intentional, add a comment or use a helper variable.
</instructions>

<examples>
## Bad
```go
_, _, _, _, err := parseConfig(data)
```

## Good
```go
result, err := parseConfig(data)
// result has named fields: Host, Port, Timeout, Retry
```
</examples>

<patterns>
- Multi-return functions where only the error is used
- Parsing functions returning many optional fields
- Legacy APIs with large return tuples
- Swallowed return values from type assertion or map operations
</patterns>

<related>
unparam, errcheck, funlen
</related>
