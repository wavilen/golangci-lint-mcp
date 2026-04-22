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
- Add trailing periods to single-line doc comments
- Check multi-sentence comments for missing punctuation on the final sentence
- End package-level comments with proper sentence punctuation
- Annotate TODO/FIXME with trailing periods unless excluded by configuration
</patterns>

<related>
dupword, godoclint, godox
</related>
