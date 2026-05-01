# revive: unhandled-error

<instructions>
Detects error return values discarded without checking. Unchecked errors produce silent failures that corrupt data and make root-cause debugging extremely difficult. Assign the error to a variable and handle it with `if err != nil`, or document intentional ignores with `//nolint` and a reason.
</instructions>

<examples>
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
- Check error returns instead of discarding them with the blank identifier `_`
- Use and check errors from `fmt.Fprint`/`fmt.Fprintf` calls on `io.Writer`
- Handle errors from `os` operations like `Chmod`, `Mkdir`, and `Remove`
- Use and handle `io.Closer.Close()` errors in defer statements
- Check errors from `encoding/json` marshal/unmarshal operations
</patterns>

<related>
revive/error-return, revive/error-naming, errcheck
</related>
