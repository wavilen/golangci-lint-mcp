# revive: comments-density

<instructions>
Detects packages or functions with insufficient comment density relative to a configured threshold. Well-commented code is easier to maintain and onboard new developers. This rule ensures a minimum ratio of comments to code.

Add meaningful doc comments to exported functions, types, and constants. Focus on explaining why, not what — the code already shows what it does.
</instructions>

<examples>
## Bad
```go
func Process(data []byte, flags int) (Result, error) {
    if flags&0x01 != 0 {
        return transform(data)
    }
    return analyze(data)
}
```

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
- Exported functions or types without doc comments
- Packages lacking a package-level comment
- Complex algorithms with no inline explanation of the approach
- Configuration or magic values without contextual comments
- Large functions where individual steps need clarification
</patterns>

<related>
exported, comment-spacings, godoclint
