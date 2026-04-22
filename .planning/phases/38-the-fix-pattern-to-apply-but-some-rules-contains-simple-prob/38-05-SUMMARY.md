---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 05
subsystem: guides
tags: [gocritic, patterns, imperative-verbs, content-editing]

# Dependency graph
requires:
  - phase: 38
    provides: "D-01 imperative-verb rule and D-05 scope constraints"
provides:
  - "105 gocritic guides with imperative-first <patterns> bullets"
  - "Zero-failure automated verification across all gocritic guides"
affects: [38-verification]

# Tech tracking
tech-stack:
  added: []
  patterns: ["Imperative-first pattern bullets: every bullet starts with an action verb (Replace, Remove, Use, Avoid, etc.)"]

key-files:
  created: []
  modified:
    - "guides/gocritic/*.md (105 files — all gocritic guides with <patterns> sections)"

key-decisions:
  - "Used Replace/Remove/Avoid/Use as primary imperative verbs for gocritic patterns"
  - "Kept transformation arrows (→) as context within bullets rather than removing them"
  - "Ensured every bullet gives actionable fix direction, not just problem description"

patterns-established:
  - "Imperative-verb-first pattern bullets across all 105 gocritic guides"

requirements-completed: []

# Metrics
duration: 57min
completed: 2026-04-21
---

# Phase 38 Plan 05: Gocritic Pattern Bullets Rewrite Summary

**Rewrote ~390 non-imperative pattern bullets across all 105 gocritic guides to start with imperative verbs providing actionable fix direction**

## Performance

- **Duration:** 57 min
- **Started:** 2026-04-21T06:08:06Z
- **Completed:** 2026-04-21T07:05:14Z
- **Tasks:** 2
- **Files modified:** 105

## Accomplishments
- Audited all 108 gocritic guides, identified 105 with non-empty `<patterns>` sections (3 skipped: appendAssign, badCall, commentedOutCode)
- Rewrote every non-imperative pattern bullet to start with an imperative verb with actionable fix direction
- Automated verification passes with 0 failures across all 105 guides
- All spot-checked guides (argOrder, badCond, appendCombine, assignOp, wrapperFunc) confirmed correct

## Task Commits

1. **task 1: Audit gocritic guides** — analysis only, no file changes
2. **task 2: Rewrite non-imperative pattern bullets** - `5dfba75` (feat)

## Files Created/Modified
- `guides/gocritic/argOrder.md` through `guides/gocritic/zeroByteRepeat.md` — 105 files with rewritten `<patterns>` sections

## Decisions Made
- Used verbs from the approved list: Replace, Remove, Avoid, Use, Simplify, Combine, Eliminate, Add, Move, Unexport, Inline, Extract, Rename, Assign, Dereference, Examine, Pair, Swap, Compile, Reorder, Convert, Guard, Identify, Ensure, Check, Preallocate, Propagate, Separate, Return, Validate
- Kept inline code examples within bullets for context — they reinforce the fix direction
- Preserved the `→` arrow notation where it helps show the before→after transformation

## Deviations from Plan

None — plan executed exactly as written. All 105 guides with non-empty patterns rewritten, 3 without patterns skipped.

## Issues Encountered
None

## User Setup Required
None — no external service configuration required.

## Next Phase Readiness
- All 105 gocritic guides verified with 0 failures on automated imperative-verb check
- Ready for overall phase 38 verification across all linter categories

## Self-Check: PASSED

- All key files exist (argOrder.md, appendCombine.md, badCond.md, wrapperFunc.md)
- Commit 5dfba75 found in git log
- Automated verification: 105 guides checked, 0 failures

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*
