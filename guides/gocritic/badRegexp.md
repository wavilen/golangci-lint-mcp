# gocritic: badRegexp

<instructions>
Detects regular expressions that will never match any string or contain syntax errors detectable at lint time. Common issues include contradictory character classes, unreachable alternatives, and patterns that are inherently empty matches.

Fix the regex pattern to correctly express the intended match. Review character classes, anchors, and alternations for logical consistency.
</instructions>

<examples>
## Good
```go
re := regexp.MustCompile(`[a-z]+`)
```
</examples>

<patterns>
- Remove duplicate regex alternatives like `a|a`
- Remove empty or contradictory character classes that match nothing
- Remove anchors that make the pattern impossible to match
- Remove redundant flags or repetition on zero-width assertions
</patterns>

<related>
gocritic/regexpPattern, gocritic/dupSubExpr
</related>
