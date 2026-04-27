# revive: comments-density

<instructions>
Detects packages or functions with insufficient comment density relative to a configured threshold. Well-commented code is easier to maintain and onboard new developers. This rule ensures a minimum ratio of comments to code.

Add meaningful doc comments to exported functions, types, and constants. Focus on explaining why, not what — the code already shows what it does.
</instructions>

<examples>
## Good
```go
// Process applies the transformation indicated by flags to the input data.
// When bit 0 of flags is set, the data is transformed; otherwise it is analyzed.
func Process(data []byte, flags int) (Result, error) {
    if flags&flagTransform != 0 {
        return transform(data)
    }
    return analyze(data)
}
```
</examples>

<patterns>
- Add doc comments to exported functions and types explaining their purpose
- Add a package-level comment starting with "Package {name}" for every package
- Document complex algorithms with inline comments explaining the approach
- Annotate configuration or magic values with comments describing their purpose
- Add targeted inline comments to clarify individual steps in large functions
</patterns>

<related>
revive/exported, revive/comment-spacings, godoclint
</related>
