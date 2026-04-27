# revive: unreachable-code

<instructions>
Detects code that follows a `return`, `break`, `continue`, `panic`, or `goto` statement. Such code can never execute and indicates a logic error, leftover debugging, or incomplete refactoring. Unreachable code confuses readers and may hide bugs.

Remove the dead code after the terminating statement. If the code should execute, move it before the `return`/`break`/`panic`.
</instructions>

<examples>
## Good
```go
func load() ([]byte, error) {
    log.Println("loading data")
    return nil, errors.New("not implemented")
}
```
</examples>

<patterns>
- Remove code placed after `return` in a function or branch
- Remove statements after `break` or `continue` in a loop — they never execute
- Remove code following `panic()` that would never run
- Eliminate fallthrough code in the final case of a type switch
- Remove dead code after `log.Fatal` or `os.Exit` calls
</patterns>

<related>
revive/unnecessary-stmt, revive/unconditional-recursion, revive/deep-exit
</related>
