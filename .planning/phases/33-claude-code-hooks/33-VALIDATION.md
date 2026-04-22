---
phase: 33
slug: claude-code-hooks
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 33 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — hook script + installer phase |
| **Config file** | none |
| **Quick run command** | `node -c hooks/golangci-lint-post.js` (syntax check) |
| **Full suite command** | `node -c hooks/golangci-lint-post.js && node -c bin/install.js` |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Syntax check with `node -c hooks/golangci-lint-post.js`
- **After every plan wave:** Verify installer still works with `node -c bin/install.js`
- **Before `/gsd-verify-work`:** Hook script syntax valid, installer exits 0
- **Max feedback latency:** ~1 second

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 33-01-T1 | 01 | 1 | CCHOK-01, CCHOK-03 | T-33-01 / T-33-04 | Hook always exits 0; nudge-only approach (never calls MCP binary); try/catch wraps all logic | manual-only | `node -c hooks/golangci-lint-post.js` | ✅ | ⬜ pending |
| 33-01-T2 | 01 | 1 | CCHOK-02 | T-33-02 / T-33-05 | Non-destructive JSON merge for settings.json and .mcp.json; $PATH lookup for binary (no auto-download) | manual-only | `grep -q 'hooks/' package.json` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Hook fires correctly in Claude Code PostToolUse pipeline | CCHOK-01 | Hook behavior depends on Claude Code runtime's PostToolUse event system — cannot simulate outside Claude Code | Run `golangci-lint run --output.json.path stdout ./...` via Bash in Claude Code; verify additionalContext nudge appears after output |
| Hook applies adaptive depth correctly (≤10 Strategy A, >10 Strategy B) | CCHOK-01 | Requires live Claude Code environment with real golangci-lint output | Test with project having ≤10 and >10 unique diagnostics; verify correct strategy nudge |
| Non-destructive merge of settings.json preserves existing hooks | CCHOK-02 | Requires filesystem state with pre-existing hooks config | Run installer on project with existing `.claude/settings.json` containing other hooks; verify no data loss |
| Non-destructive merge of .mcp.json preserves existing servers | CCHOK-02 | Requires filesystem state with pre-existing MCP server config | Run installer on project with existing `.mcp.json` containing other mcpServers; verify no data loss |
| Hook exits 0 on all error paths | CCHOK-03 | All paths verified via code inspection; runtime verification requires Claude Code environment | Feed malformed JSON, empty input, non-Bash tools to hook; verify exit code 0 in each case |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — hook runtime phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: hooks only work within Claude Code runtime; cannot be automated outside it

**Approval:** pending (retroactive validation — phase already completed)
