# godot

<instructions>
Godot checks that comments end with a period. Go doc conventions require sentences to be properly punctuated for godoc rendering and readability.

Add a period at the end of each comment sentence. For multi-sentence comments, ensure every sentence ends with punctuation.
</instructions>

<examples>
## Bad
```go
// Parse reads the input and returns a structured result
func Parse(input string) (*Result, error) {
```

## Good
```go
// Parse reads the input and returns a structured result.
func Parse(input string) (*Result, error) {
```
</examples>

<patterns>
- Single-line doc comments missing trailing period
- Multi-sentence comments where only the last sentence lacks punctuation
- Package-level comments without proper sentence endings
- TODO/FIXME comments excluded by default configuration
</patterns>

<related>
dupword, godoclint, godox
</related>
