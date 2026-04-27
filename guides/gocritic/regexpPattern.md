# gocritic: regexpPattern

<instructions>
Detects common mistakes in regular expression patterns, such as using `[\d]` instead of `\d`, `[[:digit:]]` mixed with `\d`, or other redundant or incorrect pattern constructs. These don't change behavior but indicate confusion about regex syntax.

Simplify the regex pattern to use the canonical form. Prefer `\d`, `\w`, `\s` over character class equivalents.
</instructions>

<examples>
## Good
```go
re := regexp.MustCompile(`\d+`)
```
</examples>

<patterns>
- Replace `[\d]` with `\d` — remove unnecessary character class wrapping
- Replace `[[:digit:]]` with `\d` for consistency
- Remove unnecessary nested character classes like `[[a-z]]` — use `[a-z]`
- Remove redundant anchors or grouping in regex patterns
</patterns>

<related>
gocritic/badRegexp, gocritic/dynamicFmtString
</related>
