# nilnesserr

<instructions>
Nilnesserr reports incorrect nil checks on errors after they have been used. It detects when code checks `if err != nil` after already dereferencing or using the error, meaning the nil check is redundant or misplaced.

Reorder code so the nil check happens before the error is used, or remove the redundant check.
</instructions>

<examples>
## Bad
```go
func read() ([]byte, error) {
    data, err := os.ReadFile("f.txt")
    fmt.Println(string(data))
    if err != nil {
        return nil, err
    }
    return data, nil
}
```

## Good
```go
func read() ([]byte, error) {
    data, err := os.ReadFile("f.txt")
    if err != nil {
        return nil, err
    }
    return data, nil
}
```
</examples>

<patterns>
- Check the error before using the value returned alongside it
- Check for nil before dereferencing a pointer returned with an error
- Validate the error before accessing slice or map data from the same call
</patterns>

<related>
nilerr, errcheck, staticcheck
