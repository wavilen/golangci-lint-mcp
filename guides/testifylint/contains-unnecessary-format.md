# testifylint: contains-unnecessary-format

<instructions>
Detects usage of formatted assertion variants (`assert.Containsf`) when the format arguments are used only for the test message and not needed for special formatting. The non-format variants (`assert.Contains`) accept a variadic message the same way, making the `f` suffix unnecessary.
</instructions>

<examples>
## Bad
```go
assert.Containsf(t, haystack, needle, "msg %d", 1)
```

## Good
```go
assert.Contains(t, haystack, needle, "msg", 1)
```
</examples>

<patterns>
- `assert.Containsf(t, a, b, "msg", args...)` where format is trivial — use `assert.Contains`
- Using `Xxxf` variant without actual format verbs — use non-f variant
- `assert.Equalf` with no `%` in message — use `assert.Equal`
</patterns>

<related>
formatter, expected-actual
