# revive: increment-decrement

<instructions>
Enforces using `i++` and `i--` instead of `i += 1` and `i -= 1`. Go has dedicated increment and decrement operators for the common case of adding or subtracting 1. Using the compound assignment form for this is unnecessarily verbose.

Replace `x += 1` with `x++` and `x -= 1` with `x--`.
</instructions>

<examples>
## Good
```go
i++
count--
```
</examples>

<patterns>
- Use `i++` instead of `i += 1` for single increments
- Replace `i -= 1` with `i--` for single decrements
- Simplify auto-generated `+= 1` to the `++` operator
- Replace verbose `+= n` with `++` when `n` is always 1
- Use `++` for loop counter increments instead of the verbose form
</patterns>

<related>
revive/bool-literal-in-expr
</related>
