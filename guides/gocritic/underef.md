# gocritic: underef

<instructions>
Detects expressions where a pointer is explicitly dereferenced for a method call or field access that Go handles automatically. For example, `(*p).Method()` can be written as `p.Method()` because Go dereferences pointers automatically for method calls.

Remove the explicit dereference and let Go handle the pointer indirection.
</instructions>

<examples>
## Bad
```go
(*ptr).Method()
val := (*ptr).Field
```

## Good
```go
ptr.Method()
val := ptr.Field
```
</examples>

<patterns>
- `(*p).Method()` → `p.Method()`
- `(*p).Field` → `p.Field`
- Explicit dereference for read operations on struct pointers
</patterns>

<related>
newDeref, methodExprCall, ptrToRefParam
</related>
