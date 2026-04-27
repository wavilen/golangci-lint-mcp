# gocritic: badSyncOnceFunc

<instructions>
Detects incorrect usage of `sync.OnceValue` or `sync.OnceValues`, introduced in Go 1.21. Common mistakes include calling `sync.OnceValue` inside a function on every invocation (defeating the memoization) or capturing the result incorrectly. The function returned by `sync.OnceValue` should be stored and reused.

Store the result of `sync.OnceValue(fn)` in a package-level or struct-level variable and call that stored function, rather than wrapping on every call.
</instructions>

<examples>
## Good
```go
var getConfig = sync.OnceValue(loadConfig)

// Usage: getConfig() — memoized after first call
```
</examples>

<patterns>
- Assign `sync.OnceValue(fn)` to a package-level variable — avoid inline calls on every invocation
- Ensure the `OnceValue` wrapper is actually called after storing it
- Replace manual `sync.Once`+flag patterns with `sync.OnceValue` or `sync.OnceValues`
</patterns>

<related>
gocritic/badLock, gocritic/unnecessaryDefer
</related>
