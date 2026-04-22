---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 04
subsystem: guides
tags: [linting, guides, patterns, imperative-verbs, documentation]

requires:
  - phase: 38
    provides: "Plan 01-03 pattern rewrites established imperative-verb convention"
provides:
  - "All 39 perf-testing-remaining guides with imperative-first pattern bullets"
  - "~150 pattern bullets rewritten from problem-description to fix-oriented imperative style"
affects: [phase-38-verification]

tech-stack:
  added: []
  patterns: [imperative-verb-pattern-bullets]

key-files:
  created: []
  modified:
    - guides/arangolint.md
    - guides/asasalint.md
    - guides/copyloopvar.md
    - guides/durationcheck.md
    - guides/embeddedstructfieldcheck.md
    - guides/exptostd.md
    - guides/fatcontext.md
    - guides/forbidigo.md
    - guides/gocheckcompilerdirectives.md
    - guides/gochecknoglobals.md
    - guides/gochecknoinits.md
    - guides/gochecksumtype.md
    - guides/godot.md
    - guides/godox.md
    - guides/gosmopolitan.md
    - guides/intrange.md
    - guides/ireturn.md
    - guides/lll.md
    - guides/loggercheck.md
    - guides/misspell.md
    - guides/musttag.md
    - guides/noctx.md
    - guides/nolintlint.md
    - guides/nonamedreturns.md
    - guides/paralleltest.md
    - guides/perfsprint.md
    - guides/prealloc.md
    - guides/promlinter.md
    - guides/protogetter.md
    - guides/reassign.md
    - guides/recvcheck.md
    - guides/sloglint.md
    - guides/testableexamples.md
    - guides/testpackage.md
    - guides/thelper.md
    - guides/tparallel.md
    - guides/unqueryvet.md
    - guides/wsl_v5.md
    - guides/zerologlint.md

key-decisions:
  - "Used context-specific imperative verbs: Preallocate for prealloc, Propagate for noctx, Capture for paralleltest"
  - "Added fix direction with specific code suggestions inline (e.g., make([]T, 0, n) for prealloc)"
  - "Preserved code backtick references in bullets for discoverability while leading with action verb"

patterns-established:
  - "Pattern bullets start with imperative verb: Use, Replace, Avoid, Add, Remove, Preallocate, etc."
  - "Each bullet provides actionable fix direction, not just problem description"
  - "Code snippets included inline when they directly show the fix"

requirements-completed: []

duration: 13min
completed: 2026-04-21
---

# Phase 38 Plan 04: Perf, Testing & Remaining Pattern Rewrites Summary

**Rewrote ~150 pattern bullets across 39 simple linter guides to start with imperative fix-oriented verbs per D-01**

## Performance

- **Duration:** 13 min
- **Started:** 2026-04-21T06:08:21Z
- **Completed:** 2026-04-21T06:21:34Z
- **Tasks:** 2
- **Files modified:** 39

## Accomplishments

- All 39 guides in the Perf, Testing & Remaining category now have imperative-first `<patterns>` bullets
- Every bullet provides actionable fix direction with specific code suggestions where applicable
- All guides remain under the 200-word limit
- Verification command confirms 100% imperative-verb compliance across all pattern bullets

## Task Commits

Each task was committed atomically:

1. **task 1: Audit perf-testing-remaining guides for non-imperative pattern bullets** — read-only audit, no source changes (report saved to phase directory)
2. **task 2: Rewrite non-imperative pattern bullets in perf-testing-remaining guides** - `9d1a848` (feat)

## Files Created/Modified

All 39 guide files modified (only `<patterns>` sections changed):

- `guides/arangolint.md` — Use bind parameters instead of string concatenation
- `guides/asasalint.md` — Match format verbs to argument types
- `guides/copyloopvar.md` — Remove shadow copies, eliminate workarounds
- `guides/durationcheck.md` — Avoid multiplying two Duration values
- `guides/embeddedstructfieldcheck.md` — Use named fields instead of embedding
- `guides/exptostd.md` — Replace x/exp imports with stdlib
- `guides/fatcontext.md` — Assign derived contexts to new variables
- `guides/forbidigo.md` — Replace fmt with structured logging
- `guides/gocheckcompilerdirectives.md` — Remove space in directives
- `guides/gochecknoglobals.md` — Move globals into struct fields
- `guides/gochecknoinits.md` — Replace init() with explicit setup
- `guides/gochecksumtype.md` — Handle all sum-type implementations
- `guides/godot.md` — Add trailing periods to comments
- `guides/godox.md` — Resolve TODO/FIXME comments
- `guides/gosmopolitan.md` — Use explicit time.UTC
- `guides/intrange.md` — Replace three-clause loops with range-over-int
- `guides/ireturn.md` — Return concrete types from factories
- `guides/lll.md` — Split long lines
- `guides/loggercheck.md` — Ensure even key-value arguments
- `guides/misspell.md` — Fix common typos
- `guides/musttag.md` — Add struct tags for marshaling
- `guides/noctx.md` — Replace http.NewRequest with NewRequestWithContext
- `guides/nolintlint.md` — Specify linter names in nolint directives
- `guides/nonamedreturns.md` — Remove named returns
- `guides/paralleltest.md` — Capture range variables, add t.Parallel()
- `guides/perfsprint.md` — Replace fmt.Sprintf with strconv equivalents
- `guides/prealloc.md` — Preallocate slices with make()
- `guides/promlinter.md` — Add unit suffixes to metric names
- `guides/protogetter.md` — Use generated getters on proto messages
- `guides/reassign.md` — Convert package-level vars to const
- `guides/recvcheck.md` — Use consistent receiver types
- `guides/sloglint.md` — Use consistent slog attribute style
- `guides/testableexamples.md` — Add Output comments to examples
- `guides/testpackage.md` — Move tests to external test packages
- `guides/thelper.md` — Add t.Helper() to test helpers
- `guides/tparallel.md` — Add t.Parallel() consistently
- `guides/unqueryvet.md` — Handle Query() error returns
- `guides/wsl_v5.md` — Add/remove blank lines per style rules
- `guides/zerologlint.md` — End zerolog chains with Send()/Msg()

## Decisions Made

- Used context-specific verbs (Preallocate for prealloc, Propagate for noctx, Capture for paralleltest) for maximum clarity
- Included inline code suggestions showing the fix (e.g., `make([]T, 0, n)`) rather than just describing the approach
- Balanced conciseness with specificity — some bullets combine the action with the "why" in a single line

## Deviations from Plan

None — plan executed exactly as written.

## Issues Encountered

None.

## User Setup Required

None — no external service configuration required.

## Next Phase Readiness

- All 39 perf-testing-remaining guides complete with imperative pattern bullets
- Combined with plans 01-03, all simple linter guides in the error/correctness, complexity/quality, style/formatting, and perf/testing/remaining categories are complete
- Ready for compound linter guide rewrites (gocritic, staticcheck, revive, gosec, govet)

## Self-Check: PASSED

- All 39 guide files exist and verified
- Commit `9d1a848` found in git log
- SUMMARY.md exists at expected path
- All pattern bullets verified as imperative-verb first
- All guides under 200-word limit
- No files accidentally deleted

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*
