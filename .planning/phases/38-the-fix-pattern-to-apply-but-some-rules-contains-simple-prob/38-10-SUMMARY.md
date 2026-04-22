---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 10
subsystem: guides
tags: [errorlint, ginkgolinter, grouper, modernize, testifylint, patterns, imperative-verbs]

# Dependency graph
requires:
  - phase: 38
    provides: imperative-verb pattern style decision (D-01)
provides:
  - All 49 minor compound linter guides with imperative-first <patterns> bullets
affects: [guide-validation, desloppify]

# Tech tracking
tech-stack:
  added: []
  patterns: [imperative-first-pattern-bullets]

key-files:
  created: []
  modified:
    - guides/errorlint/asserts.md
    - guides/errorlint/comparison.md
    - guides/errorlint/errorf.md
    - guides/ginkgolinter/async-assertion.md
    - guides/ginkgolinter/async-intervals.md
    - guides/ginkgolinter/compare-assertion.md
    - guides/ginkgolinter/error-assertion.md
    - guides/ginkgolinter/expect-to.md
    - guides/ginkgolinter/focus-container.md
    - guides/ginkgolinter/have-len-zero.md
    - guides/ginkgolinter/len-assertion.md
    - guides/ginkgolinter/nil-assertion.md
    - guides/ginkgolinter/spec-pollution.md
    - guides/ginkgolinter/succeed-matcher.md
    - guides/ginkgolinter/type-compare.md
    - guides/grouper/const.md
    - guides/grouper/import.md
    - guides/grouper/type.md
    - guides/grouper/var.md
    - guides/modernize/errorf.md
    - guides/modernize/loopvar.md
    - guides/modernize/maprange.md
    - guides/modernize/mapval.md
    - guides/modernize/reloop.md
    - guides/modernize/simplifyrange.md
    - guides/modernize/sliceclear.md
    - guides/modernize/slicesort.md
    - guides/modernize/sortfunc.md
    - guides/modernize/stringappend.md
    - guides/testifylint/blank-import.md
    - guides/testifylint/bool-compare.md
    - guides/testifylint/compares.md
    - guides/testifylint/contains-unnecessary-format.md
    - guides/testifylint/empty.md
    - guides/testifylint/error-as.md
    - guides/testifylint/error-nil.md
    - guides/testifylint/expected-actual.md
    - guides/testifylint/float-compare.md
    - guides/testifylint/formatter.md
    - guides/testifylint/go-require.md
    - guides/testifylint/len.md
    - guides/testifylint/nil-compare.md
    - guides/testifylint/require-error.md
    - guides/testifylint/suite-broken-parallel.md
    - guides/testifylint/suite-dont-use-pkg.md
    - guides/testifylint/suite-extra-assert-call.md
    - guides/testifylint/suite-method-signature.md
    - guides/testifylint/suite-thelper.md
    - guides/testifylint/useless-assert.md

key-decisions:
  - "Rewrote all 150 pattern bullets to lead with imperative verbs (Use, Replace, Remove, Add, Group, etc.)"
  - "Preserved semantic meaning — only changed sentence structure, not technical content"

patterns-established:
  - "Imperative-first pattern bullets: every <patterns> bullet starts with an imperative verb giving fix direction"

requirements-completed: []

# Metrics
duration: 51min
completed: 2026-04-21
---

# Phase 38 Plan 10: Minor Compound Linter Guides Summary

**Rewrote 150 pattern bullets across 49 minor compound guides (errorlint, ginkgolinter, grouper, modernize, testifylint) to imperative-first style**

## Performance

- **Duration:** 51 min
- **Started:** 2026-04-21T06:25:31Z
- **Completed:** 2026-04-21T07:16:44Z
- **Tasks:** 2 (1 audit-only, 1 rewrite)
- **Files modified:** 49

## Accomplishments
- Audited all 49 minor compound guides, identified 149 non-imperative pattern bullets (all bullets were non-imperative)
- Rewrote every bullet to start with imperative verbs (Use, Replace, Remove, Add, Group, etc.)
- Verification passes: 0 non-imperative bullets across all 49 guides
- All guides remain under 500-word compound linter limit
- No changes to `<instructions>`, `<examples>`, or `<related>` sections

## Task Commits

1. **task 1: Audit minor compound guides** - No commit (analysis-only task, no file changes)
2. **task 2: Rewrite non-imperative pattern bullets** - `71964e8` (feat)

## Files Created/Modified
All 49 files in 5 directories:
- `guides/errorlint/*.md` (3 files) — Error assertion, comparison, and formatting patterns
- `guides/ginkgolinter/*.md` (12 files) — Ginkgo assertion patterns
- `guides/grouper/*.md` (4 files) — Declaration grouping patterns
- `guides/modernize/*.md` (10 files) — Go modernization patterns
- `guides/testifylint/*.md` (20 files) — Testify assertion patterns

## Decisions Made
- Used "Simplify" instead of "Collapse" for succeed-matcher (matches accepted verb list)
- Used "Separate" instead of "Isolate" for spec-pollution (matches accepted verb list)
- Used "Move" instead of "Swap" for expected-actual (clearer imperative)
- Used "Replace" instead of "Compare" for useless-assert (matches accepted verb list)
- Used "Remove" instead of "Drop" for simplifyrange (matches accepted verb list)

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 3 - Blocking] File write tool batch failures**
- **Found during:** task 2 (rewrite)
- **Issue:** Parallel edit tool calls reported success but changes didn't persist on disk for first batch
- **Fix:** Re-wrote all files using sequential write tool calls with inter-bash verification
- **Files modified:** All 49 guide files (required multiple write attempts)
- **Verification:** Full imperative-verb scan passes with 0 failures
- **Committed in:** 71964e8 (task 2 commit)

---

**Total deviations:** 1 auto-fixed (1 blocking)
**Impact on plan:** Tooling issue required re-writing files; no scope creep or design changes.

## Issues Encountered
- File write tool reported success for parallel writes but content didn't persist — resolved by writing files sequentially with verification between batches

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- All 49 minor compound guides now have imperative-first pattern bullets
- Phase 38 pattern standardization is complete across all linter categories
- Ready for final verification across all 630 guides

## Self-Check: PASSED

- Commit 71964e8: FOUND
- SUMMARY.md: FOUND
- Files in commit: 49 (expected: 49)
- Verification: 49 guides checked, 0 failures

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*
