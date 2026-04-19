# govet: loopclosure

<instructions>
Detects goroutines or deferred closures that capture loop variables by reference. All iterations share the same variable, so the closure may see the final value instead of the intended one. In Go 1.22+, loop variables are scoped per-iteration, but older code is still vulnerable.

Use a local copy inside the loop (`v := v`) or upgrade to Go 1.22+ per-iteration scoping.
</instructions>

<examples>
## Bad
```go
for _, v := range values {
    go func() {
        process(v) // captures loop variable — may see last value
    }()
}
```

## Good
```go
for _, v := range values {
    v := v // create per-iteration copy
    go func() {
        process(v)
    }()
}
```
</examples>

<patterns>
- Goroutine launched inside loop referencing loop variable
- Defer inside loop capturing loop variable by reference
- Closure passed to `go` or `defer` with loop variable in body
</patterns>

<related>
testinggoroutine, defers
</related>
