# gocritic: regexpMust

<instructions>
Detects `regexp.MustCompile` calls with a string literal that could be replaced with `regexp.MustCompile` assigned to a package-level variable. Compiling a regex on every function call is wasteful — regex compilation is expensive and the result is immutable.

Move `regexp.MustCompile` to a package-level `var` so the pattern is compiled once at init time.
</instructions>

<examples>
## Bad
```go
func match(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}
```

## Good
```go
var digitPattern = regexp.MustCompile(`^\d+$`)

func match(s string) bool {
	return digitPattern.MatchString(s)
}
```
</examples>

<patterns>
- `regexp.MustCompile` called inside a function body
- Regex pattern is a constant string literal
- `regexp.MustCompile` called in a loop or hot path
</patterns>

<related>
regexpSimplify, regexpPattern
</related>
