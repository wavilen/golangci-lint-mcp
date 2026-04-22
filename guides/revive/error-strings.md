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
- Use lowercase for error messages — avoid uppercase at the start
- Remove trailing periods or punctuation from error messages
- Use phrases for error messages — avoid full sentences with capitalization
- Ensure wrapped errors don't include punctuation before the `%w` verb
- Use lowercase for error messages — avoid title case like "Not Found"
</patterns>

<related>
error-naming, error-return, errorf
