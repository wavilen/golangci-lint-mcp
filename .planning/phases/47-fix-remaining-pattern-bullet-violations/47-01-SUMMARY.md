---
phase: 47-fix-remaining-pattern-bullet-violations
plan: 01
subsystem: guides
tags: [staticcheck, SA2003, pattern-convention, imperative-verb]

# Dependency graph
requires:
  - phase: 38
    provides: "Pattern bullet convention D-01 (imperative verb first word)"
provides:
  - "SA2003.md with all imperative-verb pattern bullets"
  - "Zero non-imperative pattern bullets across all 626 guides"
affects: [guide-quality, pattern-convention-compliance]

# Tech tracking
tech-stack:
  added: []
  patterns: [imperative-verb-first pattern bullets]

key-files:
  created: []
  modified:
    - guides/staticcheck/SA2003.md

key-decisions:
  - "Rewrote code-start bullet to start with 'Move' imperative verb matching Phase 38 convention"

patterns-established: []

requirements-completed: []

# Metrics
duration: 2min
completed: 2026-04-22
---

# Phase 47 Plan 01: Fix Remaining Pattern Bullet Violations Summary

**Rewrote last non-imperative pattern bullet (SA2003.md) — all 626 guides now comply with Phase 38 D-01 convention**

## Performance

- **Duration:** 2 min
- **Started:** 2026-04-22T06:24:57Z
- **Completed:** 2026-04-22T06:27:18Z
- **Tasks:** 1
- **Files modified:** 1

## Accomplishments
- Fixed the single remaining non-imperative pattern bullet across all 626 guides
- SA2003.md now has all 3 pattern bullets starting with imperative verbs (Move, Fix, Place)
- Comprehensive scan confirmed zero code-start and zero gerund-start bullets remain

## Task Commits

1. **task 1: Rewrite SA2003.md non-imperative pattern bullet** - `c41354e` (fix)

## Files Created/Modified
- `guides/staticcheck/SA2003.md` - Line 24 rewritten from code-start to imperative-verb-start bullet

## Decisions Made
- Used "Move" as the imperative verb to match the fix direction (reorder defer/lock placement)

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered
None.

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- Phase 38 pattern convention D-01 is fully satisfied across all 626 guides
- No remaining pattern bullet violations exist

---
*Phase: 47-fix-remaining-pattern-bullet-violations*
*Completed: 2026-04-22*
