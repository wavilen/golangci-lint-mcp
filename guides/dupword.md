# dupword

<instructions>
Dupword detects duplicate words in comments, such as "the the" or "is is". These are typically typos introduced during editing and reduce professionalism of code documentation.

Remove the duplicate word. Proofread comments after editing or restructuring sentences.
</instructions>

<examples>
## Bad
```go
// Parse parses the the configuration file and returns a Config.
func Parse(r io.Reader) (*Config, error) {
```

## Good
```go
// Parse parses the configuration file and returns a Config.
func Parse(r io.Reader) (*Config, error) {
```
</examples>

<patterns>
- "the the", "is is", "to to" — common editing artifacts
- Duplicate words split across line wraps in comments
- Copy-paste errors in doc comments
- Adjacent identical words in multi-sentence comments
</patterns>

<related>
godot, godox, godoclint
</related>
