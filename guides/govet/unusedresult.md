# govet: unusedresult

<instructions>
Reports unused results of functions marked as "pure" by the analyzer, such as `strings.HasPrefix`, `strings.Contains`, `copy`, and `errors.New`. Ignoring these return values is almost always a bug — the function has no side effects, so calling it without using the result is pointless.

Assign the result to a variable and use it, or remove the call entirely.
</instructions>

<examples>
## Bad
```go
strings.HasPrefix(name, "test_") // result discarded — call is useless
```

## Good
```go
if strings.HasPrefix(name, "test_") {
    runTest(name)
}
```
</examples>

<patterns>
- Discarding return value of `copy()` — indicates copy count not checked
- Ignoring result of `strings.HasPrefix`/`strings.Contains`
- Not using the error value from `errors.New`
- Discarding `fmt.Sprintf` result
</patterns>

<related>
appends, assign
</related>
