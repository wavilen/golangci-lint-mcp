# revive: unhandled-error

<instructions>
Detects function return values (specifically errors) that are discarded without checking. Ignoring errors hides failures, leads to silent data corruption, and makes debugging extremely difficult. Every function that returns an error should have its return value checked.

Assign the error to a variable and check it. Use `if err != nil` to handle the error, or explicitly document why ignoring it is safe with `//nolint` and a reason.
</instructions>

<examples>
## Bad
```go
file, _ := os.Open("config.yaml")
rows, _ := db.Query("SELECT * FROM users")
fmt.Fprintf(w, "hello")
```

## Good
```go
file, err := os.Open("config.yaml")
if err != nil {
    return errors.Wrap(err, "open config")
}
rows, err := db.Query("SELECT * FROM users")
if err != nil {
    return errors.Wrap(err, "query users")
}
```
</examples>

<patterns>
- Using blank identifier `_` to discard error returns
- Calling `fmt.Fprint`/`fmt.Fprintf` on an `io.Writer` without checking the error
- Ignoring errors from `os` operations like `Chmod`, `Mkdir`, `Remove`
- Discarding `io.Closer.Close()` errors in defer statements
- Ignoring errors from `encoding/json` marshal/unmarshal operations
</patterns>

<related>
error-return, error-naming, errcheck
