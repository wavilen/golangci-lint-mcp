---
phase: 48
slug: fix-golden-config-sed-bug
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 48 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | grep + make |
| **Config file** | none |
| **Quick run command** | `grep '-\.\*/' Makefile && ! grep 'github.com/my/project' golden-config/.golangci.yml && grep 'local-prefixes: \[\]' golden-config/.golangci.yml` |
| **Full suite command** | `grep '-\.\*/' Makefile && ! grep 'github.com/my/project' golden-config/.golangci.yml && grep 'local-prefixes: \[\]' golden-config/.golangci.yml && make crosscheck` |
| **Estimated runtime** | ~10 seconds |

---

## Sampling Rate

- **After every task commit:** Run grep structural checks
- **After every plan wave:** Run full verification including make crosscheck
- **Before `/gsd-verify-work`:** All checks green
- **Max feedback latency:** ~10 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 48-01-T1 | 01 | 1 | — | — | Wildcard sed pattern prevents brittle literal matching | structural + integration | `grep '-\.\*/' Makefile && ! grep 'github.com/my/project' golden-config/.golangci.yml && grep 'local-prefixes: \[\]' golden-config/.golangci.yml` | ✅ | ✅ green |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. Grep provides structural verification. Make crosscheck provides integration testing. No new test stubs needed.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Sed pattern correctly handles various upstream placeholder formats | — | Cannot anticipate all future upstream placeholder variations | Monitor golden config updates; if sed pattern fails, verify it's due to new upstream format, then update pattern |

---

## Validation Sign-Off

- [x] All tasks have automated verify
- [x] Sampling continuity: single task with comprehensive automated coverage (structural + integration)
- [x] Wave 0 covers all references (grep + make)
- [x] No watch-mode flags
- [x] Feedback latency < 10s
- [ ] `nyquist_compliant: true` — **not set**: structural changes with upstream format dependency

**Approval:** pending (retroactive validation — phase already completed)
