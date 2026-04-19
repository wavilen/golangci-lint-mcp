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
		check       func(t *testing.T, g *Guide)
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
			wantErr: false,
			check: func(t *testing.T, g *Guide) {
				assert.Equal(t, "errcheck", g.Linter)
				assert.Equal(t, "", g.Rule)
				assert.Contains(t, g.Instructions, "Errcheck detects unchecked error returns")
				assert.Contains(t, g.Examples, "os.Open")
				assert.Contains(t, g.Patterns, "comma-ok pattern")
				assert.Contains(t, g.RawBody, "<instructions>")
				assert.Contains(t, g.RawBody, "</patterns>")
				assert.Equal(t, "errcheck", g.Key())
				assert.False(t, g.IsCompound())
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
			wantErr: false,
			check: func(t *testing.T, g *Guide) {
				assert.Equal(t, "gosimple", g.Linter)
				assert.Contains(t, g.Instructions, "Simplify code constructs")
				assert.Equal(t, "", g.Examples)
				assert.Equal(t, "", g.Patterns)
			},
		},
		{
			name:        "zero recognized tags rejected",
			filename:    "barelinter.md",
			content:     "# barelinter\n\nThis is just plain markdown.\n",
			wantErr:     true,
			errContains: "must contain at least one recognized XML tag",
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
			wantErr: false,
			check: func(t *testing.T, g *Guide) {
				assert.Contains(t, g.Instructions, "Some instructions here")
				assert.Contains(t, g.RawBody, "<tips>")
				assert.Contains(t, g.RawBody, "This is a custom tag")
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
			wantErr: false,
			check: func(t *testing.T, g *Guide) {
				assert.Contains(t, g.Examples, "```go")
				assert.Contains(t, g.Examples, "fmt.Printf")
				assert.Contains(t, g.Examples, "Multiple lines of code here")
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
			wantErr: false,
			check: func(t *testing.T, g *Guide) {
				assert.Equal(t, "gocritic", g.Linter)
				assert.Equal(t, "badcall", g.Rule)
				assert.Equal(t, "gocritic/badcall", g.Key())
				assert.True(t, g.IsCompound())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := Parse(tt.filename, []byte(tt.content))
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, g)
			if tt.check != nil {
				tt.check(t, g)
			}
		})
	}
}
