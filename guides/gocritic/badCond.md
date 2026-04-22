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
- Simplify redundant range checks like `x > 5 && x > 3` to just `x > 5`
- Eliminate tautological comparisons like `x == x` or `x != x`
- Eliminate double negation `!!cond` — use `cond` directly
- Eliminate overlapping conditions in if-else chains
- Replace contradictory conditions like `x > 10 && x < 5` — always false
</patterns>

<related>
weakCond, dupSubExpr, offBy1
</related>
