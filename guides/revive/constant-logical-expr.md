# revive: constant-logical-expr

<instructions>
Detects logical expressions that always evaluate to the same value regardless of their inputs. This includes contradictions like `x && !x`, tautologies like `x || !x`, and expressions involving constant booleans. These indicate a logic error or dead code.

Simplify or remove the constant expression. If one side is always true or false, the other side is dead code.
</instructions>

<examples>
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
- Simplify contradictory conditions like `x && !x` or `x || !x` — they indicate a logic error
- Remove hardcoded boolean literals from logical expressions (`true && x` → `x`)
- Eliminate double negation that cancels out (`!!x` → `x`)
- Replace copy-paste errors that produce always-true or always-false guards
- Remove stale conditions left after refactoring that always evaluate to a constant
</patterns>

<related>
revive/bool-literal-in-expr, revive/identical-branches
</related>
