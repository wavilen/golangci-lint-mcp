---
phase: 41
slug: npx-golangci-lint-guide-should-install-opencode-plugin-skill
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 41 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — installer writes to user filesystem |
| **Config file** | none |
| **Quick run command** | `grep -q "plugins/" package.json` |
| **Full suite command** | N/A |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Verify structural patterns in `bin/install.js` with grep
- **After every plan wave:** Verify `package.json` files array includes `plugins/`
- **Before `/gsd-verify-work`:** Structural checks pass + installer smoke test passes
- **Max feedback latency:** ~1 second

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 41-01-T1 | 01 | 1 | OPLG-03 | T-41-01 / T-41-02 / T-41-04 | Installer creates timestamped .bak backup before modifying user-level config | manual-only, structural | `grep -q "plugins/" package.json && grep -q "opencode.json" bin/install.js` | ✅ | ⬜ pending |
| 41-01-T2 | 01 | 1 | OPLG-03 | T-41-03 | Binary checks produce informational warnings only — never block installer | manual-only, structural | `grep -q "plugins/" package.json` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Installer copies plugin correctly to project and user scope directories | OPLG-03 | Installer modifies user system (`~/.config/`) — cannot safely test in automation without side effects | Run `npx golangci-lint-guide --platforms=opencode --plugin-scope=project` in clean project; verify `.opencode/plugins/golangci-lint.js` exists and matches source |
| MCP config merge is non-destructive — preserves existing entries | — | Config merge behavior with complex existing configs is hard to test exhaustively | Run installer with existing `opencode.json` containing other MCP servers; verify all entries preserved |
| Binary checks produce informational warnings only | — | Exit behavior when binaries missing needs runtime verification | Run installer on system without `golangci-lint-mcp` on PATH; verify warning printed but installer completes successfully |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — installer phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: installer writes to `~/.config/` directories, cannot safely automate; MCP config merge is destructive operation requiring manual verification; binary PATH checks are system-dependent

**Approval:** pending (retroactive validation — phase already completed)
