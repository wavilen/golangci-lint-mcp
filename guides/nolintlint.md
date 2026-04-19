# nolintlint

<instructions>
Nolintlint checks that `//nolint` directives are well-formed and justified. Bare `//nolint` comments suppress all linters without explanation, masking real issues. The linter also detects unused nolint directives (where the suppressed linter doesn't actually fire).

Always specify which linter to suppress (`//nolint:lintername`) and add an explanation.
</instructions>

<examples>
## Bad
```go
//nolint
file, _ := os.Open("config.yaml")
```

## Good
```go
//nolint:errcheck // config must exist at deploy time
file, _ := os.Open("config.yaml")
```
</examples>

<patterns>
- Bare `//nolint` without specifying linter names
- `//nolint` directives that suppress no actual warnings
- Missing explanation comment after `//nolint` directive
- File-level `//nolint` comments suppressing everything
</patterns>

<related>
gocheckcompilerdirectives, godot
</related>
