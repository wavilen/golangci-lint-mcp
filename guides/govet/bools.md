# govet: bools

<instructions>
Detects redundant boolean expressions in comparisons and logical operations. Examples include comparing a boolean to `true` or `false` (`x == true`), double negation (`!!x`), and always-true or always-false conditions (`x && x`).

Simplify boolean expressions by removing redundant comparisons and operators.
</instructions>

<examples>
## Bad
```go
if isEnabled == true {
    return flag != false
}
```

## Good
```go
if isEnabled {
    return flag
}
```
</examples>

<patterns>
- Remove explicit boolean comparisons — use `x` instead of `x == true`, `!x` instead of `x == false`
- Remove redundant boolean operations (`x && x` → `x`)
- Remove double negation (`!!x` → `x`)
- Simplify always-true/false boolean expressions to their constant value
</patterns>

<related>
assign, nilfunc
</related>
