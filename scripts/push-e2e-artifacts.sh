#!/usr/bin/env bash
# push-e2e-artifacts.sh — Commit E2E artifacts to e2e-artifacts branch via git worktree
# Uses worktree isolation: working directory is never touched.
# Requires ndjson-analysis to have already been run (blocks if diagrams missing).
# Commits locally — user pushes manually.
set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

# --- Version resolution ---
RAW_VERSION=$(git describe --tags --always 2>/dev/null | sed 's/^v//')
if [[ "$RAW_VERSION" =~ ^[0-9]+\.[0-9]+ ]]; then
    VERSION="$RAW_VERSION"
else
    VERSION="dev-${RAW_VERSION}"
fi
DATE=$(date +%Y-%m-%d)
DEST_DIR="v${VERSION}/${DATE}"

# --- Pre-flight: verify running from repo root ---
if [ "$(git rev-parse --show-toplevel)" != "$REPO_ROOT" ]; then
    echo "Error: Must run from repository root." >&2
    exit 1
fi

echo "=== Push E2E Artifacts ==="
echo "Version: ${VERSION}  Date: ${DATE}  Dest: ${DEST_DIR}"

# --- Pre-flight: detect artifact sources ---
HAS_RESULTS=false
HAS_REPORT=false
HAS_DIAGRAMS=false

if ls tmp/ndjson/*-result.json >/dev/null 2>&1; then HAS_RESULTS=true; fi
if [ -f "${REPO_ROOT}/tmp/ndjson/e2e-report.html" ]; then HAS_REPORT=true; fi
if ls tmp/ndjson_analysis/*.puml >/dev/null 2>&1; then HAS_DIAGRAMS=true; fi

if [ "$HAS_RESULTS" = "false" ] && [ "$HAS_REPORT" = "false" ]; then
    echo "Error: No artifact sources found. Run 'make integration-test' first." >&2
    exit 1
fi

if [ "$HAS_DIAGRAMS" = "false" ]; then
    echo "Error: No ndjson_analysis diagrams found in tmp/ndjson_analysis/." >&2
    echo "Run ndjson-analysis agent to generate diagrams first." >&2
    exit 1
fi

# --- Cleanup: remove worktree on any exit ---
WORKTREE=$(mktemp -d)
cleanup() {
    git worktree remove --force "$WORKTREE" 2>/dev/null || rm -rf "$WORKTREE"
}
trap cleanup EXIT

# --- Remove stale worktree if one exists ---
STALE=$(git worktree list --porcelain 2>/dev/null | grep -A1 "e2e-artifacts" | head -1 | cut -d' ' -f2 || true)
if [ -n "$STALE" ] && [ -d "$STALE" ]; then
    echo "Removing stale worktree at ${STALE}..."
    git worktree remove --force "$STALE" 2>/dev/null || rm -rf "$STALE"
fi

# --- Collect artifacts to temp staging ---
echo "Collecting artifacts..."
STAGING=$(mktemp -d)
trap 'rm -rf "$STAGING"; cleanup' EXIT
mkdir -p "$STAGING/$DEST_DIR"

if [ "$HAS_RESULTS" = "true" ]; then
    mkdir -p "$STAGING/$DEST_DIR/results"
    cp tmp/ndjson/*-result.json "$STAGING/$DEST_DIR/results/"
    echo "  $(ls tmp/ndjson/*-result.json | wc -l) result JSONs"
fi

if [ "$HAS_REPORT" = "true" ]; then
    cp "${REPO_ROOT}/tmp/ndjson/e2e-report.html" "$STAGING/$DEST_DIR/e2e-report.html"
    echo "  e2e-report.html"
fi

if [ "$HAS_DIAGRAMS" = "true" ]; then
    mkdir -p "$STAGING/$DEST_DIR/analysis"
    cp tmp/ndjson_analysis/*.puml "$STAGING/$DEST_DIR/analysis/" 2>/dev/null || true
    cp tmp/ndjson_analysis/*.png "$STAGING/$DEST_DIR/analysis/" 2>/dev/null || true
    cp tmp/ndjson_analysis/*-context.json "$STAGING/$DEST_DIR/analysis/" 2>/dev/null || true
    cp tmp/ndjson_analysis/*-analysis.md "$STAGING/$DEST_DIR/analysis/" 2>/dev/null || true
    cp tmp/ndjson_analysis/model-comparison.md "$STAGING/$DEST_DIR/analysis/" 2>/dev/null || true
    echo "  analysis outputs"
fi

# --- Create worktree ---
BRANCH_EXISTS=false
if git rev-parse --verify e2e-artifacts >/dev/null 2>&1; then
    BRANCH_EXISTS=true
fi

if [ "$BRANCH_EXISTS" = "true" ]; then
    echo "Creating worktree (incremental update)..."
    git worktree add "$WORKTREE" e2e-artifacts
else
    echo "Creating worktree (new orphan branch)..."
    git worktree add --detach "$WORKTREE"
    git -C "$WORKTREE" checkout --orphan e2e-artifacts
    git -C "$WORKTREE" rm -rf . 2>/dev/null || true
fi

# --- Copy artifacts into worktree ---
echo "Staging files into ${DEST_DIR}/..."
mkdir -p "$WORKTREE/$DEST_DIR"
cp -r "$STAGING/$DEST_DIR"/* "$WORKTREE/$DEST_DIR/"

FILE_COUNT=$(find "$WORKTREE/$DEST_DIR" -type f | wc -l)
echo "Staging ${FILE_COUNT} files..."

# --- Pruning: keep last 5 runs per version ---
PRUNED=0
dir_count=$(ls -1d "$WORKTREE/v${VERSION}"/*/ 2>/dev/null | wc -l || true)
if [ "$dir_count" -gt 5 ]; then
    pruned_dirs=$(ls -1d "$WORKTREE/v${VERSION}"/*/ 2>/dev/null | sort | head -n -5)
    echo "Pruning $((dir_count - 5)) old runs in v${VERSION}/..."
    for d in $pruned_dirs; do
        rm -rf "$d"
        PRUNED=$((PRUNED + 1))
    done
    git -C "$WORKTREE" add -A "v${VERSION}/"
fi

# --- Commit ---
git -C "$WORKTREE" add "$DEST_DIR/"

if ! git -C "$WORKTREE" diff --cached --quiet; then
    echo "Committing..."
    git -C "$WORKTREE" commit -m "e2e-artifacts: ${DEST_DIR}"
else
    echo "No changes to commit for ${DEST_DIR} — artifacts identical to last run."
fi

# --- Summary ---
DISK_SIZE=$(du -sh "$WORKTREE/$DEST_DIR" 2>/dev/null | cut -f1 || echo "?")

echo ""
echo "=== E2E Artifacts Prepared ==="
echo "Version: ${VERSION}"
echo "Date: ${DATE}"
echo "Directory: ${DEST_DIR}"
echo "Files: ${FILE_COUNT}"
echo "Size: ${DISK_SIZE}"
if [ "$PRUNED" -gt 0 ]; then echo "Pruned: ${PRUNED} old run(s)"; fi
echo ""
echo "Contents:"
(cd "$WORKTREE" && find "$DEST_DIR" -type f | sort)
echo ""
if [ "$BRANCH_EXISTS" = "true" ]; then
    echo "Mode: incremental update"
else
    echo "Mode: initial baseline (first run)"
fi
echo ""
echo "Branch: e2e-artifacts (local commit only)"
echo ""
echo "To push to remote:"
echo "  git push origin e2e-artifacts --force"
echo ""
echo "Done!"
