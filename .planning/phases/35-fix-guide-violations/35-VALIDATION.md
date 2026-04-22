---
phase: 35
slug: fix-guide-violations
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 35 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | make crosscheck (integration pipeline) |
| **Config file** | none |
| **Quick run command** | `make crosscheck` |
| **Full suite command** | `make crosscheck` |
| **Estimated runtime** | ~60 seconds |

---

## Sampling Rate

- **After every task commit:** Run `make crosscheck` to verify guide fix correctness
- **After every plan wave:** Run `make crosscheck` for full pipeline validation
- **Before `/gsd-verify-work`:** `make crosscheck` must report 0 violations
- **Max feedback latency:** ~60 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 35-01-T1 | 01 | 1 | XCHK-05 | T-35-01 | Guide files version-controlled; changes reviewed via git diff | integration | `make crosscheck` | ✅ | ⬜ pending |
| 35-01-T2 | 01 | 1 | XCHK-05 | T-35-02 | Golden config pinned and unchanged from Phase 34 | integration | `python3 -c "import json; r=json.load(open('tmp/crosscheck/violations.json')); assert r['failing']==0"` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. `make crosscheck` validates all 628 extractable guides. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Guide fix quality — fixed Good examples remain educational | XCHK-05 | Semantic correctness of educational code is a quality judgment | Review each fixed guide's Good example for clarity and correctness of the demonstrated pattern |
| Auto-wrap in pipeline doesn't mask real issues | XCHK-05 | Auto-wrapping bare statements in `func _() {}` could hide structural issues | Review pipeline's `wrapCodeForExtraction()` logic and verify excluded linter list is justified |
| 628/630 passing is acceptable (2 skipped: framepointer.md asm, struct-tag.md formatting) | XCHK-05 | Decision to skip 2 guides is a judgment call | Verify that both skipped guides are truly non-extractable (assembly code, formatting issues) |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: both tasks have `make crosscheck` integration coverage
- [x] Wave 0 covers all references (none needed — `make crosscheck` is the integration test)
- [x] No watch-mode flags
- [x] Feedback latency < 120s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: guide fix quality is semantic (educational code must remain correct demonstrations, not just lint-clean); auto-wrap logic correctness is a judgment call; excluded linter list requires justification

**Approval:** pending (retroactive validation — phase already completed)
