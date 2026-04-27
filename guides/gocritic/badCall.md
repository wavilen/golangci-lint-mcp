# gocritic: badCall

<instructions>
Detects suspicious function calls where the result is ignored despite being the only meaningful output. Common with `fmt.Sprintf` (useless without using result), `strings.Builder.Reset` called before reading, or `append` result not captured.

Ensure the function call's return value is used or explicitly discarded with justification.
</instructions>

<examples>
## Good
```go
msg := fmt.Sprintf("status: %d", code)
log.Println(msg)
```
</examples>

<patterns>
- Capture the return value from function calls that produce meaningful output
- Use `fmt.Sprintf` result directly or assign it — never discard it
- Check for ignored returns from `append`, `strings.Builder` methods, and formatter functions
</patterns>

<related>
ineffassign, wastedassign
</related>
