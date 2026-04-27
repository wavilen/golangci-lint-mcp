---
name: golangci-lint-guide
description: Fix golangci-lint issues using MCP tool guidance — run golangci-lint, look up each diagnostic, apply fixes per package
---

<objective>
Fix golangci-lint issues using MCP tools that run golangci-lint, parse results, and provide fix guidance with related context. The MCP server handles linter classification, strategy recommendation, and package breakdown automatically.
</objective>

<execution_context>
- **MCP tools:** `golangci_lint_run` (run + parse), `golangci_lint_parse` (bulk JSON), `golangci_lint_guide` (per-diagnostic), `golangci_lint_list` (discover linters), `golangci_lint_summarize` (strategy only), `gosec_ai_autofix` (conditional)
- **CLI tool:** `golangci-lint` (must be installed for golangci_lint_run)
</execution_context>

<process>

## 1. Run Per Package

Call `golangci_lint_run` with a specific package path:

```
golangci_lint_run(path="./pkg/auth/...")
```

Returns: fix guidance for all unique (linter, rule) pairs, Related Context for related issues, and strategy recommendation.

## 2. Fix All Issues

For each diagnostic in the response:
1. Use the provided guidance (instructions, patterns, examples)
2. Apply the fix directly
3. Fix related issues highlighted in Related Context — they're in the same package

## 3. Verify Package

Call `golangci_lint_run` again with the same path. If "No issues found", package is clean.

## 4. Large Output (>30 Issues): Subagent Per Package

When `golangci_lint_run` returns strategy "subagent-per-package":

**Step A:** Call `golangci_lint_run(path="./...")` for full-project summary with package breakdown.

**Step B:** For EACH package with issues, spawn a subagent scoped to that single package:

```
task(
  description="Fix golangci-lint issues in {package_path}",
  prompt="Use golangci_lint_run to fix all issues in {package_path}.",
  mode="subagent"
)
```

**Step C:** Final verification: `golangci_lint_run(path="./...")` → expect "No issues found".

**Why subagents?** Single agent at 30+ issues hits 70%+ context, producing incomplete fixes. Subagents give each package full context.

## 5. Gosec AI Autofix (Optional)

Only if `gosec_ai_autofix` is available. Group gosec diagnostics by package, call per package:

```
gosec_ai_autofix(path="./pkg/auth/...")
```

Never call with `"./..."`. On timeout, fall back to `golangci_lint_guide(linter="gosec", rule="<G-code>")`.

## 6. Final Verification

Call `golangci_lint_run(path="./...")`. Report any remaining issues.

</process>

<error_recovery>

## Error Recovery

| Error | Action |
|-------|--------|
| MCP tools unavailable | STOP — verify MCP server is running and configured |
| golangci_lint_run: "binary not found" | STOP — install golangci-lint (`brew install golangci-lint` or `curl install.sh`) |
| golangci_lint_run: timeout | Scan per-package instead of full-project |
| golangci_lint_parse: JSON parse error | Fall back to `golangci_lint_guide` per unique (linter, rule) pair |

</error_recovery>

<quick_reference>

## Quick Reference

| Tool | When | Key Parameter |
|------|------|---------------|
| `golangci_lint_run` | Run + get guidance | `path` (package or `./...`) |
| `golangci_lint_parse` | Parse existing JSON | `output` (raw JSON) |
| `golangci_lint_guide` | Single diagnostic | `linter` + `rule` |
| `golangci_lint_list` | Discover linters | (none) |
| `golangci_lint_summarize` | Strategy only | `output` (raw JSON) |

**Strategy threshold:** >30 issues or >3 packages → subagent-per-package (returned automatically by tools).

**Compound linters** (require `rule` param in `golangci_lint_guide`): staticcheck, gocritic, gosec, revive, govet, testifylint, modernize, errorlint, ginkgolinter, grouper. Call `golangci_lint_list` for full list.

</quick_reference>
