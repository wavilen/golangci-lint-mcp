# Requirements: golangci-lint-mcp

**Defined:** 2026-04-20
**Core Value:** When an agent encounters a golangci-lint diagnostic, it can call a single tool and immediately understand what the issue means and how to fix it.

## v1.1 Requirements

### opencode Plugin

- [x] **OPLG-01**: opencode plugin intercepts Bash tool calls to golangci-lint via `tool.execute.before` and ensures `--output.json.path stdout` flag is present (fixes models calling without JSON output)
- [x] **OPLG-02**: opencode plugin intercepts Bash tool calls to golangci-lint via `tool.execute.after`, parses JSON diagnostics from stdout, and injects MCP fix guidance into agent context
- [x] **OPLG-03**: opencode plugin is installed automatically via the existing npm package (npx golangci-lint-guide) or placed in `.opencode/plugins/`
- [x] **OPLG-04**: Plugin gracefully degrades when MCP server is unavailable — golangci-lint runs normally without guidance injection

### opencode Command

- [x] **OCMD-01**: opencode `/golangci-lint` command runs golangci-lint with correct flags and feeds output to MCP tools for guided fix assistance (named to avoid overlap with generic `/lint` commands from other linters)

### Claude Code Integration

- [x] **CCHOK-01**: Claude Code `PostToolUse` hook fires after any `golangci-lint` Bash call, parses JSON diagnostics, and injects MCP fix guidance
- [x] **CCHOK-02**: Claude Code hook config is installed via `.claude/settings.json` with proper `mcpServers` block wiring the MCP server
- [x] **CCHOK-03**: Claude Code hook gracefully degrades — exits 0 always (non-blocking), never prevents golangci-lint from running

### Platform Rules

- [x] **RULE-01**: Claude Code rules file (`.claude/rules/golangci-lint.md`) instructs agent to use MCP tools when running golangci-lint
- [x] **RULE-02**: Cursor rules file (`.cursor/rules/golangci-lint.mdc`) with `alwaysApply: true` instructs agent to use MCP tools when encountering golangci-lint diagnostics
- [x] **RULE-03**: opencode rules file (`.opencode/rules/golangci-lint.md`) provides fallback guidance when plugin is not active

### Golden Config Cross-Check

- [x] **XCHK-01**: Pipeline fetches maratori's golden config from `github.com/maratori/golangci-lint-config` (pin to specific tag) with all "you may want to enable" linters activated
- [x] **XCHK-02**: Pipeline extracts "Good" code blocks from all 630 guide XML-tagged `<examples>` sections, skipping "Bad" examples
- [x] **XCHK-03**: Extracted examples are written to a tmp gitignored subdirectory with minimal Go file wrapping (package declaration, import inference)
- [x] **XCHK-04**: Pipeline runs golangci-lint with golden config on all extracted examples and reports violations with guide filename and linter details
- [x] **XCHK-05**: Violating guides are fixed so their Good examples pass the golden config lint

## v2 Requirements

### Platform Enhancements

- **CCHOK-04**: Claude Code `PreToolUse` hook intercepts golangci-lint calls and augments with MCP output before Claude sees results
- **XCHK-06**: GitHub Actions workflow runs cross-check on every PR touching guides
- **XCHK-07**: README badge showing golden config pass rate (e.g. "630/630 guides pass")
- **OPLG-05**: Cross-platform install script (`npx golangci-lint-guide --setup-hooks`) that configures all platforms at once

### Future Features

- **TRANS-01**: HTTP/SSE transport in addition to stdio
- **MULTI-01**: Multi-language guide support (non-English)
- **WIND-01**: Windsurf rules file integration
- **ZED-01**: Zed editor rules file integration

## Out of Scope

| Feature | Reason |
|---------|--------|
| Blocking hooks that prevent golangci-lint from running | Violates non-breaking constraint; agents should be nudged, not blocked |
| Auto-fixing code in hooks | Risk of breaking active edits; agent should decide whether to fix |
| Persistent MCP server daemon | Against single-binary, zero-config philosophy |
| Bad example validation | Bad examples intentionally trigger linters; validation signal is noisy |
| Full compilation of guide examples | Examples are snippets; wrapping with imports is sufficient for linting |
| Linter configuration generation | Agents handle this natively |
| CI/CD integration beyond cross-check | Agents call MCP directly in their workflows |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| OPLG-01 | Phase 31 | Satisfied |
| OPLG-02 | Phase 31 | Satisfied |
| OPLG-03 | Phase 31 | Satisfied |
| OPLG-04 | Phase 31 | Satisfied |
| OCMD-01 | Phase 31 | Satisfied |
| CCHOK-01 | Phase 33/51 | Satisfied |
| CCHOK-02 | Phase 33 | Satisfied |
| CCHOK-03 | Phase 33 | Satisfied |
| RULE-01 | Phase 32/51 | Satisfied |
| RULE-02 | Phase 32/51 | Satisfied |
| RULE-03 | Phase 32/51 | Satisfied |
| XCHK-01 | Phase 34 | Satisfied |
| XCHK-02 | Phase 34 | Satisfied |
| XCHK-03 | Phase 34 | Satisfied |
| XCHK-04 | Phase 34 | Satisfied |
| XCHK-05 | Phase 46 | Satisfied |

**Coverage:**
- v1.1 requirements: 16 total
- Satisfied: 16
- Mapped to phases: 16
- Unmapped: 0 ✓

---
*Requirements defined: 2026-04-20*
*Last updated: 2026-04-22 after gap closure planning (phases 51-54)*
