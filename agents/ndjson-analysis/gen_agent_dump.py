#!/usr/bin/env python3
"""Parse opencode ndjson traces and generate human-readable agent dump files.

Shows exactly what each model agent received (prompt, strategy instructions, tool outputs)
and how it responded (reasoning text, tool calls made), enabling side-by-side debugging
of why GLM-4.7 and GLM-5-Turbo ignore the subagent strategy while GLM-5.1 follows it.

Reads:
  tmp/ndjson/*.ndjson       — opencode session traces
  tmp/ndjson/*-result.json  — test result summaries

Writes:
  tmp/ndjson_analysis/<trace>-agent-dump.txt  — per-trace human-readable dump
  tmp/ndjson_analysis/strategy-comparison.txt — compliance comparison table

Located at: agents/ndjson-analysis/gen_agent_dump.py (git-tracked)
Input/output paths are project-root-relative.
"""

import json
import os
import re
import sys
from dataclasses import dataclass, field

# The fixedPrompt from e2e/e2e_test.go — what every agent receives as its task
FIXED_PROMPT = (
    "Use the golangci-lint-guide skill to fix all golangci-lint issues in the current directory. "
    'Follow the skill process: call golangci_lint_run(path="./...") to find issues, apply fixes '
    "using the guidance returned, then verify with golangci_lint_run again. Do not add nolint "
    "directives. After the first scan, if issues remain, use per-package paths for targeted fixes.\n"
    "\n"
    "IMPORTANT: This project uses golangci-lint v2. The CLI flags changed from v1. Do NOT use "
    '--out-format json or --output=json — these are v1 flags and will fail with "unknown flag". '
    "Always use the MCP tool golangci_lint_run instead of running golangci-lint directly. "
    "If you must use the CLI, the v2 JSON output flag is: --output.json.path stdout"
)


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


def fmt_duration(ms: int) -> str:
    if ms < 1000:
        return f"{ms}ms"
    s = ms / 1000
    if s < 60:
        return f"{s:.1f}s"
    m = int(s // 60)
    sec = s % 60
    return f"{m}m{sec:.0f}s"


def extract_strategy_blocks(output: str) -> list[str]:
    """Extract <strategy_instructions>...</strategy_instructions> blocks from tool output."""
    return re.findall(
        r"<strategy_instructions>(.*?)</strategy_instructions>",
        output,
        re.DOTALL,
    )


def extract_strategy_line(summary_text: str) -> str:
    """Extract the Strategy: line from a summary section."""
    for line in summary_text.split("\n"):
        if "Strategy:" in line:
            return line.strip()
    return ""


def tool_label(event: Event) -> str:
    """Human-readable label for a tool call."""
    tool = event.tool_name or ""
    inp = event.tool_input

    if "golangci-lint" in tool:
        if "golangci_lint_list" in tool:
            return "golangci_lint_list()"
        if "golangci_lint_guide" in tool:
            linter = inp.get("linter", "")
            rule = inp.get("rule", "")
            return f"golangci_lint_guide({linter}/{rule})" if linter else "golangci_lint_guide()"
        if "golangci_lint_parse" in tool:
            return "golangci_lint_parse()"
        if "golangci_lint_summarize" in tool:
            return "golangci_lint_summarize()"
        path = inp.get("path", "")
        return f"golangci_lint_run({path})"
    if tool == "edit":
        fp = inp.get("filePath", "")
        base = os.path.basename(fp)
        return f"edit({base})"
    if tool == "read":
        fp = inp.get("filePath", "")
        base = os.path.basename(fp)
        return f"read({base})"
    if tool == "write":
        fp = inp.get("filePath", "")
        base = os.path.basename(fp)
        return f"write({base})"
    if tool == "bash":
        cmd = inp.get("command", "")
        desc = inp.get("description", "")
        if desc:
            return f"bash: {desc[:60]}"
        return f"bash: {cmd[:60]}"
    if tool == "glob":
        pat = inp.get("pattern", "")
        return f"glob({pat})"
    if tool == "todowrite":
        todos = inp.get("todos", [])
        done = sum(1 for t in todos if t.get("status") == "completed")
        return f"todowrite({done}/{len(todos)} done)"
    if tool == "skill":
        name = inp.get("name", "")
        return f"skill({name})"
    if tool == "task":
        desc = inp.get("description", "")
        return f"task: {desc[:60]}"
    if tool == "question":
        return "question()"
    return tool


def format_tokens(tokens: dict) -> str:
    """Format token dict for display."""
    parts = []
    total = tokens.get("total", 0)
    inp = tokens.get("input", 0)
    out = tokens.get("output", 0)
    reason = tokens.get("reasoning", 0)
    if total:
        parts.append(f"total={total}")
    if inp:
        parts.append(f"in={inp}")
    if out:
        parts.append(f"out={out}")
    if reason:
        parts.append(f"reason={reason}")
    cache_read = tokens.get("cache", {}).get("read", 0)
    if cache_read:
        parts.append(f"cache_read={cache_read}")
    return ", ".join(parts) if parts else "N/A"


SEPARATOR = "=" * 60
THIN_SEP = "─" * 60


def generate_agent_dump(name: str, events: list[Event], result: dict, output_path: str):
    """Generate a human-readable .txt dump file for a single trace."""
    if not events:
        return

    base_ts = events[0].timestamp
    total_ms = events[-1].timestamp - base_ts

    # Count subagent (task) calls
    task_calls = [e for e in events if e.event_type == "tool_use" and e.tool_name == "task"]

    # Collect all strategy blocks received
    strategy_blocks = []
    strategy_sources = []
    for ev in events:
        if ev.event_type == "tool_use" and ev.tool_output:
            blocks = extract_strategy_blocks(ev.tool_output)
            for block in blocks:
                strategy_blocks.append(block)
                strategy_line = extract_strategy_line(ev.tool_output)
                strategy_sources.append(
                    {
                        "tool": tool_label(ev),
                        "strategy_line": strategy_line,
                    }
                )

    lines = []

    # ═══ HEADER ═══
    lines.append(SEPARATOR)
    lines.append(f"AGENT DUMP: {name}")
    lines.append(SEPARATOR)
    lines.append(f"Model: {result.get('Model', '?')}")
    lines.append(f"Test Case: {result.get('TestCase', '?')}")
    lines.append(f"Before Issues: {result.get('BeforeIssues', 0)}")
    lines.append(f"After Issues: {result.get('AfterIssues', 0)}")
    lines.append(f"Issue Reduction: {result.get('IssueReduction', 0)}%")
    lines.append(f"Pass: {result.get('Pass', False)}")
    lines.append(f"Tool Calls: {result.get('ToolCalls', 0)}")
    lines.append(f"Retries: {result.get('Retries', 0)}")
    lines.append(f"Duration: {fmt_duration(total_ms)}")
    lines.append(f"Subagent Calls: {len(task_calls)}")
    lines.append(SEPARATOR)
    lines.append("")

    # ═══ PROMPT ═══
    lines.append(SEPARATOR)
    lines.append("PROMPT (from opencode run --format json)")
    lines.append(SEPARATOR)
    lines.append(FIXED_PROMPT)
    lines.append(SEPARATOR)
    lines.append("")

    # ═══ CONVERSATION ═══
    lines.append(SEPARATOR)
    lines.append("CONVERSATION")
    lines.append(SEPARATOR)
    lines.append("")

    # Track text events after strategy blocks for compliance reporting
    text_after_strategy = []
    strategy_received = False

    step_idx = 0
    current_step_events: list[Event] = []

    for ev in events:
        if ev.event_type == "step_start":
            current_step_events = []
            step_idx += 1
        elif ev.event_type in ("tool_use", "text"):
            current_step_events.append(ev)
        if ev.event_type == "step_finish":
            step_tools = [e for e in current_step_events if e.event_type == "tool_use"]
            step_texts = [e for e in current_step_events if e.event_type == "text"]

            if not step_tools and not step_texts:
                continue

            lines.append(f"  ┌─ STEP {step_idx} ────────────────────────────────────────")
            lines.append("")

            # Agent reasoning text
            for te in step_texts:
                txt = te.text_content.strip()
                if not txt:
                    continue
                lines.append(f"  ── AGENT REASONING {'─' * 40}")
                # Print full text content, indented
                for tline in txt.split("\n"):
                    lines.append(f"  {tline}")
                lines.append(f"  {'─' * 58}")
                lines.append("")

                if strategy_received:
                    text_after_strategy.append(txt)

            # Tool calls
            for te in step_tools:
                tl = tool_label(te)
                lines.append(f"  ── TOOL CALL {'─' * 44}")
                lines.append(f"  Tool: {tl}")
                if te.tool_input:
                    inp_str = json.dumps(te.tool_input, indent=4, ensure_ascii=False)
                    for iline in inp_str.split("\n"):
                        lines.append(f"  {iline}")
                lines.append(f"  ── TOOL OUTPUT {'─' * 41}")

                # Check for strategy_instructions in output
                out = te.tool_output or ""
                strat_blocks = extract_strategy_blocks(out)

                if strat_blocks:
                    # Print output before strategy block
                    before_strat = out.split("<strategy_instructions>")[0].rstrip()
                    if before_strat.strip():
                        for oline in before_strat.split("\n"):
                            lines.append(f"  {oline}")
                        lines.append("")

                    for sb in strat_blocks:
                        lines.append("  ╔══════════════════════════════════════════════════════════════╗")
                        lines.append("  ║  ⚡ STRATEGY INSTRUCTIONS RECEIVED BY AGENT                  ║")
                        lines.append("  ╚══════════════════════════════════════════════════════════════╝")
                        for sline in sb.strip().split("\n"):
                            lines.append(f"  {sline}")
                        lines.append("")
                        strategy_received = True

                    # Print output after strategy block
                    after_strat = out.split("</strategy_instructions>")[-1].lstrip()
                    if after_strat.strip() and "</strategy_instructions>" in out:
                        for oline in after_strat.split("\n"):
                            lines.append(f"  {oline}")
                else:
                    # Truncate very long outputs for readability, but show first 500 lines
                    out_lines = out.split("\n")
                    max_lines = 500
                    if len(out_lines) > max_lines:
                        for oline in out_lines[:max_lines]:
                            lines.append(f"  {oline}")
                        lines.append(f"  ... [{len(out_lines) - max_lines} more lines truncated] ...")
                    else:
                        for oline in out_lines:
                            lines.append(f"  {oline}")

                lines.append(f"  {'─' * 58}")
                lines.append("")

            # Step finish
            lines.append(f"  ── STEP END (reason: {ev.step_reason}, tokens: {format_tokens(ev.tokens)}) ──")
            lines.append(f"  └{'─' * 58}")
            lines.append("")

    # ═══ COMPLIANCE REPORT ═══
    lines.append("")
    lines.append(SEPARATOR)
    lines.append("STRATEGY COMPLIANCE REPORT")
    lines.append(SEPARATOR)

    strat_count = len(strategy_blocks)
    task_count = len(task_calls)

    # Determine strategy type and compliance
    strategy_type = "none"
    if strat_count == 0:
        compliance = "N/A"
    elif any("subagent" in sb.lower() for sb in strategy_blocks):
        strategy_type = "subagent-per-file"
        compliance = "FOLLOWED" if task_count > 0 else "IGNORED"
    else:
        # Extract actual strategy line for display
        if strategy_sources:
            strategy_type = strategy_sources[0]["strategy_line"].replace("Strategy: ", "").split("—")[0].strip()
        else:
            strategy_type = "other"
        compliance = "N/A"

    lines.append(f"Strategy blocks received: {strat_count}")
    lines.append(f"Subagent calls made: {task_count}")
    lines.append(f"Compliance: {compliance}")
    lines.append("")

    lines.append("Strategy blocks were received in these tool outputs:")
    if strategy_sources:
        for _i, src in enumerate(strategy_sources):
            lines.append(f"  - {src['tool']}: Strategy: {src['strategy_line']}")
    else:
        lines.append("  (none)")
    lines.append("")

    lines.append("Agent responded with these task/subagent calls:")
    if task_calls:
        for tc in task_calls:
            desc = tc.tool_input.get("description", "(no description)")
            lines.append(f"  - task: {desc}")
    else:
        lines.append("  NONE — agent did not spawn any subagents")
    lines.append("")

    lines.append("Agent reasoning after receiving strategy (first 500 chars of text events following strategy block):")
    if text_after_strategy:
        combined = " | ".join(text_after_strategy)
        lines.append(f"  {combined[:500]}")
    else:
        lines.append("  (no reasoning text after strategy)")
    lines.append("")
    lines.append(SEPARATOR)

    with open(output_path, "w") as f:
        f.write("\n".join(lines))

    print(f"  Written: {output_path}")

    return {
        "name": name,
        "model": result.get("Model", "?"),
        "strat_count": strat_count,
        "task_count": task_count,
        "compliance": compliance,
        "passed": result.get("Pass", False),
        "strategy_type": strategy_type,
    }


def generate_strategy_comparison(summaries: list[dict], output_path: str):
    """Generate strategy-comparison.txt with a summary table."""
    lines = []
    lines.append("STRATEGY COMPLIANCE COMPARISON")
    lines.append("=" * 90)
    lines.append("")

    # Table header
    hdr = f"{'Trace':<25} {'Model':<15} {'Strategy':<20} {'Tasks':<7} {'Compliance':<12} {'Pass':<6}"
    lines.append(hdr)
    lines.append("-" * 90)

    for s in summaries:
        row = (
            f"{s['name']:<25} {s['model']:<15} {s['strategy_type']:<20} "
            f"{s['task_count']:<7} {s['compliance']:<12} {s['passed']!s:<6}"
        )
        lines.append(row)

    lines.append("")
    lines.append("=" * 90)
    lines.append("")

    # Summary analysis
    ignored = [s for s in summaries if s["compliance"] == "IGNORED"]
    followed = [s for s in summaries if s["compliance"] == "FOLLOWED"]
    na = [s for s in summaries if s["compliance"] not in ("IGNORED", "FOLLOWED")]

    lines.append("SUMMARY:")
    lines.append(f"  FOLLOWED strategy: {len(followed)} traces")
    for f in followed:
        lines.append(f"    - {f['name']} ({f['model']})")
    lines.append(f"  IGNORED strategy: {len(ignored)} traces")
    for ig in ignored:
        lines.append(f"    - {ig['name']} ({ig['model']})")
    if na:
        lines.append(f"  N/A (no strategy or different strategy): {len(na)} traces")
        for n in na:
            lines.append(f"    - {n['name']} ({n['model']}): {n['compliance']}")
    lines.append("")

    # Key insight
    lines.append("KEY INSIGHT:")
    models_ignored = set(s["model"] for s in ignored)
    models_followed = set(s["model"] for s in followed)
    if models_ignored:
        lines.append(f"  Models that IGNORED subagent strategy: {', '.join(sorted(models_ignored))}")
    if models_followed:
        lines.append(f"  Models that FOLLOWED subagent strategy: {', '.join(sorted(models_followed))}")
    lines.append("")

    with open(output_path, "w") as f:
        f.write("\n".join(lines))

    print(f"  Written: {output_path}")


def main():
    project_root = os.path.abspath(os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", ".."))
    ndjson_dir = os.path.join(project_root, "tmp", "ndjson")
    out_dir = os.path.join(project_root, "tmp", "ndjson_analysis")

    if not os.path.isdir(ndjson_dir):
        print(f"Error: ndjson directory not found: {ndjson_dir}")
        sys.exit(1)

    files = sorted(f for f in os.listdir(ndjson_dir) if f.endswith(".ndjson") and "probe" not in f)

    summaries = []

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
            print("  No events found, skipping")
            continue

        dump_path = os.path.join(out_dir, stem + "-agent-dump.txt")
        summary = generate_agent_dump(stem, events, result, dump_path)
        if summary:
            summaries.append(summary)

    # Generate comparison file
    if summaries:
        comp_path = os.path.join(out_dir, "strategy-comparison.txt")
        generate_strategy_comparison(summaries, comp_path)

    print(f"\nDone. {len(summaries)} traces processed.")


if __name__ == "__main__":
    main()
