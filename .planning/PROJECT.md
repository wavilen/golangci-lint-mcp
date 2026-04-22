# Project: golangci-lint-mcp

## What This Is

A single-binary MCP server (stdio transport) that lets AI agents look up concise fix guidance for any golangci-lint diagnostic. Ships 630 guides covering all golangci-lint linters, two MCP tools (single lookup + batch JSON parse), gosec AI autofix integration, and npm/npx packaging.

## Core Value

When an agent encounters a golangci-lint diagnostic, it can call a single tool and immediately understand what the issue means and how to fix it.

## Current Milestone: v1.1 Platform Hooks + Example Validation

**Goal:** Make golangci-lint-mcp auto-activate in every major AI coding platform, and validate that all 630 guide examples are lint-clean.

**Target features:**
- Command hooks (opencode, Claude Code, Cursor) that intercept `golangci-lint` calls and auto-inject MCP guidance
- Skill-first integration — hooks route through the opencode skill, not raw MCP tool calls
- Golden config cross-check — fetch maratori's golden config, enable all recommended linters, lint all 630 good examples, fix violations

**Priority:** Hooks first, then cross-check.

## Current State

**Shipped:** v1.0 (2026-04-20)
**Codebase:** 3,139 LOC Go across 24 files + 630 markdown guides
**Tech stack:** Go 1.23, MCP SDK (mark3labs/mcp-go), embed.FS, stdio transport
**Quality:** Desloppify score 92.8/100, all tests pass
**Install:** `go install github.com/wavilen/golangci-lint-mcp@latest` or `npx golangci-lint-guide`
**Phase 31 complete** — `/golangci-lint` custom command for opencode/crush with adaptive MCP guidance and npm installer deployment
**Phase 32 complete** — Platform rules files for Claude Code, Cursor, opencode with npm installer auto-detection
**Phase 33 complete** — Claude Code PostToolUse hook + npm installer extension for hooks/settings/MCP wiring
**Phase 37 complete** — Audited and improved SKILL.md + 6 rules files with LLM prompting best practices (11 techniques applied)
**Phase 39 complete** — Added diagnostic summary block to golangci_lint_parse response with strategy recommendation; simplified Step 4 in commands/golangci-lint.md to use summary instead of manual counting
**Phase 41 complete** — Extended npx installer to configure opencode plugin, MCP server config, and binary prerequisite checks
**Phase 43 complete** — Updated all 5 documentation files (SKILL.md, commands, 3 rules) with per-package golangci-lint run emphasis and jq/python large-output workflow
**Phase 46 complete** — Verification & traceability closure: verified phases 31.1, 35, 36; marked 31.2 superseded; closed XCHK-05

## Requirements

### Validated

- ✓ MCP server registers tools over stdio transport — v1.0
- ✓ XML-tagged guide format with instructions/examples/patterns/related sections — v1.0
- ✓ 630 linter guides covering all golangci-lint linters — v1.0
- ✓ Automated validation enforcing 200-word (simple) / 500-word (compound) limits — v1.0
- ✓ Cross-referenced related guides for navigability — v1.0
- ✓ Batch JSON parse tool (golangci_lint_parse) for multi-diagnostic workflows — v1.0
- ✓ Gosec AI autofix with env-var-driven templates — v1.0
- ✓ npm/npx packaging and opencode skill — v1.0
- ✓ `/golangci-lint` custom command for opencode/crush with adaptive MCP guidance — Phase 31
- ✓ Command hooks for Claude Code that auto-inject MCP guidance on golangci-lint calls — Phase 33

### Active

- [ ] Command hooks for Cursor that auto-inject MCP guidance on golangci-lint calls
- ✓ Golden config cross-check: validate all 630 good examples against maratori's golden config — Phase 35
- ✓ Fix any guide examples that violate linter rules — Phase 35
- ✓ Golden config must enable all linters between `## you may want to enable` and `## disabled` — enforced by `update-golden-config` Makefile target via awk post-processing (added Phase 48)

### Out of Scope

- Linter configuration generation — agents handle this natively
- Code auto-fixing — golangci-lint handles this; MCP provides guidance only
- CI/CD integration — agents call MCP directly in their workflows
- HTTP/SSE transport — deferred to future milestone
- Multi-language guide support — deferred to future milestone
- Windsurf, Zed, or other minor platform hooks — defer until requested

## Context

Shipped v1.0 with 630 guides, 3,139 LOC Go, two MCP tools.
Module: github.com/wavilen/golangci-lint-mcp
Binary: golangci-lint-mcp
Desloppify: 92.8/100

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| XML-tagged guide format | Structured parsing by AI agents | ✓ Good |
| 200/500 word limits | Concise fix guidance without overwhelming context | ✓ Good |
| embed.FS for guides | Single binary, no external files | ✓ Good |
| Two MCP tools | Single lookup + batch parse for different workflows | ✓ Good |
| Environment-variable-driven gosec autofix | Per-client config without hardcoded secrets | ✓ Good |
| npm/npx packaging | Easy skill installation for opencode users | ✓ Good |
| Module path github.com/wavilen/golangci-lint-mcp | Go install compatibility | ✓ Good |

## Constraints

- Go 1.23+ required
- MCP stdio transport (no HTTP)
- Guides embedded in binary (rebuild to update)
- Word limits enforced by automated tests
- Hooks must be non-breaking — if MCP server is unavailable, golangci-lint runs normally

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `/gsd-transition`):
1. Requirements invalidated? → Move to Out of Scope with reason
2. Requirements validated? → Move to Validated with phase reference
3. New requirements emerged? → Add to Active
4. Decisions to log? → Add to Key Decisions
5. "What This Is" still accurate? → Update if drifted

**After each milestone** (via `/gsd-complete-milestone`):
1. Full review of all sections
2. Core Value check — still the right priority?
3. Audit Out of Scope — reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-04-22 after Phase 46 completion*
