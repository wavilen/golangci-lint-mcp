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
- Replace `err == someSentinel` with `errors.Is(err, someSentinel)` to walk the error chain
- Use `errors.Is` for sentinel value checks — `err != nil` is acceptable for nil checks only
- Replace `if err == io.EOF || err == io.ErrUnexpectedEOF` with separate `errors.Is` calls
</patterns>

<related>
errorf, asserts
