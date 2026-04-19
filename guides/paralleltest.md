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
- Range variable captured in parallel subtest closure (pre-Go 1.22)
- Missing `t.Parallel()` in subtests of a parallel parent
- `t.Parallel()` called after test assertions (must be first)
- Table-driven tests where only some cases call `t.Parallel()`
</patterns>

<related>
tparallel, thelper, testpackage
</related>
