---
phase: 36
slug: prepare-commands-golangci-lint-md-with-llm-prompting-best-pr
status: compliant
nyquist_compliant: false
wave_0_complete: true
created: 2026-04-22
---

# Phase 36 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | N/A — documentation-only phase |
| **Config file** | none |
| **Quick run command** | `grep -c "Strategy" commands/golangci-lint.md` |
| **Full suite command** | N/A |
| **Estimated runtime** | ~1 second |

---

## Sampling Rate

- **After every task commit:** Verify file structure with `grep -c "Strategy" commands/golangci-lint.md`
- **After every plan wave:** Run `golangci-lint run ./...` to confirm project still clean
- **Before `/gsd-verify-work`:** Structural checks pass
- **Max feedback latency:** ~5 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 36-01-T1 | 01 | 1 | — | T-36-01 / T-36-02 | Command file is version-controlled; no secrets or PII | manual-only | — | ✅ | ⬜ pending |
| 36-01-T2 | 01 | 1 | — | — | N/A — validation task | manual-only | — | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. No test stubs needed — phase is documentation-only with manual verification of prompt quality.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Prompt quality — 11 prompting techniques correctly applied | — | Semantic evaluation of LLM prompting best practices; cannot automate whether techniques improve agent behavior | Compare commands/golangci-lint.md against Phase 36 RESEARCH.md technique list; verify each technique is meaningfully applied |
| Agent correctly interprets XML-structured output format | — | Agent behavior depends on runtime interpretation of command template | Run `/golangci-lint` in opencode/crush session; verify agent follows Strategy A/B with correct output formatting |
| Exemplars are realistic and helpful | — | Quality of examples is subjective; automated checks can only verify presence not usefulness | Review `<example>` blocks for realism (actual linter names, correct MCP tool call syntax, appropriate fix patterns) |
| Error recovery paths cover all relevant failures | — | Completeness of error recovery is a design judgment | Trace each step in the command and verify failure modes have explicit recovery instructions |

---

## Validation Sign-Off

- [x] All tasks have automated verify or Wave 0 dependencies
- [x] Sampling continuity: N/A (2 tasks, both manual-only — documentation phase)
- [x] Wave 0 covers all references (none needed — retroactive validation)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter — **not set**: prompt rewrite quality is semantic; LLM agent behavior interpretation cannot be automated; 11 prompting techniques are quality judgments, not binary pass/fail

**Approval:** pending (retroactive validation — phase already completed)
