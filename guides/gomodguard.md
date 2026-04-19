# gomodguard

<instructions>
Gomodguard blocks specific modules from being imported. Teams use it to prevent deprecated packages, enforce internal library usage, or ban modules known to cause issues in their codebase.

Replace the blocked module with the approved alternative specified in `.golangci.yml` under `linters.settings.gomodguard.blocked.modules`. Common replacements: `io/ioutil` → `os` and `io` packages.
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
func readConfig(path string) ([]byte, error) {
    return os.ReadFile(path)
}
```
</examples>

<patterns>
- `io/ioutil` — replaced by `os` and `io` packages since Go 1.16
- `github.com/golang/protobuf` — replaced by `google.golang.org/protobuf`
- Deprecated testing libraries blocked by team policy
- Modules with known security vulnerabilities
</patterns>

<related>
gomoddirectives, forbidigo, importas
</related>
