---
phase: 43
slug: make-accent-in-skill-and-rule-that-golangci-lint-must-be-run
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 43 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — documentation-only phase |
| **Config file** | none |
| **Quick run command** | `grep -c "per-package" skills/golangci-lint-guide/SKILL.md` |
| **Full suite command** | N/A |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Verify structural patterns in updated documentation files
- **After every plan wave:** Verify all 5 files contain "PREFER PER-PACKAGE RUNS" or "per-package" guidance
- **Before `/gsd-verify-work`:** Structural checks pass + golangci-lint runs clean
- **Max feedback latency:** ~1 second

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 43-01-T1 | 01 | 1 | — | T-43-01 | Documentation files contain only instructional text — no PII, no credentials | manual-only, structural | `grep -q "PREFER PER-PACKAGE RUNS" skills/golangci-lint-guide/SKILL.md && grep -q "per-package" rules/claude-code.md` | ✅ | ⬜ pending |
| 43-01-T2 | 01 | 1 | — | T-43-01 | Same — documentation only | manual-only, structural | `grep -q "jq" skills/golangci-lint-guide/SKILL.md && grep -q "30 issues" rules/opencode.md` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No new test stubs needed — phase is completed and validation is retroactive.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Agent follows per-package guidance — runs on individual packages for initial diagnosis | — | Agent behavior depends on runtime interpretation of documentation | Run `/golangci-lint` on a multi-package project; verify agent uses per-package path (e.g., `./pkg/auth/...`) instead of `./...` for initial runs |
| Agent correctly uses jq/python workflow for large output (>30 issues) | — | jq/python extraction workflow quality is semantic — correct usage depends on agent's tool capability | Run golangci-lint on a project with >30 issues; verify agent extracts unique pairs before calling MCP tools |
| Documentation quality — jq/python commands are correct and copy-pasteable | — | Command correctness requires execution in real environment | Copy jq and python3 commands from SKILL.md; run against actual golangci-lint JSON output; verify correct formatting |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — documentation phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: documentation content quality is semantic; agent behavior depends on runtime; jq/python command correctness requires manual execution verification

**Approval:** pending (retroactive validation — phase already completed)
