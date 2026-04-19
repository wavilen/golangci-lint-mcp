---
name: golangci-lint-guide
description: Fix golangci-lint issues using MCP tool guidance — run golangci-lint, look up each diagnostic, apply fixes per package
---

# golangci-lint-guide

Fix golangci-lint issues in Go projects using the golangci-lint-mcp server for targeted fix guidance.

## Objective

This skill provides a structured workflow for fixing golangci-lint issues. The agent runs golangci-lint on the project, then calls the MCP tool `golangci_lint_parse` to get fix guidance for all diagnostics at once. For individual lookups, `golangci_lint_guide` is also available.

## Execution Context

- **MCP tools:** `golangci_lint_parse` (bulk), `golangci_lint_guide` (per-diagnostic), `gosec_ai_autofix` (conditional) — provided by the golangci-lint-mcp server
- **CLI tool:** `golangci-lint` (must be installed in the project)

## Context

The `golangci_lint_guide` MCP tool accepts two parameters:

- **`linter`** (string, required) — The linter name, e.g., `errcheck`, `gocritic`, `gosec`, `revive`, `staticcheck`, `govet`.
- **`rule`** (string, optional) — The specific rule ID for compound linters, e.g., `badcall` for gocritic, `G101` for gosec, `SA1000` for staticcheck.

The `golangci_lint_parse` MCP tool accepts one parameter:

- **`output`** (string, required) — The raw golangci-lint JSON output string (from `golangci-lint run --output.json.path stdout`).

It parses the JSON, deduplicates by (linter, rule) pair, and returns guidance for all unique diagnostics in a single response. Prefer this for bulk fixes.

### Compound vs Simple Linters

**Compound linters** have per-rule guides and require the `rule` parameter for specific guidance:

- `gocritic` — 108 checkers
- `staticcheck` — 172 rules (SA/S/ST/QF codes)
- `revive` — 101 rules
- `gosec` — 61 rules (G-codes)
- `govet` — 35 analyzers
- `modernize` — 10 rules
- `testifylint` — 20 rules
- `ginkgolinter` — 12 rules
- `errorlint` — 3 rules
- `grouper` — 4 rules

If a compound linter is queried without a `rule`, the tool returns a list of all available rules for that linter.

**Simple linters** (all others) just need the `linter` name.

### Response Format

The tool returns XML-tagged guidance:

- `<instructions>` — What the issue means and how to fix it
- `<examples>` — Before/after code examples
- `<patterns>` — The fix pattern to apply
- `<related>` — Related linters or rules

## Process

### Step 1: Check Prerequisites

Verify golangci-lint is installed:

```bash
golangci-lint version
```

If not installed, tell the user to install it first:

```bash
# macOS
brew install golangci-lint

# Linux
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

### Step 2: Run golangci-lint

Execute golangci-lint with JSON output for structured parsing:

```bash
golangci-lint run --output.json.path stdout ./...
```

### Step 3: Get Fix Guidance

**Preferred: bulk parsing** — Call MCP tool `golangci_lint_parse` with the full JSON output:

```
golangci_lint_parse(output="<raw JSON output>")
```

This returns guidance for all unique (linter, rule) pairs in a single call. Deduplicates automatically — no need to call per diagnostic.

**Alternative: per-diagnostic lookup** — For individual issues, call `golangci_lint_guide` with `linter` and optionally `rule`.

If no issues are found, report success and stop.

### Step 4: Fix Issues Per Package

Process packages one at a time. For each package:

**a) For each diagnostic in the package:**

1. Use the guidance returned by `golangci_lint_parse` (or look up individually with `golangci_lint_guide`).
2. Apply the fix to the source file based on the guidance.

**b) Verify the package is clean:**

```bash
golangci-lint run <package-path>
```

**c) If issues remain**, repeat the fix cycle for that package.

**d) Move to the next package.**

### Step 5: Fix Gosec Issues with AI Autofix (Optional)

This step only applies if the `gosec_ai_autofix` tool is available (server started with `--gosec-ai` flag and `GOSEC_AI_API_KEY` env var). Skip to Step 6 if the tool is not registered.

**Why per-package batching:** gosec requires Go package context for type resolution — imports, type definitions, interface satisfaction, etc. Calling `gosec_ai_autofix` on the entire project (`"./..."`) may fail or produce incomplete fixes because gosec cannot resolve cross-package type information in a single pass. Each package must be analyzed in its own scope.

**a) Identify gosec-bearing packages:**

Group the golangci-lint JSON output (from Step 2) by `Pos.Filename`. For each diagnostic with `FromLinter: "gosec"`, extract the package directory (the directory containing the `.go` file). Collect unique package directories.

Example:
- `Pos.Filename: "pkg/auth/jwt.go"` with `FromLinter: "gosec"` → package directory = `"./pkg/auth/..."`
- `Pos.Filename: "internal/server/handler.go"` with `FromLinter: "gosec"` → package directory = `"./internal/server/..."`

**b) Call gosec_ai_autofix per package:**

For each gosec-bearing package directory, call:

```
gosec_ai_autofix(path="./pkg/auth/...")
```

**Do NOT call** `gosec_ai_autofix(path="./...")` — use the individual package path.

**c) Fallback on failure:**

If `gosec_ai_autofix` fails or times out (60s limit) for a package, fall back to manual fixes using `golangci_lint_guide(linter="gosec", rule="<G-code>")` for each failed diagnostic. Apply fixes based on the guide's `<instructions>` and `<examples>` sections.

### Step 6: Final Verification

After all packages pass, run a full project check:

```bash
golangci-lint run ./...
```

### Step 7: Report

Summarize:
- Which linters were fixed
- How many issues per linter
- Any remaining issues that couldn't be auto-fixed

## Notes

### Compound Linter Rule Extraction

The `golangci_lint_parse` tool automatically extracts rule IDs from compound linter text (the prefix before `: ` in the `Text` field). When using `golangci_lint_guide` directly, you need to extract rules manually:

- **gocritic:** checker name appears in the message (e.g., `appendAssign`, `badCall`, `dupImport`)
- **gosec:** rule ID is the G-code (e.g., `G101`, `G201`, `G401`)
- **staticcheck:** rule ID is the SA/S/ST/QF code (e.g., `SA1000`, `S1003`)
- **revive:** rule name appears in the message (e.g., `exported`, `unused-parameter`, `var-naming`)
- **govet:** analyzer name from `FromLinter` field (already `govet/<analyzer>`)
- **modernize, testifylint, ginkgolinter, errorlint, grouper:** rule name in `Text` field

### --gosec-ai Flag

If the golangci-lint-mcp server was started with the `--gosec-ai` flag and `GOSEC_AI_API_KEY` environment variable, a third MCP tool is available:

**`gosec_ai_autofix`** — runs gosec with AI-powered autofix on a file or directory.

Parameters:
- **`path`** (string, required) — File or directory path to scan (e.g., `"./pkg/..."`, `"main.go"`)

**Important:** Call `gosec_ai_autofix` per package directory (e.g., `"./pkg/auth/..."`), **not** on the entire project (`"./..."`). gosec needs Go package context for type resolution — imports, type definitions, and interface satisfaction are resolved per-package. Calling on the whole project at once may fail or produce incomplete fixes.

This tool wraps gosec's built-in AI autofix internally. The API key is passed to the gosec subprocess server-side — it is never exposed in tool responses or available to the agent. When this tool is available, gosec guide responses include an `<autofix>` section pointing to the `gosec_ai_autofix` tool instead of hardcoded CLI commands.

If the tool is not available (no API key configured), the `<autofix>` section still appears in gosec guide responses, but the tool itself is not registered. If the tool times out (60s limit), fall back to the guide's `<instructions>` and `<examples>` sections for manual remediation.
