---
phase: quick
plan: smo
subsystem: plugins/golangci-lint
tags: [plugin, command-filtering, subcommand-discrimination, tdd]
dependency_graph:
  requires: []
  provides: [isGolangciLintCommand-subcommand-filter]
  affects: [plugins/golangci-lint.js]
tech_stack:
  added: [NON_RUN_SUBCOMMANDS constant, flag-value skipping heuristic]
  patterns: [subcommand discrimination via first non-flag token]
key_files:
  created: []
  modified:
    - plugins/golangci-lint.js
    - plugins/golangci-lint.test.js
decisions:
  - "Flag-value pairs handled by skipping next token after flags without = sign"
  - "Unknown subcommands default to intercept (safer to over-intercept than miss run)"
  - "NON_RUN_SUBCOMMANDS exported for testability"
metrics:
  duration: 4min
  completed: "2026-04-21T17:53:10Z"
  tasks: 1
  files: 2
  tests_added: 20
  tests_total: 62
---

# Quick Task 260421-smo: Subcommand Discrimination Summary

Restrict `isGolangciLintCommand` to only return `true` for `golangci-lint run` (or bare `golangci-lint` which defaults to run), not for other subcommands like cache, version, linters, etc.

## What Changed

### `plugins/golangci-lint.js`
- Added `NON_RUN_SUBCOMMANDS` array: `['cache', 'completion', 'config', 'custom', 'help', 'linters', 'version']`
- Modified `isGolangciLintCommand` to extract the first non-flag token after the command and check it against known subcommands
- Added flag-value pair skipping: when a flag doesn't contain `=`, the next token is treated as its value and skipped
- Exported `NON_RUN_SUBCOMMANDS` for testability

### `plugins/golangci-lint.test.js`
- Added new describe block `isGolangciLintCommand — only run subcommand` with 20 test cases
- 8 positive cases (run, bare, flags-only, path+run, env+run)
- 12 negative cases (cache, version, linters, help, config, completion, custom + path/env/flags variants)

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 1 - Bug] Flag-value pair handling not in plan's implementation approach**
- **Found during:** GREEN phase — test `golangci-lint --timeout 5m cache clean` failed
- **Issue:** Plan's implementation found `5m` as first non-flag token (it's the value of `--timeout`), treating it as unknown subcommand → returned true instead of false
- **Fix:** Added `skipNext` heuristic — after a flag without `=`, skip the next token as its value
- **Files modified:** `plugins/golangci-lint.js`
- **Commit:** `0f7271a`

## Commits

| Hash | Type | Description |
|------|------|-------------|
| `beee779` | test | Add 20 failing tests for subcommand discrimination (RED) |
| `0f7271a` | feat | Implement subcommand discrimination in isGolangciLintCommand (GREEN) |

## Verification

- All 62 tests pass (42 existing + 20 new)
- golangci-lint run: 0 issues
- Done criteria verified:
  - `isGolangciLintCommand('golangci-lint cache')` → `false` ✓
  - `isGolangciLintCommand('golangci-lint version')` → `false` ✓
  - `isGolangciLintCommand('golangci-lint run')` → `true` ✓
  - `isGolangciLintCommand('golangci-lint')` → `true` ✓
  - `isGolangciLintCommand('golangci-lint -E errcheck')` → `true` ✓

## Self-Check: PASSED

- plugins/golangci-lint.js: FOUND
- plugins/golangci-lint.test.js: FOUND
- Commit beee779: FOUND
- Commit 0f7271a: FOUND
