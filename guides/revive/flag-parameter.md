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
- Boolean parameters at the end of argument lists controlling flow
- Functions like `Do(x, true)` where callers must look up what `true` means
- Boolean flags added over time as the function gains optional behavior
- API methods like `Get(path, async bool)` mixing two modes in one function
- Constructor functions with boolean flags for optional features
</patterns>

<related>
function-result-limit, argument-limit
