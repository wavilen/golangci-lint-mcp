---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 09
subsystem: guides
tags: [govet, patterns, imperative-verbs, content-editing]

# Dependency graph
requires:
  - phase: 38
    provides: govet guide files with <patterns> sections
provides:
  - All 35 govet guides with imperative-first <patterns> bullets
affects: [govet, patterns, guide-quality]

# Tech tracking
tech-stack:
  added: []
  patterns: [imperative-verb-first pattern bullets]

key-files:
  created: []
  modified:
    - guides/govet/assign.md
    - guides/govet/atomic.md
    - guides/govet/copylocks.md
    - guides/govet/loopclosure.md
    - guides/govet/appends.md
    - guides/govet/asmdecl.md
    - guides/govet/bools.md
    - guides/govet/buildtag.md
    - guides/govet/cgocall.md
    - guides/govet/composites.md
    - guides/govet/defers.md
    - guides/govet/directive.md
    - guides/govet/errorsas.md
    - guides/govet/framepointer.md
    - guides/govet/hostport.md
    - guides/govet/httpresponse.md
    - guides/govet/ifaceassert.md
    - guides/govet/lostcancel.md
    - guides/govet/nilfunc.md
    - guides/govet/printf.md
    - guides/govet/shift.md
    - guides/govet/sigchanyzer.md
    - guides/govet/slog.md
    - guides/govet/stdmethods.md
    - guides/govet/stdversion.md
    - guides/govet/stringintconv.md
    - guides/govet/structtag.md
    - guides/govet/testinggoroutine.md
    - guides/govet/tests.md
    - guides/govet/timeformat.md
    - guides/govet/unmarshal.md
    - guides/govet/unreachable.md
    - guides/govet/unsafeptr.md
    - guides/govet/unusedresult.md
    - guides/govet/waitgroup.md

key-decisions:
  - "All 120 pattern bullets across 35 govet guides rewritten from problem-descriptive to imperative-first fix-oriented style"
  - "Two initial failures (httpresponse 'Read', loopclosure 'Capture') rephrased with verbs from approved list"

patterns-established:
  - "Govet patterns now use imperative verbs: Use, Remove, Pass, Avoid, Add, Fix, Ensure, Match, Call, Store, Place, Save, Copy, Complete, Close, Split, Pair, Check, Validate, Allocate, Move, Wrap, Provide, Capitalize, Mask"

requirements-completed: []

# Metrics
duration: 14min
completed: 2026-04-21
---

# Phase 38 Plan 09: Govet Imperative Pattern Bullets Summary

**Rewrote all 120 pattern bullets across 35 govet guides from problem-descriptive to imperative-first fix-oriented style**

## Performance

- **Duration:** 14 min
- **Started:** 2026-04-21T07:16:57Z
- **Completed:** 2026-04-21T07:31:33Z
- **Tasks:** 2
- **Files modified:** 35

## Accomplishments
- Audited all 35 govet guides — found 100% of pattern bullets were problem-descriptive (no imperative verbs)
- Rewrote all 120 bullets to start with imperative verbs providing actionable fix direction
- Verified 0 failures in imperative-verb check across all guides
- Confirmed all guides remain under the 500-word compound linter limit

## Task Commits

Each task was committed atomically:

1. **task 1: Audit govet guides** — audit only, no file changes (identified all 120 bullets as non-imperative)
2. **task 2: Rewrite non-imperative pattern bullets** - `50a3f32` (feat)

## Files Created/Modified
- `guides/govet/assign.md` - Self-assignment patterns rewritten with "Remove" imperative
- `guides/govet/atomic.md` - Atomic access patterns rewritten with "Use", "Pass" imperatives
- `guides/govet/copylocks.md` - Lock copy patterns rewritten with "Use pointer" imperatives
- `guides/govet/loopclosure.md` - Loop capture patterns rewritten with "Copy", "Pass" imperatives
- `guides/govet/appends.md` - Append result patterns rewritten with "Assign", "Use" imperatives
- `guides/govet/asmdecl.md` - Assembly mismatch patterns rewritten with "Add", "Match", "Use" imperatives
- `guides/govet/bools.md` - Boolean redundancy patterns rewritten with "Remove", "Simplify" imperatives
- `guides/govet/buildtag.md` - Build tag patterns rewritten with "Place", "Use", "Fix", "Add" imperatives
- `guides/govet/cgocall.md` - Cgo pointer patterns rewritten with "Avoid" imperatives
- `guides/govet/composites.md` - Unkeyed literal patterns rewritten with "Use keyed" imperatives
- `guides/govet/defers.md` - Loop defer patterns rewritten with "Move", "Wrap", "Use" imperatives
- `guides/govet/directive.md` - Directive patterns rewritten with "Place", "Use", "Fix" imperatives
- `guides/govet/errorsas.md` - Errors.As patterns rewritten with "Pass" imperatives
- `guides/govet/framepointer.md` - Frame pointer patterns rewritten with "Save", "Add", "Allocate" imperatives
- `guides/govet/hostport.md` - Host:port patterns rewritten with "Split", "Avoid", "Separate" imperatives
- `guides/govet/httpresponse.md` - HTTP body patterns rewritten with "Add", "Ensure", "Close" imperatives
- `guides/govet/ifaceassert.md` - Interface assertion patterns rewritten with "Remove" imperatives
- `guides/govet/lostcancel.md` - Cancel function patterns rewritten with "Store", "Call", "Ensure" imperatives
- `guides/govet/nilfunc.md` - Nil function patterns rewritten with "Remove" imperatives
- `guides/govet/printf.md` - Format verb patterns rewritten with "Match", "Add", "Complete" imperatives
- `guides/govet/shift.md` - Shift overflow patterns rewritten with "Use", "Validate" imperatives
- `guides/govet/sigchanyzer.md` - Signal channel patterns rewritten with "Use", "Create" imperatives
- `guides/govet/slog.md` - Structured log patterns rewritten with "Provide", "Use", "Pair", "Ensure" imperatives
- `guides/govet/stdmethods.md` - Standard method patterns rewritten with "Ensure" imperatives
- `guides/govet/stdversion.md` - Version constraint patterns rewritten with "Use", "Fix" imperatives
- `guides/govet/stringintconv.md` - Int-to-string patterns rewritten with "Use", "Replace" imperatives
- `guides/govet/structtag.md` - Struct tag patterns rewritten with "Remove", "Provide", "Fix" imperatives
- `guides/govet/testinggoroutine.md` - Test goroutine patterns rewritten with "Call", "Avoid" imperatives
- `guides/govet/tests.md` - Test signature patterns rewritten with "Add", "Remove", "Capitalize" imperatives
- `guides/govet/timeformat.md` - Time format patterns rewritten with "Use" imperatives
- `guides/govet/unmarshal.md` - Unmarshal target patterns rewritten with "Pass" imperatives
- `guides/govet/unreachable.md` - Dead code patterns rewritten with "Remove" imperatives
- `guides/govet/unsafeptr.md` - Unsafe pointer patterns rewritten with "Convert", "Use", "Avoid" imperatives
- `guides/govet/unusedresult.md` - Unused result patterns rewritten with "Check", "Use" imperatives
- `guides/govet/waitgroup.md` - WaitGroup patterns rewritten with "Use", "Pass", "Call", "Store" imperatives

## Decisions Made
- Used "Avoid" (not "Don't use") for anti-patterns in cgocall, testinggoroutine, unsafeptr — more concise and standard
- Used "Ensure" for stdmethods verification patterns — clearer than "Make sure"
- Used "Remove" for self-assignment/boolean/dead-code patterns — the fix is deletion

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered
None

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- All 35 govet guides have imperative-first pattern bullets
- Ready for any further guide quality audits or cross-check plans

## Self-Check: PASSED

- All 5 key files exist on disk
- Commit `50a3f32` found in git history
- Imperative-verb verification: 35 guides checked, 0 failures

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*
