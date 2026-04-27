# testpackage

<instructions>
Testpackage detects tests that import the package under test using the internal `.` or `_` import alias, or tests in the same package that test unexported symbols. Tests in a separate package (`xxx_test`) better simulate real usage and prevent coupling to internal details.

Move test files to an external test package (suffix `_test` in package name) and import the package normally.
</instructions>

<examples>
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
- Move tests to an external `xxx_test` package to test only exported symbols
- Replace dot imports (`import . "my/pkg"`) with normal imports in test files
- Test through the public API rather than relying on unexported internal state
- Add `//nolint:testpackage` for integration tests that genuinely need internal access
</patterns>

<related>
paralleltest, thelper, testableexamples
</related>
