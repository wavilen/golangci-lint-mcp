package server

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/wavilen/golangci-lint-mcp/internal/guides"

	"github.com/mark3labs/mcp-go/mcp"
)

type linterEntry struct {
	name       string
	rules      []string
	isCompound bool
}

func makeListHandler(store *guides.Store) func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		names := store.LinterNames()
		if len(names) == 0 {
			return mcp.NewToolResultText("No linters found."), nil
		}

		entries := make([]linterEntry, 0, len(names))
		for _, name := range names {
			rules := store.ListRules(name)
			entries = append(entries, linterEntry{
				name:       name,
				rules:      rules,
				isCompound: len(rules) > 0,
			})
		}

		sort.Slice(entries, func(i, j int) bool {
			return entries[i].name < entries[j].name
		})

		var compoundEntries []linterEntry
		var simpleEntries []linterEntry
		for _, entry := range entries {
			if entry.isCompound {
				compoundEntries = append(compoundEntries, entry)
			} else {
				simpleEntries = append(simpleEntries, entry)
			}
		}

		var builder strings.Builder
		fmt.Fprintf(&builder, "## Supported Linters (%d)\n\n", len(entries))

		if len(compoundEntries) > 0 {
			builder.WriteString("### Compound Linters (require rule parameter)\n")
			builder.WriteString("| Linter | Rules |\n")
			builder.WriteString("|--------|-------|\n")
			for _, entry := range compoundEntries {
				fmt.Fprintf(&builder, "| %s | %d rules |\n", entry.name, len(entry.rules))
			}
			builder.WriteString("\n")
		}

		if len(simpleEntries) > 0 {
			builder.WriteString("### Simple Linters (linter name only)\n")
			simpleNames := make([]string, 0, len(simpleEntries))
			for _, entry := range simpleEntries {
				simpleNames = append(simpleNames, entry.name)
			}
			builder.WriteString(strings.Join(simpleNames, ", "))
			builder.WriteString("\n")
		}

		return mcp.NewToolResultText(builder.String()), nil
	}
}
