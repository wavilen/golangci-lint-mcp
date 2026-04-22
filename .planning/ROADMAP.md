# Roadmap: golangci-lint-mcp

## Milestones

- ✅ **v1.0 MVP** — Phases 1-30 (shipped 2026-04-20)
- 🚧 **v1.1 Platform Hooks + Example Validation** — Phases 31-50 (in progress)

## Phases

<details>
<summary>✅ v1.0 MVP (Phases 1-30) — SHIPPED 2026-04-20</summary>

- [x] Phase 1: Server Core (2/2 plans) — completed 2026-04-17
- [x] Phase 2: Template & Validation (1/1 plan) — completed 2026-04-17
- [x] Phase 3: Simple Linters — Error & Correctness (1/1) — completed 2026-04-17
- [x] Phase 4: Simple Linters — Complexity & Quality (1/1) — completed 2026-04-17
- [x] Phase 5: Simple Linters — Style & Formatting (1/1) — completed 2026-04-17
- [x] Phase 6: Simple Linters — Perf, Testing & Remaining (1/1) — completed 2026-04-17
- [x] Phase 7: gocritic (3/3 plans) — completed 2026-04-17
- [x] Phase 8: staticcheck (3/3 plans) — completed 2026-04-17
- [x] Phase 9: revive (2/2 plans) — completed 2026-04-17
- [x] Phase 10: gosec (2/2 plans) — completed 2026-04-17
- [x] Phase 11: govet (1/1 plan) — completed 2026-04-17
- [x] Phase 12: Minor Compound Linters (1/1 plan) — completed 2026-04-17
- [x] Phase 13: Docs & Tracking Reconciliation (1/1) — completed 2026-04-18
- [x] Phase 14: Simple Linter Verification (1/1) — completed 2026-04-18
- [x] Phase 15: Compound Linter Verification (1/1) — completed 2026-04-18
- [x] Phase 15.1: Update rule — errors.Wrap (INSERTED) — completed 2026-04-18
- [x] Phase 15.2: Skill + README (INSERTED) — completed 2026-04-18
- [x] Phase 15.3: Gosec AI flag fix (INSERTED) — completed 2026-04-18
- [x] Phase 16: Add golangci_lint_parse tool (1/1) — completed 2026-04-18
- [x] Phase 17: SOLID principle guidance (1/1) — completed 2026-04-18
- [x] Phase 18: Fix README docs (1/1) — completed 2026-04-18
- [x] Phase 19: Fix golangci-lint v2 flag syntax (1/1) — completed 2026-04-19
- [x] Phase 20: Gosec AI env-var templates (1/1) — completed 2026-04-19
- [x] Phase 21: Run desloppify scan (1/1) — completed 2026-04-19
- [x] Phase 21.1: Desloppify scorecard fix (INSERTED) — completed 2026-04-19
- [x] Phase 22: fmt.Println → slog.Info (1/1) — completed 2026-04-19
- [x] Phase 23: Gosec AI per-package batching (1/1) — completed 2026-04-19
- [x] Phase 24: npm/npx packaging (1/1) — completed 2026-04-19
- [x] Phase 25: Fix desloppify JSON parsing (1/1) — completed 2026-04-19
- [x] Phase 26: Run desloppify + update badge (1/1) — completed 2026-04-19
- [x] Phase 27: Pre-publish cleanup (1/1) — completed 2026-04-19
- [x] Phase 28: Full desloppify scan (1/1) — completed 2026-04-19
- [x] Phase 29: Close Tech Debt (1/1) — completed 2026-04-19
- [x] Phase 30: Module rename + git remote (1/1) — completed 2026-04-19

</details>

### 🚧 v1.1 Platform Hooks + Example Validation (In Progress)

**Milestone Goal:** Make golangci-lint-mcp auto-activate in every major AI coding platform, and validate that all 630 guide examples are lint-clean.

- [x] **Phase 31: opencode Plugin & Command** — Before/after hooks + `/golangci-lint` command for opencode (completed 2026-04-20)
- [x] **Phase 32: Platform Rules Files** — Static rules for Claude Code, Cursor, and opencode fallback (completed 2026-04-20)
- [x] **Phase 33: Claude Code Hooks** — PostToolUse hook + settings config for Claude Code (completed 2026-04-20)
- [x] **Phase 34: Golden Config Pipeline** — Extract Good examples, fetch golden config, lint all 630 guides (completed 2026-04-20)
- [ ] **Phase 35: Fix Guide Violations** — Fix any guides whose Good examples fail golden config lint
- [ ] **Phase 36: Prepare commands/golangci-lint.md** — LLM prompting best practices referencing gsd code and arxiv 2406.06608 summary

## Phase Details

### Phase 31: opencode Plugin & Command
**Goal**: opencode users get automatic MCP guidance injection when running golangci-lint, plus a dedicated `/golangci-lint` command
**Depends on**: Phase 30 (v1.0 complete)
**Requirements**: OPLG-01, OPLG-02, OPLG-03, OPLG-04, OCMD-01
**Success Criteria** (what must be TRUE):
  1. User runs `golangci-lint run` via Bash tool in opencode and the `--output.json.path stdout` flag is automatically injected before execution
  2. After golangci-lint completes, MCP fix guidance for each diagnostic appears in the agent's context without manual tool calls
  3. User can run `/golangci-lint` command and get guided fix assistance with correct flags pre-configured
  4. Plugin installs via npm package or `.opencode/plugins/` placement without manual configuration
  5. When MCP server is unavailable, golangci-lint runs normally with no errors or broken behavior
**Plans**: 1 plan

Plans:
- [x] 31-01-PLAN.md — Create /golangci-lint custom command + extend npm installer

### Phase 31.2: smart flag replacement for golangci-lint output flags (INSERTED)

**Goal:** Fix naive `--output.json.path stdout` injection to intelligently strip all conflicting output-format flags (text, tab, html, json, checkstyle, code-climate, junit-xml, teamcity, sarif, show-stats, color, verbose, plus legacy flags) before injecting JSON output flag — across plugin, command template, and rules files
**Requirements**: TBD
**Depends on:** Phase 31
**Plans:** 1 plan

Plans:
- [ ] 31.2-01-PLAN.md — Smart flag replacement in plugin + update command template and rules files

### Phase 31.1: opencode Plugin Hooks (INSERTED)

**Goal:** opencode plugin with tool.execute.before/after hooks that auto-inject --output.json.path stdout flag and MCP fix guidance when golangci-lint is run — equivalent to Phase 33's Claude Code PostToolUse hook
**Requirements**: OPLG-01, OPLG-02, OPLG-03, OPLG-04
**Depends on:** Phase 31
**Plans:** 1 plan

Plans:
- [ ] 31.1-01-PLAN.md — Create opencode plugin with tool.execute hooks + extend npm installer

### Phase 32: Platform Rules Files
**Goal**: Every platform has instructional rules that guide agents toward MCP tools when working with golangci-lint
**Depends on**: Phase 30 (v1.0 complete)
**Requirements**: RULE-01, RULE-02, RULE-03
**Success Criteria** (what must be TRUE):
  1. Claude Code loads `.claude/rules/golangci-lint.md` and follows instructions to use MCP tools when running golangci-lint
  2. Cursor loads `.cursor/rules/golangci-lint.mdc` with `alwaysApply: true` and routes golangci-lint diagnostics to MCP tools
  3. opencode loads `.opencode/rules/golangci-lint.md` as fallback guidance when the plugin is not active
**Plans**: 1 plan

Plans:
- [x] 32-01-PLAN.md — Create platform rules files + extend npm installer

### Phase 33: Claude Code Hooks
**Goal**: Claude Code automatically injects MCP fix guidance after every golangci-lint run via PostToolUse hooks
**Depends on**: Phase 32 (rules files establish MCP wiring pattern)
**Requirements**: CCHOK-01, CCHOK-02, CCHOK-03
**Success Criteria** (what must be TRUE):
  1. User runs `golangci-lint` via Bash in Claude Code and MCP fix guidance for each diagnostic is injected into the conversation after execution
  2. `.claude/settings.json` contains `mcpServers` block wiring the MCP server and hook configuration that survives Claude Code restarts
  3. Hook always exits 0 — golangci-lint output is never blocked, even when MCP server is down or hook encounters an error
**Plans**: 1 plan

Plans:
- [x] 33-01-PLAN.md — Create PostToolUse hook script + extend npm installer for hooks/settings/MCP config

### Phase 34: Golden Config Pipeline
**Goal**: All 630 guide Good examples can be linted against maratori's golden config and violations are reported
**Depends on**: Phase 30 (v1.0 complete — guides must exist)
**Requirements**: XCHK-01, XCHK-02, XCHK-03, XCHK-04
**Success Criteria** (what must be TRUE):
  1. Pipeline fetches maratori's golden config at a pinned tag and enables all recommended linters
  2. All 630 guides have their Good code blocks extracted to a tmp gitignored subdirectory as compilable Go files with inferred imports
  3. `make crosscheck` (or equivalent) runs golangci-lint with golden config on all extracted examples and reports a pass/fail summary with guide filenames and linter details
  4. Bad code blocks are excluded — only Good examples are linted
**Plans**: 2 plans

Plans:
- [x] 34-01-PLAN.md — Vendor golden config + update-golden-config Makefile target
- [x] 34-02-PLAN.md — Build crosscheck pipeline + Makefile crosscheck target + validate

### Phase 35: Fix Guide Violations
**Goal**: All 630 guide Good examples pass the golden config lint with zero violations
**Depends on**: Phase 34 (pipeline must identify violations first)
**Requirements**: XCHK-05
**Success Criteria** (what must be TRUE):
  1. Every guide's Good example code passes golden config lint with zero violations
  2. Re-running the crosscheck pipeline reports 630/630 guides passing
**Plans**: 1 plan

Plans:
- [ ] 35-01-PLAN.md — Fix 2 failing guide Good examples + verify crosscheck pipeline passes

## Progress

**Execution Order:**
Phases execute in numeric order: 31 → 32 → 33 → 34 → 35 → 36
(Phases 31 and 32 are independent of each other; 33 depends on 32; 34 is independent of 31-33; 35 depends on 34; 36 depends on 35)

| Phase | Milestone | Plans Complete | Status | Completed |
|-------|-----------|----------------|--------|-----------|
| 31. opencode Plugin & Command | v1.1 | 1/1 | Complete    | 2026-04-20 |
| 32. Platform Rules Files | v1.1 | 1/1 | Complete    | 2026-04-20 |
| 33. Claude Code Hooks | v1.1 | 1/1 | Complete    | 2026-04-20 |
| 34. Golden Config Pipeline | v1.1 | 2/2 | Complete    | 2026-04-20 |
| 35. Fix Guide Violations | v1.1 | 0/1 | Planned | - |
| 36. Prepare commands/golangci-lint.md | v1.1 | 0/1 | Planned | - |
| 38. Fix problem-description patterns | v1.1 | 10/10 | Complete    | 2026-04-21 |
| 40. Strip pipe filters in plugin | v1.1 | 1/1 | Complete    | 2026-04-21 |
| 41. npx installer full setup | v1.1 | 1/1 | Complete    | 2026-04-21 |
| 42. opencode plugin filters | v1.1 | 1/1 | Complete    | 2026-04-21 |
| 43. per-package + jq/python guidance | v1.1 | 1/1 | Complete    | 2026-04-21 |
| 46. Verification & Traceability Closure | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 47. Fix Remaining Pattern Bullet Violations | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 48. Fix Golden Config Sed Bug | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 49. Hook Detection & Deduplication | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 50. Nyquist Validation for All Phases | v1.1 | 2/2 | Complete    | 2026-04-22 |
| 51. Deployed File Sync | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 52. Makefile Install Resources | v1.1 | 1/1 | Complete    | 2026-04-22 |
| 53. Fix Remaining Pattern Bullets | v1.1 | 2/2 | Complete    | 2026-04-22 |
| 54. Nyquist Validation Completion | v1.1 | 1/1 | Complete   | 2026-04-22 |

**Goal:** Improve commands/golangci-lint.md by incorporating empirically-validated LLM prompting techniques from arxiv 2406.06608 (The Prompt Report) and GSD code patterns — making the command more reliable, less ambiguous, and more effective when executed by LLM agents
**Requirements**: TBD
**Depends on:** Phase 35
**Plans:** 2/2 plans complete

Plans:
- [ ] 36-01-PLAN.md — Rewrite command with prompting best practices + validate

### Phase 37: Audit and improve project prompts with LLM prompting best practices from phase 36 research excluding linter guides

**Goal:** Apply 11 LLM prompting techniques from Phase 36 research (arxiv 2406.06608 + Anthropic best practices) to SKILL.md and all 6 platform rules files — enriching role definitions, adding exemplars, constraint reinforcement, error recovery, self-checks, and proactive action defaults while preserving all existing MCP tool documentation
**Requirements**: TBD
**Depends on:** Phase 36
**Plans:** 2/2 plans complete

Plans:
- [x] 37-01-PLAN.md — Improve SKILL.md with prompting best practices
- [x] 37-02-PLAN.md — Improve all 6 rules files with prompting best practices

### Phase 38: Fix problem-description patterns across all 630 guides

**Goal:** Rewrite all `<patterns>` bullets across 630 guides to start with imperative verbs providing actionable fix direction instead of merely describing what triggers the linter
**Requirements**: TBD
**Depends on:** Phase 37
**Plans:** 10/10 plans complete

Plans:
- [x] 38-01-PLAN.md — Simple linters: Error & Correctness (16 guides)
- [x] 38-02-PLAN.md — Simple linters: Complexity & Quality (17 guides)
- [x] 38-03-PLAN.md — Simple linters: Style & Formatting (28 guides)
- [x] 38-04-PLAN.md — Simple linters: Perf, Testing & Remaining (39 guides)
- [x] 38-05-PLAN.md — gocritic compound linter (108 guides)
- [x] 38-06-PLAN.md — staticcheck compound linter (172 guides)
- [x] 38-07-PLAN.md — revive compound linter (101 guides)
- [x] 38-08-PLAN.md — gosec compound linter (61 guides)
- [x] 38-09-PLAN.md — govet compound linter (35 guides)
- [x] 38-10-PLAN.md — Minor compound linters (49 guides)

### Phase 39: analyse commands/golangci-lint.md - collect info - could we change Step 4 from the textual description to programmatic parsing inside of the golangci_lint_parse tool call?

**Goal:** Add a diagnostic summary block to `golangci_lint_parse` response and simplify Step 4 of commands/golangci-lint.md to reference it — eliminating manual diagnostic counting by the LLM agent
**Requirements**: TBD
**Depends on:** Phase 38
**Plans:** 1/1 plans complete

Plans:
- [x] 39-01-PLAN.md — Add summary block to parse handler + simplify Step 4 in command template

### Phase 40: plugins/golangci-lint.js should also react to bash commands with golangci-lint where output handled by head, tail, grep commands - if they appear just drop it and run command without these filters

**Goal:** Add pipe/filter stripping to the opencode plugin's `tool.execute.before` hook so that golangci-lint commands piped through `head`, `tail`, `grep`, and similar filter utilities are executed WITHOUT those filters
**Requirements**: TBD
**Depends on:** Phase 39
**Plans:** 1/1 plans complete

Plans:
- [x] 40-01-PLAN.md — Add stripOutputFilters function + integrate into tool.execute.before hook

### Phase 41: npx golangci-lint-guide should install opencode plugin, skill and other resources needed to start using the MCP server

**Goal:** Extend bin/install.js so that `npx golangci-lint-guide` installs the opencode plugin, merges MCP server config into opencode.json, and checks for required binaries — completing the one-command setup experience
**Requirements**: TBD
**Depends on:** Phase 40
**Plans:** 1/1 plans complete

Plans:
- [x] 41-01-PLAN.md — Extend npm installer with opencode plugin + MCP config + binary checks

### Phase 42: opencode plugin filters golangci-lint command too widely, it must check specifically command (first in the chain or just a command), it false-positively reacts on any command with golangci-lint text in any part (phase description)

**Goal:** Fix the opencode plugin's command detection to match only actual golangci-lint invocations (first meaningful token), not any Bash command containing the string "golangci-lint" anywhere
**Requirements**: TBD
**Depends on:** Phase 41
**Plans:** 1/1 plans complete

Plans:
- [x] 42-01-PLAN.md — Fix command detection with isGolangciLintCommand regex + add test suite

### Phase 43: make accent in skill and rule that golangci-lint must be run on individual packages, not whole project. and use jq and python to parse json output for size bigger than 30 items.

**Goal:** Update SKILL.md, commands/golangci-lint.md, and 3 rules files to emphasize per-package golangci-lint runs for initial diagnosis and add jq/python workflow for handling large JSON output (>30 issues)
**Requirements**: TBD
**Depends on:** Phase 42
**Plans:** 1/1 plans complete

Plans:
- [x] 43-01-PLAN.md — Update all 5 docs with per-package run guidance + jq/python large-output workflow

### Phase 44: add to Makefile command to install individual opencode resources during local development

**Goal:** [To be planned]
**Requirements**: TBD
**Depends on:** Phase 43
**Plans:** 0 plans

Plans:
- [ ] TBD (run /gsd-plan-phase 44 to break down)

### Phase 45: something prevent opencode plugins from detecting real golangci-lint calls, may be shell prefix or something, use opencode source code and calling as subprocess to figure out issues better, and plan fix

**Goal:** Diagnose and fix why the opencode plugin (plugins/golangci-lint.js) fails to detect and intercept real golangci-lint calls — add shell-wrapper detection, fix export format compatibility, add debug logging, and create Makefile install target
**Requirements**: TBD
**Depends on:** Phase 44
**Plans:** 1/1 plans complete

Plans:
- [x] 45-01-PLAN.md — Fix plugin export format + add shell-wrapper detection + Makefile install target

### 🔧 Gap Closure Phases (from v1.1 audit)

### Phase 46: Verification & Traceability Closure

**Goal:** Close all verification gaps by creating VERIFICATION.md for unverified phases and marking Phase 31.2 as superseded — formally closing XCHK-05 (partial)
**Requirements**: XCHK-05
**Gap Closure:** Closes requirement gap XCHK-05 (partial → verified), Phase 31.1 unverified, Phase 35 unverified, Phase 36 unverified, Phase 31.2 unstarted
**Depends on:** Phase 45
**Plans:** 1/1 plans complete

Plans:
- [x] 46-01-PLAN.md — Create VERIFICATION.md for phases 31.1, 35, 36 + mark 31.2 superseded + update REQUIREMENTS.md traceability

### Phase 47: Fix Remaining Pattern Bullet Violations

**Goal:** Fix the last remaining non-imperative pattern bullet (SA2003.md line 24) — 24 of 25 originally identified bullets were already fixed by commit cff1dad
**Requirements**: TBD
**Gap Closure:** Closes Phase 38 tech debt — final non-imperative bullet in SA2003.md
**Depends on:** Phase 46
**Plans:** 1/1 plans complete

Plans:
- [x] 47-01-PLAN.md — Rewrite SA2003.md non-imperative bullet + verify zero violations across all 626 guides

### Phase 48: Fix Golden Config Sed Bug

**Goal:** Fix the update-golden-config Makefile target sed pattern that fails to match `github.com/my/project` — ensuring golden config fetching works correctly
**Requirements**: TBD
**Gap Closure:** Closes Phase 34 tech debt — sed bug in update-golden-config target
**Depends on:** Phase 47
**Plans:** 1/1 plans complete

Plans:
- [x] 48-01-PLAN.md — Fix sed pattern to wildcard match + verify end-to-end

### Phase 49: Hook Detection & Deduplication

**Goal:** Improve Claude Code hook command detection from indexOf to isGolangciLintCommand, and extract shared nudge logic from ~100 duplicated lines between hook and plugin
**Requirements**: TBD
**Gap Closure:** Closes Phase 33 tech debt — imprecise detection + duplicated nudge logic
**Depends on:** Phase 48
**Plans:** 1/1 plans complete

Plans:
- [x] 49-01-PLAN.md — Extract shared nudge module + simplify hook + plugin + update installer

### Phase 50: Nyquist Validation for All Phases

**Goal:** Create VALIDATION.md for all 14 v1.1 phases (31-43) to satisfy Nyquist coverage requirements
**Requirements**: TBD
**Gap Closure:** Closes cross-phase tech debt — Nyquist validation MISSING for all 14 phases
**Depends on:** Phase 49
**Plans:** 2/2 plans complete

Plans:
- [x] 50-01-PLAN.md — Create VALIDATION.md for phases 31-36 (platform integration: 6 files, nyquist_compliant: false)
- [x] 50-02-PLAN.md — Create VALIDATION.md for phases 37-43 (content + code: 7 files, 2 nyquist_compliant: true)

### Phase 51: Deployed File Sync

**Goal:** Sync all deployed files in .claude/, .opencode/, and .cursor/ to match current source — fixing stale hook, missing Phase 43 content in rules, and missing shared/nudge.js
**Requirements:** CCHOK-01, RULE-01, RULE-02, RULE-03
**Gap Closure:** Closes integration gaps — 5 stale/missing deployed files, degraded Claude Code hook flow
**Depends on:** Phase 50
**Plans:** 1/1 plans complete

Plans:
- [x] 51-01-PLAN.md — Sync deployed files + add Makefile target for .claude/ deployment

### Phase 52: Makefile Install Resources

**Goal:** Complete unstarted Phase 44 — add Makefile targets to install individual opencode resources during local development
**Requirements:** TBD
**Gap Closure:** Closes Phase 44 (UNSTARTED) tech debt
**Depends on:** Phase 51
**Plans:** 1/1 plans complete

Plans:
- [x] 52-01-PLAN.md — Add Makefile install-resources target + verify end-to-end

### Phase 53: Fix Remaining Pattern Bullets

**Goal:** Add `<patterns>` and `<related>` sections to 3 gocritic guides (appendAssign, badCall, commentedOutCode) that are missing them — completing the 4-section structure for all 629 guides
**Requirements:** TBD
**Gap Closure:** Closes Phase 38 tech debt — 3 guides missing `<patterns>` and `<related>` sections; closes VERIFICATION.md gap for errcheck.md missing `<related>` section
**Depends on:** Phase 52
**Plans:** 2/2 plans complete

Plans:
- [x] 53-01-PLAN.md — Add patterns and related sections to appendAssign, badCall, commentedOutCode
- [x] 53-02-PLAN.md — Add <related> section to errcheck.md and update SUMMARY with accurate guide count (gap closure)

### Phase 54: Nyquist Validation Completion

**Goal:** Create VALIDATION.md for 5 phases that lack it (44, 45, 46, 48, 50) — achieving full Nyquist coverage for v1.1
**Requirements:** TBD
**Gap Closure:** Closes Nyquist coverage gap — 5 phases missing VALIDATION.md
**Depends on:** Phase 53
**Plans:** 1/1 plans complete

Plans:
- [x] 54-01-PLAN.md — Create VALIDATION.md for phases 44, 45, 46, 48, 50
