# gocritic: builtinShadow

<instructions>
Detects when variables or type parameters shadow built-in identifiers such as `len`, `cap`, `append`, `new`, `make`, `copy`, `delete`, `close`, `panic`, `recover`, `print`, `println`, `complex`, `real`, `imag`, `true`, `false`, `nil`, `iota`, or `error`.

Rename the shadowing identifier to avoid confusion and preserve access to the built-in.
</instructions>

<examples>
## Bad
```go
func process(len int) error {
	for i := 0; i < len; i++ {
		_ = i
	}
	return nil
}
```

## Good
```go
func process(length int) error {
	for i := 0; i < length; i++ {
		_ = i
	}
	return nil
}
```
</examples>

<patterns>
- Function parameter named `len`, `cap`, `copy`, or `new`
- Local variable named `true`, `false`, or `nil`
- Receiver named `error` or `recover`
- Loop variable shadowing `iota` or `append`
</patterns>

<related>
builtinShadowDecl, importShadow
</related>
