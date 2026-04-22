---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 07
subsystem: guides
tags: [revive, linter, patterns, imperative-verbs, content-editing]

# Dependency graph
requires:
  - phase: 38
    provides: phase context and decisions D-01 through D-10
provides:
  - All 101 revive guides with imperative-first pattern bullets
affects: []

# Tech tracking
tech-stack:
  added: []
  patterns: []

key-files:
  created: []
  modified:
    - guides/revive/*.md (101 files)

key-decisions:
  - "Used only approved imperative verbs from plan's verification regex to ensure consistency"

patterns-established:
  - "Imperative-first pattern bullets: every <patterns> bullet starts with an imperative verb and provides fix direction"

requirements-completed: []

# Metrics
duration: 51min
completed: 2026-04-21
---

# Phase 38 Plan 07: Revive Guides Imperative Patterns Summary

**Rewrote all 500 pattern bullets across 101 revive guides from problem-descriptive to imperative-first fix direction**

## Performance

- **Duration:** 51 min
- **Started:** 2026-04-21T06:07:36Z
- **Completed:** 2026-04-21T06:58:32Z
- **Tasks:** 2 (1 audit, 1 rewrite)
- **Files modified:** 101

## Accomplishments
- Audited all 101 revive guides identifying ~458 non-imperative pattern bullets out of 500 total
- Rewrote every pattern bullet to start with an imperative verb and provide actionable fix direction
- All 101 guides verified: 0 failures against the imperative-verb verification check
- No changes to non-patterns sections (instructions, examples, related)

## Task Commits

1. **task 1: Audit revive guides** — No file changes (purely analytical)
2. **task 2: Rewrite non-imperative pattern bullets** - `02888bc` (fix)

## Files Created/Modified
- `guides/revive/*.md` (101 files) - All pattern bullets rewritten with imperative-first fix direction

## Decisions Made
- Used only approved imperative verbs from the plan's verification regex (Use, Replace, Remove, Add, Move, Ensure, etc.) to guarantee 0 verification failures
- Transformed problem-descriptive bullets (e.g., "Parameters added for future use") to fix-oriented (e.g., "Remove parameters added for future use that are never referenced")

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 3 - Blocking] Bulk fix accidentally modified non-revive guide files**
- **Found during:** task 2
- **Issue:** awk bulk-fix command ran on all guides/*.md instead of only guides/revive/*.md
- **Fix:** Reverted all non-revive files using `git checkout --`
- **Files modified:** ~91 non-revive guide files (all reverted)
- **Verification:** `git status --short | grep -v 'guides/revive/'` shows 0 changes
- **Committed in:** 02888bc (reverted before commit)

---

**Total deviations:** 1 auto-fixed (1 blocking)
**Impact on plan:** No scope creep — all non-revive changes reverted before commit.

## Issues Encountered
- Initial rewrites used verbs outside the plan's verification regex (e.g., Fix, Split, Delete), requiring iterative fixes to use only approved verbs
- Bulk sed/awk fix was efficient but over-scoped; resolved by reverting non-target files

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- All 101 revive guides have imperative-first pattern bullets
- Ready for other linter category plans to apply the same transformation

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*
