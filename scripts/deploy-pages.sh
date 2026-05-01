#!/usr/bin/env bash
# deploy-pages.sh — Deploy documentation site to GitHub Pages
# Creates/updates an orphan gh-pages branch with landing page, linter graph,
# diagrams page, optional e2e report, image assets, and diagram PNGs.
# Force-pushes to origin, then returns to the original branch.
set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"

# --- Source file paths ---
INDEX_HTML="${REPO_ROOT}/pages/index.html"
GRAPH_HTML="${REPO_ROOT}/graphify-out/graph.html"
DIAGRAMS_HTML="${REPO_ROOT}/pages/diagrams.html"
E2E_REPORT="${REPO_ROOT}/e2e-report.html"

# --- Pre-flight checks (required files) ---
for f in "$INDEX_HTML" "$GRAPH_HTML" "$DIAGRAMS_HTML"; do
    if [ ! -f "$f" ]; then
        echo "Error: Required file not found: $f" >&2
        exit 1
    fi
done

# --- E2E report check ---
if [ ! -f "$E2E_REPORT" ]; then
    echo "Error: e2e-report.html not found. Run 'make integration-test' first." >&2
    exit 1
fi

# --- Guard: refuse to run with uncommitted changes ---
if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "Error: Uncommitted changes detected. Commit or stash before deploying." >&2
    git status --short >&2
    exit 1
fi

echo "Deploying documentation site to GitHub Pages..."

# --- Save current branch ---
CURRENT=$(git branch --show-current)

# --- Temp staging directory ---
TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

# --- Stage deploy content ---
mkdir -p "$TMPDIR/assets" "$TMPDIR/diagrams"

# Pages
cp "$INDEX_HTML" "$TMPDIR/index.html"
cp "$GRAPH_HTML" "$TMPDIR/graph.html"
cp "$DIAGRAMS_HTML" "$TMPDIR/diagrams.html"

# Optional e2e report
if [ -n "$E2E_REPORT" ]; then
    cp "$E2E_REPORT" "$TMPDIR/e2e-report.html"
fi

# Assets
cp "${REPO_ROOT}/assets/graphify.png" "$TMPDIR/assets/graphify.png"
cp "${REPO_ROOT}/assets/desloppify-scorecard.png" "$TMPDIR/assets/scorecard.png"

# Diagrams
cp "${REPO_ROOT}/assets/ndjson_analysis/"*.png "$TMPDIR/diagrams/"

# Nojekyll
touch "$TMPDIR/.nojekyll"

echo "Staged deploy content: index.html, graph.html, diagrams.html, assets/, diagrams/, .nojekyll"

# --- Cleanup function: always return to original branch ---
cleanup() {
    echo "Restoring original branch: $CURRENT"
    git checkout "$CURRENT" 2>/dev/null || true
    git branch -D gh-pages 2>/dev/null || true
}
trap cleanup EXIT

# --- Handle gh-pages branch ---
git branch -D gh-pages 2>/dev/null || true
echo "Created orphan gh-pages branch"

git checkout --orphan gh-pages
git rm -rf . 2>/dev/null || true

# Copy staged files into worktree root
cp -r "$TMPDIR"/* .

git add .nojekyll index.html graph.html diagrams.html assets/ diagrams/
if [ -f e2e-report.html ]; then
    git add e2e-report.html
fi
git commit -m "deploy: golangci-lint-mcp documentation site"

echo "Pushing to origin/gh-pages..."
git push origin gh-pages --force

echo "Pushed to origin/gh-pages"

# cleanup trap will restore branch
echo "Done! Site deployed with:"
echo "  - index.html (landing page)"
echo "  - graph.html (linter relationship graph)"
echo "  - diagrams.html (sequence diagrams)"
if [ -n "$E2E_REPORT" ]; then
    echo "  - e2e-report.html (integration test report)"
fi
echo "Visit: https://wavilen.github.io/golangci-lint-mcp/"
