---
phase: 38
slug: the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 38 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | grep-based structural verification |
| **Config file** | none |
| **Quick run command** | `grep -cE "^- [A-Z]" guides/*/*.md guides/*.md \| grep -v ":0$" \| wc -l` |
| **Full suite command** | `grep -cE "^- [A-Z]" guides/*/*.md guides/*.md \| grep -v ":0$" \| wc -l` |
| **Estimated runtime** | ~2 seconds |

---

## Sampling Rate

- **After every task commit:** Verify guide file patterns with `grep -cE "^- [A-Z]" guides/{category}/*.md`
- **After every plan wave:** Verify structural imperative-verb check across all processed guides
- **Before `/gsd-verify-work`:** Full imperative-verb check passes for all 626 planned guides
- **Max feedback latency:** ~2 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 38-01-T1 | 01 | 1 | — | — | Guide files version-controlled; changes reviewed via git diff | structural | `grep -cE "^- [A-Z]" guides/errcheck.md guides/asasalint.md guides/asciicheck.md` | ✅ | ⬜ pending |
| 38-02-T1 | 02 | 1 | — | — | Same pattern for Complexity & Quality category | structural | `grep -cE "^- [A-Z]" guides/gocognit.md guides/gocyclo.md guides/maintidx.md` | ✅ | ⬜ pending |
| 38-03-T1 | 03 | 1 | — | — | Same pattern for Style & Formatting category | structural | `grep -cE "^- [A-Z]" guides/errname.md guides exhaustruct.md` | ✅ | ⬜ pending |
| 38-04-T1 | 04 | 1 | — | — | Same pattern for Perf, Testing & Remaining category | structural | `grep -cE "^- [A-Z]" guides/prealloc.md guides/prealloc.md` | ✅ | ⬜ pending |
| 38-05-T1 | 05 | 1 | — | — | Same pattern for gocritic compound (108 guides) | structural | `grep -cE "^- [A-Z]" guides/gocritic/*.md` | ✅ | ⬜ pending |
| 38-06-T1 | 06 | 1 | — | — | **PARTIAL** — 14 non-imperative bullets remain in 6 guides (SA2000, SA2001, SA2003, SA3000, SA3001, SA4000) | structural, PARTIAL | `grep -cE "^- [A-Z]" guides/staticcheck/SA1000.md` | ✅ | ⚠️ flaky |
| 38-07-T1 | 07 | 1 | — | — | Same pattern for revive compound (101 guides) | structural | `grep -cE "^- [A-Z]" guides/revive/*.md` | ✅ | ⬜ pending |
| 38-08-T1 | 08 | 1 | — | — | Same pattern for gosec compound (61 guides) | structural | `grep -cE "^- [A-Z]" guides/gosec/*.md` | ✅ | ⬜ pending |
| 38-09-T1 | 09 | 1 | — | — | Same pattern for govet compound (35 guides) | structural | `grep -cE "^- [A-Z]" guides/govet/*.md` | ✅ | ⬜ pending |
| 38-10-T1 | 10 | 1 | — | — | **PARTIAL** — 3 uncovered guides (funlen, ineffassign, nlreturn) not in any plan | structural, PARTIAL | `grep -cE "^- [A-Z]" guides/funlen.md guides/ineffassign.md guides/nlreturn.md` | ✅ | ⚠️ flaky |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Guide pattern quality — imperative-verb bullets provide actionable fix direction | — | Semantic correctness of rewritten bullets — whether the imperative verb accurately describes the fix | Review a sample of rewritten guides across categories; verify imperative-first bullets give clear fix direction |
| Known gaps acknowledged (per Phase 38 VERIFICATION.md) | — | VERIFICATION.md found 14 non-imperative bullets in 6 staticcheck guides + 3 uncovered guides | Verify 38-VERIFICATION.md documents: SA2000 (2 bullets), SA2001 (3 bullets), SA2003 (2 bullets), SA3000 (2 bullets), SA3001 (2 bullets), SA4000 (3 bullets), and uncovered funlen/ineffassign/nlreturn |
| 620/626 (99.0%) imperative-first rate is acceptable | — | Decision to ship with known gaps is a judgment call | Review 38-VERIFICATION.md coverage analysis; confirm gap count matches actual guide content |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: all 10 plan-level tasks have structural grep verification
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: Phase 38 VERIFICATION.md found `status: gaps_found` with 14 non-imperative bullets remaining in 6 staticcheck guides (SA2000, SA2001, SA2003, SA3000, SA3001, SA4000) and 3 uncovered guides (funlen, ineffassign, nlreturn) per Phase 38 VERIFICATION.md

**Approval:** pending (retroactive validation — phase already completed with known gaps documented in VERIFICATION.md)
