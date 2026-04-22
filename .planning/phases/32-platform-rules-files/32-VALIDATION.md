---
phase: 32
slug: platform-rules-files
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 32 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — documentation/configuration phase |
| **Config file** | none |
| **Quick run command** | `test -f rules/claude-code.md && test -f rules/cursor.mdc && test -f rules/opencode.md` |
| **Full suite command** | `node bin/install.js` (exits 0) |
| **Estimated runtime** | ~2 seconds |

---

## Sampling Rate

- **After every task commit:** Verify rules file existence with `test -f rules/claude-code.md`
- **After every plan wave:** Run `node bin/install.js` to verify installer still works
- **Before `/gsd-verify-work`:** All 3 rules files exist, installer exits 0
- **Max feedback latency:** ~2 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 32-01-T1 | 01 | 1 | RULE-01, RULE-02, RULE-03 | T-32-02 / T-32-03 | Rules files contain correct MCP tool names and graceful degradation | structural | `test -f rules/claude-code.md && grep -q "golangci_lint_parse" rules/claude-code.md` | ✅ | ⬜ pending |
| 32-01-T2 | 01 | 1 | RULE-01, RULE-02, RULE-03 | T-32-01 | Installer only writes to known safe directories; non-destructive | manual-only | `grep '"rules/"' package.json` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Rules content quality — agents correctly interpret MCP tool nudge | RULE-01, RULE-02, RULE-03 | Semantic quality of rule instructions cannot be validated programmatically | Load rules in Claude Code/Cursor/opencode and run golangci-lint; verify agent calls MCP tools |
| Platform auto-detection correctness | RULE-01, RULE-02 | Installer checks for platform directories — detection logic depends on user's environment | Run `npx golangci-lint-guide` on systems with different platform installs; verify correct detection |
| Cursor `.mdc` format with `alwaysApply: true` loads unconditionally | RULE-02 | Requires running Cursor environment to verify rule loading | Open Cursor on project with installed rules; verify rule auto-activates for Go files |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, structural + manual-only — documentation phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: rules content quality is semantic, installer writes to user dirs

**Approval:** pending (retroactive validation — phase already completed)
