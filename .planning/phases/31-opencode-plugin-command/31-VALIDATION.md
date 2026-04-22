---
phase: 31
slug: opencode-plugin-command
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 31 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — documentation-only phase |
| **Config file** | none |
| **Quick run command** | `grep -c "golangci_lint_parse" commands/golangci-lint.md` |
| **Full suite command** | N/A |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Verify file existence with `test -f commands/golangci-lint.md`
- **After every plan wave:** Verify package.json files array with `grep '"commands/"' package.json`
- **Before `/gsd-verify-work`:** Structural checks pass
- **Max feedback latency:** ~1 second

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 31-01-T1 | 01 | 1 | OCMD-01, OPLG-01, OPLG-02, OPLG-04 | T-31-01 / T-31-03 | Command template includes `--output.json.path stdout` and graceful MCP degradation | manual-only | — | ✅ | ⬜ pending |
| 31-01-T2 | 01 | 1 | OPLG-03 | T-31-02 | Installer writes to user ~/.config — no privilege escalation | manual-only | `grep '"commands/"' package.json` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Agent interprets `/golangci-lint` command template correctly and follows the 7-step workflow | OCMD-01 | Command template is a prompt interpreted by AI agent — behavior depends on agent's prompt comprehension | Run `/golangci-lint` in opencode/crush session on a Go project with known lint issues; verify agent runs golangci-lint with correct flags and calls `golangci_lint_parse` |
| Agent applies adaptive depth correctly (≤10 inline, >10 TODO splitting) | OPLG-02 | Agent's interpretation of threshold instruction and presentation strategy cannot be verified without a running session | Run `/golangci-lint` on a project with >10 unique diagnostics; verify agent recommends splitting into TODO list |
| npm installer copies command to correct user config directories | OPLG-03 | Installer writes to user's `~/.config/` — cannot safely test in automation without side effects | Run `npx golangci-lint-guide` in clean environment; verify files at `~/.config/opencode/commands/golangci-lint.md` and `~/.config/crush/commands/golangci-lint.md` |
| Agent handles MCP server unavailability gracefully | OPLG-04 | Requires running session without MCP server configured | Run `/golangci-lint` without MCP server; verify golangci-lint output shown normally with actionable warning |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — documentation phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: command template is a prompt, installer writes to user dirs; neither has automated test coverage

**Approval:** pending (retroactive validation — phase already completed)
