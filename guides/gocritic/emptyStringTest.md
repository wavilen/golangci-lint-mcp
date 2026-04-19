# gocritic: emptyStringTest

<instructions>
Detects verbose empty-string checks such as `len(s) == 0` or `s == ""` where a simpler idiom exists. For strings, `s == ""` is preferred over `len(s) == 0`. For byte slices, `len(b) == 0` is correct.

Use `s == ""` for string emptiness checks instead of `len(s) == 0`.
</instructions>

<examples>
## Bad
```go
if len(name) == 0 {
	return errors.New("name required")
}
```

## Good
```go
if name == "" {
	return errors.New("name required")
}
```
</examples>

<patterns>
- `len(s) == 0` for strings → `s == ""`
- `len(s) != 0` for strings → `s != ""`
- `s == ""` is already optimal — no further simplification needed
</patterns>

<related>
stringConcatSimplify, stringsCompare
</related>
