# govet: cgocall

<instructions>
Detects incorrect `cgo` calls that violate Go's pointer passing rules. Go pointers must not be passed to C if C retains them after the call returns, and Go pointers to Go pointers cannot be passed to C at all.

Use `C.malloc`/`C.free` for C-side memory, or ensure C does not store the Go pointer beyond the call.
</instructions>

<examples>
## Bad
```go
/*
void store(void **p) { *p = malloc(1); }
*/
import "C"

var buf unsafe.Pointer
C.store((**C.void)(unsafe.Pointer(&buf))) // passing Go pointer to pointer
```

## Good
```go
/*
void store(void **p) { *p = malloc(1); }
*/
import "C"

p := C.malloc(C.size_t(unsafe.Sizeof(C.void{})))
C.store((**C.void)(p)) // C-allocated memory is safe
```
</examples>

<patterns>
- Avoid passing Go pointers to C functions that retain them after the call returns
- Avoid passing `**C.type` where the inner pointer is Go-allocated — use `C.malloc` instead
- Avoid using `unsafe.Pointer` to bypass cgo pointer rules — allocate C-side memory
</patterns>

<related>
unsafeptr, asmdecl
</related>
