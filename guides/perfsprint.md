# perfsprint

<instructions>
Perfsprint detects `fmt.Sprintf` calls that can be replaced with faster alternatives. `fmt.Sprintf` uses reflection and allocates — for simple conversions, strconv functions are significantly faster and allocate less.

Use `strconv.Itoa`, `strconv.FormatInt`, `strconv.FormatFloat`, or simple string concatenation instead of `fmt.Sprintf` for non-formatting cases.
</instructions>

<examples>
## Bad
```go
s := fmt.Sprintf("%d", port)
s := fmt.Sprintf("%s", name)
```

## Good
```go
s := strconv.Itoa(port)
s := name // no conversion needed
```
</examples>

<patterns>
- Replace `fmt.Sprintf("%d", x)` with `strconv.Itoa` or `strconv.FormatInt`
- Use the string directly instead of `fmt.Sprintf("%s", s)`
- Replace `fmt.Sprintf("%t", b)` with `strconv.FormatBool`
- Use `strconv.FormatX` or `fmt.Sprint` instead of `fmt.Sprintf("%v", x)`
</patterns>

<related>
nosprintfhostport, govet, errcheck
</related>
