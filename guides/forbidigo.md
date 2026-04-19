# forbidigo

<instructions>
Forbidigo forbids usage of specific functions, methods, or identifiers configured by the user. Teams use it to prevent deprecated patterns, discourage print statements, or block dangerous standard library functions.

Replace the forbidden identifier with the approved alternative. Check project configuration (`.golangci.yml`) for the specific `forbid` rules and suggested replacements.
</instructions>

<examples>
## Bad
```go
func main() {
    fmt.Println("server started")
}
```

## Good
```go
func main() {
    slog.Info("server started")
}
```
</examples>

<patterns>
- `fmt.Println` or `fmt.Printf` instead of structured logging
- `ioutil` package functions deprecated since Go 1.16
- Deprecated functions configured in `.golangci.yml` forbid list
- Testing helpers like `t.Log` used outside test files
</patterns>

<related>
godot, gomoddirectives, predeclared
</related>
