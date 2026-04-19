# gocritic: caseOrder

<instructions>
Detects switch cases where a later case can never be reached because an earlier case already covers it. This commonly occurs in type switches or value switches with overlapping conditions. The first matching case always wins, so subsequent overlapping cases are dead code.

Reorder cases so more specific cases come before more general ones, or remove the unreachable case entirely.
</instructions>

<examples>
## Bad
```go
switch v := x.(type) {
case io.Reader:
    handleReader(v)
case io.ReadCloser: // unreachable — ReadCloser is a Reader
    handleReadCloser(v)
}
```

## Good
```go
switch v := x.(type) {
case io.ReadCloser:
    handleReadCloser(v)
case io.Reader:
    handleReader(v)
}
```
</examples>

<patterns>
- Type switches with interface subtypes in wrong order
- Value switches with `case 1:` before `case 1, 2, 3:`
- Overlapping string patterns in switch cases
- Fallthrough making later cases redundant
</patterns>

<related>
dupCase, dupBranchBody, evalOrder
</related>
