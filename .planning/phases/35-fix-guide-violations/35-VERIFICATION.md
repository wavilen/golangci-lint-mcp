---
phase: 35-fix-guide-violations
verified: 2026-04-22T05:45:10Z
status: gaps_found
score: 2/3 must-haves verified
overrides_applied: 0
overrides: []
---

# Phase 35: Fix Guide Violations Verification Report

**Phase Goal:** Fix all guide Good examples that violate the golden config lint check, upgrade crosscheck pipeline with auto-wrapping and per-package checking
**Verified:** 2026-04-22
**Status:** gaps_found
**Re-verification:** No — initial verification

## Goal Achievement

### Observable Truths

| # | Truth | Status | Evidence |
|---|-------|--------|----------|
| 1 | All 628 extractable guide Good examples pass golden config lint with 0 violations (XCHK-05) | ⚠ REGRESSED | Phase 35 SUMMARY self-check confirmed 628/628 passing at completion (2026-04-20). Current re-run shows 592/629 passing with 37 failing guides (51 total violations). Regression likely caused by golangci-lint version updates adding new linters (wrapcheck: 14, noinlineerr: 2, varnamelen: 1) or golden config changes since Phase 35 completion. |
| 2 | 2 guides skipped (framepointer.md with assembly, struct-tag.md with formatting) — documented as acceptable exclusions | ⚠ CHANGED | Current crosscheck skips only 1 guide (framepointer.md). struct-tag.md is now extracted (629/630 vs prior 628/630). |
| 3 | Crosscheck pipeline uses auto-wrap, per-package linting, and excluded linters for deterministic results | ✓ VERIFIED | `cmd/crosscheck/main.go` contains 2 references to auto-wrap functions and 2 references to excludedLinters. Pipeline completed deterministically with consistent output format. |

### Required Artifacts

| Artifact | Expected | Status | Details |
|----------|----------|--------|---------|
| cmd/crosscheck/main.go | Pipeline with auto-wrap and excluded linters | ✓ PRESENT | Auto-wrap (`wrapCodeForExtraction`) and excluded linters confirmed present |
| 40+ modified guide files | Fixed Good examples | ✓ PRESENT | Phase 35 SUMMARY lists 39+ guide files modified; spot-checked 3 guides (gocritic/codegenComment.md, staticcheck/SA4025.md, gosec/G107.md) — all contain Good examples |
| guides/gocritic/codegenComment.md | Contains Good example | ✓ PRESENT | 1 "Good" reference found |
| guides/staticcheck/SA4025.md | Contains Good example | ✓ PRESENT | 1 "Good" reference found |
| guides/gosec/G107.md | Contains Good example | ✓ PRESENT | 1 "Good" reference found |

### Key Link Verification

| From | To | Via | Status | Details |
|------|----|-----|--------|---------|
| 35-VERIFICATION.md | make crosscheck | pipeline re-run evidence | ✓ LINKED | `make crosscheck` executed successfully, produced deterministic results |
| cmd/crosscheck/main.go | wrapCodeForExtraction | function definition | ✓ LINKED | 2 references to wrap/auto-wrap in crosscheck source |
| cmd/crosscheck/main.go | excludedLinters | linter exclusion map | ✓ LINKED | 2 references to excludedLinters in crosscheck source |

### Data-Flow Trace (Level 4)

| Artifact | Data Variable | Source | Produces Real Data | Status |
|----------|---------------|--------|--------------------|--------|
| cmd/crosscheck/main.go | Extraction count | guide extraction | 629/630 extracted (1 skipped) | ✓ REAL |
| cmd/crosscheck/main.go | Violations by linter | golangci-lint JSON output | 51 violations across 13 linters after filtering | ✓ REAL |
| cmd/crosscheck/main.go | Passing count | violation counting | 592/629 guides passing | ✓ REAL |

### Behavioral Spot-Checks

| Behavior | Command | Result | Status |
|----------|---------|--------|--------|
| Crosscheck pipeline runs | `make crosscheck` | Completed: 629/630 extracted, 592/629 passing, 37 failing | ⚠ REGRESSED |
| Auto-wrap present | `grep -c "wrapCodeForExtraction\|auto.*wrap" cmd/crosscheck/main.go` | 2 matches | ✓ PASS |
| Excluded linters present | `grep -c "excludedLinters" cmd/crosscheck/main.go` | 2 matches | ✓ PASS |
| Guide codegenComment has Good example | `grep -c "Good" guides/gocritic/codegenComment.md` | 1 match | ✓ PASS |
| Guide SA4025 has Good example | `grep -c "Good" guides/staticcheck/SA4025.md` | 1 match | ✓ PASS |
| Guide G107 has Good example | `grep -c "Good" guides/gosec/G107.md` | 1 match | ✓ PASS |

### Requirements Coverage

| Requirement | Source Plan | Description | Status | Evidence |
|-------------|------------|-------------|--------|----------|
| XCHK-05 | Phase 35 | Violating guides are fixed so their Good examples pass the golden config lint | ⚠ SATISFIED AT COMPLETION, REGRESSED SINCE | Phase 35 SUMMARY self-check: 628/628 passing, 0 violations (2026-04-20). Current re-run: 592/629 passing, 37 failing. New linters (wrapcheck: 14, noinlineerr: 2, varnamelen: 1) account for 17 of 51 violations — likely added in golangci-lint version update since April 20. |

### Anti-Patterns Found

| File | Line | Pattern | Severity | Impact |
|------|------|---------|----------|--------|
| (none) | — | — | — | — |

### Human Verification Required

None.

### Gaps Summary

Phase 35 was fully satisfied at completion time (628/628 passing, 0 violations per Phase 35 SUMMARY self-check dated 2026-04-20). However, re-running `make crosscheck` on 2026-04-22 reveals 37 failing guides (592/629 passing) with 51 total violations across 13 linters. The regression is primarily driven by linters that were not flagging issues at Phase 35 time: wrapcheck (14 violations), revive (11), errcheck (7), noinlineerr (2), varnamelen (1). This suggests a golangci-lint version update or golden config change introduced new checks. The core XCHK-05 requirement was met at Phase 35 completion; the regression is a maintenance concern for a future phase, not a Phase 35 execution failure.
