# govet: unsafeptr

<instructions>
Reports invalid `unsafe.Pointer` conversions. Go's garbage collector relies on `unsafe.Pointer` safety rules: a `uintptr` may not be stored in a variable and later converted back to `unsafe.Pointer`, because the GC may move the underlying object. Arithmetic on `unsafe.Pointer` values is also invalid.

Perform pointer arithmetic in a single expression: `unsafe.Pointer(uintptr(p) + offset)`. Do not store intermediate `uintptr` values.
</instructions>

<examples>
## Bad
```go
p := unsafe.Pointer(&x)
offset := uintptr(p) + 16 // stored uintptr — GC may invalidate
q := unsafe.Pointer(offset)
```

## Good
```go
p := unsafe.Pointer(&x)
q := unsafe.Pointer(uintptr(p) + 16) // single expression — safe
```
</examples>

<patterns>
- Storing `uintptr` in a variable then converting back to `unsafe.Pointer`
- Using arithmetic on `unsafe.Pointer` directly
- Converting `reflect.Value.Pointer()` result back to `unsafe.Pointer`
</patterns>

<related>
cgocall, framepointer
</related>
