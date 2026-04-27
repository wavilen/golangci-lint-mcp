# gocritic: uncheckedInlineErr

<instructions>
Detects inline error returns that are not checked before continuing. When a function returns multiple values including an error, assigning the error inline (e.g., in a multi-assignment) without immediately checking it is error-prone. The code may proceed with invalid data from the failed call.

Check the error immediately after the call. Do not defer error handling when the subsequent code depends on the success of the call.
</instructions>

<examples>
## Good
```go
data, err := os.ReadFile(path)
if err != nil {
    return errors.Wrap(err, "reading file")
}
lines := strings.Split(string(data), "\n")
```
</examples>

<patterns>
- Check all returned errors — avoid blank identifier `_` for error values
- Check error before using other return values from multi-assignment calls
- Return or handle errors from potentially failed calls — avoid proceeding with nil/zero values
- Return errors from failed calls — avoid logging with `fmt.Fprintln` instead of propagating
</patterns>

<related>
gocritic/externalErrorReassign, gocritic/nilValReturn, gocritic/sqlQuery
</related>
