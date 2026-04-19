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
- Database or cache connections established in `init()`
- Global flag parsing inside `init()`
- Registration patterns using `init()` to register handlers
- Configuration loading triggered implicitly at import time
</patterns>

<related>
gochecknoglobals, decorder, funcorder
</related>
