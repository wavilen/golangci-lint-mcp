# gochecknoinits

<instructions>
Gochecknoinits checks that no `init()` functions are declared. Init functions run implicitly, have unpredictable ordering across files, and make dependencies invisible to callers.

Replace `init()` with an explicit setup function called from `main()` or a constructor. This makes initialization order visible and testable.
</instructions>

<examples>
## Bad
```go
func init() {
    db.Connect("postgres://...")
}
```

## Good
```go
func NewApp() (*App, error) {
    db, err := db.Connect("postgres://...")
    if err != nil {
        return nil, err
    }
    return &App{db: db}, nil
}
```
</examples>

<patterns>
- Replace `init()` with an explicit setup function called from `main()` or a constructor
- Move flag parsing into `main()` instead of global `init()`
- Use explicit registration functions instead of `init()` handler registration
- Load configuration in a constructor rather than implicitly at import time
</patterns>

<related>
gochecknoglobals, decorder, funcorder
</related>
