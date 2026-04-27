# gocritic: methodExprCall

<instructions>
Detects method expression calls like `Type.Method(receiver, args)` that could be written as regular method calls `receiver.Method(args)`. Method expression syntax is harder to read and unnecessary when you have a value or pointer of the correct type.

Use the standard method call syntax on the receiver instead of the method expression form.
</instructions>

<examples>
## Good
```go
myHandler.ServeHTTP(w, r)
```
</examples>

<patterns>
- Replace `Type.Method(instance, args)` with `instance.Method(args)`
- Replace method expressions with method calls when the receiver is available as a value
- Replace `(*Type).Method(ptr, args)` with `ptr.Method(args)`
</patterns>

<related>
gocritic/underef, gocritic/unlambda
</related>
