---
phase: 37
slug: audit-and-improve-project-prompts-with-llm-prompting-best-pr
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 37 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — documentation-only phase (prompt quality) |
| **Config file** | none |
| **Quick run command** | `grep -c "Style\|Role\|exemplar" skills/golangci-lint-guide/SKILL.md` |
| **Full suite command** | N/A |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Verify file existence with `test -f skills/golangci-lint-guide/SKILL.md` or `test -f rules/claude-code.md`
- **After every plan wave:** Verify structural checks pass (Role sections, golangci_lint_parse references)
- **Before `/gsd-verify-work`:** Structural checks pass + golangci-lint runs clean on project
- **Max feedback latency:** ~1 second

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 37-01-T1 | 01 | 1 | — | T-37-01 | Prompt file contains only instructional text — no PII, no credentials | manual-only, structural | `grep -q "## Style" skills/golangci-lint-guide/SKILL.md` | ✅ | ⬜ pending |
| 37-02-T1 | 02 | 1 | — | T-37-02 | Rules files load into every agent session — must not break existing workflows | manual-only, structural | `grep -q "## Role" rules/claude-code.md` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| SKILL.md applies all 11 prompting techniques from Phase 36 research | — | Prompt technique application quality is qualitative — semantic evaluation of whether enriched role, style, exemplars, constraint reinforcement, etc. are effective | Compare SKILL.md against Phase 36 RESEARCH.md technique list; verify each technique is present and correctly applied |
| Agent correctly interprets improved SKILL.md prompt | — | Agent behavior depends on runtime (opencode, Claude Code, Cursor) — cannot test interpretation in isolation | Run `/golangci-lint` in opencode session on a Go project; verify agent follows the 7-step workflow with improved guidance |
| Rules files produce consistent agent behavior across platforms | — | Rules files are consumed by different AI platforms with different prompt parsing | Run golangci-lint workflow on each platform (opencode, Claude Code, Cursor); verify consistent behavior for JSON output flag injection and MCP tool usage |
| Rules source templates match deployed copies | — | Content consistency between `rules/` and `.xxx/rules/` directories | `diff rules/claude-code.md .claude/rules/golangci-lint.md` — should show no differences |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — documentation phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: prompt quality is qualitative, agent interpretation depends on runtime, neither has automated test coverage

**Approval:** pending (retroactive validation — phase already completed)
