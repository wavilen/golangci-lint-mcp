#!/usr/bin/env python3
"""
Curate <related> tags across guide files based on CLUSTER-ANALYSIS.md directives.

This script applies hardcoded curation actions (add/replace/remove/set) to ~63 guide files.
It ONLY modifies content between <related> and </related> tags — never touching other sections.
"""

import os
import re
import sys

# ─── Helpers ───────────────────────────────────────────────────────────────────

GUIDES_ROOT = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "guides")


def guide_path(rel: str) -> str:
    """Resolve a guide path like 'guides/gosec/G107.md' to full path."""
    # Strip leading 'guides/' if present since GUIDES_ROOT already points to guides/
    if rel.startswith("guides/"):
        rel = rel[len("guides/"):]
    return os.path.join(GUIDES_ROOT, rel)


def get_linter_dir(rel_path: str) -> str:
    """Get the linter subdirectory for a guide (e.g., 'gosec' from 'gosec/G107.md'), or '' for flat files."""
    # Strip 'guides/' prefix if present
    p = rel_path.replace("\\", "/")
    if p.startswith("guides/"):
        p = p[len("guides/"):]
    parts = p.split("/")
    if len(parts) == 2:
        return parts[0]
    return ""


def normalize_ref(ref: str, target_rel_path: str) -> str:
    """
    Normalize a reference for the target guide.
    If the ref is compound (e.g., 'gocritic/rangeValCopy') and the target is in the same
    linter subdirectory, use the bare form ('rangeValCopy').
    """
    target_linter = get_linter_dir(target_rel_path)
    if "/" in ref:
        ref_linter, _ = ref.split("/", 1)
        if ref_linter == target_linter:
            return ref.split("/", 1)[1]
    return ref


def parse_related_tags(content: str):
    """
    Find <related>...</related> section in file content.
    Handles both closed tags (<related>..</related>) and open tags (<related>...EOF).
    Returns (start, end, refs_list, has_close_tag) where start/end are character positions
    of the content between the tags, and refs_list is the parsed references.
    """
    # Try closed form first: <related>...</related>
    pattern_closed = re.compile(r"(<related>)\s*(.*?)\s*(</related>)", re.DOTALL)
    m = pattern_closed.search(content)
    if m:
        tag_content = m.group(2).strip()
        refs = [r.strip() for r in tag_content.split(",") if r.strip()] if tag_content else []
        return m.start(2), m.end(2), refs, True

    # Try open form: <related>... (no closing tag, goes to end of file)
    pattern_open = re.compile(r"<related>\s*\n?(.*?)$", re.DOTALL)
    m = pattern_open.search(content)
    if m:
        tag_content = m.group(1).strip()
        refs = [r.strip() for r in tag_content.split(",") if r.strip()] if tag_content else []
        return m.start(1), m.end(1), refs, False

    return None


def update_related_section(filepath: str, new_refs: list) -> bool:
    """
    Update the <related> section in a file with the new reference list.
    Returns True if the file was modified, False otherwise.
    """
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    result = parse_related_tags(content)
    if result is None:
        print(f"  WARNING: No <related> tag found in {filepath}")
        return False

    start, end, old_refs, has_close_tag = result
    new_content_str = ", ".join(new_refs)

    if has_close_tag:
        replacement = f"\n{new_content_str}\n"
    else:
        # Open tag (no </related>) — content is at end of file
        replacement = f"{new_content_str}\n"

    new_file_content = content[:start] + replacement + content[end:]

    if new_file_content == content:
        return False

    with open(filepath, "w", encoding="utf-8") as f:
        f.write(new_file_content)
    return True


def do_add(rel_path: str, refs_to_add: list) -> tuple:
    """Add references to a guide's <related> section. Skip duplicates."""
    filepath = guide_path(rel_path)
    if not os.path.exists(filepath):
        return (False, f"File not found: {filepath}")

    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    result = parse_related_tags(content)
    if result is None:
        return (False, f"No <related> tag in {rel_path}")

    _, _, existing_refs, _ = result
    added = []
    for ref in refs_to_add:
        normalized = normalize_ref(ref, rel_path)
        # Check for both compound and bare forms
        if normalized not in existing_refs and ref not in existing_refs:
            existing_refs.append(normalized)
            added.append(normalized)

    if not added:
        return (False, f"No new refs to add (all duplicates): {rel_path}")

    modified = update_related_section(filepath, existing_refs)
    return (modified, f"Added: {', '.join(added)} → {rel_path}")


def do_replace(rel_path: str, old_ref: str, new_ref: str) -> tuple:
    """Replace a specific reference in a guide's <related> section."""
    filepath = guide_path(rel_path)
    if not os.path.exists(filepath):
        return (False, f"File not found: {filepath}")

    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    result = parse_related_tags(content)
    if result is None:
        return (False, f"No <related> tag in {rel_path}")

    _, _, existing_refs, _ = result

    # Try exact match first, then try matching bare form
    normalized_new = normalize_ref(new_ref, rel_path)
    found = False

    for i, ref in enumerate(existing_refs):
        if ref == old_ref:
            existing_refs[i] = normalized_new
            found = True
            break

    if not found:
        # Try bare form of old_ref
        old_bare = old_ref.split("/")[-1] if "/" in old_ref else old_ref
        for i, ref in enumerate(existing_refs):
            if ref == old_bare:
                existing_refs[i] = normalized_new
                found = True
                break

    if not found:
        return (False, f"WARNING: Replace target '{old_ref}' not found in {rel_path} (has: {', '.join(existing_refs)})")

    modified = update_related_section(filepath, existing_refs)
    return (modified, f"Replaced: {old_ref} → {normalized_new} in {rel_path}")


def do_remove(rel_path: str, ref_to_remove: str) -> tuple:
    """Remove a specific reference from a guide's <related> section."""
    filepath = guide_path(rel_path)
    if not os.path.exists(filepath):
        return (False, f"File not found: {filepath}")

    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    result = parse_related_tags(content)
    if result is None:
        return (False, f"No <related> tag in {rel_path}")

    _, _, existing_refs, _ = result

    # Try exact match, then bare form
    removed = False
    new_refs = []
    for ref in existing_refs:
        if ref == ref_to_remove or ref == ref_to_remove.split("/")[-1]:
            removed = True
            continue
        new_refs.append(ref)

    if not removed:
        return (False, f"WARNING: Remove target '{ref_to_remove}' not found in {rel_path}")

    modified = update_related_section(filepath, new_refs)
    return (modified, f"Removed: {ref_to_remove} from {rel_path}")


def do_set(rel_path: str, refs_to_set: list) -> tuple:
    """
    Set the entire <related> content. If tag already has content, convert to add instead.
    """
    filepath = guide_path(rel_path)
    if not os.path.exists(filepath):
        return (False, f"File not found: {filepath}")

    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    result = parse_related_tags(content)
    if result is None:
        return (False, f"No <related> tag in {rel_path}")

    _, _, existing_refs, _ = result

    if existing_refs:
        # Tag has content — convert to add instead of set to avoid destroying content
        normalized_refs = [normalize_ref(r, rel_path) for r in refs_to_set]
        added = []
        for ref in normalized_refs:
            if ref not in existing_refs:
                existing_refs.append(ref)
                added.append(ref)
        if not added:
            return (False, f"Set → Add (no new refs, all present): {rel_path}")
        modified = update_related_section(filepath, existing_refs)
        return (modified, f"Set → Add (tag had content): Added {', '.join(added)} to {rel_path}")

    # Tag is truly empty — set the content
    normalized_refs = [normalize_ref(r, rel_path) for r in refs_to_set]
    modified = update_related_section(filepath, normalized_refs)
    return (modified, f"Set: {', '.join(normalized_refs)} in {rel_path}")


# ─── Action List (hardcoded from CLUSTER-ANALYSIS.md) ──────────────────────────

ACTIONS = [
    # ── Security Auditing cluster ──────────────────────────────────────────────
    # 1. gosec/G107 — add refs: noctx, gosec/G204
    ("guides/gosec/G107.md", "add", {"refs": ["noctx", "gosec/G204"]}),
    # 2. gosec/G101 — add ref: bidichk
    ("guides/gosec/G101.md", "add", {"refs": ["bidichk"]}),
    # 3. durationcheck — set refs: gosec/G109 (expected empty)
    ("guides/durationcheck.md", "set", {"refs": ["gosec/G109"]}),

    # ── Error Handling cluster ─────────────────────────────────────────────────
    # 4. errcheck — add refs: nilerr, rowserrcheck
    ("guides/errcheck.md", "add", {"refs": ["nilerr", "rowserrcheck"]}),
    # 5. wrapcheck — set refs: errcheck, err113 (expected empty)
    ("guides/wrapcheck.md", "set", {"refs": ["errcheck", "err113"]}),
    # 6. nilerr — set refs: errcheck, nilnesserr (expected empty)
    ("guides/nilerr.md", "set", {"refs": ["errcheck", "nilnesserr"]}),
    # 7. nilnesserr — set refs: nilerr, errcheck (expected empty)
    ("guides/nilnesserr.md", "set", {"refs": ["nilerr", "errcheck"]}),
    # 8. nilnil — set refs: nilerr (expected empty)
    ("guides/nilnil.md", "set", {"refs": ["nilerr"]}),
    # 9. errorlint/errorf — add ref: modernize/errorf
    ("guides/errorlint/errorf.md", "add", {"refs": ["modernize/errorf"]}),
    # 10. errorlint/asserts — add ref: govet/errorsas
    ("guides/errorlint/asserts.md", "add", {"refs": ["govet/errorsas"]}),
    # 11. errorlint/comparison — add ref: err113
    ("guides/errorlint/comparison.md", "add", {"refs": ["err113"]}),
    # 12. revive/unhandled-error — add ref: errcheck
    ("guides/revive/unhandled-error.md", "add", {"refs": ["errcheck"]}),
    # 13. revive/error-strings — add ref: errname
    ("guides/revive/error-strings.md", "add", {"refs": ["errname"]}),

    # ── Concurrency Safety cluster ─────────────────────────────────────────────
    # 14. bodyclose — add refs: contextcheck, noctx
    ("guides/bodyclose.md", "add", {"refs": ["contextcheck", "noctx"]}),
    # 15. contextcheck — replace revive with noctx, add refs: bodyclose, govet/lostcancel
    ("guides/contextcheck.md", "replace", {"old": "revive", "new": "noctx"}),
    ("guides/contextcheck.md", "add", {"refs": ["bodyclose", "govet/lostcancel"]}),
    # 16. makezero — add ref: prealloc
    ("guides/makezero.md", "add", {"refs": ["prealloc"]}),
    # 17. containedctx — replace revive with staticcheck/SA1013
    ("guides/containedctx.md", "replace", {"old": "revive", "new": "staticcheck/SA1013"}),
    # 18. govet/lostcancel — add ref: contextcheck
    ("guides/govet/lostcancel.md", "add", {"refs": ["contextcheck"]}),
    # 19. govet/loopclosure — add ref: revive/range-val-in-closure
    ("guides/govet/loopclosure.md", "add", {"refs": ["revive/range-val-in-closure"]}),
    # 20. govet/waitgroup — add ref: revive/waitgroup-by-value
    ("guides/govet/waitgroup.md", "add", {"refs": ["revive/waitgroup-by-value"]}),
    # 21. govet/atomic — add ref: gocritic/badLock
    ("guides/govet/atomic.md", "add", {"refs": ["gocritic/badLock"]}),

    # ── Style and Formatting cluster ───────────────────────────────────────────
    # 22. lll — replace revive with revive/line-length-limit
    ("guides/lll.md", "replace", {"old": "revive", "new": "revive/line-length-limit"}),
    # 23. nakedret — add ref: errcheck
    ("guides/nakedret.md", "add", {"refs": ["errcheck"]}),
    # 24. sloglint — add ref: zerologlint
    ("guides/sloglint.md", "add", {"refs": ["zerologlint"]}),
    # 25. loggercheck — add refs: sloglint, zerologlint
    ("guides/loggercheck.md", "add", {"refs": ["sloglint", "zerologlint"]}),

    # ── Dead Code Detection cluster ────────────────────────────────────────────
    # 26. unused — add ref: unparam
    ("guides/unused.md", "add", {"refs": ["unparam"]}),
    # 27. unparam — add ref: unused
    ("guides/unparam.md", "add", {"refs": ["unused"]}),
    # 28. govet/unreachable — add ref: staticcheck/SA4032
    ("guides/govet/unreachable.md", "add", {"refs": ["staticcheck/SA4032"]}),
    # 29. govet/errorsas — add ref: errorlint/asserts
    ("guides/govet/errorsas.md", "add", {"refs": ["errorlint/asserts"]}),
    # 30. staticcheck/SA4006 — add ref: ineffassign
    ("guides/staticcheck/SA4006.md", "add", {"refs": ["ineffassign"]}),

    # ── Code Complexity cluster ────────────────────────────────────────────────
    # 31. gocyclo — add ref: revive/cyclomatic
    ("guides/gocyclo.md", "add", {"refs": ["revive/cyclomatic"]}),
    # 32. funlen — add ref: revive/function-length
    ("guides/funlen.md", "add", {"refs": ["revive/function-length"]}),
    # 33. nestif — add ref: revive/max-control-nesting
    ("guides/nestif.md", "add", {"refs": ["revive/max-control-nesting"]}),
    # 34. gocognit — add ref: revive/cognitive-complexity
    ("guides/gocognit.md", "add", {"refs": ["revive/cognitive-complexity"]}),
    # 35. interfacebloat — replace revive with ireturn
    ("guides/interfacebloat.md", "replace", {"old": "revive", "new": "ireturn"}),
    # 36. ireturn — replace revive with revive/max-public-structs, replace gocritic with gocritic/unnamedResult
    ("guides/ireturn.md", "replace", {"old": "revive", "new": "revive/max-public-structs"}),
    ("guides/ireturn.md", "replace", {"old": "gocritic", "new": "gocritic/unnamedResult"}),
    # 37. dogsled — remove refs: unparam, errcheck; add ref: gocritic/tooManyResultsChecker
    ("guides/dogsled.md", "remove", {"ref": "unparam"}),
    ("guides/dogsled.md", "remove", {"ref": "errcheck"}),
    ("guides/dogsled.md", "add", {"refs": ["gocritic/tooManyResultsChecker"]}),

    # ── Performance Optimization cluster ───────────────────────────────────────
    # 38. prealloc — add refs: wastedassign, ineffassign
    ("guides/prealloc.md", "add", {"refs": ["wastedassign", "ineffassign"]}),
    # 39. ineffassign — add refs: wastedassign, prealloc
    ("guides/ineffassign.md", "add", {"refs": ["wastedassign", "prealloc"]}),
    # 40. wastedassign — add refs: ineffassign, prealloc
    ("guides/wastedassign.md", "add", {"refs": ["ineffassign", "prealloc"]}),
    # 41. unconvert — add ref: perfsprint
    ("guides/unconvert.md", "add", {"refs": ["perfsprint"]}),
    # 42. perfsprint — replace govet with govet/printf
    ("guides/perfsprint.md", "replace", {"old": "govet", "new": "govet/printf"}),
    # 43. nosprintfhostport — replace govet with govet/hostport
    ("guides/nosprintfhostport.md", "replace", {"old": "govet", "new": "govet/hostport"}),
    # 44. govet/copylocks — add ref: gocritic/badLock
    ("guides/govet/copylocks.md", "add", {"refs": ["gocritic/badLock"]}),
    # 45. fatcontext — add refs: gocritic/hugeParam, gocritic/rangeValCopy
    ("guides/fatcontext.md", "add", {"refs": ["gocritic/hugeParam", "gocritic/rangeValCopy"]}),
    # 46. gocritic/hugeParam — add refs: gocritic/rangeValCopy, prealloc
    ("guides/gocritic/hugeParam.md", "add", {"refs": ["gocritic/rangeValCopy", "prealloc"]}),
    # 47. gocritic/rangeValCopy — add refs: gocritic/rangeExprCopy, gocritic/hugeParam
    ("guides/gocritic/rangeValCopy.md", "add", {"refs": ["gocritic/rangeExprCopy", "gocritic/hugeParam"]}),

    # ── Code Simplification cluster ────────────────────────────────────────────
    # 48. gocritic/unlambda — add ref: gocritic/deferUnlambda
    ("guides/gocritic/unlambda.md", "add", {"refs": ["gocritic/deferUnlambda"]}),
    # 49. gocritic/unslice — add ref: gocritic/typeUnparen
    ("guides/gocritic/unslice.md", "add", {"refs": ["gocritic/typeUnparen"]}),

    # ── Deprecated API Patterns cluster ────────────────────────────────────────
    # 50. staticcheck/SA1000 — remove ref: gosec/G204
    ("guides/staticcheck/SA1000.md", "remove", {"ref": "gosec/G204"}),

    # ── Testing Frameworks cluster ─────────────────────────────────────────────
    # 51. paralleltest — add ref: usetesting
    ("guides/paralleltest.md", "add", {"refs": ["usetesting"]}),
    # 52. tparallel — add ref: usetesting
    ("guides/tparallel.md", "add", {"refs": ["usetesting"]}),
    # 53. thelper — add refs: usetesting, testifylint/suite-thelper
    ("guides/thelper.md", "add", {"refs": ["usetesting", "testifylint/suite-thelper"]}),
    # 54. testifylint/error-as — add refs: testifylint/error-nil, errorlint/asserts
    ("guides/testifylint/error-as.md", "add", {"refs": ["testifylint/error-nil", "errorlint/asserts"]}),
    # 55. testifylint/error-nil — add refs: testifylint/error-as, errcheck
    ("guides/testifylint/error-nil.md", "add", {"refs": ["testifylint/error-as", "errcheck"]}),
    # 56. testifylint/require-error — add ref: testifylint/error-nil
    ("guides/testifylint/require-error.md", "add", {"refs": ["testifylint/error-nil"]}),
    # 57. testifylint/expected-actual — add refs: testifylint/bool-compare, testifylint/nil-compare
    ("guides/testifylint/expected-actual.md", "add", {"refs": ["testifylint/bool-compare", "testifylint/nil-compare"]}),
    # 58. testifylint/useless-assert — add ref: testifylint/expected-actual
    ("guides/testifylint/useless-assert.md", "add", {"refs": ["testifylint/expected-actual"]}),
    # 59. testifylint/nil-compare — add ref: testifylint/bool-compare
    ("guides/testifylint/nil-compare.md", "add", {"refs": ["testifylint/bool-compare"]}),
    # 60. ginkgolinter/error-assertion — add ref: ginkgolinter/nil-assertion
    ("guides/ginkgolinter/error-assertion.md", "add", {"refs": ["ginkgolinter/nil-assertion"]}),
    # 61. ginkgolinter/nil-assertion — add ref: ginkgolinter/error-assertion
    ("guides/ginkgolinter/nil-assertion.md", "add", {"refs": ["ginkgolinter/error-assertion"]}),
    # 62. ginkgolinter/compare-assertion — add ref: ginkgolinter/nil-assertion
    ("guides/ginkgolinter/compare-assertion.md", "add", {"refs": ["ginkgolinter/nil-assertion"]}),
    # 63. ginkgolinter/succeed-matcher — add ref: ginkgolinter/error-assertion
    ("guides/ginkgolinter/succeed-matcher.md", "add", {"refs": ["ginkgolinter/error-assertion"]}),
]


# ─── Main ──────────────────────────────────────────────────────────────────────

def main():
    modified_count = 0
    skipped_count = 0
    warning_count = 0
    errors = []

    print("=" * 70)
    print("Related Tag Curation Script")
    print(f"Processing {len(ACTIONS)} actions across {len(set(a[0] for a in ACTIONS))} guide files")
    print("=" * 70)

    for rel_path, action_type, params in ACTIONS:
        filepath = guide_path(rel_path)

        # Quick existence check
        if not os.path.exists(filepath):
            msg = f"ERROR: File not found: {filepath}"
            print(f"  {msg}")
            errors.append(msg)
            continue

        if action_type == "add":
            ok, msg = do_add(rel_path, params["refs"])
        elif action_type == "replace":
            ok, msg = do_replace(rel_path, params["old"], params["new"])
        elif action_type == "remove":
            ok, msg = do_remove(rel_path, params["ref"])
        elif action_type == "set":
            ok, msg = do_set(rel_path, params["refs"])
        else:
            msg = f"ERROR: Unknown action type: {action_type}"
            print(f"  {msg}")
            errors.append(msg)
            continue

        if ok:
            modified_count += 1
            print(f"  ✓ {msg}")
        else:
            skipped_count += 1
            if "WARNING" in msg:
                warning_count += 1
            print(f"  - {msg}")

    print()
    print("=" * 70)
    print(f"Results: {modified_count} modified, {skipped_count} skipped ({warning_count} warnings), {len(errors)} errors")

    if errors:
        print("\nErrors:")
        for e in errors:
            print(f"  - {e}")
        sys.exit(1)

    sys.exit(0)


if __name__ == "__main__":
    main()
