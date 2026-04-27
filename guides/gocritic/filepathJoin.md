# gocritic: filepathJoin

<instructions>
Detects `filepath.Join` calls where one of the arguments is an absolute path. When an absolute path is passed to `filepath.Join`, it discards all previous path components, which is almost certainly unintended. For example, `filepath.Join("/data", "/etc/passwd")` results in `/etc/passwd`.

Ensure all path components are relative. If an absolute path is intentional, use it directly without `filepath.Join`.
</instructions>

<examples>
## Good
```go
path := filepath.Join("/data", "config.yaml")
// Result: "/data/config.yaml"
```
</examples>

<patterns>
- Remove leading slash in `filepath.Join` arguments — `filepath.Join(baseDir, filename)` not `"/"+filename`
- Validate user-provided paths — reject absolute paths before joining
- Remove hardcoded `/` separators — use `filepath.Join` for all path construction
- Replace hardcoded separators with `filepath.Join` for cross-platform compatibility
</patterns>

<related>
gocritic/argOrder, gocritic/badCall
</related>
