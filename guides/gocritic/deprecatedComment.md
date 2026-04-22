# gocritic: deprecatedComment

<instructions>
Detects malformed deprecation comments. Go conventions require the exact format `// Deprecated: message` (with capital P and colon) for `go vet` and IDEs to recognize deprecation notices. Missing colons, wrong capitalization, or non-standard formats cause tooling to silently ignore the deprecation.

Use the exact `// Deprecated: explanation` format. The colon and capitalization are required.
</instructions>

<examples>
## Bad
```go
// deprecated — use NewClient instead
func OldClient() *Client { ... }
```

## Good
```go
// Deprecated: Use NewClient instead.
func OldClient() *Client { ... }
```
</examples>

<patterns>
- Replace `// deprecated` with `// Deprecated` — capitalize the P
- Add colon separator — use `// Deprecated: reason`
- Replace `// DEPRECATED` all-caps with standard `// Deprecated:`
- Replace `// @deprecated` JSDoc-style with `// Deprecated:`
- Add a deprecation reason and replacement suggestion after the colon
</patterns>

<related>
codegenComment, commentedOutCode
</related>
