# gocritic: dupArg

<instructions>
Detects function calls where the same argument is passed multiple times to parameters that expect different values. Passing identical values to distinct parameters is usually a copy-paste mistake rather than intentional.

Pass the correct distinct arguments to each parameter. If the duplication is intentional, assign the argument to a named variable for clarity.
</instructions>

<examples>
## Bad
```go
strings.ReplaceAll(s, "-", "-") // old and new are identical
```

## Good
```go
strings.ReplaceAll(s, "-", "_")
```
</examples>

<patterns>
- `strings.Replace(s, old, old, n)` — same old and new
- `strings.Trim(s, s)` — same string and cutset
- `copy(a, a)` — source and destination identical
- `fmt.Sprintf("%d %d", x, x)` — same value for different format verbs
- `math.Max(a, a)` — comparing a value with itself
</patterns>

<related>
dupSubExpr, dupBranchBody, argOrder
</related>
