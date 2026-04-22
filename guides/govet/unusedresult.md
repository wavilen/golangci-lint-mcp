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
- Check the return value of `copy()` to verify the number of bytes copied
- Use the return value of string predicates (`strings.HasPrefix`, `strings.Contains`) — never discard
- Use the error value from `errors.New` — assign or check the result
- Use the return value of `fmt.Sprintf` — assign it to a variable
</patterns>

<related>
appends, assign
</related>
