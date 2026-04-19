# modernize: loopvar

<instructions>
Detects loop variable capture patterns where `for i, v := range` variables are referenced by closures or goroutines. In Go versions before 1.22, loop variables are shared across iterations, causing all closures to capture the final value. Go 1.22+ creates per-iteration copies automatically. If targeting Go 1.22+, this pattern is safe by default.
</instructions>

<examples>
## Bad
```go
for _, v := range values {
    go func() {
        process(v) // all goroutines see last value (pre-Go 1.22)
    }()
}
```

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
- Closure capturing loop variable — copy with `v := v` if pre-Go 1.22
- Goroutine launched inside loop using loop variable — pass as argument or copy
- `go func()` referencing `i` or `v` from range — per-iteration copy or Go 1.22+
</patterns>

<related>
simplifyrange, reloop
