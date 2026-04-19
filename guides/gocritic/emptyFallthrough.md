# gocritic: emptyFallthrough

<instructions>
Detects `fallthrough` statements in `switch` `case` clauses that do nothing before falling through — the `fallthrough` is the only statement. An empty fallthrough makes the case identical to having no body, which is misleading.

Remove the empty case body and `fallthrough`, or merge the case label with the next one.
</instructions>

<examples>
## Bad
```go
switch x {
case 1:
	fallthrough
case 2:
	handleTwo()
}
```

## Good
```go
switch x {
case 1, 2:
	handleTwo()
}
```
</examples>

<patterns>
- Case containing only `fallthrough`
- Empty case body followed by `fallthrough`
- Sequential cases all falling through with empty bodies
</patterns>

<related>
singleCaseSwitch, defaultCaseOrder
</related>
