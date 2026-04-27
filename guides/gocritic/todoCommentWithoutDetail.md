# gocritic: todoCommentWithoutDetail

<instructions>
Detects `TODO` or `FIXME` comments that lack context — no author, no issue tracker reference, or no description of what needs to be done. Bare `// TODO` comments are noise without actionable information.

Add descriptive detail to the comment: what needs to be done, who is responsible, or a link to an issue.
</instructions>

<examples>
## Good
```go
// TODO(john): migrate to structured logger by 2026-06 (#1234).
// FIXME: race condition when concurrent writers access the cache.
```
</examples>

<patterns>
- Add description after `// TODO` — explain what needs to be done
- Add description and owner to `// FIXME` comments
- Add explanation to `// HACK` comments — describe why the workaround exists
- Add context to `// XXX` comments — explain the concern
</patterns>

<related>
gocritic/commentFormatting, gocritic/docStub
</related>
