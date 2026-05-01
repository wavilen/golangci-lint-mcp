# goheader

<instructions>
Checks that source files contain a required header comment, typically a copyright or license notice. Files missing headers cause license audit tools (FOSSA, Snyk) to flag them as unlicensed, blocking compliance approvals and third-party adoption. Add the required header comment at the top of each file, configured via `.golangci.yml` under `linters.settings.goheader`.
</instructions>

<examples>
## Good
```go
// Copyright 2026 My Company. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package main

import "fmt"
```
</examples>

<patterns>
- Add the configured copyright header to every new source file
- Update file headers to match the configured template (year, company name)
- Add headers to generated files or exclude them in config
- Fix header spacing and formatting to match the template exactly
</patterns>

<related>
gocheckcompilerdirectives, godot, gomoddirectives
</related>
