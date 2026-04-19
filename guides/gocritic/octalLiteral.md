# gocritic: octalLiteral

<instructions>
Detects old-style octal literals (e.g., `0755`) that should use the modern `0o755` form introduced in Go 1.13. The `0o` prefix is more explicit and avoids confusion with decimal or hex literals.

Replace `0`-prefixed octal literals with the `0o` prefix form.
</instructions>

<examples>
## Bad
```go
os.MkdirAll(dir, 0755)
os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
```

## Good
```go
os.MkdirAll(dir, 0o755)
os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
```
</examples>

<patterns>
- File permissions: `0644`, `0755` → `0o644`, `0o755`
- Any `0`-prefixed number intended as octal
- `0o` prefix works identically but is self-documenting
</patterns>

<related>
hexLiteral
</related>
