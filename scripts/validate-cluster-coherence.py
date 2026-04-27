#!/usr/bin/env python3
"""
Validate cluster coherence of curated <related> tags.

Compares curated <related> tags against graph.json cluster structure to confirm
that curation produced meaningful cluster groupings without noise or trivial
relationships (RCUR-03 validation).

Steps:
1. Parse CLUSTER-ANALYSIS.md cluster membership (10 clusters)
2. Load graph.json relationship data for intra-cluster link counts
3. Read curated <related> tags from all 629 guides
4. Compute per-cluster coherence metrics
5. Report PASS/FAIL verdict

Output: Per-cluster table + overall summary + PASS/FAIL verdict.
Exit 0 for PASS, exit 1 for FAIL.
"""

import json
import os
import re
import sys


SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
PROJECT_ROOT = os.path.dirname(SCRIPT_DIR)
GUIDES_ROOT = os.path.join(PROJECT_ROOT, "guides")
GRAPH_JSON = os.path.join(PROJECT_ROOT, "graphify-out", "graph.json")
CLUSTER_ANALYSIS = os.path.join(PROJECT_ROOT, "graphify-out", "CLUSTER-ANALYSIS.md")


def parse_cluster_membership(filepath):
    """Parse CLUSTER-ANALYSIS.md to extract cluster name -> set of member IDs.

    Only reads from ### Member List sections — ignores Curation Guidance
    and other sections that also have backtick-quoted entries.

    Returns dict: {cluster_name: set_of_member_ids}
    Member IDs are normalized to use slashes (e.g., 'gosec/G101').
    """
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    clusters = {}
    current_cluster = None
    in_member_list = False

    for line in content.split("\n"):
        # Match cluster headers: ## Cluster: Security Auditing
        header = re.match(r"^## Cluster:\s+(.+)$", line)
        if header:
            current_cluster = header.group(1).strip()
            clusters[current_cluster] = set()
            in_member_list = False
            continue

        # Track when we enter the Member List section
        if current_cluster and line.strip() == "### Member List":
            in_member_list = True
            continue

        # Any other ### header ends the member list section
        if line.startswith("### "):
            in_member_list = False
            continue

        # Match member list entries: - `gosec/G101` — description
        if in_member_list and line.startswith("- `"):
            member_match = re.match(r"^-\s+`([^`]+)`", line)
            if member_match:
                member_id = member_match.group(1).strip()
                clusters[current_cluster].add(member_id)

    return clusters


def normalize_to_node_id(ref, compound_linters):
    """Normalize a related-tag reference to a graph.json node ID.

    graph.json uses underscores: gosec_G101, gocritic_hugeParam
    Related tags use slashes: gosec/G101, gocritic/hugeParam
    Simple linters: errcheck, prealloc (same in both)

    Returns the node ID (underscore form).
    """
    if "/" in ref:
        # Already in linter/rule format — convert to underscore
        return ref.replace("/", "_", 1)
    return ref


def node_id_to_display(node_id):
    """Convert graph.json node ID to display format (slash for compound)."""
    if "_" in node_id:
        parts = node_id.split("_", 1)
        # Check if it's a known compound linter pattern
        # Graph.json uses e.g. gosec_G101, gocritic_hugeParam
        return parts[0] + "/" + parts[1]
    return node_id


def load_graph_links(filepath):
    """Load graph.json and return all links as list of (source, target, weight) tuples."""
    with open(filepath, "r", encoding="utf-8") as f:
        data = json.load(f)

    links = []
    for link in data.get("links", []):
        src = link.get("source", "")
        tgt = link.get("target", "")
        weight = link.get("weight", 0)
        links.append((src, tgt, weight))

    return links


def build_node_community_map(filepath):
    """Build node_id -> community number from graph.json nodes."""
    with open(filepath, "r", encoding="utf-8") as f:
        data = json.load(f)

    node_comm = {}
    for node in data.get("nodes", []):
        node_comm[node["id"]] = node.get("community", -1)

    return node_comm


def discover_compound_linters():
    """Scan guides/ for subdirectories containing .md files.

    Returns dict: {linter_name: set_of_rule_names}
    """
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


def parse_related(content):
    """Extract <related> section content. Returns inner text or None."""
    m = re.search(r"<related>(.*?)</related>", content, re.DOTALL)
    if m:
        return m.group(1)
    # Open form (no closing tag — tag is last section in file)
    m = re.search(r"<related>\s*\n?(.*?)$", content, re.DOTALL)
    if m:
        return m.group(1)
    return None


def parse_refs(inner_content):
    """Parse comma-separated references from <related> inner content."""
    refs = []
    for r in inner_content.split(","):
        r = r.strip()
        if r:
            refs.append(r)
    return refs


def get_guide_node_id(filepath):
    """Determine the graph.json node ID for a guide file.

    guides/errcheck.md -> 'errcheck'
    guides/gosec/G101.md -> 'gosec_G101'
    """
    relpath = os.path.relpath(filepath, GUIDES_ROOT)
    parts = relpath.split(os.sep)
    if len(parts) >= 2:
        # Compound linter: guides/gosec/G101.md -> gosec_G101
        return parts[0] + "_" + parts[1].replace(".md", "")
    else:
        # Simple linter: guides/errcheck.md -> errcheck
        return parts[0].replace(".md", "")


def read_all_guide_related_tags():
    """Read <related> tags from all guide files.

    Returns dict: {node_id: set_of_related_node_ids}
    Related refs are normalized to node IDs (underscore form).
    """
    guide_related = {}

    for entry in os.listdir(GUIDES_ROOT):
        entry_path = os.path.join(GUIDES_ROOT, entry)
        if os.path.isdir(entry_path):
            # Compound linter directory
            for fname in os.listdir(entry_path):
                if fname.endswith(".md") and fname != "_template.md":
                    filepath = os.path.join(entry_path, fname)
                    node_id = entry + "_" + fname[:-3]
                    _read_guide_related(filepath, node_id, guide_related)
        elif entry.endswith(".md") and entry != "_template.md":
            # Simple linter file
            filepath = entry_path
            node_id = entry[:-3]
            _read_guide_related(filepath, node_id, guide_related)

    return guide_related


def _read_guide_related(filepath, node_id, result_dict):
    """Read a single guide file and extract its related refs."""
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            content = f.read()
    except (IOError, OSError):
        return

    inner = parse_related(content)
    if inner is None:
        result_dict[node_id] = set()
        return

    refs = parse_refs(inner)
    normalized = set()
    for ref in refs:
        # Normalize to node ID (underscore form for compound)
        node_ref = normalize_to_node_id(ref, None)
        normalized.add(node_ref)

    result_dict[node_id] = normalized


def compute_cluster_member_node_ids(cluster_members_display):
    """Convert cluster member IDs from display format to node IDs.

    CLUSTER-ANALYSIS uses slash format: gosec/G101
    graph.json uses underscore format: gosec_G101

    Returns set of node IDs.
    """
    node_ids = set()
    for member in cluster_members_display:
        if "/" in member:
            node_ids.add(member.replace("/", "_", 1))
        else:
            node_ids.add(member)
    return node_ids


def compute_intra_cluster_graph_links(links, cluster_node_ids):
    """Count intra-cluster links from graph.json.

    A link is intra-cluster if both source and target are in the cluster.
    Returns count of such links.
    """
    count = 0
    for src, tgt, weight in links:
        if src in cluster_node_ids and tgt in cluster_node_ids:
            count += 1
    return count


def compute_tagged_intra_cluster_refs(guide_related, cluster_node_ids):
    """Compute intra-cluster related-tag cross-references.

    For each guide in the cluster, count how many of its <related> refs
    point to other members of the same cluster.

    Returns:
        tagged_pairs: count of directed intra-cluster related-tag refs
        isolated_members: set of node IDs with zero intra-cluster refs
        per_member_refs: dict {node_id: count_of_intra_cluster_refs}
    """
    tagged_pairs = 0
    isolated_members = set()
    per_member_refs = {}

    for node_id in cluster_node_ids:
        refs = guide_related.get(node_id, set())
        intra_refs = refs & cluster_node_ids
        per_member_refs[node_id] = len(intra_refs)
        tagged_pairs += len(intra_refs)
        if len(intra_refs) == 0:
            isolated_members.add(node_id)

    return tagged_pairs, isolated_members, per_member_refs


def main():
    print("=== Cluster Coherence Validation ===")
    print("RCUR-03: Re-analyze curated guides to confirm coherent clusters\n")

    # Step 1: Parse cluster membership from CLUSTER-ANALYSIS.md
    print("Step 1: Parsing CLUSTER-ANALYSIS.md cluster membership...")
    clusters = parse_cluster_membership(CLUSTER_ANALYSIS)
    print(f"  Found {len(clusters)} clusters")
    for name, members in sorted(clusters.items(), key=lambda x: -len(x[1])):
        print(f"    {name}: {len(members)} members")

    if len(clusters) != 10:
        print(f"  WARNING: Expected 10 clusters, found {len(clusters)}")

    # Step 2: Load graph.json relationship data
    print("\nStep 2: Loading graph.json relationship data...")
    links = load_graph_links(GRAPH_JSON)
    print(f"  Total links in graph.json: {len(links)}")

    node_comm = build_node_community_map(GRAPH_JSON)
    print(f"  Total nodes in graph.json: {len(node_comm)}")

    # Step 3: Read curated <related> tags from all guides
    print("\nStep 3: Reading <related> tags from all guides...")
    guide_related = read_all_guide_related_tags()
    guides_with_refs = sum(1 for refs in guide_related.values() if refs)
    guides_empty = sum(1 for refs in guide_related.values() if not refs)
    total_guides = len(guide_related)
    print(f"  Total guides: {total_guides}")
    print(f"  Guides with related refs: {guides_with_refs}")
    print(f"  Guides with empty related: {guides_empty}")

    # Step 4: Compute per-cluster coherence metrics
    print("\nStep 4: Computing per-cluster coherence metrics...\n")

    cluster_results = []
    for cluster_name, members_display in sorted(clusters.items(), key=lambda x: -len(x[1])):
        member_count = len(members_display)
        cluster_node_ids = compute_cluster_member_node_ids(members_display)

        # Intra-cluster graph links
        graph_links = compute_intra_cluster_graph_links(links, cluster_node_ids)

        # Intra-cluster related-tag cross-references
        tagged_pairs, isolated_members, per_member_refs = compute_tagged_intra_cluster_refs(
            guide_related, cluster_node_ids
        )

        # Coverage: what fraction of graph-identified relationships are captured
        coverage = tagged_pairs / max(graph_links, 1)

        # Average refs per member
        avg_refs = tagged_pairs / max(member_count, 1)

        # Find guides in cluster that we have data for
        guides_in_cluster_with_data = sum(
            1 for nid in cluster_node_ids if nid in guide_related
        )

        cluster_results.append({
            "name": cluster_name,
            "members": member_count,
            "graph_links": graph_links,
            "tagged_links": tagged_pairs,
            "coverage": coverage,
            "isolated_count": len(isolated_members),
            "isolated_members": isolated_members,
            "avg_refs": avg_refs,
            "guides_found": guides_in_cluster_with_data,
        })

    # Print per-cluster table
    header = f"{'Cluster':<28} {'Members':>7} {'GLinks':>7} {'TLinks':>7} {'Cover%':>7} {'Isol':>5} {'AvgRef':>7}"
    print(header)
    print("-" * len(header))

    for r in cluster_results:
        cover_pct = f"{r['coverage']*100:.1f}%"
        print(
            f"{r['name']:<28} {r['members']:>7} {r['graph_links']:>7} "
            f"{r['tagged_links']:>7} {cover_pct:>7} {r['isolated_count']:>5} "
            f"{r['avg_refs']:>7.2f}"
        )

    # Print overall summary
    total_members = sum(r["members"] for r in cluster_results)
    total_graph_links = sum(r["graph_links"] for r in cluster_results)
    total_tagged_links = sum(r["tagged_links"] for r in cluster_results)
    total_isolated = sum(r["isolated_count"] for r in cluster_results)
    avg_coverage = total_tagged_links / max(total_graph_links, 1)

    print()
    print("=" * 60)
    print("OVERALL SUMMARY")
    print("=" * 60)
    print(f"Total clusters: {len(cluster_results)}")
    print(f"Total cluster members: {total_members}")
    print(f"Total intra-cluster graph links: {total_graph_links}")
    print(f"Total intra-cluster tagged links: {total_tagged_links}")
    print(f"Average coverage: {avg_coverage*100:.1f}%")
    print(f"Total isolated members: {total_isolated}")

    # Print isolated members details
    for r in cluster_results:
        if r["isolated_members"]:
            print(f"\n  Isolated in '{r['name']}' ({r['isolated_count']} members):")
            for m in sorted(r["isolated_members"])[:15]:
                display = node_id_to_display(m)
                print(f"    {display}")
            if r["isolated_count"] > 15:
                print(f"    ... and {r['isolated_count'] - 15} more")

    # Step 5: Verdict
    print()
    print("=" * 60)
    print("VERDICT")
    print("=" * 60)

    # Check PASS conditions:
    # 1. All 10 clusters have at least some related-tag cross-refs
    empty_clusters = [r["name"] for r in cluster_results if r["tagged_links"] == 0]

    # 2. Clusters with >0 graph.json intra-cluster links show >0% coverage
    zero_coverage_clusters = [
        r["name"] for r in cluster_results
        if r["graph_links"] > 0 and r["tagged_links"] == 0
    ]

    # 3. No cluster has >50% isolated members
    high_isolation_clusters = [
        r["name"] for r in cluster_results
        if r["isolated_count"] > r["members"] * 0.5
    ]

    passed = True
    reasons = []

    if empty_clusters:
        passed = False
        reasons.append(
            f"Empty clusters (zero tagged links): {', '.join(empty_clusters)}"
        )

    if zero_coverage_clusters:
        passed = False
        reasons.append(
            f"Zero coverage clusters (graph links but no tagged refs): "
            f"{', '.join(zero_coverage_clusters)}"
        )

    if high_isolation_clusters:
        passed = False
        reasons.append(
            f"High isolation clusters (>50% isolated): "
            f"{', '.join(high_isolation_clusters)}"
        )

    if passed:
        print("PASS — All clusters show coherent related-tag structure")
        print()
        for r in cluster_results:
            cover_pct = f"{r['coverage']*100:.1f}%"
            iso_pct = f"{r['isolated_count']/max(r['members'],1)*100:.0f}%"
            print(f"  {r['name']}: {cover_pct} coverage, {iso_pct} isolated")
    else:
        print("FAIL — Cluster coherence issues detected:")
        for reason in reasons:
            print(f"  - {reason}")

    print()

    return 0 if passed else 1


if __name__ == "__main__":
    sys.exit(main())
