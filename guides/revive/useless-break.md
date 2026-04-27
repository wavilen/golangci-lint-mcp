# revive: useless-break

<instructions>
Detects unnecessary `break` statements at the end of `case` or `default` clauses in `switch` and `select` statements. Unlike C-family languages, Go's `case` clauses do not fall through by default, so a `break` before the next `case` is redundant.

Remove the trailing `break` from the case clause. Only keep `break` when you need to exit early from a loop inside a case.
</instructions>

<examples>
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
- Remove trailing `break` at the end of every `case` in a switch — Go cases don't fall through
- Remove `break` at the end of `default` clause
- Remove `break` from `select` case clauses
- Eliminate break statements in type switch cases
- Remove `break` from the last case of a switch where it serves no purpose
</patterns>

<related>
revive/useless-fallthrough, revive/unnecessary-stmt, revive/unnecessary-if
</related>
