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
- HTTP status codes: `200` → `http.StatusOK`, `404` → `http.StatusNotFound`
- HTTP methods: `"GET"` → `http.MethodGet`
- OS paths: `"/dev/null"` → `os.DevNull`
- Time layouts: `"2006-01-02"` → `time.DateOnly`
</patterns>

<related>
perfsprint, goconst
</related>
