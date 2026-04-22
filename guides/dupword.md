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
- Remove common editing artifacts like "the the", "is is", "to to"
- Check for duplicate words that span line wraps in comments
- Proofread doc comments for copy-paste errors introducing duplicate words
- Eliminate adjacent identical words in multi-sentence comments
</patterns>

<related>
godot, godox, godoclint
</related>
