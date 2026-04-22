---
phase: 43-make-accent-in-skill-and-rule-that-golangci-lint-must-be-run
plan: 01
subsystem: documentation
tags: [golangci-lint, per-package, jq, python, mcp-tools, large-output]

# Dependency graph
requires:
  - phase: 37
    provides: SKILL.md and rules files established during skill improvement
  - phase: 39
    provides: commands/golangci-lint.md with adaptive guidance depth
provides:
  - Per-package run guidance in SKILL.md and commands/golangci-lint.md
  - jq/python large-output (>30 issues) extraction workflow in SKILL.md Step 3.5
  - Large-output handling in commands/golangci-lint.md Step 3
  - Brief large-output notes in all 3 rules files (step 4)
  - JSON dump file handling note in SKILL.md Notes
affects: [documentation, agent-behavior, linting-workflow]

# Tech tracking
tech-stack:
  added: [jq, python3]
  patterns: [per-package-linting, large-output-extraction]

key-files:
  created: []
  modified:
    - skills/golangci-lint-guide/SKILL.md
    - commands/golangci-lint.md
    - rules/opencode.md
    - rules/claude-code.md
    - rules/cursor.mdc

key-decisions:
  - "PREFER PER-PACKAGE RUNS warning placed before code blocks in SKILL.md and commands"
  - "30-issue threshold triggers jq/python extraction instead of golangci_lint_parse"
  - "JSON dump files must always be processed through MCP enrichment before code modifications"
  - "Rules files get concise step 4 note; SKILL.md and commands get detailed workflows"

patterns-established:
  - "Per-package initial runs with ./... only for final verification"
  - "jq/python extraction for large diagnostic output (>30 issues)"

requirements-completed: []

# Metrics
duration: 3min
completed: 2026-04-21
---

# Phase 43 Plan 01: Per-Package and Large-Output Guidance Summary

**Per-package golangci-lint run emphasis and jq/python large-output (>30 issues) extraction workflow added to all 5 documentation files**

## Performance

- **Duration:** 3 min
- **Started:** 2026-04-21T18:46:15Z
- **Completed:** 2026-04-21T18:49:15Z
- **Tasks:** 2
- **Files modified:** 5

## Accomplishments
- Added PREFER PER-PACKAGE RUNS warnings with callout boxes to SKILL.md Step 2 and commands/golangci-lint.md Step 1
- Inserted new SKILL.md Step 3.5 with full jq/python large-output extraction workflow (check count, extract pairs, format summary, call golangci_lint_guide per pair)
- Added large-output handling block to commands/golangci-lint.md Step 3
- Added brief step 4 large-output note to all 3 rules files (opencode.md, claude-code.md, cursor.mdc)
- Added JSON Dump File Handling subsection to SKILL.md Notes

## Task Commits

Each task was committed atomically:

1. **task 1: Update SKILL.md and commands with per-package run emphasis** - `dc90c66` (feat)
2. **task 2: Add jq/python large-output workflow to SKILL.md, commands, and rules files** - `2b77f36` (feat)

## Files Created/Modified
- `skills/golangci-lint-guide/SKILL.md` - Per-package warning in Step 2, Step 3 example update, new Step 3.5 (jq/python workflow), Steps 4-7 renumbered to 5-8, JSON Dump File Handling in Notes
- `commands/golangci-lint.md` - Per-package default in Step 1, large-output handling block in Step 3
- `rules/opencode.md` - Step 4 large-output brief note
- `rules/claude-code.md` - Step 4 large-output brief note
- `rules/cursor.mdc` - Step 4 large-output brief note

## Decisions Made
- PREFER PER-PACKAGE RUNS warning placed before code blocks (not inline) for maximum visibility
- Step 3.5 numbering used (instead of renumbering to Step 4) to preserve the semantic flow from Step 3's MCP parsing into large-output handling
- Rules files get concise single-step note rather than detailed workflow — proportionate to their 17-22 line length

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

None

## User Setup Required

None - no external service configuration required.

## Next Phase Readiness
- All 5 documentation files updated with per-package and large-output guidance
- No blockers or concerns

## Self-Check: PASSED

- All 5 modified files exist on disk
- SUMMARY.md exists at expected path
- Both task commits (dc90c66, 2b77f36) found in git log

---
*Phase: 43-make-accent-in-skill-and-rule-that-golangci-lint-must-be-run*
*Completed: 2026-04-21*
