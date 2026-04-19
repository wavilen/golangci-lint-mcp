# gocritic: regexpPattern

<instructions>
Detects common mistakes in regular expression patterns, such as using `[\d]` instead of `\d`, `[[:digit:]]` mixed with `\d`, or other redundant or incorrect pattern constructs. These don't change behavior but indicate confusion about regex syntax.

Simplify the regex pattern to use the canonical form. Prefer `\d`, `\w`, `\s` over character class equivalents.
</instructions>

<examples>
## Bad
```go
re := regexp.MustCompile(`[\d]+`) // unnecessary character class
```

## Good
```go
re := regexp.MustCompile(`\d+`)
```
</examples>

<patterns>
- `[\d]` instead of `\d`
- `[[:digit:]]` mixed with `\d`
- Unnecessary nested character classes: `[[a-z]]`
- Redundant anchors or grouping
</patterns>

<related>
badRegexp, dynamicFmtString
</related>
