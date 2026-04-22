# gocritic: dupSubExpr

<instructions>
Detects identical sub-expressions on both sides of a binary operator or within a compound expression. Operations like `x == x`, `x && x`, `x - x`, or `a[i] - a[i]` are almost always copy-paste mistakes. The second operand was likely meant to be different.

Replace the duplicate sub-expression with the correct intended operand.
</instructions>

<examples>
## Bad
```go
if x.Min - x.Min > threshold { // x.Min subtracted from itself
    return true
}
```

## Good
```go
if x.Max - x.Min > threshold {
    return true
}
```
</examples>

<patterns>
- Remove self-comparisons like `x == x` or `x != x`
- Remove self-operations like `x - x` or `x / x`
- Replace `a[i] == a[i]` with different indices like `a[i] == a[j]`
- Replace `s[i] - s[i]` in loops with `s[i] - s[j]`
- Remove redundant boolean like `x && x` or `x || x` — use `x` directly
</patterns>

<related>
dupArg, badCond, weakCond
</related>
