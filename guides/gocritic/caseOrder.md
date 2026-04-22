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
- Reorder type switches — place more specific interface types before general ones
- Reorder value switches — place broader cases like `case 1, 2, 3:` before `case 1:`
- Eliminate overlapping string patterns in switch cases
- Remove fallthrough that makes later cases unreachable
</patterns>

<related>
dupCase, dupBranchBody, evalOrder
</related>
