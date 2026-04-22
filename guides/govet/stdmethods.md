# govet: stdmethods

<instructions>
Reports methods that look like standard interface methods but have incorrect signatures. For example, a `String()` method that returns `int` instead of `string`, or an `Error()` method that takes parameters. These methods will never satisfy the intended interface.

Fix the method signature to match the standard interface exactly.
</instructions>

<examples>
## Bad
```go
func (m MyType) String() int { // should return string
    return m.id
}
```

## Good
```go
func (m MyType) String() string {
    return fmt.Sprintf("MyType(%d)", m.id)
}
```
</examples>

<patterns>
- Ensure `String()` methods return `string` — fix the return type
- Ensure `Error()` methods return `string` — fix the return type
- Ensure `Read(p []byte)` returns `(int, error)` — match the `io.Reader` signature
- Ensure `Write(p []byte)` returns `(int, error)` — match the `io.Writer` signature
- Ensure `MarshalJSON()` returns `([]byte, error)` — match the `json.Marshaler` signature
</patterns>

<related>
composites, structtag
</related>
