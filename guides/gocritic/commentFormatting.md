# gocritic: commentFormatting

<instructions>
Detects comments that don't follow Go's conventional formatting. This includes comments missing the space after `//`, comments that start with a newline inside `/* */` blocks, or `//` comments followed by more than one space before text.

Start every `//` comment with `// ` (double-slash followed by a single space). Keep `/* */` comments on one line or open with `/*\n` on its own line.
</instructions>

<examples>
## Bad
```go
//this is missing a space
//  this has extra spaces
/*this block comment lacks spacing*/
```

## Good
```go
// this is properly formatted
// this has a single space
/* this block comment has spacing */
```
</examples>

<patterns>
- `//text` without space after slashes
- `//  text` with more than one space
- `/*text*/` without space inside delimiters
- Missing space in `//nolint` directives
</patterns>

<related>
todoCommentWithoutDetail, docStub
</related>
