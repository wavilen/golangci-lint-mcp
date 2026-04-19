# exptostd

<instructions>
Exptostd detects functionally identical code in third-party packages that can be replaced by standard library equivalents. Using the standard library reduces dependencies and improves long-term maintainability.

Replace the third-party import with the equivalent standard library function or type.
</instructions>

<examples>
## Bad
```go
import "golang.org/x/exp/slices"

sorted := slices.Sort(nums)
```

## Good
```go
import "slices"

sorted := slices.Sort(nums)
```
</examples>

<patterns>
- Using `golang.org/x/exp/slices` instead of `slices` (Go 1.21+)
- Using `golang.org/x/exp/maps` instead of `maps` (Go 1.21+)
- Using third-party string utilities duplicated by `strings` or `strconv`
- Using deprecated third-party packages with stdlib equivalents
</patterns>

<related>
depguard, gomodguard, usestdlibvars
