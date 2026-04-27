# gocritic: octalLiteral

<instructions>
Detects old-style octal literals (e.g., `0755`) that should use the modern `0o755` form introduced in Go 1.13. The `0o` prefix is more explicit and avoids confusion with decimal or hex literals.

Replace `0`-prefixed octal literals with the `0o` prefix form.
</instructions>

<examples>
## Good
```go
os.MkdirAll(dir, 0o755)
os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
```
</examples>

<patterns>
- Replace `0644` with `0o644`, `0755` with `0o755` — use `0o` prefix for clarity
- Replace `0`-prefixed numbers with `0o` prefix when octal is intended
- Use `0o` prefix for octal literals — works identically but is self-documenting
</patterns>

<related>
gocritic/hexLiteral
</related>
