# gocritic: uncheckedInlineErr

<instructions>
Detects inline error returns that are not checked before continuing. When a function returns multiple values including an error, assigning the error inline (e.g., in a multi-assignment) without immediately checking it is error-prone. The code may proceed with invalid data from the failed call.

Check the error immediately after the call. Do not defer error handling when the subsequent code depends on the success of the call.
</instructions>

<examples>
## Bad
```go
data, _ := os.ReadFile(path) // error ignored
lines := strings.Split(string(data), "\n")
```

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
- Blank identifier `_` assigned to errors
- Multi-assignment where error is not checked before using other values
- Proceeding with nil/zero values after a potentially failed call
- Using `fmt.Fprintln` or `log.Println` instead of returning errors
</patterns>

<related>
externalErrorReassign, nilValReturn, sqlQuery
</related>
