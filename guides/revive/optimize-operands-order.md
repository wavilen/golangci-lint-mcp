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
- Comparisons with constant values on the right side
- Empty string or nil checks with the variable on the left
- Numeric comparisons where the expected value could come first
- Boolean expressions in conditions where ordering is inconsistent
- Mixed operand ordering style across a codebase
</patterns>

<related>
bool-literal-in-expr, unnecessary-if
