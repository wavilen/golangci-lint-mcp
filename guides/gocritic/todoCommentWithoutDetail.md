# gocritic: todoCommentWithoutDetail

<instructions>
Detects `TODO` or `FIXME` comments that lack context — no author, no issue tracker reference, or no description of what needs to be done. Bare `// TODO` comments are noise without actionable information.

Add descriptive detail to the comment: what needs to be done, who is responsible, or a link to an issue.
</instructions>

<examples>
## Bad
```go
// TODO
// FIXME
// HACK
```

## Good
```go
// TODO(john): migrate to structured logger by 2026-06 (#1234)
// FIXME: race condition when concurrent writers access the cache
```
</examples>

<patterns>
- `// TODO` with no text after it
- `// FIXME` without description or owner
- `// HACK` without explaining why the hack exists
- `// XXX` without context
</patterns>

<related>
commentFormatting, docStub
</related>
