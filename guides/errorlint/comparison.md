# errorlint: comparison

<instructions>
Detects direct error comparison like `err == io.EOF` or `err != nil` using `==`/`!=`. Wrapped errors will not match a direct comparison because the outer error is a different value. Use `errors.Is(err, io.EOF)` which walks the error chain via `Unwrap()` to find the target error.
</instructions>

<examples>
## Bad
```go
if err == io.EOF {
    return nil
}
```

## Good
```go
if errors.Is(err, io.EOF) {
    return nil
}
```
</examples>

<patterns>
- `err == someSentinel` — use `errors.Is(err, someSentinel)`
- `err != nil` is acceptable for nil check but `errors.Is` is safer for sentinel values
- `if err == io.EOF || err == io.ErrUnexpectedEOF` — use `errors.Is` for each
</patterns>

<related>
errorf, asserts
