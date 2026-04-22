---
phase: 34
slug: golden-config-pipeline
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 34 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | Go testing + make crosscheck |
| **Config file** | none |
| **Quick run command** | `test -f golden-config/.golangci.yml` |
| **Full suite command** | `make crosscheck` |
| **Estimated runtime** | ~60 seconds |

---

## Sampling Rate

- **After every task commit:** Verify config file existence with `test -f golden-config/.golangci.yml`
- **After every plan wave:** Run `make crosscheck` for full pipeline validation
- **Before `/gsd-verify-work`:** `make crosscheck` must complete with informational report
- **Max feedback latency:** ~60 seconds (golangci-lint across 629 packages)

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 34-01-T1 | 01 | 1 | XCHK-01 | — | Vendored config pinned to specific tag; no remote fetch at runtime | structural | `test -f golden-config/.golangci.yml && grep -q "linters:" golden-config/.golangci.yml` | ✅ | ⬜ pending |
| 34-02-T1 | 02 | 2 | XCHK-02, XCHK-03 | — | Extraction pipeline reads trusted local content only | integration | `make crosscheck` | ✅ | ⬜ pending |
| 34-02-T2 | 02 | 2 | XCHK-04 | — | Pipeline exits 0 regardless of violations (informational only) | integration | `make crosscheck && python3 -c "import json; json.load(open('tmp/crosscheck/violations.json'))"` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. `make crosscheck` provides integration-level coverage for the pipeline. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Golden config has correct modifications for example linting | XCHK-01 | Config modifications (local-prefixes cleared, depguard commented out) are correctness judgments | Review `golden-config/.golangci.yml` for empty `local-prefixes: []` and commented `"non-main files"` rule |
| `make update-golden-config` produces correct result on fresh download | XCHK-01 | Known sed pattern bug in local-prefixes clearing (`github/my/project` vs `github.com/my/project`) | Run `make update-golden-config` and verify local-prefixes is cleared correctly in downloaded config |
| Violation report accurately maps back to original guides | XCHK-04 | Mapping correctness requires manual spot-check of JSON report fields | Run `make crosscheck`, review `tmp/crosscheck/violations.json`, verify guide paths match actual guide file names |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: no 3 consecutive tasks without automated verify
- [x] Wave 0 covers all references (none needed — `make crosscheck` covers integration)
- [x] No watch-mode flags
- [x] Feedback latency < 120s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: golden config vendoring task (34-01-T1) has only structural check, not integration; also `make update-golden-config` has known sed bug requiring manual verification

**Approval:** pending (retroactive validation — phase already completed)
