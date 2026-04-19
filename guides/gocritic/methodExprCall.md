# gocritic: methodExprCall

<instructions>
Detects method expression calls like `Type.Method(receiver, args)` that could be written as regular method calls `receiver.Method(args)`. Method expression syntax is harder to read and unnecessary when you have a value or pointer of the correct type.

Use the standard method call syntax on the receiver instead of the method expression form.
</instructions>

<examples>
## Bad
```go
http.HandlerFunc(myHandler).ServeHTTP(w, r)
```

## Good
```go
myHandler.ServeHTTP(w, r)
```
</examples>

<patterns>
- `Type.Method(instance, args)` → `instance.Method(args)`
- Using method expressions when the receiver is available as a value
- `(*Type).Method(ptr, args)` → `ptr.Method(args)`
</patterns>

<related>
underef, unlambda
</related>
