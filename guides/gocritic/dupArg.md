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
- Replace `strings.Replace` duplicate old/new arguments — pass distinct values
- Replace `strings.Trim(s, s)` — use a separate cutset, not the string itself
- Replace `copy(a, a)` — pass distinct source and destination slices
- Replace `fmt.Sprintf` duplicate values for different format verbs — pass correct arguments
- Replace `math.Max(a, a)` — pass distinct values to compare
</patterns>

<related>
dupSubExpr, dupBranchBody, argOrder
</related>
