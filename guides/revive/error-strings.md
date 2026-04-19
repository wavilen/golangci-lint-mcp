# revive: error-strings

<instructions>
Enforces Go error message conventions: error strings should not start with a capital letter and should not end with punctuation. This convention follows the principle that error strings may be combined or wrapped, so they should read as a continuation of a sentence, not standalone messages.

Lowercase the first letter of error messages and remove trailing punctuation (periods, exclamation marks, colons).
</instructions>

<examples>
## Bad
```go
return fmt.Errorf("File not found.")
return errors.New("Connection refused!")
return fmt.Errorf("Invalid input: %v", err)
```

## Good
```go
return fmt.Errorf("file not found")
return errors.New("connection refused")
return errors.Wrap(err, "invalid input")
```
</examples>

<patterns>
- Error messages starting with uppercase (often copied from user-facing messages)
- Error messages ending with periods or other punctuation
- Error messages that are full sentences with proper capitalization
- Wrapped errors with punctuation before `%w`
- Error messages using title case like "Not Found"
</patterns>

<related>
error-naming, error-return, errorf
