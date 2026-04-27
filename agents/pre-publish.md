---
description: Pre-publishing workflow: verify docs, backup branch, clean ignored files, squash milestone commits, verify tests
mode: subagent
temperature: 0.1
permission:
  bash: allow
  edit: allow
---

<objective>
Execute the pre-publishing workflow for the golangci-lint-mcp project. This agent runs AFTER `gsd-complete-milestone` (which creates the release tag). It performs a 5-step sequence: verify documentation reflects the milestone, create a backup branch, clean all gitignored files, squash milestone commits into one, and verify all tests pass. Each step must complete before the next begins. The agent aborts safely on any failure.
</objective>

<execution_context>
- **Project:** golangci-lint-mcp — a Go MCP server + npm package with JS hooks/plugins
- **When to run:** After `gsd-complete-milestone` has created the release tag
- **Tools available:** bash (full access), edit (full access)
- **Git state expected:** Clean working tree on main or master branch
- **Files involved:** `.planning/PROJECT.md`, `.planning/REQUIREMENTS.md`, `.planning/ROADMAP.md`, `.gitignore`
</execution_context>

<safety>
## Pre-flight Checks (MUST run before any other step)

You MUST perform these checks before proceeding. If either fails, ABORT immediately and report the issue.

**Check 1: Working tree must be clean**

Run:
```bash
git status --porcelain
```

If there is ANY output (any uncommitted changes, untracked files, etc.), ABORT with the message:
> "Working tree has uncommitted changes. Commit or stash before running pre-publish."

**Check 2: Must be on main or master branch**

Run:
```bash
git branch --show-current
```

If the output is NOT `main` or `master`, ABORT with the message:
> "Not on main/master branch. Current branch: {branch}. Switch to main before running pre-publish."

Only proceed to Step 1 if BOTH checks pass.
</safety>

<process>

## Step 1: Verify Documentation

Read and verify that project documentation reflects the current milestone.

**1a. Verify PROJECT.md**

Read `.planning/PROJECT.md` and check:
- "What This Is" section reflects the current milestone
- "Current Milestone" heading names the correct milestone
- "Validated" requirements section lists all shipped features with phase references
- "Active" requirements section contains only in-progress work
- Completed phases are reflected in the Context section

**1b. Verify REQUIREMENTS.md**

Read `.planning/REQUIREMENTS.md` and check:
- Validated requirements match shipped features
- No stale Active requirements remain for completed work

**Action:** If either file is outdated or inconsistent, STOP and report exactly what needs updating. You may make minor factual corrections (e.g., updating phase references, moving completed items from Active to Validated) but do NOT rewrite sections — flag larger issues for manual review.

---

## Step 2: Create Backup Branch

Create a local backup branch before any destructive operations.

```bash
BRANCH=$(git branch --show-current)
TIMESTAMP=$(date -u +"%Y%m%dT%H%M%S")
git branch "${BRANCH}-backup-${TIMESTAMP}"
```

Example: creates `main-backup-20260426T221500`.

- Do NOT push to remote — this is a local safety net only
- Report the backup branch name to the user
- This branch is used for manual recovery if anything goes wrong in later steps

---

## Step 3: Untrack Gitignored Files from Git Index

Remove all tracked-but-gitignored files from the git index, keeping them on disk. This cleans up `.planning/`, `.claude/`, `.cursor/`, `.opencode/`, `tmp/`, `node_modules/`, `.venv/`, and any other gitignored content — without deleting the actual files.

```bash
git rm -r --cached . && git add .
```

How this works:
- `git rm -r --cached .` — recursively removes ALL entries from the git index without touching files on disk
- `git add .` — re-stages only the files that match current `.gitignore` rules. Any file that was tracked but is now gitignored will be removed from the index but remain on disk.

Report what was untracked from the index (git will list removed files).

**Warn the user:** "Tracked files that should be gitignored have been removed from the git index. Files remain on disk. Backup branch from Step 2 is available for recovery if needed."

---

## Step 4: Squash Milestone Commits

Squash all commits between the previous release tag and the current release tag into a single commit.

**4a. Identify tags**

```bash
CURRENT_TAG=$(git tag --sort=-version:refname | head -1)
PREV_TAG=$(git tag --sort=-version:refname | head -2 | tail -1)
```

- `CURRENT_TAG` — the newest tag (created by gsd-complete-milestone)
- `PREV_TAG` — the second-newest tag (previous release)

**4b. Count commits between tags**

```bash
git log ${PREV_TAG}..${CURRENT_TAG} --oneline
```

Review the commit list to understand what phases were included.

**4c. Generate squash commit message**

Read `.planning/ROADMAP.md` (if it still exists — it may have been cleaned in Step 3) or use the commit log to build a milestone summary. Format:

```
{current-tag}: {phase-name-1}, {phase-name-2}, {phase-name-3}
```

Example: `v1.2: compound command parsing, graphify research, related-tag curation, MCP expansion, ESLint, golangci-lint fixes`

If ROADMAP.md was cleaned, extract phase names from commit messages.

Write the commit message to a temp file:
```bash
echo "{milestone-summary}" > /tmp/squash-commit-msg
```

**4d. Squash using programmatic interactive rebase**

```bash
GIT_SEQUENCE_EDITOR="sed -i '2,\$s/^pick/squash/'" GIT_EDITOR="cp /tmp/squash-commit-msg" git rebase -i ${PREV_TAG}
```

How this works:
- `GIT_SEQUENCE_EDITOR` — sed changes all lines after the first from `pick` to `squash`, squashing everything into the first commit
- `GIT_EDITOR` — uses the pre-written commit message file instead of opening an editor

**4e. Force-move the current tag to the squashed commit**

```bash
git tag -f ${CURRENT_TAG}
```

**4f. Report the result**

Report:
- Number of commits that were squashed
- New commit hash (short)
- Tag moved to new commit

---

## Step 5: Verify Tests

Run all 4 verification command suites in sequence. ALL must pass.

**Test 1: Go unit tests**
```bash
go test ./...
```

**Test 2: golangci-lint**
```bash
golangci-lint run ./...
```

**Test 3: ESLint**
```bash
npx eslint plugins/ shared/ hooks/ bin/install.js
```

**Test 4: Node.js tests**
```bash
node --test plugins/golangci-lint.test.js shared/nudge.test.js
```

**Failure handling:**
- If ANY test fails: ABORT immediately
- Report which test(s) failed with the full error output
- Do NOT auto-restore — the backup branch from Step 2 is available for manual recovery
- Tell the user: "Tests failed. Fix issues and re-run @pre-publish. Backup branch {name} is available for recovery."

---

## Success Output

When all 5 steps complete successfully, report:

```
Pre-publish workflow complete.

Documentation verified ✓
Backup branch: {name} ✓
Gitignored files untracked from index ✓
{N} commits squashed into {tag} ✓
All 4 test suites passed ✓
```

The repository is now ready for publishing.

</process>
