#!/usr/bin/env bash
# test-push-e2e-artifacts.sh — Verify push-e2e-artifacts.sh meets all requirements
set -euo pipefail

SCRIPT="scripts/push-e2e-artifacts.sh"
PASS=0
FAIL=0

pass() { PASS=$((PASS + 1)); echo "  PASS: $1"; }
fail() { FAIL=$((FAIL + 1)); echo "  FAIL: $1"; }

echo "Testing: $SCRIPT"
echo ""

# --- Structural checks ---

# Must use git worktree
if grep -q "git worktree" "$SCRIPT"; then
    pass "uses git worktree"
else
    fail "uses git worktree"
fi

# Must NOT contain executable "git push" (no auto-push)
# Allow 'git push' inside echo/print strings, but not as an executable command
if grep -qE '^\s*git\s+push' "$SCRIPT"; then
    fail "no auto-push (contains executable 'git push')"
else
    pass "no auto-push"
fi

# Must NOT use broad staging
if grep -qE "git add \." "$SCRIPT"; then
    fail "no 'git add .' (broad staging)"
else
    pass "no 'git add .'"
fi

if grep -qE 'git\s+(-C\s+\S+\s+)?add\s+-A(\s+\.|\s*$)' "$SCRIPT"; then
    fail "no 'git add -A' with broad staging (adds everything)"
else
    pass "no 'git add -A' with broad staging"
fi

# Must NOT have uncommitted-changes guard
if grep -q "uncommitted" "$SCRIPT"; then
    fail "no uncommitted-changes guard"
else
    pass "no uncommitted-changes guard"
fi

# Must NOT use git rm -rf in the main path (only in orphan creation inside worktree)
# We check: no `git rm -rf .` outside a worktree context
RM_COUNT=$(grep -cE 'git\s+(-C\s+\S+\s+)?rm\s+-rf' "$SCRIPT" || true)
if [ "$RM_COUNT" -le 1 ]; then
    pass "git rm -rf only in orphan creation (count: $RM_COUNT)"
else
    fail "too many git rm -rf calls (count: $RM_COUNT)"
fi

# Must echo progress at each step
if grep -q "Collecting" "$SCRIPT"; then
    pass "has 'Collecting' progress echo"
else
    fail "has 'Collecting' progress echo"
fi

if grep -qi "worktree" "$SCRIPT"; then
    pass "has 'worktree' in output/echo"
else
    fail "has 'worktree' in output/echo"
fi

# Must have commit message format
if grep -q "e2e-artifacts:" "$SCRIPT"; then
    pass "has commit message format 'e2e-artifacts:'"
else
    fail "has commit message format"
fi

# Must check ndjson_analysis diagrams
if grep -q "ndjson_analysis" "$SCRIPT"; then
    pass "checks ndjson_analysis diagrams"
else
    fail "checks ndjson_analysis diagrams"
fi

# Must use DEST_DIR variable for explicit paths
if grep -q "DEST_DIR" "$SCRIPT"; then
    pass "uses DEST_DIR variable"
else
    fail "uses DEST_DIR variable"
fi

# Must have cleanup trap
if grep -q "trap" "$SCRIPT"; then
    pass "has trap for cleanup"
else
    fail "has trap for cleanup"
fi

# Must use mktemp for worktree location
if grep -q "mktemp" "$SCRIPT"; then
    pass "uses mktemp for worktree"
else
    fail "uses mktemp for worktree"
fi

# Must NOT use git checkout for branch switching (worktree replaces that)
CHECKOUT_COUNT=$(grep -cE 'git\s+(-C\s+\S+\s+)?checkout' "$SCRIPT" || true)
if [ "$CHECKOUT_COUNT" -le 1 ]; then
    pass "git checkout only in orphan creation (count: $CHECKOUT_COUNT)"
else
    fail "too many git checkout calls — found $CHECKOUT_COUNT instances"
fi

# Must NOT save/restore current branch (not needed with worktree)
if grep -q "CURRENT.*git branch" "$SCRIPT"; then
    fail "does not save current branch (not needed with worktree)"
else
    pass "does not save current branch"
fi

# Script must parse without syntax errors
if bash -n "$SCRIPT" 2>/dev/null; then
    pass "bash syntax check passes"
else
    fail "bash syntax check passes"
fi

# --- Summary ---
echo ""
echo "Results: $PASS passed, $FAIL failed"
if [ "$FAIL" -gt 0 ]; then
    echo "STATUS: RED (tests fail)"
    exit 1
else
    echo "STATUS: GREEN (all tests pass)"
    exit 0
fi
