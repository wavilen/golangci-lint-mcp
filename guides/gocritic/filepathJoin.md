# gocritic: filepathJoin

<instructions>
Detects `filepath.Join` calls where one of the arguments is an absolute path. When an absolute path is passed to `filepath.Join`, it discards all previous path components, which is almost certainly unintended. For example, `filepath.Join("/data", "/etc/passwd")` results in `/etc/passwd`.

Ensure all path components are relative. If an absolute path is intentional, use it directly without `filepath.Join`.
</instructions>

<examples>
## Bad
```go
path := filepath.Join("/data", "/config.yaml")
// Result: "/config.yaml" — absolute path discards "/data"
```

## Good
```go
path := filepath.Join("/data", "config.yaml")
// Result: "/data/config.yaml"
```
</examples>

<patterns>
- `filepath.Join(baseDir, "/"+filename)` with leading slash
- User-provided paths that may be absolute
- `filepath.Join` where second argument starts with `/`
- Cross-platform path construction with hardcoded separators
</patterns>

<related>
argOrder, badCall
</related>
