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
- `String()` not returning `string`
- `Error()` not returning `string`
- `Read(p []byte)` not returning `(int, error)`
- `Write(p []byte)` not returning `(int, error)`
- `MarshalJSON()` not returning `([]byte, error)`
</patterns>

<related>
composites, structtag
</related>
