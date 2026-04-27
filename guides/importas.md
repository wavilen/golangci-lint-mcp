# importas

<instructions>
Importas enforces consistent alias names for imported packages. Teams use it to prevent inconsistent aliases like `httpClient` vs `hc` for the same package across the codebase.

Set the enforced alias in `.golangci.yml` under `linters.settings.importas.alias`. Use the same alias everywhere the package is imported.
</instructions>

<examples>
## Good
```go
// file_a.go
import "net/http"
// file_b.go
import "net/http"
```
</examples>

<patterns>
- Use the same alias for each imported package across all files
- Replace non-standard aliases with the configured canonical alias
- Configure explicit aliases for auto-generated versioned API packages
- Standardize alias casing to the project convention
</patterns>

<related>
gomodguard, forbidigo, decorder
</related>
