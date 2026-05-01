# gocritic: uncheckedInlineErr

<instructions>
Detects error return values assigned inline that are not checked before subsequent code runs. Proceeding without checking lets invalid data from the failed call flow into downstream logic, producing wrong results or panics. Check the error immediately after the call — do not defer error handling when subsequent code depends on the call's success.
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
