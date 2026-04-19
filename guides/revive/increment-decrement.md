# revive: increment-decrement

<instructions>
Enforces using `i++` and `i--` instead of `i += 1` and `i -= 1`. Go has dedicated increment and decrement operators for the common case of adding or subtracting 1. Using the compound assignment form for this is unnecessarily verbose.

Replace `x += 1` with `x++` and `x -= 1` with `x--`.
</instructions>

<examples>
## Bad
```go
i += 1
count -= 1
```

## Good
```go
i++
count--
```
</examples>

<patterns>
- Developers coming from languages where `++` is discouraged (e.g., Go style guides)
- Auto-generated code using `+= 1` instead of `++`
- Code formatted by tools that expand increment operators
- Consistent use of `+= n` even when `n` is always 1
- Manual loop counter increments using verbose form
</patterns>

<related>
bool-literal-in-expr
