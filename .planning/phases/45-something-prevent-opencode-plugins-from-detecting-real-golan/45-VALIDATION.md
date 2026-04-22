---
phase: 45
slug: something-prevent-opencode-plugins-from-detecting-real-golan
status: compliant
nyquist_compliant: true
wave_0_complete: true
created: 2026-04-22
---

# Phase 45 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | node:test (built-in) |
| **Config file** | none |
| **Quick run command** | `node --test plugins/golangci-lint.test.js` |
| **Full suite command** | `node --test plugins/golangci-lint.test.js` |
| **Estimated runtime** | ~2 seconds |

---

## Sampling Rate

- **After every task commit:** Run `node --test plugins/golangci-lint.test.js`
- **After every plan wave:** Run `node --test plugins/golangci-lint.test.js`
- **Before `/gsd-verify-work`:** Full suite green
- **Max feedback latency:** ~2 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 45-01-T1 | 01 | 1 | — | T-45-01 / T-45-02 | Shell-wrapper regex prevents arbitrary command injection | unit | `node --test plugins/golangci-lint.test.js` | ✅ | ✅ green |
| 45-01-T2 | 01 | 1 | — | — | Makefile targets exist and plugin installs correctly | structural | `grep -E "install-plugin|verify-plugin" Makefile` | ✅ | ✅ green |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. Test suite in `plugins/golangci-lint.test.js` provides 76+ tests covering shell-wrapper detection (7 new tests), plugin export format, and existing functionality. No new test stubs needed.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Plugin hooks integrate correctly with opencode runtime | — | Plugin behavior depends on opencode tool execution framework — cannot simulate in unit tests | Run various golangci-lint commands in opencode session; verify positive commands get JSON flag injection and negative commands are ignored |

---

## Validation Sign-Off

- [x] All tasks have automated verify
- [x] Sampling continuity: both tasks have automated coverage (unit + structural)
- [x] Wave 0 covers all references (test suite covers all exported functions)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [x] `nyquist_compliant: true` — comprehensive test suite with 76+ tests via `node --test plugins/golangci-lint.test.js` covering: shell-wrapper detection (7 new tests), plugin export format, Makefile targets

**Approval:** pending (retroactive validation — phase already completed)
