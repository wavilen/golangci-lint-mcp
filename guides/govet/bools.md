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
- Comparing boolean to `true` or `false` (`x == true`, `x != false`)
- Redundant logical AND of same expression (`x && x`)
- Double negation (`!!x`)
- Tautological boolean expression (always true/false)
</patterns>

<related>
assign, nilfunc
</related>
