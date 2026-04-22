# gocritic: typeUnparen

<instructions>
Detects unnecessary parentheses around type expressions. Go's type syntax doesn't require parentheses in most contexts — `func()(int)` should be `func() int`, and `[](int)` should be `[]int`.

Remove unnecessary parentheses from type expressions.
</instructions>

<examples>
## Bad
```go
func getValue() (int) {
	return 42
}
var nums [](int)
var data map[string](int)
```

## Good
```go
func getValue() int {
	return 42
}
var nums []int
var data map[string]int
```
</examples>

<patterns>
- Remove unnecessary parentheses: `func() (int)` → `func() int`
- Remove unnecessary parentheses: `[](int)` → `[]int`
- Remove unnecessary parentheses: `map[string](int)` → `map[string]int`
- Remove unnecessary parentheses: `*(int)` → `*int`
</patterns>

<related>
typeDefFirst, underef
</related>
