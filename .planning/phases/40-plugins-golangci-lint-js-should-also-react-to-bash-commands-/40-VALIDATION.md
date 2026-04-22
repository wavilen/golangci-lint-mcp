---
phase: 40
slug: plugins-golangci-lint-js-should-also-react-to-bash-commands-
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 40 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | node:test (indirect — Phase 42) |
| **Config file** | none |
| **Quick run command** | `node --test plugins/golangci-lint.test.js` |
| **Full suite command** | `node --test plugins/golangci-lint.test.js` |
| **Estimated runtime** | ~2 seconds |

---

## Sampling Rate

- **After every task commit:** Verify `stripOutputFilters` function exists in plugin file
- **After every plan wave:** Run `node --test plugins/golangci-lint.test.js` (Phase 42 tests include stripOutputFilters coverage)
- **Before `/gsd-verify-work`:** Phase 42 test suite green
- **Max feedback latency:** ~2 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 40-01-T1 | 01 | 1 | — | T-40-01 / T-40-02 | Only strips pipe segments and redirects AFTER golangci-lint token — never modifies the golangci-lint command itself | manual-only, indirect | `node --test plugins/golangci-lint.test.js` (42 tests from Phase 42 include stripOutputFilters coverage) | ✅ | ✅ green (indirect) |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. Phase 42's test suite (`plugins/golangci-lint.test.js`) includes 8 stripOutputFilters regression tests that cover the filter stripping logic added in Phase 40. No new test stubs needed.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Plugin behavior in opencode runtime — stripOutputFilters integrates correctly with before-hook pipeline | — | Plugin hooks depend on opencode runtime — cannot test hook execution in isolation | Run golangci-lint via `opencode` with pipe filters (e.g., `golangci-lint run ./... \| head -20`); verify filters stripped before execution |
| Filter stripping doesn't over-strip valid golangci-lint commands | — | Edge cases like quoted pipes or complex shell syntax may behave differently in runtime | Test with various pipe/filter combinations in opencode session |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: single task with indirect Phase 42 test coverage
- [x] Wave 0 covers all references (Phase 42 test suite covers stripOutputFilters)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: filter stripping function tested indirectly through Phase 42's test suite (8 stripOutputFilters tests), but no direct Phase 40-specific test suite; plugin hook integration depends on opencode runtime

**Approval:** pending (retroactive validation — phase already completed; indirect coverage via Phase 42 tests)
