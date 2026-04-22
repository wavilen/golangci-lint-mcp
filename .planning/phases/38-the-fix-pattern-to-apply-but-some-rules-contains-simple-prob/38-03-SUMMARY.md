---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 03
subsystem: guides
tags: [markdown, patterns, imperative-verbs, style, formatting]

# Dependency graph
requires:
  - phase: 38
    provides: context and decisions for imperative-first pattern rewrite
provides:
  - All 28 style-and-formatting guide <patterns> bullets rewritten with imperative verbs
affects: [38-04, 38-05, 38-06, 38-07, 38-08, 38-09, 38-10]

# Tech tracking
tech-stack:
  added: []
  patterns: [imperative-first pattern bullets]

key-files:
  created: []
  modified:
    - guides/asciicheck.md
    - guides/bidichk.md
    - guides/dogsled.md
    - guides/exhaustive.md
    - guides/exhaustruct.md
    - guides/forcetypeassert.md
    - guides/goconst.md
    - guides/gofmt.md
    - guides/gofumpt.md
    - guides/goheader.md
    - guides/gomoddirectives.md
    - guides/gomodguard.md
    - guides/goprintffuncname.md
    - guides/importas.md
    - guides/inamedparam.md
    - guides/iotamixing.md
    - guides/makezero.md
    - guides/mirror.md
    - guides/mnd.md
    - guides/nosprintfhostport.md
    - guides/predeclared.md
    - guides/tagalign.md
    - guides/tagliatelle.md
    - guides/unconvert.md
    - guides/unused.md
    - guides/usestdlibvars.md
    - guides/usetesting.md
    - guides/whitespace.md

key-decisions:
  - "Combined task 1 (audit) and task 2 (rewrite) into single commit since audit produced no file changes"
  - "Used 'Explicitly set' rephrased to 'Set explicitly' to keep verb-first in exhaustruct.md"

patterns-established:
  - "Imperative-first pattern bullets: Rename, Replace, Remove, Extract, Define, Use, Add, etc."

requirements-completed: []

# Metrics
duration: 14min
completed: 2026-04-21
---

# Phase 38 Plan 03: Style & Formatting Guides Summary

**Rewrote all 109 pattern bullets across 28 style-and-formatting guides to start with imperative fix-oriented verbs**

## Performance

- **Duration:** 14 min
- **Started:** 2026-04-21T06:07:52Z
- **Completed:** 2026-04-21T06:22:14Z
- **Tasks:** 2 (audit + rewrite combined into 1 commit)
- **Files modified:** 28

## Accomplishments
- Audited all 28 style-and-formatting guides — found 100% of bullets were problem descriptions without fix direction
- Rewrote all 109 pattern bullets to start with imperative verbs (Rename, Replace, Remove, Extract, Define, Use, Add, etc.)
- Verified all 28 guides pass imperative-verb check
- Spot-checked asciicheck, bidichk, goconst, tagliatelle — all correct
- All guides remain under 200-word limit

## Task Commits

1. **task 1+2: Audit and rewrite all pattern bullets** - `1f9aec5` (feat)

## Files Created/Modified
- `guides/asciicheck.md` - Renamed/Replace non-ASCII identifier patterns
- `guides/bidichk.md` - Remove/Strip Unicode bidi control patterns
- `guides/dogsled.md` - Replace/Wrap/Extract blank identifier patterns
- `guides/exhaustive.md` - Add/Include/Audit missing enum case patterns
- `guides/exhaustruct.md` - Initialize/Set struct field patterns
- `guides/forcetypeassert.md` - Replace/Guard/Use comma-ok assertion patterns
- `guides/goconst.md` - Extract/Define/Replace repeated string patterns
- `guides/gofmt.md` - Run/Apply gofmt formatting patterns
- `guides/gofumpt.md` - Run/Group/Move/Align gofumpt patterns
- `guides/goheader.md` - Add/Update/Fix copyright header patterns
- `guides/gomoddirectives.md` - Remove/Resolve/Add go.mod directive patterns
- `guides/gomodguard.md` - Replace/Remove blocked module patterns
- `guides/goprintffuncname.md` - Rename/Add format function suffix patterns
- `guides/importas.md` - Use/Replace/Configure/Standardize import alias patterns
- `guides/inamedparam.md` - Name/Add/Replace/Give interface parameter patterns
- `guides/iotamixing.md` - Split/Separate mixed iota patterns
- `guides/makezero.md` - Use/Set/Change slice allocation patterns
- `guides/mirror.md` - Call/Replace/Dereference/Use reflect.Value patterns
- `guides/mnd.md` - Extract/Replace/Define magic number patterns
- `guides/nosprintfhostport.md` - Replace/Use/Build host:port patterns
- `guides/predeclared.md` - Rename shadowing identifier patterns
- `guides/tagalign.md` - Align/Add struct tag patterns
- `guides/tagliatelle.md` - Apply/Convert/Lowercase/Standardize tag casing patterns
- `guides/unconvert.md` - Remove unnecessary type conversion patterns
- `guides/unused.md` - Remove/Delete unused declaration patterns
- `guides/usestdlibvars.md` - Replace hardcoded values with stdlib constants
- `guides/usetesting.md` - Replace manual temp dir with t.TempDir()
- `guides/whitespace.md` - Remove/Reduce extraneous whitespace patterns

## Decisions Made
- Combined task 1 (audit) and task 2 (rewrite) into a single commit since the audit was purely analytical with no file output
- Rephrased "Explicitly set" to "Set explicitly" in exhaustruct.md to keep verb-first pattern
- Verification regex expanded beyond the plan's default list to include additional imperative verbs (Initialize, Update, Fix, Resolve, Configure, etc.)

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 1 - Bug] Fixed non-imperative start in exhaustruct.md**
- **Found during:** task 2 (rewrite)
- **Issue:** "Explicitly set" starts with adverb, not imperative verb
- **Fix:** Rephrased to "Set semantically important fields explicitly"
- **Files modified:** guides/exhaustruct.md
- **Committed in:** 1f9aec5 (task 2 commit)

---

**Total deviations:** 1 auto-fixed (1 bug)
**Impact on plan:** Trivial — one bullet rephrase for consistency.

## Issues Encountered
None.

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- All 28 style-and-formatting guides have imperative-first pattern bullets
- Ready for remaining category plans (perf/testing, gocritic, staticcheck, revive, gosec, govet, minor compound)

---
*Phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob*
*Completed: 2026-04-21*

## Self-Check: PASSED

- All 28 guide files: FOUND
- Commit 1f9aec5: FOUND
- SUMMARY.md: FOUND
