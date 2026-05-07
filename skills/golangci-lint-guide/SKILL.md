---
name: golangci-lint-guide
description: Validate and fix Go code using golangci-lint. Use this skill whenever working on Go projects — run golangci-lint, diagnose issues, and apply fixes per package.
---

<objective>
Fix golangci-lint issues using MCP tools that run golangci-lint, parse results, and provide fix guidance with related context. The MCP server handles linter classification, strategy recommendation, and package breakdown automatically.
</objective>

<execution_context>
- **MCP tools:** `golangci_lint_run` (run + parse), `golangci_lint_parse` (bulk JSON), `golangci_lint_guide` (per-diagnostic), `golangci_lint_list` (discover linters), `golangci_lint_summarize` (strategy only), `gosec_ai_autofix` (conditional)
- **CLI tool:** `golangci-lint` (must be installed for golangci_lint_run)
- **Claude Code hooks:** PreToolUse (command modification) + PostToolUse (MCP nudge injection) auto-configure golangci-lint for JSON output
- **Cursor hooks:** Same PreToolUse + PostToolUse automation via `.cursor/hooks.json` — auto-installed by `npx golangci-lint-guide` when Cursor is detected
- **Response format:** Tools return XML-tagged sections: `<summary>` (stats + breakdowns), `<guidance>` (per-issue fix directions), `<related_context>` (related linters with fix hints). Tools omit sections they don't have.
- **CLI version:** golangci-lint v2 — v1 flags like `--out-format` and `--output` do not work; use `--output.json.path stdout` for JSON output
</execution_context>

<process>

## 1. Run Per Package

Call `golangci_lint_run` with a specific package path:

```
golangci_lint_run(path="./pkg/auth/...")
```

Returns: `<summary>` with stats, `<guidance>` with fix directions for all unique (linter, rule) pairs, `<related_context>` for related issues, and strategy recommendation.

## 2. Fix All Issues

For each diagnostic in the response:
1. Use the provided guidance (instructions, patterns, examples)
2. Apply the fix directly
3. Fix related issues highlighted in Related Context — they're in the same package

## 3. Verify Package

Call `golangci_lint_run` again with the same path. If "No issues found", package is clean.

## 4. Strategy Instructions (Auto-Provided by MCP)

When `golangci_lint_run` returns a `<strategy_instructions>` block, you **MUST follow it exactly**.

The MCP tool automatically determines the correct strategy based on issue count and package count:
- **single-agent** — fix issues yourself (≤30 issues)
- **subagent-per-file** — spawn one subagent per file (>30 issues, few packages)
- **subagent-per-package** — spawn one subagent per package (>3 packages)

The `<strategy_instructions>` block contains step-by-step commands with the exact package/file paths to use. Do NOT override or ignore these instructions — they prevent context-limit exhaustion.

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
| golangci_lint_run: "binary not found" | Use `golangci_lint_guide(linter="...", rule="...")` for per-diagnostic guidance as fallback |
| golangci_lint_run: timeout with partial results | Note partial issue count, scan per-package paths (e.g., `./pkg/auth/...`) |
| golangci_lint_parse: "invalid JSON" | Try `golangci_lint_run` on a specific package, or use `golangci_lint_guide` per diagnostic |
| golangci_lint_guide: "Unknown linter" | Check for typos; may be from newer/older golangci-lint version — use `golangci_lint_list` to verify |
| golangci-lint CLI: "unknown flag: --out-format" | You are using v1 flags with golangci-lint v2. Use `golangci_lint_run` MCP tool instead. If CLI is needed: `--output.json.path stdout` is the v2 equivalent |

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

**Strategy:** MCP tools auto-detect strategy and include `<strategy_instructions>` block in their response. Always follow the instructions in that block. Do NOT attempt to fix >30 issues without subagents.

**Compound linters** (require `rule` param in `golangci_lint_guide`): staticcheck, gocritic, gosec, revive, govet, testifylint, modernize, errorlint, ginkgolinter, grouper. Call `golangci_lint_list` for full list.

**Response format:** All tools use XML tags — `<summary>`, `<guidance>`, `<related_context>`. Markdown content (tables, lists, code blocks) appears inside these tags.

</quick_reference>

<cli>

## CLI: intercept

The `intercept` subcommand runs golangci-lint from the terminal without an MCP client:

```bash
golangci-lint-mcp intercept [--raw] <path>
```

- **`--raw`**: Output raw golangci-lint JSON without parsing or summarization
- Auto-fix is always enabled (runs golangci-lint with `--fix`)

Use intercept when you need terminal-based output or want to pipe results to other tools. Uses the same guide-enrichment pipeline as MCP tools.

</cli>
