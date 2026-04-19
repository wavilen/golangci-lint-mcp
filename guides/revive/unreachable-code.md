# revive: unreachable-code

<instructions>
Detects code that follows a `return`, `break`, `continue`, `panic`, or `goto` statement. Such code can never execute and indicates a logic error, leftover debugging, or incomplete refactoring. Unreachable code confuses readers and may hide bugs.

Remove the dead code after the terminating statement. If the code should execute, move it before the `return`/`break`/`panic`.
</instructions>

<examples>
## Bad
```go
func load() ([]byte, error) {
    return nil, errors.New("not implemented")
    log.Println("loading data") // unreachable
}
```

## Good
```go
func load() ([]byte, error) {
    log.Println("loading data")
    return nil, errors.New("not implemented")
}
```
</examples>

<patterns>
- Code after `return` in a function or branch
- Statements after `break` or `continue` in a loop
- Code following `panic()` that would never execute
- Fallthrough code in the final case of a type switch
- Dead code after `log.Fatal` or `os.Exit` calls
</patterns>

<related>
unnecessary-stmt, unconditional-recursion, deep-exit
