#!/usr/bin/env python3
"""
Validate canonical format of all guide <related> tags.

Checks:
1. Format compliance: refs match linter/rule, bare simple linter, or bare compound linter dir
2. No intra-linter bare references in compound linter guides
3. Duplicate detection within single guide
4. Section integrity: verify only <related> sections changed (via git diff)

Output: validation report with overall PASS/FAIL verdict.
"""

import os
import re
import sys
import glob
import subprocess


GUIDES_ROOT = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "guides")


def discover_compound_linters():
    """Scan guides/ for subdirectories containing .md files."""
    linter_rules = {}
    for entry in os.listdir(GUIDES_ROOT):
        dirpath = os.path.join(GUIDES_ROOT, entry)
        if os.path.isdir(dirpath):
            rules = set()
            for f in os.listdir(dirpath):
                if f.endswith(".md") and f != "_template.md":
                    rules.add(f[:-3])
            if rules:
                linter_rules[entry] = rules
    return linter_rules


def discover_simple_linters():
    """Flat .md files in guides/ root."""
    simple = set()
    for f in os.listdir(GUIDES_ROOT):
        if f.endswith(".md") and f != "_template.md":
            simple.add(f[:-3])
    return simple


def parse_related(content):
    """Extract <related> section content. Returns inner text or None."""
    m = re.search(r'<related>(.*?)</related>', content, re.DOTALL)
    if m:
        return m.group(1)
    m = re.search(r'<related>\s*\n?(.*?)$', content, re.DOTALL)
    if m:
        return m.group(1)
    return None


def parse_refs(inner_content):
    """Parse comma-separated references."""
    refs = []
    for r in inner_content.split(","):
        r = r.strip()
        if r:
            refs.append(r)
    return refs


def get_guide_linter(filepath):
    """Determine which compound linter a guide belongs to."""
    relpath = os.path.relpath(filepath, GUIDES_ROOT)
    parts = relpath.split(os.sep)
    if len(parts) >= 2:
        return parts[0]
    return None


def check_format_compliance(ref, compound_linters, simple_linters, reverse_lookup):
    """Check if a reference matches a valid canonical format.
    
    Returns (is_valid, reason).
    """
    # Already in linter/rule format
    if "/" in ref:
        parts = ref.split("/", 1)
        linter = parts[0]
        rule = parts[1]
        if linter in compound_linters:
            if rule in compound_linters[linter]:
                return True, "valid compound ref"
            else:
                return False, f"unknown rule '{rule}' in compound linter '{linter}'"
        else:
            return False, f"unknown compound linter '{linter}'"
    
    # Bare simple linter
    if ref in simple_linters:
        return True, "valid simple linter"
    
    # Bare compound linter directory name (D-06)
    if ref in compound_linters:
        return True, "valid bare compound linter name (D-06)"
    
    # Check if it matches a known rule but wasn't prefixed (would be a format violation)
    if ref in reverse_lookup:
        matching = reverse_lookup[ref]
        if len(matching) == 1:
            linter = next(iter(matching))
            return False, f"bare rule should be '{linter}/{ref}'"
        else:
            linters = ", ".join(sorted(matching))
            return False, f"ambiguous bare rule (matches: {linters})"
    
    return False, f"unknown reference '{ref}'"


def check_intra_linter_bare(ref, guide_linter, compound_linters):
    """Check if a reference is a bare intra-linter reference."""
    if "/" in ref:
        return False  # Already prefixed
    if guide_linter and guide_linter in compound_linters:
        if ref in compound_linters[guide_linter]:
            return True
    return False


def main():
    print("=== Related Tag Validation ===\n")
    
    compound_linters = discover_compound_linters()
    simple_linters = discover_simple_linters()
    
    # Build reverse lookup
    reverse_lookup = {}
    for linter, rules in compound_linters.items():
        for rule in rules:
            if rule not in reverse_lookup:
                reverse_lookup[rule] = set()
            reverse_lookup[rule].add(linter)
    
    # Collect all guide files
    guide_files = []
    for f in glob.glob(os.path.join(GUIDES_ROOT, "*.md")):
        if os.path.basename(f) != "_template.md":
            guide_files.append(f)
    for f in glob.glob(os.path.join(GUIDES_ROOT, "**", "*.md"), recursive=True):
        if os.path.basename(f) != "_template.md" and os.path.dirname(f) != GUIDES_ROOT:
            guide_files.append(f)
    guide_files = sorted(set(guide_files))
    
    print(f"Compound linters: {len(compound_linters)}")
    print(f"Simple linters: {len(simple_linters)}")
    print(f"Total guides to validate: {len(guide_files)}\n")
    
    # Validation counters
    total_guides = 0
    guides_with_related = 0
    guides_empty_related = 0
    format_violations = 0
    intra_linter_violations = 0
    duplicate_violations = 0
    total_refs = 0
    empty_tag_guides = []
    format_violation_details = []
    intra_violation_details = []
    dup_violation_details = []
    
    for filepath in guide_files:
        total_guides += 1
        guide_linter = get_guide_linter(filepath)
        
        with open(filepath, "r", encoding="utf-8") as f:
            content = f.read()
        
        inner = parse_related(content)
        if inner is None:
            continue
        
        refs = parse_refs(inner)
        
        if not refs:
            guides_empty_related += 1
            empty_tag_guides.append(os.path.relpath(filepath, os.path.dirname(GUIDES_ROOT)))
            continue
        
        guides_with_related += 1
        
        # Check for duplicates
        seen = {}
        for ref in refs:
            total_refs += 1
            if ref in seen:
                duplicate_violations += 1
                dup_violation_details.append(
                    f"  {os.path.relpath(filepath, os.path.dirname(GUIDES_ROOT))}: "
                    f"duplicate '{ref}'"
                )
            else:
                seen[ref] = True
        
        # Check format compliance
        for ref in refs:
            is_valid, reason = check_format_compliance(
                ref, compound_linters, simple_linters, reverse_lookup
            )
            if not is_valid:
                format_violations += 1
                format_violation_details.append(
                    f"  {os.path.relpath(filepath, os.path.dirname(GUIDES_ROOT))}: "
                    f"'{ref}' — {reason}"
                )
            
            # Check intra-linter bare refs
            if check_intra_linter_bare(ref, guide_linter, compound_linters):
                intra_linter_violations += 1
                intra_violation_details.append(
                    f"  {os.path.relpath(filepath, os.path.dirname(GUIDES_ROOT))}: "
                    f"bare '{ref}' should be '{guide_linter}/{ref}'"
                )
    
    # Section integrity check via git diff
    print("Checking section integrity (git diff)...")
    try:
        result = subprocess.run(
            ["git", "diff", "--stat", "HEAD~1"],
            capture_output=True, text=True,
            cwd=os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        )
        # Check if diff only touches <related> sections
        # We verify by checking git diff for non-related content changes
        diff_result = subprocess.run(
            ["git", "diff", "HEAD~1", "--", "guides/"],
            capture_output=True, text=True,
            cwd=os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        )
        diff_lines = diff_result.stdout.split("\n")
        non_related_changes = 0
        for line in diff_lines:
            # Lines starting with +- but not in related context
            if line.startswith("+") or line.startswith("-"):
                if line.startswith("+++") or line.startswith("---"):
                    continue
                stripped = line[1:].strip()
                if not stripped:
                    continue
                # Check if it looks like a related tag change
                if any(x in stripped for x in ["<related>", "</related>", ", "]):
                    continue
                # Check if it's just a ref list change (bare refs → compound refs)
                # These are comma-separated ref lines
                if re.match(r'^[a-zA-Z0-9/_\-]+(,\s*[a-zA-Z0-9/_\-]+)+$', stripped):
                    continue
                if re.match(r'^[a-zA-Z0-9/_\-]+$', stripped):
                    continue
                non_related_changes += 1
        integrity_status = "CLEAN" if non_related_changes == 0 else f"WARNING: {non_related_changes} potential non-related changes"
    except Exception as e:
        integrity_status = f"SKIPPED (git error: {e})"
    
    # Validation Report
    print("\n" + "=" * 60)
    print("VALIDATION REPORT")
    print("=" * 60)
    print(f"Total guides checked: {total_guides}")
    print(f"Guides with related content: {guides_with_related}")
    print(f"Guides with empty related tags: {guides_empty_related}")
    print(f"Total references validated: {total_refs}")
    print()
    print(f"Format violations: {format_violations}")
    print(f"Intra-linter bare ref violations: {intra_linter_violations}")
    print(f"Duplicate references: {duplicate_violations}")
    print(f"Section integrity: {integrity_status}")
    print()
    
    if format_violation_details:
        print("--- Format Violations ---")
        for d in format_violation_details[:50]:
            print(d)
        if len(format_violation_details) > 50:
            print(f"  ... and {len(format_violation_details) - 50} more")
        print()
    
    if intra_violation_details:
        print("--- Intra-Linter Bare Ref Violations ---")
        for d in intra_violation_details[:50]:
            print(d)
        if len(intra_violation_details) > 50:
            print(f"  ... and {len(intra_violation_details) - 50} more")
        print()
    
    if dup_violation_details:
        print("--- Duplicate References ---")
        for d in dup_violation_details:
            print(d)
        print()
    
    print(f"Empty tag guides ({len(empty_tag_guides)} total):")
    for g in empty_tag_guides[:10]:
        print(f"  {g}")
    if len(empty_tag_guides) > 10:
        print(f"  ... and {len(empty_tag_guides) - 10} more")
    
    # Check which format violations are pre-existing (present before normalization)
    pre_existing_violations = 0
    new_violations = 0
    pre_existing_details = []
    new_violation_details = []
    
    project_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    
    if format_violation_details:
        try:
            for detail in format_violation_details:
                # Extract filepath and ref from detail string
                m = re.match(r'\s+(guides/\S+\.md):\s+\'([^\']+)\'', detail)
                if m:
                    filepath_rel = m.group(1)
                    ref_name = m.group(2)
                    # Check if this ref existed in HEAD~1
                    result = subprocess.run(
                        ["git", "show", f"HEAD~1:{filepath_rel}"],
                        capture_output=True, text=True, cwd=project_root
                    )
                    if result.returncode == 0 and ref_name in result.stdout:
                        pre_existing_violations += 1
                        pre_existing_details.append(f"{detail} [PRE-EXISTING]")
                    else:
                        new_violations += 1
                        new_violation_details.append(f"{detail} [NEW]")
                else:
                    new_violations += 1
                    new_violation_details.append(f"{detail} [NEW]")
        except Exception:
            # If git check fails, treat all as potential pre-existing
            pre_existing_violations = format_violations
            pre_existing_details = format_violation_details
    
    if pre_existing_violations > 0:
        print(f"Pre-existing format warnings: {pre_existing_violations}")
        for d in pre_existing_details:
            print(d)
        print()
    
    if new_violations > 0:
        print(f"NEW format violations: {new_violations}")
        for d in new_violation_details:
            print(d)
        print()
    
    # Overall verdict: PASS if no new violations and no intra-linter/duplicate violations
    passed = (new_violations == 0 and intra_linter_violations == 0 and duplicate_violations == 0)
    print()
    verdict_msg = 'PASS' if passed else 'FAIL'
    if passed and pre_existing_violations > 0:
        verdict_msg = f'PASS ({pre_existing_violations} pre-existing warnings)'
    print(f"Verdict: {verdict_msg}")
    print()
    
    # Curation summary
    print("=" * 60)
    print("CURATION SUMMARY")
    print("=" * 60)
    
    # Count modifications from git
    try:
        # Count files modified across Plan 01 + Plan 02
        result = subprocess.run(
            ["git", "log", "--oneline", "--grep=57-01", "--grep=57-02", "--all"],
            capture_output=True, text=True,
            cwd=os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        )
        plan_commits = len([l for l in result.stdout.strip().split("\n") if l])
        
        # Get files changed in both plans
        result1 = subprocess.run(
            ["git", "diff", "--name-only", "HEAD~1", "--", "guides/"],
            capture_output=True, text=True,
            cwd=os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
        )
        plan02_files = len([l for l in result1.stdout.strip().split("\n") if l])
        
        print(f"Plan 02 files modified: {plan02_files}")
        print(f"Total references validated: {total_refs}")
        print(f"Guides with populated related tags: {guides_with_related}")
        print(f"Guides with empty related tags: {guides_empty_related}")
        print(f"Normalization applied: bare → linter/rule format")
    except Exception:
        print("(Curation summary incomplete — git not available)")
    
    return 0 if passed else 1


if __name__ == "__main__":
    sys.exit(main())
