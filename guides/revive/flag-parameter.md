# revive: flag-parameter

<instructions>
Detects boolean parameters that control function behavior (flag parameters). A function like `Log(msg string, verbose bool)` essentially does two different things based on the flag. This makes the function harder to understand, test, and extend.

Split the function into two clearly named functions, or use an options struct. If the boolean represents a mode, consider a typed enum or a separate method.
</instructions>

<examples>
## Bad
```go
func render(data []byte, pretty bool) string {
    if pretty {
        return formatIndented(data)
    }
    return formatCompact(data)
}
```

## Good
```go
func render(data []byte) string {
    return formatCompact(data)
}

func renderPretty(data []byte) string {
    return formatIndented(data)
}
```
</examples>

<patterns>
- Replace boolean parameters controlling flow with separate methods for each behavior
- Separate functions like `Do(x, true)` into clearly named methods — callers shouldn't need to look up what `true` means
- Extract boolean flags added over time into an options struct or separate functions
- Provide distinct methods for each mode instead of mixing two modes in one function
- Use an options struct for constructor functions with boolean flags for optional features
</patterns>

<related>
function-result-limit, argument-limit
