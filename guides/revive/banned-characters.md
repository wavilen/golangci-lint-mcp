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
- Smart quotes or em-dashes copied from documentation into string literals
- Non-ASCII Unicode characters in identifiers
- Invisible zero-width characters pasted from external editors
- Emoji characters in comments or strings when disallowed by team policy
- Language-specific characters in variable names
</patterns>

<related>
filename-format
