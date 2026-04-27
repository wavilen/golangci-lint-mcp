# revive: comment-spacings

<instructions>
Detects comments that lack a space after the `//` marker. Go convention requires a space between the comment marker and the comment text for readability. The only exception is comments used for code sections like `//nolint`.

Add a space after `//` in all comments. For special directives, follow the specific directive format.
</instructions>

<examples>
## Good
```go
// This has a space
x := compute() // also here
// TODO: fix this later
```
</examples>

<patterns>
- Add a space after `//` in all comments — write `// comment` not `//comment`
- Use `//nolint` and similar directives in their directive format as exceptions to spacing
- Ensure block comments (`/* */`) include proper internal spacing
- Replace auto-generated comments from tools that omit the space after `//`
- Ensure code section separators have proper spacing after the comment marker
</patterns>

<related>
revive/comments-density, godot
</related>
