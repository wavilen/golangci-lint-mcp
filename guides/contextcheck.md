# contextcheck

<instructions>
Contextcheck verifies that `context.Context` is properly propagated through function call chains. It flags places where a context is created but not passed down, or where a function accepting a context receives `context.Background()` instead of the caller's context.

Pass the parent context as the first argument to every function that needs it.
</instructions>

<examples>
## Bad
```go
func (s *Service) Fetch(ctx context.Context) error {
    req, _ := http.NewRequest("GET", url, nil)
    // context not passed to the request
    resp, err := http.DefaultClient.Do(req)
    return err
}
```

## Good
```go
func (s *Service) Fetch(ctx context.Context) error {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }
    resp, err := http.DefaultClient.Do(req)
    return err
}
```
</examples>

<patterns>
- HTTP requests without using `http.NewRequestWithContext`
- Goroutines that discard the parent context instead of deriving a child
- Functions that accept `context.Context` but call helpers without passing it
- Using `context.Background()` where the caller's context should propagate
</patterns>

<related>
errcheck, govet, revive
