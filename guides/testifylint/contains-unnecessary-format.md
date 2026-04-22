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
- Use `assert.Contains` instead of `assert.Containsf(t, a, b, "msg", args...)` when format is trivial
- Use non-f assertion variants when `Xxxf` is used without actual format verbs
- Use `assert.Equal` instead of `assert.Equalf` when the message has no `%` format verbs
</patterns>

<related>
formatter, expected-actual
