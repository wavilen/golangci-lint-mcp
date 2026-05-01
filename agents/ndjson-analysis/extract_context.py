#!/usr/bin/env python3
"""Extract per-trace structured context from ndjson traces and result JSON files.

Reads:
  tmp/ndjson/*.ndjson       — opencode session traces
  tmp/ndjson/*-result.json  — test result summaries

Writes:
  tmp/ndjson_analysis/<trace>-context.json — per-trace structured context

Located at: agents/ndjson-analysis/extract_context.py (git-tracked)
Input/output paths are project-root-relative.
"""

import json
import os
import re
import sys
from dataclasses import dataclass, field


@dataclass
class Event:
    timestamp: int
    event_type: str  # "step_start", "step_finish", "tool_use", "text"
    tool_name: str | None = None
    tool_input: dict = field(default_factory=dict)
    tool_output: str = ""  # Full tool output text (NO truncation)
    text_content: str = ""  # Agent reasoning text (NO truncation)
    tokens: dict = field(default_factory=dict)
    step_reason: str = ""
    session_id: str = ""


def parse_ndjson_full(filepath: str) -> list[Event]:
    """Parse NDJSON file with NO truncation of text or output content."""
    events = []
    with open(filepath) as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                obj = json.loads(line)
            except json.JSONDecodeError:
                continue

            t = obj.get("type", "")
            ts = obj.get("timestamp", 0)
            part = obj.get("part", {})
            sid = obj.get("sessionID", "")

            if t == "step_start":
                events.append(
                    Event(
                        timestamp=ts,
                        event_type="step_start",
                        session_id=sid,
                    )
                )
            elif t == "step_finish":
                tokens = part.get("tokens", {})
                reason = part.get("reason", "")
                events.append(
                    Event(
                        timestamp=ts,
                        event_type="step_finish",
                        tokens=tokens,
                        step_reason=reason,
                        session_id=sid,
                    )
                )
            elif t == "tool_use":
                tool = part.get("tool", "")
                state = part.get("state", {})
                inp = state.get("input", {})
                out = state.get("output", "")
                # NO truncation — keep full content
                if not isinstance(out, str):
                    out = str(out)
                events.append(
                    Event(
                        timestamp=ts,
                        event_type="tool_use",
                        tool_name=tool,
                        tool_input=inp,
                        tool_output=out,
                        session_id=sid,
                    )
                )
            elif t == "text":
                text = part.get("text", "")
                # NO truncation — keep full text
                events.append(
                    Event(
                        timestamp=ts,
                        event_type="text",
                        text_content=text,
                        session_id=sid,
                    )
                )
    return events


def extract_strategy_blocks(output: str) -> list[str]:
    """Extract <strategy_instructions>...</strategy_instructions> blocks from tool output."""
    return re.findall(
        r"<strategy_instructions>(.*?)</strategy_instructions>",
        output,
        re.DOTALL,
    )


def normalize_tool_name(tool_name: str) -> str:
    """Normalize tool name for breakdown: strip golangci-lint_ prefix for MCP tools."""
    if tool_name.startswith("golangci-lint_"):
        # golangci-lint_golangci_lint_run → golangci_lint_run
        return tool_name.replace("golangci-lint_", "", 1)
    return tool_name


def extract_strategy(events: list[Event]) -> str:
    """Extract strategy type from first tool_use output containing strategy info.

    Looks for:
    1. <strategy_instructions> block → extract Strategy: line
    2. Fallback: any "Strategy:" line in tool outputs
    """
    for ev in events:
        if ev.event_type != "tool_use" or not ev.tool_output:
            continue

        # Check for <strategy_instructions> blocks
        blocks = extract_strategy_blocks(ev.tool_output)
        for block in blocks:
            for line in block.split("\n"):
                stripped = line.strip()
                if stripped.lower().startswith("strategy:"):
                    return stripped.replace("Strategy:", "").strip()
                if "strategy" in stripped.lower() and ("subagent" in stripped.lower() or "single" in stripped.lower()):
                    # Take a meaningful strategy description
                    return stripped.split("—")[0].strip()

        # Also check raw output for Strategy: line (from MCP tool summary)
        for line in ev.tool_output.split("\n"):
            stripped = line.strip()
            # Handle "- Strategy:" and "Strategy:" prefixes
            if "Strategy:" in stripped:
                strat_part = stripped[stripped.index("Strategy:") :]
                return strat_part.replace("Strategy:", "").strip()

    return "none"


def extract_context(base: str, events: list[Event], result: dict) -> dict:
    """Extract per-trace structured context from events and result JSON.

    Produces all D-06 fields:
    - model, test_case, duration_ms, event_counts, tool_breakdown,
      token_totals, issue_counts, timestamps, subagent_count, strategy, pass
    """
    if not events:
        return {}

    # Timestamps
    start_ts = events[0].timestamp
    end_ts = events[-1].timestamp
    duration_ms = end_ts - start_ts

    # Event counts
    event_counts = {"step_start": 0, "tool_use": 0, "text": 0, "step_finish": 0}
    for ev in events:
        if ev.event_type in event_counts:
            event_counts[ev.event_type] += 1

    # Tool breakdown: tool_name → count (normalized)
    tool_breakdown = {}
    subagent_count = 0
    for ev in events:
        if ev.event_type == "tool_use" and ev.tool_name:
            if ev.tool_name == "task":
                subagent_count += 1
            norm_name = normalize_tool_name(ev.tool_name)
            tool_breakdown[norm_name] = tool_breakdown.get(norm_name, 0) + 1

    # Token totals: sum across all step_finish events
    token_totals = {"input": 0, "output": 0, "reasoning": 0, "cache_read": 0}
    for ev in events:
        if ev.event_type == "step_finish":
            tokens = ev.tokens or {}
            token_totals["input"] += tokens.get("input", 0)
            token_totals["output"] += tokens.get("output", 0)
            token_totals["reasoning"] += tokens.get("reasoning", 0)
            cache_read = tokens.get("cache", {}).get("read", 0)
            token_totals["cache_read"] += cache_read

    # Issue counts from result JSON
    issue_counts = {
        "before": result.get("BeforeIssues", 0),
        "after": result.get("AfterIssues", 0),
        "reduction_pct": result.get("IssueReduction", 0),
    }

    # Strategy extraction
    strategy = extract_strategy(events)

    ctx = {
        "model": result.get("Model", "unknown"),
        "test_case": result.get("TestCase", "unknown"),
        "duration_ms": duration_ms,
        "event_counts": event_counts,
        "tool_breakdown": tool_breakdown,
        "token_totals": token_totals,
        "issue_counts": issue_counts,
        "timestamps": {"start": start_ts, "end": end_ts},
        "subagent_count": subagent_count,
        "strategy": strategy,
        "pass": result.get("Pass", False),
    }

    return ctx


def main():
    project_root = os.path.abspath(os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", ".."))
    ndjson_dir = os.path.join(project_root, "tmp", "ndjson")
    out_dir = os.path.join(project_root, "tmp", "ndjson_analysis")

    if not os.path.isdir(ndjson_dir):
        print(f"Error: ndjson directory not found: {ndjson_dir}")
        sys.exit(1)

    files = sorted(f for f in os.listdir(ndjson_dir) if f.endswith(".ndjson") and "probe" not in f)

    count = 0

    for fname in files:
        ndjson_path = os.path.join(ndjson_dir, fname)
        stem = fname.replace(".ndjson", "")

        result_path = os.path.join(ndjson_dir, stem + "-result.json")
        # Retry traces (e.g., large-glm-4-7-r1) use the base trace's result file
        if not os.path.exists(result_path) and stem.endswith("-r1"):
            base_stem = stem[:-3]  # Remove "-r1"
            base_result = os.path.join(ndjson_dir, base_stem + "-result.json")
            if os.path.exists(base_result):
                result_path = base_result

        result = {}
        if os.path.exists(result_path):
            with open(result_path) as f:
                result = json.load(f)

        print(f"Processing: {fname}")
        events = parse_ndjson_full(ndjson_path)
        if not events:
            print("  Warning: No events found, skipping")
            continue

        ctx = extract_context(stem, events, result)
        if not ctx:
            print("  Warning: Could not extract context, skipping")
            continue

        ctx_path = os.path.join(out_dir, f"{stem}-context.json")
        with open(ctx_path, "w") as f:
            json.dump(ctx, f, indent=2, ensure_ascii=False)

        count += 1
        print(f"  Written: {ctx_path}")

    print(f"\nExtracted context for {count} traces.")


if __name__ == "__main__":
    main()
