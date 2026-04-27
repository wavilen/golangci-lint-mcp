# nlreturn

<instructions>
Nlreturn enforces a blank line before `return` statements. Grouping the return visually separates the final logic from the exit point, improving scanability in longer blocks.

Insert a blank line before each `return` statement. Short single-line functions are typically excluded by configuration.
</instructions>

<examples>
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
- Add a blank line before `return` statements that immediately follow logic without visual separation
- Insert blank lines before early returns in guard clauses for readability
- Separate multiple return paths in `switch`/`case` blocks with blank lines
- Add blank lines before returns at the end of nested `if`-`else` chains
</patterns>

<related>
wsl_v5, whitespace, nakedret
</related>
