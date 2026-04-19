# bidichk

<instructions>
Bidichk detects dangerous bidirectional Unicode characters in Go source files. These invisible characters can hide malicious code that appears different in editors than it does to the compiler, posing a supply-chain security risk.

Remove any bidirectional control characters (U+202A–U+202E, U+2066–U+2069, U+200F, U+200E) from source files.
</instructions>

<examples>
## Bad
```go
// Source contains hidden bidirectional override characters
// "admin" ← contains U+202E RTL override before text
if user == "admin" {
```

## Good
```go
// Clean source with no hidden Unicode control characters
if user == "admin" {
```
</examples>

<patterns>
- Files committed with copy-pasted code from external sources containing invisible characters
- RTL (right-to-left) override characters hiding logic in string comparisons
- Bidirectional embeddings in comments that mask code behavior
- Trojan source attacks using Unicode bidi controls to reorder tokens
</patterns>

<related>
asciicheck, goheader, godot
