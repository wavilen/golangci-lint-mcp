# testifylint: blank-import

<instructions>
Detects blank imports of testify packages (`_ "github.com/stretchr/testify/..."`). Blank imports execute the package's `init()` function but testify is a testing library meant to be used explicitly through its `assert`, `require`, and `mock` packages. Blank importing it serves no purpose and likely indicates a mistake.
</instructions>

<examples>
## Bad
```go
import (
    _ "github.com/stretchr/testify/assert"
    "testing"
)
```

## Good
```go
import (
    "github.com/stretchr/testify/assert"
    "testing"
)
```
</examples>

<patterns>
- `_ "github.com/stretchr/testify/assert"` — import explicitly and use
- `_ "github.com/stretchr/testify/require"` — same, use require directly
- Blank import leftover from debugging — remove it
</patterns>

<related>
useless-assert, suite-dont-use-pkg
