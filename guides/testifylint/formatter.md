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
- `assert.Xxx(t, ..., fmt.Sprintf(...))` — use `assert.Xxxf` with format args
- `fmt.Sprintf` in assertion message parameter — pass format string and args directly
- Building complex message strings for assertions — use formatted assertion variant
</patterns>

<related>
contains-unnecessary-format, expected-actual
