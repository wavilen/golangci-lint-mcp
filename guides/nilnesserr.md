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
- Using error-returning values before checking the error
- Dereferencing a pointer before checking if it is nil
- Accessing slice/map data before validating the error
</patterns>

<related>
nilerr, errcheck, staticcheck
