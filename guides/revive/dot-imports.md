# revive: dot-imports

<instructions>
Detects dot imports (`import . "pkg"`) which bring the imported package's identifiers into the current file's namespace without qualification. This makes it unclear where a name comes from and can cause name collisions.

Use qualified imports (`import pkg "..."`) unless in test files importing the package under test, or in DSL-style code where dot imports are an established convention (e.g., `table-driven tests` with `is` packages).
</instructions>

<examples>
## Bad
```go
package handler

import . "net/http"

func setup() {
    Handle("/api", HandlerFunc(mux)) // where does Handle come from?
}
```

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
- Dot-importing utility packages to save typing
- Test files dot-importing packages other than the one under test
- DSL packages where dot imports are not an established convention
- Migrating code and leaving dot imports from copy-paste
- Circular import workarounds using dot imports
</patterns>

<related>
blank-imports, duplicated-imports, staticcheck ST1001
