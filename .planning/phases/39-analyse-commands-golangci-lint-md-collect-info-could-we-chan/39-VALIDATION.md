---
phase: 39
slug: analyse-commands-golangci-lint-md-collect-info-could-we-chan
status: compliant
nyquist_compliant: true
wave_0_complete: true
created: 2026-04-22
---

# Phase 39 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | Go testing (stdlib) |
| **Config file** | none |
| **Quick run command** | `go test ./internal/server/... -run TestSummary -v` |
| **Full suite command** | `go test ./internal/server/...` |
| **Estimated runtime** | ~5 seconds |

---

## Sampling Rate

- **After every task commit:** Run `go test ./internal/server/... -run TestSummary -v`
- **After every plan wave:** Run `go test ./internal/server/...`
- **Before `/gsd-verify-work`:** Full suite green
- **Max feedback latency:** ~5 seconds

---

## Per-task Verification Map

| task ID | Plan | Wave | Requirement | Threat Ref | Secure Behavior | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|------------|-----------------|-----------|-------------------|-------------|--------|
| 39-01-T1 | 01 | 1 | — | T-39-01 / T-39-02 | Summary generated server-side from validated data — no tampering vector | unit | `go test ./internal/server/... -run TestSummary -v` | ✅ | ✅ green |
| 39-01-T2 | 01 | 1 | — | T-39-01 | Step 4 references summary block instead of manual counting — reduces LLM error surface | structural | `grep -c "Summary" commands/golangci-lint.md` | ✅ | ✅ green |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

Existing infrastructure covers all phase requirements. Go test suite in `internal/server/parse_handler_test.go` provides 14 tests (11 existing + 3 new summary tests). No new test stubs needed.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Step 4 simplification produces correct agent behavior | — | Agent's interpretation of Strategy A/B from summary block depends on runtime | Run `/golangci-lint` on a project with >10 unique diagnostics; verify agent reads summary strategy and follows correct approach |
| Summary block format is clear and actionable for LLM agents | — | Formatting clarity is a quality judgment | Inspect `golangci_lint_parse` response format; verify Summary block is parseable by LLM |

---

## Validation Sign-Off

- [x] All tasks have automated verify
- [x] Sampling continuity: both tasks have automated coverage (unit + structural)
- [x] Wave 0 covers all references (Go test suite covers core logic)
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [x] `nyquist_compliant: true` — core logic has Go unit tests (14 tests in `internal/server/parse_handler_test.go` covering summary block computation, strategy thresholds, and linter breakdown)

**Approval:** pending (retroactive validation — phase already completed)
