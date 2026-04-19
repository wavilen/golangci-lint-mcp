# testpackage

<instructions>
Testpackage detects tests that import the package under test using the internal `.` or `_` import alias, or tests in the same package that test unexported symbols. Tests in a separate package (`xxx_test`) better simulate real usage and prevent coupling to internal details.

Move test files to an external test package (suffix `_test` in package name) and import the package normally.
</instructions>

<examples>
## Bad
```go
package mypkg

import "testing"

func TestInternal(t *testing.T) {
    result := helperFunc() // testing unexported symbol
    // ...
}
```

## Good
```go
package mypkg_test

import (
    "testing"
    "example.com/mymypkg"
)

func TestPublic(t *testing.T) {
    result := mypkg.PublicFunc()
    // ...
}
```
</examples>

<patterns>
- Test files in the same package accessing unexported functions
- Dot imports of the package under test: `import . "my/pkg"`
- Tests that rely on internal state not visible to callers
- Integration tests that need internal access (may warrant `//nolint:testpackage`)
</patterns>

<related>
paralleltest, thelper, testableexamples
</related>
