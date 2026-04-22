---
phase: 49-hook-detection-and-deduplication
plan: 01
subsystem: hooks
tags: [shared-module, deduplication, hook-detection, cjs, esm, npm-installer]

requires:
  - phase: 45
    provides: Shell-wrapper detection in isGolangciLintCommand()
  - phase: 42
    provides: Precise command detection (first-token, subcommand filtering)

provides:
  - "shared/nudge.js: CJS module with all 13 shared symbols for detection, nudging, output flag injection"
  - "shared/nudge.test.js: 106 tests covering all shared functions"
  - "Simplified hook (26 lines) importing from shared module with precise isGolangciLintCommand()"
  - "Simplified plugin (49 lines) importing from shared module with backward-compatible re-exports"
  - "Updated installer deploying shared module to .claude/shared/ and .opencode/shared/"
  - "Makefile verify-shared target for module load verification"

affects: [hooks, plugins, installer, npm-package]

tech-stack:
  added: []
  patterns:
    - "CJS shared module pattern: module.exports for hook (CJS) + ESM re-export for plugin"
    - "Single source of truth: all detection + nudge logic in shared/nudge.js"

key-files:
  created:
    - shared/nudge.js
    - shared/nudge.test.js
  modified:
    - hooks/golangci-lint-post.js
    - plugins/golangci-lint.js
    - bin/install.js
    - package.json
    - Makefile

key-decisions:
  - "CJS module.exports for shared/nudge.js — hooks use require(), plugins use ESM import"
  - "Plugin re-exports all shared functions for backward compatibility with existing tests"
  - "Hook uses path.join(__dirname, '..', 'shared', 'nudge.js') for relative require"

patterns-established:
  - "Shared module pattern: CJS exports for hook compatibility, ESM import destructuring for plugin"

requirements-completed: []

duration: 13min
completed: 2026-04-22
---

# Phase 49 Plan 01: Hook Detection and Deduplication Summary

**Extracted shared nudge + detection logic into CJS module, replaced 80+ lines of duplicated code between hook (163→26 lines) and plugin (321→49 lines), upgraded hook from indexOf to precise isGolangciLintCommand() detection**

## Performance

- **Duration:** 13 min
- **Started:** 2026-04-22T12:17:58Z
- **Completed:** 2026-04-22T12:31:25Z
- **Tasks:** 3
- **Files modified:** 7

## Accomplishments
- Created shared/nudge.js CJS module with all 13 exported symbols (constants + functions)
- Eliminated all duplicated nudge function bodies between hook and plugin — single source of truth
- Hook upgraded from imprecise `indexOf('golangci-lint')` to precise `isGolangciLintCommand()` detection
- All 199 tests pass: 106 in shared module + 93 in plugin (re-export backward compatibility)
- Installer deploys shared module to both .claude/shared/ and .opencode/shared/
- package.json files array includes shared/ for npm tarball distribution

## Task Commits

Each task was committed atomically:

1. **Task 1 (RED): Failing test for shared nudge module** - `26d523c` (test)
2. **Task 1 (GREEN): Create shared nudge module** - `f97be7b` (feat)
3. **Task 2: Simplify hook + plugin to import from shared** - `f848492` (feat)
4. **Task 3: Update installer, package.json, Makefile** - `c853924` (feat)

_Note: Task 1 followed TDD — RED (failing test) then GREEN (implementation)._

## Files Created/Modified
- `shared/nudge.js` - CJS shared module with 13 exported symbols (detection, nudge, output injection)
- `shared/nudge.test.js` - 106 tests migrated from plugin tests, CJS require imports
- `hooks/golangci-lint-post.js` - Simplified from 163→26 lines, uses shared.isGolangciLintCommand()
- `plugins/golangci-lint.js` - Simplified from 321→49 lines, imports from ../shared/nudge.js
- `bin/install.js` - Copies shared/nudge.js to .claude/shared/ and .opencode/shared/
- `package.json` - Added "shared/" to files array
- `Makefile` - Updated install-plugin to copy shared module, added verify-shared target

## Decisions Made
- Used CJS module.exports for shared module — hooks require CJS, plugins use ESM import from same file
- Plugin re-exports all shared functions for backward compatibility with existing test suite
- Hook path uses `path.join(__dirname, '..', 'shared', 'nudge.js')` for safe relative resolution

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 1 - Bug] Fixed hook require path — missing `..` parent directory**
- **Found during:** Task 2 (hook verification)
- **Issue:** Used `path.join(__dirname, 'shared', 'nudge.js')` which resolved to hooks/shared/ instead of project-root/shared/
- **Fix:** Changed to `path.join(__dirname, '..', 'shared', 'nudge.js')` as specified in the plan
- **Files modified:** hooks/golangci-lint-post.js
- **Verification:** Hook end-to-end test passes with golangci-lint run and rejects golangci-lint version
- **Committed in:** f848492 (part of task 2 commit)

**2. [Rule 3 - Blocking] Task 2 make verify-plugin deferred to task 3**
- **Found during:** Task 2 acceptance criteria verification
- **Issue:** `make verify-plugin` failed because the Makefile didn't yet copy shared/nudge.js alongside the plugin
- **Fix:** This was expected — task 3 updates the Makefile to copy the shared module. Verified manually that the plugin logic was correct, deferred `make verify-plugin` to task 3
- **Verification:** After task 3, `make verify-plugin` passes with shared module deployment
- **Committed in:** c853924 (task 3 commit)

---

**Total deviations:** 2 auto-fixed (1 bug, 1 blocking)
**Impact on plan:** Both fixes necessary for correctness. No scope creep.

## Issues Encountered
None

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- Phase 49 complete — shared module pattern established for all detection + nudge logic
- Ready for any follow-up work requiring changes to nudge behavior (single file to modify)
- No blockers or concerns

---
*Phase: 49-hook-detection-and-deduplication*
*Completed: 2026-04-22*

## Self-Check: PASSED

All 8 key files verified present. All 4 commits verified in git log.
