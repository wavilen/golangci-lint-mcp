# gocritic: boolExprSimplify

<instructions>
Detects boolean expressions that can be simplified. This includes double negations (`!!b`), redundant comparisons (`b == true`, `b != false`), and expressions where De Morgan's laws can reduce complexity.

Remove redundant boolean operations and use the boolean value directly.
</instructions>

<examples>
## Bad
```go
if enabled == true {
	return
}
if !(!ok) {
	return
}
if a || (!a && b) {
	return
}
```

## Good
```go
if enabled {
	return
}
if ok {
	return
}
if a || b {
	return
}
```
</examples>

<patterns>
- Replace `x == true` with `x`
- Replace `x == false` with `!x`
- Replace `x != true` with `!x`
- Replace `!!x` with `x`
- Simplify `x || (!x && y)` to `x || y`
</patterns>

<related>
yodaStyleExpr, elseif
</related>
