# forbidigo

<instructions>
Forbidigo forbids usage of specific functions, methods, or identifiers configured by the user. Teams use it to prevent deprecated patterns, discourage print statements, or block dangerous standard library functions.

Replace the forbidden identifier with the approved alternative. Check project configuration (`.golangci.yml`) for the specific `forbid` rules and suggested replacements.
</instructions>

<examples>
## Good
```go
func main() {
    slog.Info("server started")
}
```
</examples>

<patterns>
- Replace `fmt.Println`/`fmt.Printf` with `slog.Info`/`slog.Error` for structured logging
- Use `os`/`io` equivalents instead of deprecated `ioutil` functions (Go 1.16+)
- Check `.golangci.yml` forbid list for approved replacements of deprecated functions
- Move `t.Log` calls into test files only, or use structured logging outside tests
</patterns>

<related>
godot, gomoddirectives, predeclared
</related>
