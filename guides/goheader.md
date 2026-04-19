# goheader

<instructions>
Goheader checks that source files contain a required header comment, typically a copyright or license notice. Projects enforce this to ensure legal compliance and consistent file attribution.

Add the required header comment at the top of the file. Configure the expected header template in `.golangci.yml` under `linters.settings.goheader.values` and `template`.
</instructions>

<examples>
## Bad
```go
package main

import "fmt"
```

## Good
```go
// Copyright 2026 My Company. All rights reserved.
// Use of this source code is governed by a BSD-style license.

package main

import "fmt"
```
</examples>

<patterns>
- New files created without the required copyright header
- Headers that don't match the configured template (wrong year, company name)
- Generated files missing headers when not excluded by config
- Headers with incorrect spacing or formatting
</patterns>

<related>
gocheckcompilerdirectives, godot, gomoddirectives
</related>
