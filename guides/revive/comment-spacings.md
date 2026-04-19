# revive: comment-spacings

<instructions>
Detects comments that lack a space after the `//` marker. Go convention requires a space between the comment marker and the comment text for readability. The only exception is comments used for code sections like `//nolint`.

Add a space after `//` in all comments. For special directives, follow the specific directive format.
</instructions>

<examples>
## Bad
```go
//this lacks a space
x := compute() //also here
//TODO: fix this later
```

## Good
```go
// This has a space
x := compute() // also here
// TODO: fix this later
```
</examples>

<patterns>
- Inline comments starting directly after `//` without a space
- Comment directives like `//nolint` that are exceptions to the spacing rule
- Block comments (`/* */`) with missing internal spacing
- Auto-generated comments from tools that omit the space
- Code sections or separators without proper spacing
</patterns>

<related>
comments-density, godot
