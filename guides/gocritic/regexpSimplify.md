# gocritic: regexpSimplify

<instructions>
Detects regular expressions that can be simplified without changing their meaning. This includes using character classes `[0-9]` instead of `\d`, using `?` instead of `{0,1}`, or replacing verbose alternations with character classes.

Simplify the regex pattern to its most concise equivalent form.
</instructions>

<examples>
## Bad
```go
re := regexp.MustCompile(`[0-9]`)
re := regexp.MustCompile(`a{0,1}b`)
```

## Good
```go
re := regexp.MustCompile(`\d`)
re := regexp.MustCompile(`a?b`)
```
</examples>

<patterns>
- Replace `[0-9]` with `\d`
- Replace `[a-zA-Z0-9_]` with `\w`
- Replace `{0,1}` with `?`, `{1,}` with `+`, `{0,}` with `*`
- Simplify character classes where possible — some like `[0-9a-fA-F]` may not simplify further
</patterns>

<related>
regexpMust, regexpPattern
</related>
