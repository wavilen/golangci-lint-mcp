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
- Remove invisible Unicode characters from copy-pasted code before committing
- Strip RTL override characters (U+202E) that hide logic in string comparisons
- Remove bidirectional embedding characters from comments
- Eliminate Unicode bidi controls that reorder tokens for trojan source attacks
</patterns>

<related>
asciicheck, goheader, godot
