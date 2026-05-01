#!/usr/bin/env bash
# deploy-pages.sh — Deploy documentation site to GitHub Pages
# Uses git worktree for safe orphan branch management.
# Fetches missing assets from e2e-artifacts branch where possible.
# All assets required — no partial deploys (per D-07).
set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

# --- Source file paths ---
INDEX_HTML="${REPO_ROOT}/pages/index.html"
DIAGRAMS_HTML="${REPO_ROOT}/pages/diagrams.html"
GRAPH_HTML="${REPO_ROOT}/graphify-out/graph.html"
E2E_REPORT="${REPO_ROOT}/tmp/ndjson_analysis/e2e-report.html"
NDJSON_PNGS="${REPO_ROOT}/tmp/ndjson_analysis"

# --- Guard: refuse to run with uncommitted changes ---
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "Error: Uncommitted changes detected. Commit or stash before deploying." >&2
    git status --short >&2
    exit 1
fi

# --- Asset recovery: fetch missing assets from branches (D-06) ---
# Try to recover e2e-report.html from e2e-artifacts branch
if [ ! -f "$E2E_REPORT" ]; then
    echo "e2e-report.html not found locally. Attempting to fetch from e2e-artifacts branch..."
    if git fetch origin e2e-artifacts --depth=1 2>/dev/null; then
        LATEST_REPORT=$(git ls-tree -r --name-only origin/e2e-artifacts | grep -F 'e2e-report.html' | sort -r | head -1)
        if [ -n "$LATEST_REPORT" ]; then
            git show "origin/e2e-artifacts:${LATEST_REPORT}" > "$E2E_REPORT"
            echo "  Recovered e2e-report.html from: ${LATEST_REPORT}"
        fi
    fi
fi

# Try to recover ndjson analysis PNGs from e2e-artifacts branch
_NDJSON_PNG_LOCAL=0
if [ -d "$NDJSON_PNGS" ]; then
    _NDJSON_PNG_LOCAL=$(find "$NDJSON_PNGS" -maxdepth 1 -name '*.png' -type f 2>/dev/null | wc -l | tr -d ' ')
fi
if [ "$_NDJSON_PNG_LOCAL" -eq 0 ]; then
    echo "No ndjson analysis PNGs found locally. Attempting to fetch from e2e-artifacts branch..."
    if git fetch origin e2e-artifacts --depth=1 2>/dev/null; then
        LATEST_ANALYSIS=$(git ls-tree -r --name-only origin/e2e-artifacts | grep -E '/analysis/.*\.png$' | sort -r | head -1)
        if [ -n "$LATEST_ANALYSIS" ]; then
            ANALYSIS_PREFIX=$(dirname "$LATEST_ANALYSIS")
            mkdir -p "$NDJSON_PNGS"
            for png_path in $(git ls-tree --name-only origin/e2e-artifacts "$ANALYSIS_PREFIX/" 2>/dev/null | grep '\.png$'); do
                png_file=$(basename "$png_path")
                git show "origin/e2e-artifacts:${png_path}" > "${NDJSON_PNGS}/${png_file}"
            done
            RECOVERED_PNGS=$(find "$NDJSON_PNGS" -maxdepth 1 -name '*.png' -type f | wc -l | tr -d ' ')
            echo "  Recovered ${RECOVERED_PNGS} PNG(s) from: ${ANALYSIS_PREFIX}/"
        fi
    fi
fi

# --- Fail-fast: validate ALL required assets exist (D-05, D-07, D-08) ---
ERRORS=0
for f in "$INDEX_HTML" "$DIAGRAMS_HTML"; do
    if [ ! -f "$f" ]; then
        echo "Error: Required file not found: $f" >&2
        ERRORS=$((ERRORS + 1))
    fi
done

for f in "${REPO_ROOT}/assets/graphify.png" "${REPO_ROOT}/assets/desloppify-scorecard.png"; do
    if [ ! -f "$f" ]; then
        echo "Error: Required file not found: $f" >&2
        ERRORS=$((ERRORS + 1))
    fi
done

if [ ! -f "$GRAPH_HTML" ]; then
    echo "Error: graph.html not found. Generate it with graphify (requires conda) before deploying." >&2
    echo "  Expected at: $GRAPH_HTML" >&2
    ERRORS=$((ERRORS + 1))
fi

if [ ! -f "$E2E_REPORT" ]; then
    echo "Error: e2e-report.html not found. Run 'make integration-test' then 'make push-e2e-artifacts' first." >&2
    ERRORS=$((ERRORS + 1))
fi

# Check ndjson_analysis PNGs
NDJSON_PNG_COUNT=0
if [ -d "$NDJSON_PNGS" ]; then
    NDJSON_PNG_COUNT=$(find "$NDJSON_PNGS" -maxdepth 1 -name '*.png' -type f 2>/dev/null | wc -l | tr -d ' ')
fi
if [ "$NDJSON_PNG_COUNT" -eq 0 ]; then
    echo "Error: No ndjson analysis PNGs found in ${NDJSON_PNGS}/" >&2
    echo "  Run 'make push-e2e-artifacts' to generate diagrams, or run ndjson-analysis agent." >&2
    ERRORS=$((ERRORS + 1))
fi

if [ "$ERRORS" -gt 0 ]; then
    echo "" >&2
    echo "Deploy aborted: ${ERRORS} required asset(s) missing. All assets required (per D-07)." >&2
    exit 1
fi

echo "Deploying documentation site to GitHub Pages..."

# --- Stage deploy content to temp directory ---
STAGING=$(mktemp -d)
mkdir -p "$STAGING/assets" "$STAGING/diagrams"

cp "$INDEX_HTML" "$STAGING/index.html"
cp "$DIAGRAMS_HTML" "$STAGING/diagrams.html"
cp "$GRAPH_HTML" "$STAGING/graph.html"
cp "$E2E_REPORT" "$STAGING/e2e-report.html"
cp "${REPO_ROOT}/assets/graphify.png" "$STAGING/assets/graphify.png"
cp "${REPO_ROOT}/assets/desloppify-scorecard.png" "$STAGING/assets/scorecard.png"
cp "$NDJSON_PNGS"/*.png "$STAGING/diagrams/"
touch "$STAGING/.nojekyll"

FILE_COUNT=$(find "$STAGING" -type f | wc -l | tr -d ' ')
echo "Staged deploy content: ${FILE_COUNT} files"

# --- Git worktree for orphan branch (D-10, D-11) ---
WORKTREE=$(mktemp -d)

# Cleanup: remove worktree and staging on any exit
cleanup() {
    cd "$REPO_ROOT" 2>/dev/null || true
    git worktree remove --force "$WORKTREE" 2>/dev/null || rm -rf "$WORKTREE"
    rm -rf "$STAGING"
}
trap cleanup EXIT

# Handle existing gh-pages branch
BRANCH_EXISTS=false
if git show-ref --verify --quiet "refs/heads/gh-pages" 2>/dev/null; then
    BRANCH_EXISTS=true
fi
if git show-ref --verify --quiet "refs/remotes/origin/gh-pages" 2>/dev/null; then
    BRANCH_EXISTS=true
    if ! git show-ref --verify --quiet "refs/heads/gh-pages" 2>/dev/null; then
        git branch gh-pages origin/gh-pages 2>/dev/null || true
    fi
fi

if [ "$BRANCH_EXISTS" = "true" ]; then
    # Existing branch: add worktree pointing to it
    git worktree add "$WORKTREE" gh-pages 2>/dev/null
else
    # New orphan branch: create worktree with orphan
    git worktree add --orphan -b gh-pages "$WORKTREE"
fi

# Copy staged content into worktree (remove old content first for clean slate)
rm -rf "${WORKTREE:?}"/*
for item in "${WORKTREE:?}"/.[!.]*; do
    [ -e "$item" ] || continue
    case "$(basename "$item")" in
        .git) ;;
        *) rm -rf "$item" ;;
    esac
done
cp -r "$STAGING"/* "$WORKTREE/"
cp "$STAGING"/.nojekyll "$WORKTREE/"

# Remove stale files from worktree (keep only intended site content)
cd "$WORKTREE"
for item in *; do
    case "$item" in
        index.html|diagrams.html|graph.html|e2e-report.html|.nojekyll|assets|diagrams) ;;
        *) rm -rf "$item" ;;
    esac
done

# Remove hidden items except .nojekyll and .git (worktree link)
for item in .[!.]*; do
    [ -e "$item" ] || continue
    case "$item" in
        .nojekyll|.git) ;;
        *) rm -rf "$item" ;;
    esac
done

# Stage intended site files and handle deletions of stale files
git add index.html diagrams.html graph.html e2e-report.html .nojekyll assets/ diagrams/
# Remove any previously tracked files that are no longer in the allowlist
git ls-files --deleted -z | grep -zv '^\.git$' | xargs -0 -r git add -- || true

# Verify only intended files are staged
ALLOWLIST="index.html diagrams.html graph.html e2e-report.html .nojekyll assets/ diagrams/"
STAGED=$(git diff --cached --diff-filter=ACMR --name-only)
for f in $STAGED; do
    case "$f" in
        index.html|diagrams.html|graph.html|e2e-report.html|.nojekyll|assets/*|diagrams/*) ;;
        *)
            echo "Warning: Unexpected file staged: $f — unstaging" >&2
            git reset HEAD "$f" 2>/dev/null || true
            ;;
    esac
done

if ! git diff --cached --quiet; then
    git commit -m "deploy: golangci-lint-mcp documentation site ($(date +%Y-%m-%d))"
else
    echo "No changes to commit — site content is identical to previous deploy."
fi

echo ""  # blank line before summary

# --- Summary (git-based, shows actual committed files) ---
COMMITTED_FILES=$(git ls-tree -r --name-only HEAD)
COMMIT_FILE_COUNT=$(echo "$COMMITTED_FILES" | wc -l | tr -d ' ')
DISK_SIZE=$(du -sh . 2>/dev/null | cut -f1 || echo "?")

cd "$REPO_ROOT"

echo ""
echo "=== Pages Deploy Prepared ==="
echo "Files: ${COMMIT_FILE_COUNT}"
echo "Size: ${DISK_SIZE}"
echo ""
echo "Contents:"
echo "$COMMITTED_FILES" | sed 's|^|./|' | sort
echo ""
echo "Branch: gh-pages (local commit only)"
echo ""
echo "To push to remote:"
echo "  git push origin gh-pages --force"
echo ""
echo "Done!"
