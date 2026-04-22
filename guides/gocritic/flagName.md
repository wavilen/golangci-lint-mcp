# gocritic: flagName

<instructions>
Detects flag names that don't follow Go conventions. Flag names should use lowercase letters and hyphens (e.g., `max-retries`), not underscores or CamelCase. Inconsistent naming confuses users on the command line.

Use lowercase kebab-case for flag names: `my-flag-name` instead of `my_flag_name` or `myFlagName`.
</instructions>

<examples>
## Bad
```go
flag.Int("max_retries", 3, "maximum retry count")
```

## Good
```go
flag.Int("max-retries", 3, "maximum retry count")
```
</examples>

<patterns>
- Replace underscores in flag names with hyphens — use `db-host` not `db_host`
- Use `kebab-case` for flag names — avoid `maxRetries`, use `max-retries`
- Use consistent naming convention across all flags in the same program
- Use lowercase for flag names — avoid uppercase starting letters
</patterns>

<related>
flagDeref
</related>
