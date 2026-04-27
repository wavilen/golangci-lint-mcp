# govet: unreachable

<instructions>
Reports unreachable code — statements that follow a `return`, `panic`, `break`, `continue`, or `goto` in the same block. This code can never execute and is either dead logic or a mistake.

Remove the unreachable code or restructure the control flow.
</instructions>

<examples>
## Good
```go
func status() int {
    log.Println("returning status")
    return 200
}
```
</examples>

<patterns>
- Remove code after `return` in the same block — it can never execute
- Remove code after `panic()` in the same block — move it before the panic
- Remove code after `break`/`continue` in the same loop block
- Remove code after `goto` in the same block — it is never reached
</patterns>

<related>
govet/tests, govet/defers, staticcheck/SA4032
</related>
