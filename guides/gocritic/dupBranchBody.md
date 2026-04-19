# gocritic: dupBranchBody

<instructions>
Detects `if-else` or `switch` branches with identical (or nearly identical) bodies. When two branches do the same thing, the conditional logic is either unnecessary or contains a bug where one branch should behave differently.

Combine the branches or fix the body of one branch to perform the correct distinct action.
</instructions>

<examples>
## Bad
```go
if useCache {
    return fetchFromDB(key)
} else {
    return fetchFromDB(key) // identical body
}
```

## Good
```go
if useCache {
    return fetchFromCache(key)
}
return fetchFromDB(key)
```
</examples>

<patterns>
- `if`/`else` with identical return statements
- Switch cases with duplicated logic
- Ternary-equivalent branches that always produce the same result
- Copy-paste errors in error handling branches
</patterns>

<related>
dupCase, dupArg, dupSubExpr
</related>
