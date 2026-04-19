# revive: constant-logical-expr

<instructions>
Detects logical expressions that always evaluate to the same value regardless of their inputs. This includes contradictions like `x && !x`, tautologies like `x || !x`, and expressions involving constant booleans. These indicate a logic error or dead code.

Simplify or remove the constant expression. If one side is always true or false, the other side is dead code.
</instructions>

<examples>
## Bad
```go
if isValid && !isValid { // always false
    handleBoth()
}
if !hasPermission || hasPermission { // always true
    proceed()
}
if true && isEnabled { // equivalent to just isEnabled
    activate()
}
```

## Good
```go
if isValid {
    handleBoth()
}
if !hasPermission {
    requestAccess()
}
proceed() // always reached anyway
```
</examples>

<patterns>
- Contradictory conditions like `x && !x` or `x || !x`
- Hardcoded boolean literals in logical expressions (`true && x`)
- Double negation that cancels out (`!!x` is just `x`)
- Copy-paste errors producing always-true or always-false guards
- Stale conditions left after refactoring
</patterns>

<related>
bool-literal-in-expr, identical-branches
