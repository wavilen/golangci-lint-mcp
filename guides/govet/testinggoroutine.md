# govet: testinggoroutine

<instructions>
Reports calls to `t.Fatal`, `t.Skip`, `t.Log`, and other `testing.T` methods from goroutines launched within tests. These methods must be called from the test goroutine that owns the `testing.T` instance. Calling them from another goroutine causes panics or race conditions.

Use channels to communicate test results back to the test goroutine, or restructure to use `t.Parallel()`.
</instructions>

<examples>
## Bad
```go
func TestParallel(t *testing.T) {
    go func() {
        result := doWork()
        if result != expected {
            t.Fatal("wrong result") // called from wrong goroutine
        }
    }()
}
```

## Good
```go
func TestParallel(t *testing.T) {
    errCh := make(chan error, 1)
    go func() {
        result := doWork()
        if result != expected {
            errCh <- fmt.Errorf("wrong result")
            return
        }
        errCh <- nil
    }()
    if err := <-errCh; err != nil {
        t.Fatal(err.Error())
    }
}
```
</examples>

<patterns>
- `t.Fatal`/`t.Fatalf` called from goroutine
- `t.Skip`/`t.Skipf` called from goroutine
- `t.Log`/`t.Errorf` used in goroutine-launched closure
</patterns>

<related>
loopclosure, tests
</related>
