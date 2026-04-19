# govet: unreachable

<instructions>
Reports unreachable code — statements that follow a `return`, `panic`, `break`, `continue`, or `goto` in the same block. This code can never execute and is either dead logic or a mistake.

Remove the unreachable code or restructure the control flow.
</instructions>

<examples>
## Bad
```go
func status() int {
    return 200
    log.Println("returned") // unreachable
}
```

## Good
```go
func status() int {
    log.Println("returning status")
    return 200
}
```
</examples>

<patterns>
- Code after `return` in same block
- Code after `panic()` in same block
- Code after `break` or `continue` in loop body
- Statements after `goto` label
</patterns>

<related>
tests, defers
</related>
