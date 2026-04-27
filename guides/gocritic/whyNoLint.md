# gocritic: whyNoLint

<instructions>
Detects `//nolint` directives that lack an explanation. Suppressing linter warnings without explaining why makes the codebase harder to maintain and hides potential issues. Future readers won't know if the suppression is still valid.

Add a reason after the `//nolint` directive: `//nolint:gocritic // reason for suppression`.
</instructions>

<examples>
## Good
```go
//nolint:errcheck // error is logged inside doSomething, safe to ignore
result, _ = doSomething()
```
</examples>

<patterns>
- Add a reason after every `//nolint` directive — explain why the suppression is needed
- Add justification after `//nolint:gocritic` — describe why the rule is intentionally disabled
- Add specific linter names in `//nolint` directives — avoid blanket suppression of all linters
</patterns>

<related>
gocritic/todoCommentWithoutDetail, gocritic/commentFormatting
</related>
