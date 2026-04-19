# gocritic: whyNoLint

<instructions>
Detects `//nolint` directives that lack an explanation. Suppressing linter warnings without explaining why makes the codebase harder to maintain and hides potential issues. Future readers won't know if the suppression is still valid.

Add a reason after the `//nolint` directive: `//nolint:gocritic // reason for suppression`.
</instructions>

<examples>
## Bad
```go
//nolint
result, _ = doSomething()
```

## Good
```go
//nolint:errcheck // error is logged inside doSomething, safe to ignore
result, _ = doSomething()
```
</examples>

<patterns>
- `//nolint` without any comment explaining the reason
- `//nolint:gocritic` with no justification
- Blanket `//nolint` suppressing all linters without specificity
</patterns>

<related>
todoCommentWithoutDetail, commentFormatting
</related>
