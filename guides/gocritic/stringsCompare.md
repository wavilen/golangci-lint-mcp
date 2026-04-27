# gocritic: stringsCompare

<instructions>
Detects uses of `strings.Compare` that can be replaced with simpler operators. `strings.Compare(a, b) == 0` is equivalent to `a == b`, and `strings.Compare(a, b) < 0` is equivalent to `a < b`. The `Compare` function exists for interop with other systems and should not be used for ordinary comparisons.

Use the standard comparison operators `==`, `<`, `>` instead of `strings.Compare`.
</instructions>

<examples>
## Good
```go
if a == b {
	return true
}
```
</examples>

<patterns>
- Replace `strings.Compare(a, b) == 0` with `a == b`
- Replace `strings.Compare(a, b) < 0` with `a < b`
- Replace `strings.Compare(a, b) > 0` with `a > b`
- Replace `strings.Compare(a, b) != 0` with `a != b`
</patterns>

<related>
gocritic/stringConcatSimplify, gocritic/emptyStringTest
</related>
