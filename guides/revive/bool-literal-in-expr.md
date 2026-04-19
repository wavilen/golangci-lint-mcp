# revive: bool-literal-in-expr

<instructions>
Detects redundant boolean literals in expressions, such as `== true`, `== false`, `!= true`, or wrapping a boolean in `if` when the expression itself is already boolean. These are noisy and suggest the author was unsure of the type.

Remove the comparison to the boolean literal and use the boolean expression directly, or apply logical negation with `!`.
</instructions>

<examples>
## Bad
```go
if isActive == true {
    process()
}
if isEnabled == false {
    disable()
}
if ! (isValid == true) {
    handleError()
}
```

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
- Comparing a boolean variable to `true` or `false`
- Using `!= true` instead of the `!` operator
- Double negation patterns like `!(!flag)`
- Wrapping boolean function returns in redundant comparisons
- Ternary-style expressions rewritten as `== true` conditions
</patterns>

<related>
constant-logical-expr, simplify-computation
