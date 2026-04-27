# revive: bool-literal-in-expr

<instructions>
Detects redundant boolean literals in expressions, such as `== true`, `== false`, `!= true`, or wrapping a boolean in `if` when the expression itself is already boolean. These are noisy and suggest the author was unsure of the type.

Remove the comparison to the boolean literal and use the boolean expression directly, or apply logical negation with `!`.
</instructions>

<examples>
## Good
```go
if isActive {
    process()
}
if !isEnabled {
    disable()
}
if !isValid {
    handleError()
}
```
</examples>

<patterns>
- Use the boolean variable directly instead of comparing to `true` or `false`
- Use the `!` operator instead of writing `!= true`
- Simplify double negation `!(!flag)` to just `flag`
- Return boolean expressions directly instead of wrapping in `== true` comparisons
- Replace ternary-style `== true` conditions with the expression itself
</patterns>

<related>
revive/constant-logical-expr
</related>
