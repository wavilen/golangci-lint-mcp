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
- Bitmask values like `255`, `65535` → hex `0xFF`, `0xFFFF`
- Permission modes like `493` → octal `0755`
- Page sizes and alignment: `4096` → `0x1000`
</patterns>

<related>
octalLiteral
</related>
