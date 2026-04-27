# govet: unsafeptr

<instructions>
Reports invalid `unsafe.Pointer` conversions. Go's garbage collector relies on `unsafe.Pointer` safety rules: a `uintptr` may not be stored in a variable and later converted back to `unsafe.Pointer`, because the GC may move the underlying object. Arithmetic on `unsafe.Pointer` values is also invalid.

Perform pointer arithmetic in a single expression: `unsafe.Pointer(uintptr(p) + offset)`. Do not store intermediate `uintptr` values.
</instructions>

<examples>
## Good
```go
p := unsafe.Pointer(&x)
q := unsafe.Pointer(uintptr(p) + 16) // single expression — safe
```
</examples>

<patterns>
- Convert `uintptr` to `unsafe.Pointer` in the same expression — never store intermediate `uintptr` values
- Use `uintptr` for pointer arithmetic then immediately convert back — never compute on `unsafe.Pointer` directly
- Avoid converting `reflect.Value.Pointer()` results to `unsafe.Pointer` — use the reflect API instead
</patterns>

<related>
govet/cgocall, govet/framepointer
</related>
