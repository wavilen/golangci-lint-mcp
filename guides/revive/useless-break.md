# revive: useless-break

<instructions>
Detects unnecessary `break` statements at the end of `case` or `default` clauses in `switch` and `select` statements. Unlike C-family languages, Go's `case` clauses do not fall through by default, so a `break` before the next `case` is redundant.

Remove the trailing `break` from the case clause. Only keep `break` when you need to exit early from a loop inside a case.
</instructions>

<examples>
## Bad
```go
switch color {
case "red":
    handleRed()
    break // unnecessary
case "blue":
    handleBlue()
    break
}
```

## Good
```go
switch color {
case "red":
    handleRed()
case "blue":
    handleBlue()
}
```
</examples>

<patterns>
- `break` at the end of every `case` in a switch (C-style habit)
- `break` at the end of `default` clause
- `break` in `select` case clauses
- Break statements in type switch cases
- Break in the last case of a switch where it serves no purpose
</patterns>

<related>
useless-fallthrough, unnecessary-stmt, unnecessary-if
