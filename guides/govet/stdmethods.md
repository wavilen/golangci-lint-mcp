# govet: stdmethods

<instructions>
Reports methods with names matching standard Go interfaces (`String`, `Error`, `Read`, `Write`, `MarshalJSON`) but with incorrect signatures. A wrong signature means the method silently fails to satisfy the interface — `fmt.Stringer`, `error`, `io.Reader` — without any compiler error. Fix the method signature to match the standard interface exactly.
</instructions>

<examples>
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
govet/composites, govet/structtag
</related>
