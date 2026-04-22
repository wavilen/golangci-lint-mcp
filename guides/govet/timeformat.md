# govet: timeformat

<instructions>
Reports `time.Format` and `time.Parse` calls with format strings that don't use Go's reference time `Mon Jan 2 15:04:05 MST 2006`. Common mistakes include using YYYY-MM-DD or other C-style format specifiers, which produce incorrect results since Go formats are mnemonic-based.

Use Go's reference time constants: `time.DateTime`, `time.DateOnly`, or the literal `"2006-01-02 15:04:05"`.
</instructions>

<examples>
## Bad
```go
t.Format("YYYY-MM-DD") // Go doesn't use YYYY — produces "YYYY-MM-DD" literally
```

## Good
```go
t.Format("2006-01-02") // Go reference time format
```
</examples>

<patterns>
- Use `2006` for year in `time.Format`, not `YYYY`
- Use `01` for month in `time.Format`, not `MM`
- Use `02` for day in `time.Format`, not `DD`
- Use `15:04:05` for time in `time.Format`, not `HH:mm:ss`
</patterns>

<related>
printf, stringintconv
</related>
