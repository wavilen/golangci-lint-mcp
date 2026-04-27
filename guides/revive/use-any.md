# revive: use-any

<instructions>
Suggests using `any` instead of `interface{}` for type declarations, function parameters, and return types. Since Go 1.18, `any` is a built-in type alias for `interface{}` — they are identical at runtime. Using `any` is more concise and idiomatic in modern Go.

Replace every occurrence of `interface{}` with `any`. They are type aliases, so no behavior changes.
</instructions>

<examples>
## Good
```go
func Process(data any) error
var items []any
func Marshal(v any) ([]byte, error)
```
</examples>

<patterns>
- Replace `interface{}` with `any` in function signatures and variable declarations
- Use `any` as the map value type instead of `interface{}`
- Replace `[]interface{}` with `[]any` for slice types
- Use `any` in generic constraint positions instead of `interface{}`
- Replace legacy pre-Go 1.18 code by replacing `interface{}` with `any`
</patterns>

<related>
revive/use-errors-new, revive/use-fmt-print, revive/use-slices-sort
</related>
