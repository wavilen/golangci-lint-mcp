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
- `func() (int)` → `func() int`
- `[](int)` → `[]int`
- `map[string](int)` → `map[string]int`
- `*(int)` → `*int`
</patterns>

<related>
typeDefFirst, underef
</related>
