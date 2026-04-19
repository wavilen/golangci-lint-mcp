# revive: blank-imports

<instructions>
Detects blank imports (`import _ "pkg"`) outside of `main` or test files. Blank imports are used solely for side effects (registering drivers, init functions). Using them in library packages forces unwanted side effects on consumers.

Move blank imports to `main` packages or test files. If a package needs a driver registered, document the requirement and let the caller import it.
</instructions>

<examples>
## Bad
```go
// In a library package (not main)
package store

import (
    _ "github.com/lib/pq" // forces PostgreSQL driver on all consumers
)
```

## Good
```go
// In cmd/main.go
package main

import (
    _ "github.com/lib/pq" // main controls which driver to register
)
```
</examples>

<patterns>
- Database driver blank imports in repository or data-layer packages
- Image format decoder registrations in utility packages
- Plugin or metric handler registrations in shared libraries
- Profiling or debugging tool blank imports left in production code
</patterns>

<related>
dot-imports, imports-blocklist
