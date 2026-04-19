# gocritic: preferFilepathJoin

<instructions>
Detects manual string concatenation for building file paths using `+` or `fmt.Sprintf` instead of `filepath.Join`. Manual concatenation is error-prone — it may produce double slashes, missing separators, or platform-specific paths.

Use `filepath.Join` to construct file paths safely across operating systems.
</instructions>

<examples>
## Bad
```go
path := dir + "/" + filename
path := fmt.Sprintf("%s/%s", dir, file)
```

## Good
```go
path := filepath.Join(dir, filename)
```
</examples>

<patterns>
- `dir + "/" + file` → `filepath.Join(dir, file)`
- `fmt.Sprintf("%s/%s", ...)` for paths → `filepath.Join(...)`
- `base + "\\" + name` on Windows → `filepath.Join(base, name)`
</patterns>

<related>
filepathJoin, stringConcatSimplify
</related>
