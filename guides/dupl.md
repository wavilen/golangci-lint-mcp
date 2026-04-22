# dupl

<instructions>
Dupl detects code clones — sequences of statements that are duplicated across your codebase. Duplicated code increases maintenance burden because fixes must be applied in multiple places.

Extract the duplicated logic into a shared function, method, or helper, then call it from each original location.
</instructions>

<examples>
## Bad
```go
func validateName(name string) error {
    if len(name) == 0 {
        return fmt.Errorf("name is required")
    }
    if len(name) > 100 {
        return fmt.Errorf("name too long")
    }
    return nil
}

func validateEmail(email string) error {
    if len(email) == 0 {
        return fmt.Errorf("email is required")
    }
    if len(email) > 100 {
        return fmt.Errorf("email too long")
    }
    return nil
}
```

## Good
```go
func validateField(name, value string, maxLen int) error {
    if len(value) == 0 {
        return fmt.Errorf("%s is required", name)
    }
    if len(value) > maxLen {
        return fmt.Errorf("%s too long", name)
    }
    return nil
}
```
</examples>

<patterns>
- Extract copy-pasted validation or error-handling logic into a shared function
- Extract repeated struct initialization into a shared factory function
- Replace duplicated iteration patterns with a shared higher-order function
- Extract repeated test setup code into a shared helper or table-driven test
</patterns>

<related>
funlen, gocyclo, gocognit
</related>
