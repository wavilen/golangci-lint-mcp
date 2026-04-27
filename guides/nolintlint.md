# nolintlint

<instructions>
Nolintlint checks that `//nolint` directives are well-formed and justified. Bare `//nolint` comments suppress all linters without explanation, masking real issues. The linter also detects unused nolint directives (where the suppressed linter doesn't actually fire).

Always specify which linter to suppress (`//nolint:lintername`) and add an explanation.
</instructions>

<examples>
## Good
```go
//nolint:errcheck // config must exist at deploy time
file, _ := os.Open("config.yaml")
```
</examples>

<patterns>
- Specify linter names in `//nolint:lintername` instead of bare `//nolint`
- Remove unused `//nolint` directives that suppress no actual warnings
- Add an explanation comment after every `//nolint` directive
- Avoid file-level `//nolint` comments — scope suppressions to specific lines
</patterns>

<related>
gocheckcompilerdirectives, godot
</related>
