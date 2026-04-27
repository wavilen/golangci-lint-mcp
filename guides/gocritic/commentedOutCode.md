# gocritic: commentedOutCode

<instructions>
Detects commented-out code that should be removed. Commented-out code clutters the codebase, confuses readers about intent, and decays without testing. Use version control to preserve history instead of commenting out code.

Remove commented-out code. If code might be needed later, rely on git history.
</instructions>

<examples>
## Good
```go
// Delete the commented code entirely.
// Git preserves the history if you need it back.
```
</examples>

<patterns>
- Delete commented-out code blocks instead of keeping them for "future reference"
- Rely on git history to retrieve deleted code — version control preserves all changes
- Remove commented functions, imports, and logic blocks to improve code clarity
</patterns>

<related>
godot
</related>
