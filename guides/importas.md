# importas

<instructions>
Importas enforces consistent alias names for imported packages. Teams use it to prevent inconsistent aliases like `httpClient` vs `hc` for the same package across the codebase.

Set the enforced alias in `.golangci.yml` under `linters.settings.importas.alias`. Use the same alias everywhere the package is imported.
</instructions>

<examples>
## Bad
```go
// file_a.go
import http "net/http"
// file_b.go
import httplib "net/http"
```

## Good
```go
// file_a.go
import "net/http"
// file_b.go
import "net/http"
```
</examples>

<patterns>
- Same package imported under different aliases across files
- Non-standard aliases for well-known packages (`proto` for `protobuf`)
- Auto-generated aliases like `v1`, `v2` from versioned APIs
- Inconsistent casing in alias names (camelCase vs snake_case)
</patterns>

<related>
gomodguard, forbidigo, decorder
</related>
