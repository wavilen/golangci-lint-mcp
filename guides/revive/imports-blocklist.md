# revive: imports-blocklist

<instructions>
Enforces that certain packages are not imported. Teams blocklist packages that are deprecated, unsafe, or do not follow project standards (e.g., `io/ioutil`).

Remove the blocked import and use the recommended alternative. Check the revive configuration for the specific blocklist and suggested replacements.
</instructions>

<examples>
## Bad
```go
import "io/ioutil" // deprecated since Go 1.16

func readConfig(path string) ([]byte, error) {
    return ioutil.ReadFile(path)
}
```

## Good
```go
import "os"

func readConfig(path string) ([]byte, error) {
    return os.ReadFile(path)
}
```
</examples>

<patterns>
- Deprecated standard library packages replaced by newer equivalents
- Internal packages that should not be used outside their module
- Known problematic or insecure third-party libraries
- Test-only packages imported in production code
</patterns>

<related>
blank-imports, dot-imports, duplicated-imports
