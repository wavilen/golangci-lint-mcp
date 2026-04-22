# revive: optimize-operands-order

<instructions>
Suggests reordering operands in comparison expressions to place the stable or constant value on the left side. This follows the "Yoda conditions" principle and helps catch accidental assignment (`=`) instead of comparison (`==`) errors in some languages. In Go, the readability benefit is primary — placing the expected value first makes the intent clearer.

Swap the operands so the constant or invariant value appears on the left of the comparison operator.
</instructions>

<examples>
## Bad
```go
if x == 42 {
    doSomething()
}
if name != "" {
    process(name)
}
```

## Good
```go
if 42 == x {
    doSomething()
}
if "" != name {
    process(name)
}
```
</examples>

<patterns>
- Move constant values on the left side of comparisons (e.g., `42 == x`)
- Move empty string or nil checks to put the constant on the left (e.g., `"" != name`)
- Reorder numeric comparisons to put the expected value first
- Use consistent operand ordering in boolean expressions
- Use a consistent operand ordering style across the codebase
</patterns>

<related>
bool-literal-in-expr, unnecessary-if
