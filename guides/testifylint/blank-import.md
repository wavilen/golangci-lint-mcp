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
- Import `github.com/stretchr/testify/assert` explicitly instead of blank-importing it
- Use `require` directly instead of blank-importing `github.com/stretchr/testify/require`
- Remove blank import leftovers from debugging
</patterns>

<related>
useless-assert, suite-dont-use-pkg
