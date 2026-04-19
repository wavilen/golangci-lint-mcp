# revive: use-any

<instructions>
Suggests using `any` instead of `interface{}` for type declarations, function parameters, and return types. Since Go 1.18, `any` is a built-in type alias for `interface{}` — they are identical at runtime. Using `any` is more concise and idiomatic in modern Go.

Replace every occurrence of `interface{}` with `any`. They are type aliases, so no behavior changes.
</instructions>

<examples>
## Bad
```go
func Process(data interface{}) error
var items []interface{}
func Marshal(v interface{}) ([]byte, error)
```

## Good
```go
func Process(data any) error
var items []any
func Marshal(v any) ([]byte, error)
```
</examples>

<patterns>
- `interface{}` in function signatures and variable declarations
- Map types using `interface{}` as the value type
- Slice types `[]interface{}` instead of `[]any`
- Empty interface in generic constraint positions
- Legacy code written before Go 1.18 using the old spelling
</patterns>

<related>
use-errors-new, use-fmt-print, use-slices-sort
