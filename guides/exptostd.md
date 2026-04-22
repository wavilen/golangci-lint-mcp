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
- Replace `golang.org/x/exp/slices` with `slices` (stdlib since Go 1.21)
- Replace `golang.org/x/exp/maps` with `maps` (stdlib since Go 1.21)
- Use `strings` or `strconv` equivalents instead of third-party string utilities
- Switch deprecated third-party packages to their standard library equivalents
</patterns>

<related>
depguard, gomodguard, usestdlibvars
