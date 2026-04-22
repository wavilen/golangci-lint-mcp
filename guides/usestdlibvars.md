# usestdlibvars

<instructions>
Usestdlibvars detects opportunities to use standard library constants and variables instead of hard-coded values. Using stdlib vars like `http.StatusOK` instead of `200` improves readability and prevents typos.

Replace magic numbers and strings with their stdlib equivalents: HTTP status codes, OS dev/null paths, time layouts, and more.
</instructions>

<examples>
## Bad
```go
w.WriteHeader(200)
fmt.Sprintf("200 OK")
os.Open("/dev/null")
```

## Good
```go
w.WriteHeader(http.StatusOK)
fmt.Sprintf("%d OK", http.StatusOK)
os.Open(os.DevNull)
```
</examples>

<patterns>
- Replace hardcoded HTTP status codes with `http.Status*` constants
- Replace string HTTP methods with `http.Method*` constants
- Replace hardcoded OS paths with `os.*` constants
- Replace hardcoded time layout strings with `time.*` constants
</patterns>

<related>
perfsprint, goconst
</related>
