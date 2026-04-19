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
- `x == x` or `x != x` — self-comparison
- `x - x` or `x / x` — self-operation
- `a[i] == a[i]` — same index on both sides
- `s[i] - s[i]` in loop bodies — likely meant `s[j]`
- `x && x` or `x || x` — redundant boolean
</patterns>

<related>
dupArg, badCond, weakCond
</related>
