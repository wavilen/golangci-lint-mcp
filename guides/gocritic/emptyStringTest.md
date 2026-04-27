# gocritic: emptyStringTest

<instructions>
Detects verbose empty-string checks such as `len(s) == 0` or `s == ""` where a simpler idiom exists. For strings, `s == ""` is preferred over `len(s) == 0`. For byte slices, `len(b) == 0` is correct.

Use `s == ""` for string emptiness checks instead of `len(s) == 0`.
</instructions>

<examples>
## Good
```go
if name == "" {
	return errors.New("name required")
}
```
</examples>

<patterns>
- Replace `len(s) == 0` with `s == ""` for string empty checks
- Replace `len(s) != 0` with `s != ""` for string non-empty checks
- Use `s == ""` directly — no further simplification needed
</patterns>

<related>
gocritic/stringConcatSimplify, gocritic/stringsCompare
</related>
