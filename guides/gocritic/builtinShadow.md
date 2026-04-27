# gocritic: builtinShadow

<instructions>
Detects when variables or type parameters shadow built-in identifiers such as `len`, `cap`, `append`, `new`, `make`, `copy`, `delete`, `close`, `panic`, `recover`, `print`, `println`, `complex`, `real`, `imag`, `true`, `false`, `nil`, `iota`, or `error`.

Rename the shadowing identifier to avoid confusion and preserve access to the built-in.
</instructions>

<examples>
## Good
```go
func process(length int) error {
	for i := range length {
		_ = i
	}
	return nil
}
```
</examples>

<patterns>
- Rename parameters that shadow builtins — avoid `len`, `cap`, `copy`, `new` as parameter names
- Avoid `true`, `false`, `nil` as local variable names
- Rename receivers that shadow builtins — avoid `error` or `recover` as receiver names
- Avoid loop variables named `iota` or `append` — they shadow predeclared identifiers
</patterns>

<related>
gocritic/builtinShadowDecl, gocritic/importShadow
</related>
