# gocritic: hexLiteral

<instructions>
Detects integer literals that could be written as hex literals for clarity. When a value represents a bitmask, permission flags, or byte value, hex notation is more readable than decimal.

Replace the decimal literal with its hex equivalent (e.g., `255` → `0xFF`, `4096` → `0x1000`).
</instructions>

<examples>
## Bad
```go
const mask = 255
const perm = 493 // octal 0755 in decimal
const page = 4096
```

## Good
```go
const mask = 0xFF
const perm = 0755
const page = 0x1000
```
</examples>

<patterns>
- Replace bitmask values like `255` with hex `0xFF`, `65535` with `0xFFFF`
- Replace permission modes like `493` with octal `0o644` or `0755`
- Replace page sizes and alignment constants like `4096` with hex `0x1000`
</patterns>

<related>
octalLiteral
</related>
