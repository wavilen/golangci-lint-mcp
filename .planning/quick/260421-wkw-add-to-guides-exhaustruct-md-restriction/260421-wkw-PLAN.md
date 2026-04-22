---
phase: quick
plan: 1
type: execute
wave: 1
depends_on: []
files_modified:
  - guides/exhaustruct.md
autonomous: true
requirements: [guide-enhancement]
must_haves:
  truths:
    - "Guide warns agents not to suppress exhaustruct via .golangci.yml config changes"
    - "Guide recommends functional options pattern for structs with many fields"
    - "Functional options example shows a realistic Go code snippet"
  artifacts:
    - path: "guides/exhaustruct.md"
      provides: "Updated exhaustruct guide with restriction and recommendation"
      contains: "functional options"
  key_links: []
---

<objective>
Add a restriction note and a functional options pattern recommendation to guides/exhaustruct.md.

Purpose: Prevent agents from suppressing exhaustruct diagnostics via config changes, and guide them toward the idiomatic Go solution (functional options) for structs with many fields.
Output: Updated guides/exhaustruct.md
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@guides/exhaustruct.md
@guides/_template.md
</context>

<tasks>

<task type="auto">
  <name>task 1: add restriction and functional options recommendation to exhaustruct guide</name>
  <files>guides/exhaustruct.md</files>
  <action>
Edit `guides/exhaustruct.md` to add two additions while preserving the existing XML-tagged section structure:

**1. Restriction in `<instructions>`:** Append a sentence after the existing instructions text (before the closing `</instructions>` tag):

> Do not suppress exhaustruct diagnostics by adding struct types to the exclusion list in `.golangci.yml`. Fix the code instead — unlisted struct types can silently accumulate zero-value bugs.

**2. New `<recommendation>` section after `<examples>` and before `<patterns>`:** Add a section with a brief intro and a code example showing the functional options pattern. Use this structure:

```markdown
<recommendation>
## Functional Options Pattern

For structs with many exported fields where initializing every field in a literal is impractical, use a constructor with the functional options pattern. This lets callers set only the fields they need while providing sensible defaults for the rest.

```go
type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.Timeout = d }
}

func WithLogger(l Logger) Option {
    return func(s *Server) { s.Logger = l }
}

func NewServer(addr string, opts ...Option) *Server {
    s := &Server{
        Addr:    addr,
        Timeout: 30 * time.Second, // sensible default
        Logger:  defaultLogger,    // sensible default
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

Usage:

```go
// Only specify what differs from defaults
srv := NewServer(":8080", WithTimeout(10 * time.Second))
```
</recommendation>
```

Note: The nested code blocks inside the `<recommendation>` section should use standard markdown fenced code blocks. Since the guide uses XML tags for section boundaries (not markdown nesting), the inner code fences will work fine — they're parsed as markdown content within the `<recommendation>` XML element.

The final section order should be: `<instructions>` → `<examples>` → `<recommendation>` → `<patterns>` → `<related>`.
  </action>
  <verify>
    <automated>grep -c "functional options\|Do not suppress\|exclusion list" guides/exhaustruct.md</automated>
  </verify>
  <done>
    - guides/exhaustruct.md contains a restriction warning against modifying .golangci.yml in the `<instructions>` section
    - guides/exhaustruct.md contains a new `<recommendation>` section with a functional options pattern code example
    - Section order is: instructions → examples → recommendation → patterns → related
    - Existing content is preserved unchanged
  </done>
</task>

</tasks>

<verification>
```bash
# Verify the file has both additions
grep -q "Do not suppress" guides/exhaustruct.md
grep -q "functional options" guides/exhaustruct.md
grep -q "WithTimeout" guides/exhaustruct.md
# Verify section order
grep -n "^<" guides/exhaustruct.md
```
</verification>

<success_criteria>
- guides/exhaustruct.md updated with restriction note and functional options recommendation
- Existing sections preserved
- File is valid and reads naturally
</success_criteria>

<output>
After completion, create `.planning/quick/260421-wkw-add-to-guides-exhaustruct-md-restriction/260421-wkw-SUMMARY.md`
</output>
