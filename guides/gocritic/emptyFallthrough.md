# gocritic: emptyFallthrough

<instructions>
Detects `fallthrough` statements in `switch` `case` clauses that do nothing before falling through — the `fallthrough` is the only statement. An empty fallthrough makes the case identical to having no body, which is misleading.

Remove the empty case body and `fallthrough`, or merge the case label with the next one.
</instructions>

<examples>
## Good
```go
switch x {
case 1, 2:
	handleTwo()
}
```
</examples>

<patterns>
- Replace empty `fallthrough` with merged case bodies or explicit logic
- Remove empty case bodies that only contain `fallthrough`
- Combine sequential empty fallthrough cases into a single case body
</patterns>

<related>
gocritic/singleCaseSwitch, gocritic/defaultCaseOrder
</related>
