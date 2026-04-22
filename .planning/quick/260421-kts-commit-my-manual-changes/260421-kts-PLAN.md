---
phase: quick
plan: 01
type: execute
wave: 1
depends_on: []
files_modified:
  - commands/golangci-lint.md
  - guides/wrapcheck.md
autonomous: true
requirements: [content-cleanup]
must_haves:
  truths:
    - "commands/golangci-lint.md no longer mentions editing the command file to change threshold"
    - "guides/wrapcheck.md recommends only errors.Wrap (not fmt.Errorf) for error wrapping"
  artifacts:
    - path: "commands/golangci-lint.md"
      provides: "Simplified threshold note"
    - path: "guides/wrapcheck.md"
      provides: "Consistent error wrapping guidance"
  key_links: []
---

<objective>
Commit the user's manual content edits to two files.

Purpose: The user made editorial corrections — removing an irrelevant instruction from the threshold note and simplifying wrapcheck guidance to recommend only `errors.Wrap`.
Output: Git commit with both changes.
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
User's manual changes (already applied to working tree):

1. `commands/golangci-lint.md` — removed "You may adjust this by editing the command file and changing this number." from the threshold note
2. `guides/wrapcheck.md` — changed error wrapping patterns from `fmt.Errorf("...: %w", err)` or `errors.Wrap` to only `errors.Wrap`
</context>

<tasks>

<task type="auto">
  <name>task 1: commit manual content edits</name>
  <files>commands/golangci-lint.md, guides/wrapcheck.md</files>
  <action>Stage ONLY the two content files (NOT .planning/ files). Commit with a message describing both edits:
  - commands/golangci-lint.md: remove "edit command file" from threshold note
  - guides/wrapcheck.md: simplify error wrapping patterns to errors.Wrap only

  Use: `git add commands/golangci-lint.md guides/wrapcheck.md && git commit -m "docs: simplify threshold note in golangci-lint command and prefer errors.Wrap in wrapcheck guide"`
  </action>
  <verify>
    <automated>git log -1 --oneline && git diff HEAD~1 --stat</automated>
  </verify>
  <done>Both files committed in a single commit. No .planning/ files included in the commit.</done>
</task>

</tasks>

<verification>
- `git log -1 --oneline` shows the commit
- `git diff HEAD~1 --stat` shows exactly 2 files changed
- Working tree still has .planning/ changes unstaged (not committed)
</verification>

<success_criteria>
Single git commit containing both content edits. No .planning/ files committed.
</success_criteria>

<output>
After completion, create `.planning/quick/260421-kts-commit-my-manual-changes/260421-kts-SUMMARY.md`
</output>
