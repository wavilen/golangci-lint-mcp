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
- Replace excessive blank identifiers with a struct return type or named variables
- Wrap parsing function returns in a result struct instead of discarding with `_`
- Extract legacy API returns into a dedicated struct to avoid long `_` runs
- Capture type assertion or map operation results explicitly instead of discarding
</patterns>

<related>
unparam, errcheck, funlen
</related>
