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
- Replace `dir + "/" + file` with `filepath.Join(dir, file)`
- Replace `fmt.Sprintf("%s/%s", ...)` path construction with `filepath.Join(...)`
- Replace `base + "\\" + name` with `filepath.Join(base, name)` for cross-platform support
</patterns>

<related>
filepathJoin, stringConcatSimplify
</related>
