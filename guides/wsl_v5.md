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
- Add blank lines before `if`, `for`, `switch` statements that follow assignments
- Remove blank lines between `if` and its `else` or `else if` clause
- Remove blank lines at the start or end of block bodies
- Separate statements within blocks with blank lines when required by the style config
</patterns>

<related>
whitespace, nlreturn, nakedret
</related>
