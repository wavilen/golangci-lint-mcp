# modernize: loopvar

<instructions>
Detects loop variable capture patterns where `for i, v := range` variables are referenced by closures or goroutines. In Go versions before 1.22, loop variables are shared across iterations, causing all closures to capture the final value. Go 1.22+ creates per-iteration copies automatically. If targeting Go 1.22+, this pattern is safe by default.
</instructions>

<examples>
## Good
```go
for _, v := range values {
    v := v // explicit copy for pre-1.22 compatibility
    go func() {
        process(v)
    }()
}
```
</examples>

<patterns>
- Use `v := v` to copy loop variables before closures capture them (pre-Go 1.22)
- Pass loop variables as goroutine arguments instead of capturing them directly
- Use Go 1.22+ range semantics or add per-iteration copies for `go func()` referencing `i` or `v`
</patterns>

<related>
modernize/simplifyrange, modernize/reloop
</related>
