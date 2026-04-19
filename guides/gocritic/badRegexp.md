# gocritic: badRegexp

<instructions>
Detects regular expressions that will never match any string or contain syntax errors detectable at lint time. Common issues include contradictory character classes, unreachable alternatives, and patterns that are inherently empty matches.

Fix the regex pattern to correctly express the intended match. Review character classes, anchors, and alternations for logical consistency.
</instructions>

<examples>
## Bad
```go
// Empty character class — matches nothing
re := regexp.MustCompile(`[^a]`) // fine, but `[]` is empty
re := regexp.MustCompile(`a|a`)  // duplicate alternative
```

## Good
```go
re := regexp.MustCompile(`[a-z]+`)
```
</examples>

<patterns>
- Duplicate alternatives: `a|a`
- Empty or contradictory character classes
- Anchors that make patterns impossible to match
- Redundant flags or repetition on zero-width assertions
</patterns>

<related>
regexpPattern, dupSubExpr
</related>
