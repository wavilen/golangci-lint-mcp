# decorder

<instructions>
Decorder checks declaration order and grouping in Go files. It enforces a consistent ordering of `import`, `const`, `var`, and `type` declarations, and optionally requires grouping of const/var blocks.

Reorder declarations to follow the conventional sequence: imports, then types, then constants, then variables, then functions. Group related constants and variables into single blocks.
</instructions>

<examples>
## Bad
```go
var a = 1
const b = 2
import "fmt"
var c = 3
```

## Good
```go
import "fmt"

const b = 2

var (
    a = 1
    c = 3
)
```
</examples>

<patterns>
- Move imports to the top of the file before all other declarations
- Group related const or var declarations into a single block
- Reorder declarations to separate type definitions from variable declarations
- Consolidate related const/iota groups into a single const block
</patterns>

<related>
gofmt, gofumpt, godoclint
</related>
