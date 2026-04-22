# gocritic: singleCaseSwitch

<instructions>
Detects `switch` statements that contain only a single case (no `default`, or only one case and a `default`). A single-case switch can be replaced with a simpler `if` statement for clarity.

Replace the single-case `switch` with an equivalent `if` statement.
</instructions>

<examples>
## Bad
```go
switch {
case x > 10:
	big()
}
```

## Good
```go
if x > 10 {
	big()
}
```
</examples>

<patterns>
- Replace `switch` with one `case` and no `default` with an `if` statement
- Replace `switch` with one `case` and `default` with `if`/`else`
- Replace tagged `switch` with a single `case` value with an `if` comparison
</patterns>

<related>
switchTrue, ifElseChain, emptyFallthrough
</related>
