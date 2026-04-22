---
phase: 42
slug: opencode-plugin-filters-golangci-lint-command-too-widely-it-
status: compliant
nyquist_compliant: true
wave_0_complete: true
created: 2026-04-22
---

# Phase 42 — Validation Strategy

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
| 42-01-T1 | 01 | 1 | — | T-42-01 / T-42-02 | First-token extraction + endsWith prevents false positives from argument/substring matching | unit | `node --test plugins/golangci-lint.test.js` | ✅ | ✅ green |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. Test suite in `plugins/golangci-lint.test.js` provides 42 tests covering positive detection (11), negative detection (14), stripOutputFilters regression (8), injectJsonOutputFlag regression (4), and parseDiagnostics regression (5). No new test stubs needed.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Plugin hooks integrate correctly with opencode runtime | — | Plugin behavior depends on opencode tool execution framework — cannot simulate in unit tests | Run various golangci-lint commands in opencode session; verify positive commands get JSON flag injection and negative commands are ignored |

---

## Validation Sign-Off

- [x] All tasks have automated verify
- [x] Sampling continuity: single task with comprehensive automated coverage (42 tests)
- [x] Wave 0 covers all references (test suite covers all exported functions)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [x] `nyquist_compliant: true` — comprehensive test suite with 42 tests via `node --test plugins/golangci-lint.test.js` covering: 11 positive detection cases, 14 negative detection cases, 8 stripOutputFilters regression tests, 4 injectJsonOutputFlag regression tests, 5 parseDiagnostics regression tests

**Approval:** pending (retroactive validation — phase already completed)
