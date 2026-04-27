# revive: dot-imports

<instructions>
Detects dot imports (`import . "pkg"`) which bring the imported package's identifiers into the current file's namespace without qualification. This makes it unclear where a name comes from and can cause name collisions.

Use qualified imports (`import pkg "..."`) unless in test files importing the package under test, or in DSL-style code where dot imports are an established convention (e.g., `table-driven tests` with `is` packages).
</instructions>

<examples>
## Good
```go
package handler

import "net/http"

func setup() {
    http.Handle("/api", http.HandlerFunc(mux))
}
```
</examples>

<patterns>
- Use qualified imports instead of dot-importing utility packages
- Limit dot imports in test files to the package under test only
- Avoid dot imports for DSL packages unless dot imports are an established convention
- Remove leftover dot imports when migrating code via copy-paste
- Refactor circular import workarounds that use dot imports into proper package boundaries
</patterns>

<related>
revive/blank-imports, revive/duplicated-imports, staticcheck/ST1001
</related>
