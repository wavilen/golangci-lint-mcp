# paralleltest

<instructions>
Paralleltest checks that `t.Parallel()` is used correctly in tests. It detects missing `t.Parallel()` calls in table-driven tests and incorrect usage of loop variables inside parallel subtests, which causes all subtests to reference the last iteration's value.

Call `t.Parallel()` in each subtest and capture loop variables properly (Go 1.22+ handles this automatically for `for` loops).
</instructions>

<examples>
## Bad
```go
tests := []struct{ name string }{{name: "a"}, {name: "b"}}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        // tt is captured by reference — may be "b" for both
    })
}
```

## Good
```go
tests := []struct{ name string }{{name: "a"}, {name: "b"}}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        _ = tt.name // tt is safe in Go 1.22+; pre-1.22: capture tt := tt
    })
}
```
</examples>

<patterns>
- Capture range variables before the closure in parallel subtests (pre-Go 1.22)
- Add `t.Parallel()` to all subtests within a parallel parent test
- Call `t.Parallel()` before any test assertions, not after
- Mark all table-driven test cases with `t.Parallel()` consistently
</patterns>

<related>
tparallel, thelper, testpackage
</related>
