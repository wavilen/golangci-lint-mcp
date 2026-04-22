# revive: banned-characters

<instructions>
Detects the use of configured banned characters in source files. Teams use this rule to enforce character-set restrictions, such as disallowing certain Unicode characters, smart quotes, or non-ASCII characters in identifiers and strings.

Remove or replace the banned character with an allowed alternative as configured in the revive config.
</instructions>

<examples>
## Bad
```go
// Assuming "Ω" is banned
const maxOmega = 3.14 // Ω character in identifier suffix
```

## Good
```go
const maxOmega = 3.14 // use plain ASCII identifier
```
</examples>

<patterns>
- Replace smart quotes or em-dashes copied from documentation with plain ASCII equivalents
- Use ASCII-only characters in identifiers instead of non-ASCII Unicode
- Remove invisible zero-width characters pasted from external editors
- Replace emoji characters in comments or strings with text equivalents when disallowed by team policy
- Rename variables with language-specific characters to ASCII equivalents
</patterns>

<related>
filename-format
