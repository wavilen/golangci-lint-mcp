---
phase: 45-something-prevent-opencode-plugins-from-detecting-real-golan
plan: 01
subsystem: opencode-plugin
tags: [opencode, plugin, shell-wrapper, bash, golangci-lint]

# Dependency graph
requires:
  - phase: 42
    provides: "first-token golangci-lint command detection in isGolangciLintCommand()"
provides:
  - "Shell-wrapper detection in isGolangciLintCommand() for bash -c, sh -c, zsh -c prefixes"
  - "export default GolangciLintPlugin for opencode module loader compatibility"
  - "Debug logging in tool.execute.before and tool.execute.after hooks"
  - "install-plugin and verify-plugin Makefile targets"
affects: [opencode-plugin, golangci-lint-plugin, plugin-loading]

# Tech tracking
tech-stack:
  added: []
  patterns: [shell-wrapper-regex, defensive-export-default]

key-files:
  created: [.opencode/plugins/golangci-lint.js]
  modified: [plugins/golangci-lint.js, plugins/golangci-lint.test.js, Makefile]

key-decisions:
  - "Added export default as fallback for opencode PluginLoader.getLegacyPlugins() scanning Object.values(mod)"
  - "Added defensive shell-wrapper stripping after env-var stripping for robustness against shell-wrapped commands"
  - "Added console.error debug logging for troubleshooting hook firing in opencode logs"

patterns-established:
  - "Shell-wrapper regex: /^(?:\\/[\w\\/]+\\/)?(?:bash|sh|zsh|dash|ksh)\s+-c\s+[\"']?/"
  - "Debug logging pattern: console.error with [golangci-lint-plugin] prefix"

requirements-completed: []

# Metrics
duration: 3min
completed: 2026-04-21
---

# Phase 45 Plan 01: Diagnose Plugin Loading Summary

**Shell-wrapper detection for golangci-lint commands via bash -c prefixes, export default for opencode compatibility, and Makefile install/verify targets**

## Performance

- **Duration:** ~3 min
- **Started:** 2026-04-21T20:23:55Z
- **Completed:** 2026-04-21T20:26:35Z
- **Tasks:** 2
- **Files modified:** 4

## Accomplishments
- Added shell-wrapper regex detection in `isGolangciLintCommand()` for bash/sh/zsh/dash/ksh `-c` prefixes
- Added `export default GolangciLintPlugin` for opencode PluginLoader module scanning compatibility
- Added debug `console.error` logging in both `tool.execute.before` and `tool.execute.after` hooks
- Created `install-plugin` and `verify-plugin` Makefile targets for local development workflow
- All 69 tests pass (62 existing + 7 new shell-wrapper detection tests)

## Task Commits

Each task was committed atomically:

1. **Task 1: Diagnose plugin loading and fix export format** - `b1456d6` (feat)
2. **Task 2: Create Makefile target for local plugin install + verify end-to-end** - `0217984` (chore)

## Files Created/Modified
- `plugins/golangci-lint.js` - Shell-wrapper detection regex, export default, debug logging
- `plugins/golangci-lint.test.js` - 7 new shell-wrapper detection tests
- `Makefile` - install-plugin and verify-plugin targets
- `.opencode/plugins/golangci-lint.js` - Installed plugin for local development

## Decisions Made
- Added `export default GolangciLintPlugin` as fallback — opencode's `getLegacyPlugins()` scans `Object.values(mod)` for plugin functions, so named exports already work, but default export provides belt-and-suspenders compatibility
- Shell-wrapper regex is strict: only matches known shell names (bash, sh, zsh, dash, ksh) followed by `-c` flag — prevents arbitrary command injection (T-45-01 threat mitigation)
- Debug logging uses `console.error` with `[golangci-lint-plugin]` prefix for easy filtering in opencode logs

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered
None.

## User Setup Required
None - no external service configuration required.

## Next Phase Readiness
- Plugin now has defensive shell-wrapper detection for cases where opencode wraps bash commands
- Plugin has export default for module loader compatibility
- `make install-plugin` available for local development workflow
- Pre-existing golangci-lint issues in Go test files (exhaustruct, testpackage) are out of scope for this plugin-focused phase

## Self-Check: PASSED

All files and commits verified:
- plugins/golangci-lint.js ✓
- plugins/golangci-lint.test.js ✓
- Makefile ✓
- .opencode/plugins/golangci-lint.js ✓
- 45-01-SUMMARY.md ✓
- b1456d6 (Task 1) ✓
- 0217984 (Task 2) ✓

---
*Phase: 45-something-prevent-opencode-plugins-from-detecting-real-golan*
*Completed: 2026-04-21*
