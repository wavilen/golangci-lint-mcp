# govet: loopclosure

<instructions>
Detects goroutines or deferred closures that capture loop variables by reference. All iterations share the same variable, so the closure may see the final value instead of the intended one. In Go 1.22+, loop variables are scoped per-iteration, but older code is still vulnerable.

Use a local copy inside the loop (`v := v`) or upgrade to Go 1.22+ per-iteration scoping.
</instructions>

<examples>
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
- Copy loop variables locally (`v := v`) before launching goroutines that reference them
- Pass loop variables as arguments to deferred functions inside loops
- Pass loop variables as closure parameters to `go` and `defer` calls
</patterns>

<related>
govet/testinggoroutine, govet/defers, revive/range-val-in-closure
</related>
