# nlreturn

<instructions>
Nlreturn enforces a blank line before `return` statements. Grouping the return visually separates the final logic from the exit point, improving scanability in longer blocks.

Insert a blank line before each `return` statement. Short single-line functions are typically excluded by configuration.
</instructions>

<examples>
## Bad
```go
func resolve(cfg Config) (string, error) {
    if cfg.Host == "" {
        return "", errors.New("host required")
    }
    return cfg.Host + ":" + cfg.Port, nil
}
```

## Good
```go
func resolve(cfg Config) (string, error) {
    if cfg.Host == "" {

        return "", errors.New("host required")
    }

    return cfg.Host + ":" + cfg.Port, nil
}
```
</examples>

<patterns>
- Return statements immediately following logic without a blank line separator
- Early returns in guard clauses lacking visual separation
- Multiple return paths in switch/case blocks
- Returns at the end of nested if-else chains
</patterns>

<related>
wsl_v5, whitespace, nakedret
</related>
