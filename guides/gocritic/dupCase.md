# gocritic: dupCase

<instructions>
Detects duplicate case values in switch statements. Go does not allow duplicate case expressions — the compiler already rejects integer/string duplicates, but gocritic also catches duplicates involving constants, variable expressions, or type switches that may not be caught by the compiler.

Remove the duplicate case or consolidate the logic into a single case.
</instructions>

<examples>
## Bad
```go
switch status {
case http.StatusOK:
    handleOK()
case http.StatusOK: // duplicate
    handleSuccess()
}
```

## Good
```go
switch status {
case http.StatusOK:
    handleOK()
case http.Created:
    handleCreated()
}
```
</examples>

<patterns>
- Copy-pasted case values in long switch statements
- Duplicate constant expressions across cases
- Type switch with same type appearing twice
- Integer literal duplicates in enum-based switches
</patterns>

<related>
dupBranchBody, dupArg, caseOrder
</related>
