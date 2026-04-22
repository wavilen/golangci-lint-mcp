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
- Extract identical `if`/`else` return statements into a single call before the branch
- Remove ternary-equivalent branches that always produce the same result
- Separate copy-paste errors in error handling branches — differentiate the logic
</patterns>

<related>
dupCase, dupArg, dupSubExpr
</related>
