# gocritic: deferUnlambda

<instructions>
Detects `defer` statements that wrap a function call in an unnecessary anonymous function when the call doesn't need deferred evaluation of its arguments. For example, `defer func() { f() }()` can be simplified to `defer f()` when no arguments need deferred evaluation.

Remove the unnecessary wrapper and defer the function call directly.
</instructions>

<examples>
## Bad
```go
defer func() { mu.Unlock() }()
defer func() { conn.Close() }()
```

## Good
```go
defer mu.Unlock()
defer conn.Close()
```
</examples>

<patterns>
- `defer func() { f() }()` with no arguments → `defer f()`
- `defer func() { obj.Method() }()` → `defer obj.Method()`
- Wrapping a simple call that has no parameters needing deferred evaluation
</patterns>

<related>
unlambda, unnecessaryBlock
</related>
