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
- Replace `io/ioutil` imports with `os` and `io` package equivalents
- Replace `github.com/golang/protobuf` with `google.golang.org/protobuf`
- Replace deprecated testing libraries with approved alternatives
- Remove modules with known vulnerabilities and use secure alternatives
</patterns>

<related>
gomoddirectives, forbidigo, importas
</related>
