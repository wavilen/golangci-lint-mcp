---
phase: 38-the-fix-pattern-to-apply-but-some-rules-contains-simple-prob
plan: 08
subsystem: guides
tags: [gosec, security, patterns, imperative-verbs]

# Dependency graph
requires:
  - phase: 38
    provides: decision D-01 that all pattern bullets must start with imperative verbs
provides:
  - 61 gosec guides with imperative-first pattern bullets providing security remediation
affects: [gosec, security-guides]

# Tech tracking
tech-stack:
  added: []
  patterns: [imperative-verb-first patterns, security-remediation-focused bullets]

key-files:
  created: []
  modified:
    - guides/gosec/G101.md
    - guides/gosec/G102.md
    - guides/gosec/G103.md
    - guides/gosec/G104.md
    - guides/gosec/G105.md
    - guides/gosec/G106.md
    - guides/gosec/G107.md
    - guides/gosec/G108.md
    - guides/gosec/G109.md
    - guides/gosec/G110.md
    - guides/gosec/G111.md
    - guides/gosec/G112.md
    - guides/gosec/G113.md
    - guides/gosec/G114.md
    - guides/gosec/G115.md
    - guides/gosec/G116.md
    - guides/gosec/G117.md
    - guides/gosec/G118.md
    - guides/gosec/G119.md
    - guides/gosec/G120.md
    - guides/gosec/G121.md
    - guides/gosec/G122.md
    - guides/gosec/G123.md
    - guides/gosec/G124.md
    - guides/gosec/G201.md
    - guides/gosec/G202.md
    - guides/gosec/G203.md
    - guides/gosec/G204.md
    - guides/gosec/G301.md
    - guides/gosec/G302.md
    - guides/gosec/G303.md
    - guides/gosec/G304.md
    - guides/gosec/G305.md
    - guides/gosec/G306.md
    - guides/gosec/G307.md
    - guides/gosec/G401.md
    - guides/gosec/G402.md
    - guides/gosec/G403.md
    - guides/gosec/G404.md
    - guides/gosec/G405.md
    - guides/gosec/G406.md
    - guides/gosec/G407.md
    - guides/gosec/G408.md
    - guides/gosec/G501.md
    - guides/gosec/G502.md
    - guides/gosec/G503.md
    - guides/gosec/G504.md
    - guides/gosec/G505.md
    - guides/gosec/G506.md
    - guides/gosec/G507.md
    - guides/gosec/G601.md
    - guides/gosec/G602.md
    - guides/gosec/G701.md
    - guides/gosec/G702.md
    - guides/gosec/G703.md
    - guides/gosec/G704.md
    - guides/gosec/G705.md
    - guides/gosec/G706.md
    - guides/gosec/G707.md
    - guides/gosec/G708.md
    - guides/gosec/G709.md

key-decisions:
  - "Used security-specific remediation verbs for gosec: Replace crypto, Remove insecure patterns, Validate input, Restrict access"
  - "Preserved inline code examples in pattern bullets for concrete fix direction"

patterns-established:
  - "Security linter patterns emphasize the correct fix (e.g., 'use crypto/rand' not just 'don't use math/rand')"
  - "Each gosec pattern bullet includes the specific recommended alternative or mitigation"

requirements-completed: []

# Metrics
duration: 27min
completed: 2026-04-21
---

# Phase 38 Plan 08: Gosec Pattern Bullets Rewrite Summary

**Rewrote all 204 pattern bullets across 61 gosec guides from problem-descriptive to imperative-verb-first with security remediation direction**

## Performance

- **Duration:** 27 min
- **Started:** 2026-04-21T06:09:15Z
- **Completed:** 2026-04-21T06:36:35Z
- **Tasks:** 2
- **Files modified:** 61

## Accomplishments

- Audited all 61 gosec guides — found 100% failure rate (all 204 bullets were problem-descriptive)
- Rewrote every pattern bullet to start with imperative verbs providing security remediation
- Verified 0 failures across all 61 guides with automated imperative-verb check
- Confirmed all guides remain under 500-word limit

## Task Commits

Each task was committed atomically:

1. **task 1: Audit gosec guides** — audit-only task (report in tmp/, not committed due to .gitignore)
2. **task 2: Rewrite non-imperative pattern bullets** - `8b2fb01` (fix)

**Plan metadata:** pending (docs: complete plan)

## Files Created/Modified

- `guides/gosec/G101.md` through `guides/gosec/G124.md` — Credentials, network, unsafe, errors, SSH, SSRF, profiling, overflow, decompression, filesystem, HTTP, CORS, TOCTOU, TLS, cookies
- `guides/gosec/G201.md` through `guides/gosec/G204.md` — SQL injection patterns
- `guides/gosec/G301.md` through `guides/gosec/G307.md` — File/directory permissions
- `guides/gosec/G401.md` through `guides/gosec/G408.md` — Weak crypto, TLS, RSA, random, ciphers, nonces, SSH key caching
- `guides/gosec/G501.md` through `guides/gosec/G507.md` — Deprecated imports (MD5, DES, RC4, CGI, SHA1, MD4, RIPEMD-160)
- `guides/gosec/G601.md` through `guides/gosec/G602.md` — Range variable capture, slice bounds
- `guides/gosec/G701.md` through `guides/gosec/G709.md` — Injection patterns (SQL, command, path, SSRF, XSS, log, email, template, deserialization)

## Transformation Examples

| Before (problem-descriptive) | After (imperative + remediation) |
|---|---|
| Hardcoded string values assigned to variables named password, secret, token | Avoid hardcoding credentials in variable names like `password`, `secret`, `token` — use environment variables or a secrets manager |
| `fmt.Sprintf` used to build SQL query strings with user input | Use parameterized queries (`db.QueryContext(ctx, "SELECT ... WHERE id = $1", id)`) instead of `fmt.Sprintf` for SQL |
| `math/rand` used to generate session tokens, API keys, or passwords | Use `crypto/rand` for session tokens, API keys, and passwords — `math/rand` is not cryptographically secure |
| Cookies missing the `Secure` flag, allowing transmission over HTTP | Set the `Secure` flag on all cookies to ensure they are only transmitted over HTTPS |

## Verification

- **Imperative check:** 61/61 guides pass, 0 failures
- **Word count check:** All guides under 500 words
- **Spot-check:** G101, G102, G201, G401, G601 all confirmed correct

## Deviations from Plan

### Auto-fixed Issues

**1. [Rule 3 - Blocking] Initial verification showed 19 of 61 guides with non-approved verbs**
- **Found during:** task 2 verification pass
- **Issue:** Used verbs not in the approved list (Strip, Generate, Migrate, Capture, Create, Store, Reject, Resolve, Encode, Escape, Unmarshal, Clean, Open)
- **Fix:** Replaced each non-approved verb with an approved alternative (Remove, Use, Replace, Add, Ensure, Avoid, Check, Switch, Set)
- **Files modified:** G105, G116, G118, G122, G302, G403, G405, G406, G407, G408, G501, G502, G504, G506, G601, G703, G705, G707, G709
- **Commit:** included in 8b2fb01

None otherwise — plan executed exactly as written.

## Self-Check: PASSED

- SUMMARY.md: FOUND
- Commit 8b2fb01: FOUND
- Gosec files in commit: 61 (all modified)
