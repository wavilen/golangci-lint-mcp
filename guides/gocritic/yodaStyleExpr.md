# gocritic: yodaStyleExpr

<instructions>
Detects Yoda-style conditions where the constant is placed on the left side of the comparison: `nil == x` instead of `x == nil`. While valid Go, this style goes against Go conventions where the variable comes first.

Place the variable or expression on the left side of the comparison.
</instructions>

<examples>
## Good
```go
if err == nil {
	return
}
if len(items) == 0 {
	return
}
if role == "admin" {
	grantAccess()
}
```
</examples>

<patterns>
- Reverse Yoda-style comparisons: `nil == x` → `x == nil`
- Reverse Yoda-style comparisons: `0 == len(s)` → `len(s) == 0` or `s == ""`
- Reverse constant-first comparisons: `"value" == variable` → `variable == "value"`
</patterns>

<related>
gocritic/boolExprSimplify, gocritic/emptyStringTest
</related>
