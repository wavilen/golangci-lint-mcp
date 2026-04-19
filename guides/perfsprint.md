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
- `fmt.Sprintf("%d", x)` → use `strconv.Itoa` or `strconv.FormatInt`
- `fmt.Sprintf("%s", s)` → use the string directly
- `fmt.Sprintf("%t", b)` → use `strconv.FormatBool`
- `fmt.Sprintf("%v", x)` → use `strconv.FormatX` or `fmt.Sprint`
</patterns>

<related>
nosprintfhostport, govet, errcheck
</related>
