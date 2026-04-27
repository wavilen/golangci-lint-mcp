#!/usr/bin/env python3
"""
Normalize all guide <related> tags to canonical format.

Canonical format rules:
- Compound linter refs use linter/rule (e.g., gosec/G204, staticcheck/SA4006)
- Simple linter refs stay bare (e.g., errcheck, wrapcheck)
- Bare compound linter directory names stay bare per D-06 (e.g., bare "govet")
- Space-separated refs normalized to slash format (e.g., "staticcheck SA1029" → "staticcheck/SA1029")

Discovers compound linters from directory structure in guides/.
"""

import os
import re
import sys
import glob


GUIDES_ROOT = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), "guides")


def discover_compound_linters():
    """Scan guides/ for subdirectories containing .md files → compound linters."""
    linter_rules = {}  # linter_name → set of rule_names
    for entry in os.listdir(GUIDES_ROOT):
        dirpath = os.path.join(GUIDES_ROOT, entry)
        if os.path.isdir(dirpath):
            rules = set()
            for f in os.listdir(dirpath):
                if f.endswith(".md") and f != "_template.md":
                    rules.add(f[:-3])  # strip .md
            if rules:
                linter_rules[entry] = rules
    return linter_rules


def discover_simple_linters():
    """Flat .md files in guides/ root → simple linters."""
    simple = set()
    for f in os.listdir(GUIDES_ROOT):
        if f.endswith(".md") and f != "_template.md":
            simple.add(f[:-3])
    return simple


def build_reverse_lookup(linter_rules):
    """Map bare rule name → set of linter names that contain it."""
    reverse = {}
    for linter, rules in linter_rules.items():
        for rule in rules:
            if rule not in reverse:
                reverse[rule] = set()
            reverse[rule].add(linter)
    return reverse


def parse_related(content):
    """Extract <related> section content from a guide file.
    
    Handles both closed form (<related>...</related>) and open form (<related>...$).
    Returns (start_idx, end_idx, inner_content) or None.
    """
    # Try closed form first
    m = re.search(r'(<related>)(.*?)(</related>)', content, re.DOTALL)
    if m:
        return m.start(1), m.end(3), m.group(2)
    
    # Open form: <related> ... to end of file
    m = re.search(r'(<related>)(.*?)$', content, re.DOTALL)
    if m:
        return m.start(1), m.end(2), m.group(2)
    
    return None


def parse_refs(inner_content):
    """Parse comma-separated references from <related> inner content."""
    refs = []
    for r in inner_content.split(","):
        r = r.strip()
        if r:
            refs.append(r)
    return refs


def format_related(refs):
    """Format references back into <related> tag content."""
    if not refs:
        return "\n\n"
    return "\n\n" + ", ".join(refs) + "\n\n"


def get_guide_linter(filepath):
    """Determine which compound linter a guide belongs to (if any)."""
    relpath = os.path.relpath(filepath, GUIDES_ROOT)
    parts = relpath.split(os.sep)
    if len(parts) >= 2:
        return parts[0]  # subdirectory = compound linter name
    return None


def normalize_ref(ref, guide_linter, reverse_lookup, compound_linters, simple_linters):
    """Normalize a single reference. Returns (normalized_ref, changed_bool, reason)."""
    original = ref
    
    # Already canonical: contains /
    if "/" in ref:
        return ref, False, None
    
    # Space-separated compound ref: "staticcheck SA1029" → "staticcheck/SA1029"
    # Pattern: known_compound_linter + space + something
    for cl in compound_linters:
        if ref.startswith(cl + " ") and len(ref) > len(cl) + 1:
            rule_part = ref[len(cl) + 1:].strip()
            if rule_part in compound_linters[cl]:
                return f"{cl}/{rule_part}", True, f"space-sep → slash: '{original}' → '{cl}/{rule_part}'"
            # Even if not a known rule, normalize the format
            return f"{cl}/{rule_part}", True, f"space-sep → slash: '{original}' → '{cl}/{rule_part}'"
    
    # Check if it's a known bare rule in a compound linter
    if ref in reverse_lookup:
        matching_linters = reverse_lookup[ref]
        
        # Priority 1: same compound linter as the guide
        if guide_linter and guide_linter in matching_linters:
            return f"{guide_linter}/{ref}", True, f"intra-linter bare: '{original}' → '{guide_linter}/{ref}'"
        
        # Priority 2: exactly one matching compound linter
        if len(matching_linters) == 1:
            linter = next(iter(matching_linters))
            return f"{linter}/{ref}", True, f"cross-linter bare: '{original}' → '{linter}/{ref}'"
        
        # Multiple matches and guide isn't in any of them — ambiguous
        # Keep as-is to avoid incorrect normalization
        return ref, False, f"ambiguous: '{ref}' matches rules in {matching_linters}"
    
    # Check if it's a simple linter — keep bare
    if ref in simple_linters:
        return ref, False, None
    
    # Check if it's a compound linter directory name — keep bare per D-06
    if ref in compound_linters:
        return ref, False, None
    
    # Unknown reference — keep as-is
    return ref, False, f"unknown ref: '{ref}'"


def process_guide(filepath, reverse_lookup, compound_linters, simple_linters):
    """Process a single guide file. Returns (changes_list, modified_bool)."""
    guide_linter = get_guide_linter(filepath)
    
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()
    
    parsed = parse_related(content)
    if parsed is None:
        return [], False
    
    start, end, inner = parsed
    refs = parse_refs(inner)
    
    changes = []
    new_refs = []
    seen = set()
    
    for ref in refs:
        normalized, changed, reason = normalize_ref(
            ref, guide_linter, reverse_lookup, compound_linters, simple_linters
        )
        
        if changed:
            changes.append((ref, normalized, reason))
        
        # Deduplicate
        if normalized not in seen:
            new_refs.append(normalized)
            seen.add(normalized)
        else:
            changes.append((ref, normalized, f"duplicate removed: '{ref}'"))
    
    if not changes:
        return changes, False
    
    # Reconstruct the file
    new_inner = format_related(new_refs)
    new_content = content[:start] + "<related>" + new_inner + content[end:]
    
    # For open form (no </related> tag), adjust
    if "</related>" not in content[start:end]:
        # The parse_related open form returns up to end of content
        # We need to replace from start to end of file
        new_content = content[:start] + "<related>" + new_inner
    
    with open(filepath, "w", encoding="utf-8") as f:
        f.write(new_content)
    
    return changes, True


def main():
    print("=== Related Tag Normalization ===\n")
    
    # Discover structure
    compound_linters = discover_compound_linters()
    simple_linters = discover_simple_linters()
    reverse_lookup = build_reverse_lookup(compound_linters)
    
    print(f"Compound linters: {len(compound_linters)}")
    for name in sorted(compound_linters):
        print(f"  {name}: {len(compound_linters[name])} rules")
    print(f"Simple linters: {len(simple_linters)}")
    print()
    
    # Collect all guide files
    guide_files = []
    for f in glob.glob(os.path.join(GUIDES_ROOT, "*.md")):
        if os.path.basename(f) != "_template.md":
            guide_files.append(f)
    for f in glob.glob(os.path.join(GUIDES_ROOT, "**", "*.md"), recursive=True):
        if os.path.basename(f) != "_template.md" and os.path.dirname(f) != GUIDES_ROOT:
            guide_files.append(f)
    
    # Deduplicate
    guide_files = sorted(set(guide_files))
    print(f"Total guide files: {len(guide_files)}\n")
    
    # Process
    total_modified = 0
    total_normalizations = 0
    total_duplicates_removed = 0
    all_changes = []
    unrecognized = []
    
    for filepath in guide_files:
        try:
            changes, modified = process_guide(filepath, reverse_lookup, compound_linters, simple_linters)
        except Exception as e:
            print(f"ERROR processing {filepath}: {e}")
            continue
        
        if modified:
            total_modified += 1
            relpath = os.path.relpath(filepath, os.path.dirname(GUIDES_ROOT))
            for old, new, reason in changes:
                if "duplicate" in reason.lower():
                    total_duplicates_removed += 1
                else:
                    total_normalizations += 1
                all_changes.append((relpath, old, new, reason))
                if "unknown" in reason.lower() or "ambiguous" in reason.lower():
                    unrecognized.append((relpath, old, reason))
    
    # Change report
    print("=" * 60)
    print("CHANGE REPORT")
    print("=" * 60)
    print(f"Total files processed: {len(guide_files)}")
    print(f"Total files modified: {total_modified}")
    print(f"Total references normalized (bare → canonical): {total_normalizations}")
    print(f"Total duplicates removed: {total_duplicates_removed}")
    print()
    
    if all_changes:
        print("--- Detailed Changes ---")
        for relpath, old, new, reason in all_changes:
            print(f"  {relpath}: '{old}' → '{new}' ({reason})")
    
    if unrecognized:
        print()
        print("--- Unrecognized/Ambiguous References ---")
        for relpath, ref, reason in unrecognized:
            print(f"  {relpath}: {reason}")
    
    print()
    print(f"Result: {total_normalizations} normalizations + {total_duplicates_removed} dedup across {total_modified} files")
    
    return 0


if __name__ == "__main__":
    sys.exit(main())
