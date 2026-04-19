# wsl_v5

<instructions>
WSL (Whitespace Linter, v5) enforces strict rules about when blank lines should appear in code. It requires blank lines before `if`, `for`, `switch`, and other block statements (except when directly after another block), and forbids blank lines in certain positions.

Add or remove blank lines to match the enforced style. Blocks should have a blank line before them unless they follow another block (like `else if`), and no blank lines inside the block at start/end.
</instructions>

<examples>
## Bad
```go
func process(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty")
    }
    result := parse(data)
    if result.Invalid {
        return errors.New("invalid")
    }
    return nil
}
```

## Good
```go
func process(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty")
    }

    result := parse(data)

    if result.Invalid {
        return errors.New("invalid")
    }

    return nil
}
```
</examples>

<patterns>
- Missing blank line before `if`, `for`, `switch` statements after assignments
- Unnecessary blank lines between `if` and its `else` clause
- Blank lines inside block bodies at start or end
- Multiple statements in a block without proper separation
</patterns>

<related>
whitespace, nlreturn, nakedret
</related>
