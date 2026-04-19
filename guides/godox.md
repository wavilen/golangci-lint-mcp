# godox

<instructions>
Godox detects TODO, FIXME, and BUG comments left in your code. While these markers are useful during development, they can accumulate and mask real issues if never addressed.

Either resolve the flagged item or convert it into a tracked issue in your project management system rather than leaving it as a comment.
</instructions>

<examples>
## Bad
```go
func handleSubmit(w http.ResponseWriter, r *http.Request) {
    // TODO: add input validation
    // FIXME: this crashes on empty body
    process(r.Body)
}
```

## Good
```go
func handleSubmit(w http.ResponseWriter, r *http.Request) {
    if err := validate(r); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    process(r.Body)
}
```
</examples>

<patterns>
- TODO comments added during prototyping that were never revisited
- FIXME markers for known bugs left unfixed across releases
- HACK or XXX comments indicating workarounds that need proper solutions
</patterns>

<related>
godoclint, revive, gocritic
</related>
