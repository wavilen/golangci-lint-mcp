package guides

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		content     string
		wantErr     bool
		errContains string
		check       func(t *testing.T, guide *Guide)
	}{
		{
			name:     "all three recognized tags",
			filename: "errcheck.md",
			content: `# errcheck

<instructions>
Errcheck detects unchecked error returns.
</instructions>

<examples>
## Bad
` + "```go\nfile, _ := os.Open(\"config.yaml\")\n```" + `

## Good
` + "```go\nfile, err := os.Open(\"config.yaml\")\nif err != nil {\n    return err\n}\n```" + `
</examples>

<patterns>
- Always check os.Open errors
- Use comma-ok pattern for type assertions
</patterns>
`,
			wantErr:     false,
			errContains: "",
			check: func(t *testing.T, guide *Guide) {
				assert.Equal(t, "errcheck", guide.Linter)
				assert.Empty(t, guide.Rule)
				assert.Contains(t, guide.Instructions, "Errcheck detects unchecked error returns")
				assert.Contains(t, guide.Examples, "os.Open")
				assert.Contains(t, guide.Patterns, "comma-ok pattern")
				assert.Contains(t, guide.RawBody, "<instructions>")
				assert.Contains(t, guide.RawBody, "</patterns>")
				assert.Equal(t, "errcheck", guide.Key())
				assert.False(t, guide.IsCompound())
			},
		},
		{
			name:     "one tag only accepted",
			filename: "gosimple.md",
			content: `# gosimple

<instructions>
Simplify code constructs.
</instructions>
`,
			wantErr:     false,
			errContains: "",
			check: func(t *testing.T, guide *Guide) {
				assert.Equal(t, "gosimple", guide.Linter)
				assert.Contains(t, guide.Instructions, "Simplify code constructs")
				assert.Empty(t, guide.Examples)
				assert.Empty(t, guide.Patterns)
			},
		},
		{
			name:        "zero recognized tags rejected",
			filename:    "barelinter.md",
			content:     "# barelinter\n\nThis is just plain markdown.\n",
			wantErr:     true,
			errContains: "must contain at least one recognized XML tag",
			check:       nil,
		},
		{
			name:     "custom tags preserved in raw body",
			filename: "customtag.md",
			content: `# customtag

<instructions>
Some instructions here.
</instructions>

<tips>
This is a custom tag that should be preserved.
</tips>
`,
			wantErr:     false,
			errContains: "",
			check: func(t *testing.T, guide *Guide) {
				assert.Contains(t, guide.Instructions, "Some instructions here")
				assert.Contains(t, guide.RawBody, "<tips>")
				assert.Contains(t, guide.RawBody, "This is a custom tag")
			},
		},
		{
			name:     "multiline content with code fences",
			filename: "govet.md",
			content: `# govet

<instructions>
Govet reports suspicious constructs.
</instructions>

<examples>
` + "```go\nfunc main() {\n    fmt.Printf(\"hello\")\n}\n```" + `

Multiple lines of code here.
</examples>
`,
			wantErr:     false,
			errContains: "",
			check: func(t *testing.T, guide *Guide) {
				assert.Contains(t, guide.Examples, "```go")
				assert.Contains(t, guide.Examples, "fmt.Printf")
				assert.Contains(t, guide.Examples, "Multiple lines of code here")
			},
		},
		{
			name:     "compound linter path parsing",
			filename: "gocritic/badcall.md",
			content: `# gocritic: badCall

<instructions>
Detects suspicious function calls.
</instructions>
`,
			wantErr:     false,
			errContains: "",
			check: func(t *testing.T, guide *Guide) {
				assert.Equal(t, "gocritic", guide.Linter)
				assert.Equal(t, "badcall", guide.Rule)
				assert.Equal(t, "gocritic/badcall", guide.Key())
				assert.True(t, guide.IsCompound())
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			guide, parseErr := Parse(testCase.filename, []byte(testCase.content))
			if testCase.wantErr {
				require.Error(t, parseErr)
				assert.Contains(t, parseErr.Error(), testCase.errContains)
				return
			}
			require.NoError(t, parseErr)
			require.NotNil(t, guide)
			if testCase.check != nil {
				testCase.check(t, guide)
			}
		})
	}
}
