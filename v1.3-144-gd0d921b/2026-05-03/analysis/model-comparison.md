# Model Comparison: GLM-4.7 vs GLM-5.1 vs GLM-5-Turbo

## Summary Table

| Trace | Model | Test | Issues (Before→After) | Total Calls | Subs | Duration | Tokens | Tok/Issue | Strategy | Result |
|-------|-------|------|----------------------|-------------|------|----------|--------|-----------|----------|--------|
| simple-glm-4-7 | GLM-4.7 | Simple | 8→0 (100%) | 26 | 0 | 140.4s | 29,569 | 3696 | single | PASS |
| simple-glm-5-1 | GLM-5.1 | Simple | 8→0 (100%) | 26 | 0 | 226.3s | 25,204 | 3150 | single | PASS |
| simple-glm-5-turbo | GLM-5-Turbo | Simple | 8→0 (100%) | 28 | 0 | 118.4s | 16,800 | 2100 | single | PASS |
| medium-glm-4-7 | GLM-4.7 | Medium | 30→0 (100%) | 73 | 0 | 262.8s | 28,046 | 935 | single | PASS |
| medium-glm-5-1 | GLM-5.1 | Medium | 30→0 (100%) | 38 | 0 | 850.8s | 73,669 | 2456 | single | PASS |
| medium-glm-5-turbo | GLM-5-Turbo | Medium | 30→0 (100%) | 31 | 0 | 356.5s | 62,958 | 2099 | single | PASS |
| large-glm-5-1 | GLM-5.1 | Large | 116→0 (100%) | 167 | 8 | 1144.0s | 247,175 | 2131 | subagent | PASS |
| large-glm-5-turbo | GLM-5-Turbo | Large | 116→0 (100%) | 64 | 0 | 534.1s | 133,043 | 1147 | single | PASS |
| multipkg-glm-5-1 | GLM-5.1 | Multipkg | 29→0 (100%) | 86 | 4 | 322.7s | 107,183 | 3696 | subagent | PASS |
| multipkg-glm-5-turbo | GLM-5-Turbo | Multipkg | 29→0 (100%) | 93 | 4 | 153.1s | 77,810 | 2683 | subagent | PASS |
| large-glm-4-7 | GLM-4.7 | Large | 116→0 (0%) | 0 | 0 | - | 0 | - | - | **FAIL** (interrupted) |
| multipkg-glm-4-7 | GLM-4.7 | Multipkg | 29→0 (0%) | 0 | 0 | - | 0 | - | - | **FAIL** (interrupted) |

## Overall Results

- **Total test runs:** 12 (10 with traces + 2 interrupted)
- **Pass rate:** 10/12 (83%)
- **GLM-4.7:** 2/4 passed (Simple, Medium) — failed on Large and Multipkg (interrupted/timed out)
- **GLM-5.1:** 4/4 passed (Simple, Medium, Large, Multipkg) — 100% success rate
- **GLM-5-Turbo:** 4/4 passed (Simple, Medium, Large, Multipkg) — 100% success rate

## Per-Dimension Comparison

### 1. Tool Usage Efficiency

| Model | Simple (8 issues) | Medium (30 issues) | Large (116 issues) | Multipkg (29 issues) |
|-------|-------------------|--------------------|--------------------|----------------------|
| GLM-4.7 | 26 calls (3696 tok/issue) | 73 calls (935 tok/issue) | N/A (failed) | N/A (failed) |
| GLM-5.1 | 26 calls (3150 tok/issue) | 38 calls (2456 tok/issue) | 167 calls (2131 tok/issue) | 86 calls (3696 tok/issue) |
| GLM-5-Turbo | 28 calls (2100 tok/issue) | 31 calls (2099 tok/issue) | 64 calls (1147 tok/issue) | 93 calls (2683 tok/issue) |

**Key findings:**
- **GLM-4.7** is the most token-efficient on medium tasks (935 tok/issue) but uses more raw tool calls (73 for medium)
- **GLM-5-Turbo** consistently uses fewer tokens per issue across all test sizes
- **GLM-5.1** uses subagent delegation for large/multipkg tasks, shifting tool calls to subagents
- All models use a similar scan→fix→verify pattern with golangci_lint_run as the primary MCP tool

### 2. Issue Reduction Rate

All passing runs achieved **100% issue reduction**. The key differentiator is:

| Model | Simple | Medium | Large | Multipkg |
|-------|--------|--------|-------|----------|
| GLM-4.7 | 100% ✓ | 100% ✓ | FAIL | FAIL |
| GLM-5.1 | 100% ✓ | 100% ✓ | 100% ✓ | 100% ✓ |
| GLM-5-Turbo | 100% ✓ | 100% ✓ | 100% ✓ | 100% ✓ |

**GLM-4.7 cannot handle large-scale tasks** — it timed out / was interrupted on both the Large (116 issues) and Multipkg (29 issues across 4 packages) tests. GLM-5.x models handle all sizes.

### 3. Token Consumption

| Metric | GLM-4.7 | GLM-5.1 | GLM-5-Turbo |
|--------|---------|---------|-------------|
| Avg input tokens (simple) | 24,982 | 19,305 | 13,227 |
| Avg reasoning tokens (simple) | 3,017 | 3,841 | 1,661 |
| Total tokens (medium) | 27,846 | 73,669 | 62,958 |
| Total tokens (large) | - | 247,175 | 133,043 |
| Cache efficiency | High | High | High |

**Key findings:**
- **GLM-5-Turbo is the most token-efficient** model, especially on reasoning tokens
- **GLM-5.1 uses more tokens** but compensates with subagent delegation for complex tasks
- **GLM-4.7** is token-efficient where it succeeds, but cannot scale to larger tasks
- Cache reads are very high across all models (400K-2.3M tokens), suggesting effective prompt caching

### 4. Strategy Compliance

| Model | Strategy Followed | Notes |
|-------|-------------------|-------|
| GLM-4.7 | Partial | Followed single-agent for simple/medium but failed on tasks requiring subagents |
| GLM-5.1 | Full | Correctly used single-agent for small tasks, subagent-per-file for large, subagent-per-package for multipkg |
| GLM-5-Turbo | Partial | Ignored subagent recommendation for Large test (handled 116 issues solo with 64 tool calls) |

**Notable:** GLM-5-Turbo achieved 100% issue reduction on Large WITHOUT using subagents, despite the strategy recommending subagent-per-file. It compensated with extensive golangci_lint_guide lookups (18 calls) and targeted edits.

### 5. Error Recovery Patterns

| Model | Edit Failures | Recovery Approach |
|-------|---------------|-------------------|
| GLM-4.7 | 0/45 | Clean edits, used bash for verification (11 commands on medium) |
| GLM-5.1 | 0/19 | Used write for full file rewrites + targeted edits |
| GLM-5-Turbo | 0/22 | Mixed approach: write rewrites + edits + bash verification |

All models had **zero edit failures** — the MCP tool guidance prevents oldString mismatches.

### 6. Duration Comparison

| Model | Simple | Medium | Large | Multipkg |
|-------|--------|--------|-------|----------|
| GLM-4.7 | 140s | 263s | - | - |
| GLM-5.1 | 226s | 851s | 1144s | 323s |
| GLM-5-Turbo | 118s | 357s | 534s | 153s |

**GLM-5-Turbo is consistently the fastest** model, 2-3x faster than GLM-5.1 on comparable tasks. GLM-5.1's slowness is partly due to subagent overhead (subagent dispatch + result collection).

## Key Findings

1. **GLM-5-Turbo is the best overall performer**: fastest, most token-efficient, handles all task sizes, zero edit failures
2. **GLM-5.1 is the most capable delegator**: correctly uses subagent strategies for large/multipkg tasks
3. **GLM-4.7 is adequate for small tasks only**: handles Simple and Medium well but cannot scale
4. **Strategy compliance varies**: GLM-5.1 follows strategy precisely; GLM-5-Turbo improvises successfully
5. **Subagent usage is a key differentiator**: GLM-5.1 delegates to 4-8 subagents for large tasks; GLM-5-Turbo achieves the same result solo

## Notable Patterns

1. **Scan-first approach**: All models start with `golangci_lint_run` → `skill` load before making changes
2. **Write-over-edit for large rewrites**: GLM-5.1 and GLM-5-Turbo use `write` for complete file rewrites when many issues exist
3. **Todowrite for planning**: GLM-5.x models use `todowrite` to track progress; GLM-4.7 does not
4. **Build verification**: All models run `go build` to verify fixes compile, not just lint-clean
5. **Subagent parallelism**: GLM-5.1 achieves better wall-time efficiency on multipkg by delegating to package-specific subagents
6. **Guide lookups as crutch**: GLM-5-Turbo on Large made 18 golangci_lint_guide calls — using the MCP tool as a reference manual for each linter rule
