# gocritic: yodaStyleExpr

<instructions>
Detects Yoda-style conditions where the constant is placed on the left side of the comparison: `nil == x` instead of `x == nil`. While valid Go, this style goes against Go conventions where the variable comes first.

Place the variable or expression on the left side of the comparison.
</instructions>

<examples>
## Bad
```go
if nil == err {
	return
}
if 0 == len(items) {
	return
}
if "admin" == role {
	grantAccess()
}
```

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
- `nil == x` → `x == nil`
- `0 == len(s)` → `len(s) == 0` (or use `s == ""` / `len(s) == 0`)
- Constant-first comparisons: `"value" == variable`
</patterns>

<related>
boolExprSimplify, emptyStringTest
</related>
