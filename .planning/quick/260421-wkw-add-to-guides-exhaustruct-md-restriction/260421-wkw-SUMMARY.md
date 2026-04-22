---
phase: quick
plan: 260421-wkw
subsystem: guides
tags: [docs, exhaustruct, functional-options]
dependency_graph:
  requires: []
  provides: [exhaustruct-guide-restriction]
  affects: [guides/exhaustruct.md]
tech_stack:
  added: []
  patterns: [functional-options]
key_files:
  created: []
  modified:
    - guides/exhaustruct.md
decisions:
  - Added restriction note against .golangci.yml suppression in <instructions> section
  - Added functional options pattern recommendation as new <recommendation> section
metrics:
  duration: 53s
  completed: "2026-04-21T20:30:39Z"
  tasks: 1
  files: 1
---

# Quick Task 260421-wkw: Add exhaustruct guide restriction Summary

Added a restriction warning and functional options pattern recommendation to the exhaustruct guide to prevent agents from suppressing diagnostics via config and guide them toward idiomatic Go solutions.

## Tasks Completed

| Task | Name | Commit | Files |
|------|------|--------|-------|
| 1 | Add restriction and functional options recommendation | 62bfa3c | guides/exhaustruct.md |

## Changes Made

### guides/exhaustruct.md

- **Restriction added** in `<instructions>`: "Do not suppress exhaustruct diagnostics by adding struct types to the exclusion list in `.golangci.yml`"
- **New `<recommendation>` section** between `<examples>` and `<patterns>`: Functional Options Pattern with complete Go code example showing `Option` type, `WithTimeout`/`WithLogger` option constructors, and `NewServer` constructor with defaults

## Verification

```bash
grep -q "Do not suppress" guides/exhaustruct.md        # ✅ Found
grep -q "functional options" guides/exhaustruct.md      # ✅ Found
grep -q "WithTimeout" guides/exhaustruct.md             # ✅ Found
```

Section order confirmed: instructions → examples → recommendation → patterns → related

## Deviations from Plan

None — plan executed exactly as written.

## Self-Check: PASSED

- [x] guides/exhaustruct.md exists and contains all additions
- [x] Commit 62bfa3c exists in git log
- [x] No unexpected file deletions
