# gocritic: badCond

<instructions>
Detects suspicious conditional expressions that are always true, always false, or redundant. This includes conditions like `x == x`, overlapping range checks (`x > 5 && x > 3`), and double negations. These usually indicate a typo or logic error.

Simplify or correct the condition. Replace redundant comparisons with a single check, and fix overlapping or contradictory conditions.
</instructions>

<examples>
## Bad
```go
if x > 5 && x > 3 { // second check is redundant
    process(x)
}
```

## Good
```go
if x > 5 {
    process(x)
}
```
</examples>

<patterns>
- Redundant range checks: `x > 5 && x > 3`
- Tautological comparisons: `x == x` or `x != x`
- Double negation: `!!cond`
- Overlapping conditions in if-else chains
- Contradictory conditions: `x > 10 && x < 5`
</patterns>

<related>
weakCond, dupSubExpr, offBy1
</related>
