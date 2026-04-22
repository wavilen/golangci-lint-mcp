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
- Replace deprecated standard library packages with their newer equivalents
- Remove internal packages imported outside their intended module boundary
- Replace known problematic or insecure third-party libraries with safe alternatives
- Move test-only package imports out of production code
</patterns>

<related>
blank-imports, dot-imports, duplicated-imports
