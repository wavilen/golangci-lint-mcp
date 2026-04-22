# testifylint: formatter

<instructions>
Detects `assert.Equal(t, fmt.Sprintf("msg: %s", x), y)` where `fmt.Sprintf` is used to build an assertion message. Testify assertions accept format strings natively through their `f` variants (`assert.Equalf`) or via the variadic message parameter. Remove the manual `fmt.Sprintf` and use the built-in formatting.
</instructions>

<examples>
## Bad
```go
assert.Equal(t, expected, actual, fmt.Sprintf("item %d failed", i))
```

## Good
```go
assert.Equalf(t, expected, actual, "item %d failed", i)
```
</examples>

<patterns>
- Use `assert.Xxxf` with format args instead of `assert.Xxx(t, ..., fmt.Sprintf(...))`
- Pass format string and args directly instead of `fmt.Sprintf` in assertion message parameters
- Use formatted assertion variants for building complex message strings
</patterns>

<related>
contains-unnecessary-format, expected-actual
