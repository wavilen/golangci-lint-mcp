---
phase: 42-opencode-plugin-filters-golangci-lint-command-too-widely-it
plan: 01
subsystem: testing
tags: [opencode, plugin, regex, node:test, command-detection]

# Dependency graph
requires:
  - phase: 40-strip-pipe-filters-in-plugin
    provides: stripOutputFilters function with indexOf-based token location
  - phase: 41-npx-installer-opencode-plugin
    provides: Plugin installed via npm to .opencode/plugins/
provides:
  - isGolangciLintCommand() with precise first-token detection for golangci-lint commands
  - 42-test regression suite for all plugin utility functions
  - Named exports for stripOutputFilters, injectJsonOutputFlag, parseDiagnostics
affects: [opencode-plugin, command-hooks, future-plugin-changes]

# Tech tracking
tech-stack:
  added: [node:test, node:assert/strict]
  patterns: [first-token command detection, env-var prefix stripping]

key-files:
  created:
    - plugins/golangci-lint.test.js
  modified:
    - plugins/golangci-lint.js

key-decisions:
  - "Used first-token extraction + endsWith('/golangci-lint') instead of complex regex for more readable and maintainable detection"
  - "Exported utility functions (stripOutputFilters, injectJsonOutputFlag, parseDiagnostics) for testability"

patterns-established:
  - "First-token detection: strip env vars → extract first whitespace-delimited token → check equality or path suffix"
  - "Named exports alongside default plugin export for test access without breaking module pattern"

requirements-completed: []

# Metrics
duration: 4min
completed: 2026-04-21
---

# Phase 42 Plan 01: Fix golangci-lint Command Detection Summary

**Precise first-token command detection replacing substring indexOf, with 42-test suite covering positive/negative detection and function regressions**

## Performance

- **Duration:** 4 min
- **Started:** 2026-04-21T15:18:06Z
- **Completed:** 2026-04-21T15:23:04Z
- **Tasks:** 1 (TDD)
- **Files modified:** 2

## Accomplishments
- Fixed false-positive plugin triggering on commands like `echo golangci-lint`, `cat docs/golangci-lint.md`, `git commit -m "fix golangci-lint"`
- New `isGolangciLintCommand()` function uses first-token extraction with env-var prefix stripping
- Handles path-qualified invocations (`/usr/bin/golangci-lint`, `./golangci-lint`, `~/bin/golangci-lint`)
- Handles env var prefixed commands (`FOO=bar golangci-lint run`)
- 42 tests passing: 11 positive detection, 14 negative detection, 8 stripOutputFilters, 4 injectJsonOutputFlag, 5 parseDiagnostics

## Task Commits

Each task was committed atomically (TDD):

1. **Task 1 RED: Failing test suite** - `01a46bf` (test)
2. **Task 1 GREEN: Implement isGolangciLintCommand + update hooks** - `3cc339b` (feat)

_Note: TDD task has two commits (test → implementation)_

## Files Created/Modified
- `plugins/golangci-lint.js` - Added isGolangciLintCommand(), updated both hooks to use it, updated stripOutputFilters token location, added named exports
- `plugins/golangci-lint.test.js` - New 221-line test suite with 42 test cases using node:test

## Decisions Made
- Used first-token extraction (`cmd.match(/^(\S+)/)`) + `endsWith('/golangci-lint')` instead of a complex regex — simpler to understand and verify, handles all path variants naturally
- Exported utility functions for test access without changing the plugin's default export pattern

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 1 - Bug] Simplified regex to first-token extraction for robustness**
- **Found during:** Task 1 GREEN phase
- **Issue:** Plan-specified regex `/^(?:\.\/|~\/|\/[\S]*\/)?golangci-lint(?:\s|$)/` failed to match `~/bin/golangci-lint` because `~\/` only matched `~/` (2 chars), not `~/bin/`
- **Fix:** Replaced regex approach with first-token extraction + string comparison: extract first non-whitespace token after env var stripping, then check `firstToken === 'golangci-lint' || firstToken.endsWith('/golangci-lint')`
- **Files modified:** plugins/golangci-lint.js
- **Verification:** All 42 tests pass including `~/bin/golangci-lint run`
- **Committed in:** 3cc339b (task 1 GREEN commit)

---

**Total deviations:** 1 auto-fixed (1 bug)
**Impact on plan:** Implementation more robust than plan's regex approach. All acceptance criteria met.

## Issues Encountered
None beyond the regex edge case documented above.

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- Plugin command detection is precise and well-tested
- All existing functionality (stripOutputFilters, injectJsonOutputFlag, after-hook nudge) continues working
- No `indexOf('golangci-lint')` remains in the codebase

## Self-Check: PASSED

- ✅ `plugins/golangci-lint.js` exists
- ✅ `plugins/golangci-lint.test.js` exists
- ✅ Commit `01a46bf` (RED) exists
- ✅ Commit `3cc339b` (GREEN) exists

---
*Phase: 42-opencode-plugin-filters-golangci-lint-command-too-widely-it*
*Completed: 2026-04-21*
