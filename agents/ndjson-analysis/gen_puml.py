#!/usr/bin/env python3
"""Parse opencode ndjson traces and generate PlantUML sequence diagrams.

Reads:
  tmp/ndjson/*.ndjson       — opencode session traces
  tmp/ndjson/*-result.json  — test result summaries

Writes:
  tmp/ndjson_analysis/<trace>.puml   — per-trace sequence diagram
  tmp/ndjson_analysis/comparison-overview.puml — all-traces summary

Located at: agents/ndjson-analysis/gen_puml.py (git-tracked)
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
    event_type: str
    tool_name: str | None = None
    tool_input: dict = field(default_factory=dict)
    tool_output: str = ""
    text_content: str = ""
    tokens: dict = field(default_factory=dict)
    step_reason: str = ""
    session_id: str = ""


def parse_ndjson(filepath: str) -> list[Event]:
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
                if not isinstance(out, str):
                    out = str(out)
                limit = 25000 if tool == "task" else 300
                out = out[:limit]
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
                events.append(
                    Event(
                        timestamp=ts,
                        event_type="text",
                        text_content=text[:500],
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


def escape_puml(s: str) -> str:
    """Escape a string for safe use inside PlantUML notes and labels."""
    if not s:
        return ""
    # Order matters: do < > before other replacements
    s = s.replace("&", "&amp;")
    s = s.replace("<", "&lt;")
    s = s.replace(">", "&gt;")
    s = s.replace("|", "\\|")
    s = s.replace("#", "&#35;")
    s = s.replace('"', "'")
    s = s.replace("\r", "")
    s = s.replace("\n", "\\n")
    return s


def escape_puml_inline(s: str) -> str:
    """Escape for single-line inline use (note right of X : ...)."""
    s = escape_puml(s)
    # Also strip any remaining newlines for inline
    s = s.replace("\\n", " ")
    return s[:120]


def extract_summary_info(output: str) -> str:
    """Extract key info from MCP tool output."""
    lines = output.split("\n")
    info_parts = []
    for line in lines:
        line = line.strip()
        if "Unique diagnostics:" in line or "Strategy:" in line or "TOTAL:" in line:
            info_parts.append(line.strip())
        if len(info_parts) >= 3:
            break
    if not info_parts and "No issues found" in output:
        return "No issues found"
    return "\n".join(info_parts) if info_parts else ""


def parse_task_result(output: str) -> dict:
    """Parse task tool output to extract task_id, summary, file, and result text."""
    result = {
        "task_id": "",
        "summary": "",
        "file": "",
        "result_text": "",
    }
    if not output:
        return result

    # Extract task_id
    tid_match = re.search(r"task_id:\s*(ses_\S+)", output)
    if tid_match:
        result["task_id"] = tid_match.group(1)

    # Extract <task_result> block content
    tr_match = re.search(r"<task_result>\s*(.*?)\s*</task_result>", output, re.DOTALL)
    if tr_match:
        block = tr_match.group(1).strip()
        result["result_text"] = block

        # Extract summary: first non-empty line that's not a table separator
        for line in block.split("\n"):
            line = line.strip()
            if not line:
                continue
            if line.startswith("|---") or line.startswith("| ---"):
                continue
            # Remove leading markdown bold markers
            clean = re.sub(r"^\*+\s*", "", line)
            if clean:
                result["summary"] = clean[:200]
                break

    # Extract file name from description or task_result content
    # Look for backtick-wrapped filenames first
    file_match = re.search(r"`(\w+\.\w+)`", output)
    if file_match:
        result["file"] = file_match.group(1)

    return result


def tool_label(event: Event) -> str:
    tool = event.tool_name or ""
    inp = event.tool_input

    if "golangci-lint" in tool:
        # Map specific MCP tool names
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
            return f"bash: {desc[:50]}"
        return f"bash: {cmd[:50]}"
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
        return f"task → {desc[:50]}"
    if tool == "question":
        return "question()"
    return tool


def tool_target(event: Event) -> str:
    """Determine which participant a tool call targets."""
    tool = event.tool_name or ""
    if tool in ("edit", "read", "write", "bash", "glob", "todowrite"):
        return "FS"
    if tool == "task":
        return "SUB"
    return "MCP"


def tool_note(event: Event) -> str:
    """Extract a note for the tool call from output."""
    out = event.tool_output
    if not out:
        return ""
    tool = event.tool_name or ""

    if "golangci-lint" in tool:
        return extract_summary_info(out)

    if tool == "skill":
        return "Skill loaded"

    if tool == "edit":
        if "Edit applied" in out:
            return "Edit applied"
        if "oldString not found" in out:
            return "FAILED: oldString not found"
        return out[:60]

    if tool == "bash":
        if "(no output)" in out:
            return "(no output)"
        lines = out.strip().split("\n")
        first = lines[0][:80] if lines else ""
        return first

    if tool == "glob":
        cnt = 0
        for line in out.split("\n"):
            line = line.strip()
            if line and not line.startswith('"') and not line.startswith("{"):
                cnt += 1
        if "No files found" in out:
            return "No files found"
        return f"{cnt} files" if cnt else out[:60]

    return ""


def generate_puml(
    name: str,
    events: list[Event],
    result: dict,
    output_path: str,
):
    if not events:
        return

    base_ts = events[0].timestamp
    total_ms = events[-1].timestamp - base_ts

    test_case = result.get("TestCase", "?")
    model = result.get("Model", "?")
    before = result.get("BeforeIssues", 0)
    after = result.get("AfterIssues", 0)
    tool_calls = result.get("ToolCalls", 0)
    passed = result.get("Pass", False)
    reduction = result.get("IssueReduction", 0)

    status = "PASS" if passed else "FAIL"
    status_color = "green" if passed else "red"

    # Scan for task tool calls to build subagent participants
    task_events = [e for e in events if e.event_type == "tool_use" and e.tool_name == "task"]
    has_subagents = len(task_events) > 0

    # Build subagent participant mapping: task_index -> (alias, short_desc)
    sub_participants = []
    for idx, te in enumerate(task_events):
        desc = te.tool_input.get("description", "")
        # Extract short name: use the file name or last meaningful word
        file_match = re.search(r"(\w+\.\w+)", desc)
        if file_match:
            short = file_match.group(1)
        else:
            # Take last meaningful part of description
            parts = desc.split()
            short = parts[-1] if parts else f"sub{idx + 1}"
        # Clean short for use as alias component
        short = re.sub(r"[^a-zA-Z0-9_]", "_", short)
        alias = f"SUB_{idx + 1}"
        sub_participants.append((alias, short, desc))

    lines = []
    lines.append("@startuml")
    lines.append(f"title {escape_puml(name)}")
    lines.append("")
    lines.append("skinparam backgroundColor #FEFEFE")
    lines.append("skinparam sequenceMessageAlign center")
    lines.append("skinparam maxMessageSize 150")
    lines.append("")
    lines.append('participant "Agent" as A #LightBlue')
    lines.append('participant "MCP Tools" as MCP #LightYellow')
    lines.append('participant "FileSystem" as FS #LightGreen')

    # Declare subagent participants if any
    if has_subagents:
        lines.append("")
        lines.append('box "Subagents" #LightCoral')
        for alias, short, _desc in sub_participants:
            lines.append(f'  participant "Sub: {escape_puml(short)}" as {alias}')
        lines.append("end box")

    lines.append("")

    # Header box with summary
    lines.append("note over A, FS")
    lines.append(f"  **{escape_puml(name)}**")
    lines.append(f"  Model: {escape_puml(model)}")
    lines.append(f"  Test: {test_case} / Issues: {before} to {after} ({reduction}% reduction)")
    lines.append(f"  Tool Calls: {tool_calls} / Duration: {fmt_duration(total_ms)}")
    if has_subagents:
        lines.append(f"  Subagents: {len(task_events)}")
    lines.append(f"  Result: <color:{status_color}><b>{status}</b></color>")
    lines.append("end note")
    lines.append("")

    # Group events into steps (step_start -> step_finish)
    current_step_events: list[Event] = []
    step_idx = 0

    for ev in events:
        if ev.event_type == "step_start":
            current_step_events = []
            step_idx += 1
        elif ev.event_type in ("tool_use", "text"):
            current_step_events.append(ev)
        if ev.event_type == "step_finish":
            # Process accumulated events in this step
            step_tools = [e for e in current_step_events if e.event_type == "tool_use"]
            step_texts = [e for e in current_step_events if e.event_type == "text"]

            if not step_tools and not step_texts:
                continue

            # Process tool calls
            for te in step_tools:
                target = tool_target(te)
                tl = tool_label(te)

                # Handle task tool calls with subagent participants
                if te.tool_name == "task" and has_subagents:
                    # Find which subagent index this corresponds to
                    # Count task events up to this point
                    task_idx_in_list = None
                    count = 0
                    for e in events:
                        if e is te:
                            task_idx_in_list = count
                            break
                        if e.event_type == "tool_use" and e.tool_name == "task":
                            count += 1

                    if task_idx_in_list is not None and task_idx_in_list < len(sub_participants):
                        alias, short, desc = sub_participants[task_idx_in_list]
                        label = escape_puml_inline(desc[:80])
                        lines.append(f"A -> {alias} : **{label}**")

                        # Parse task result and add note
                        tr = parse_task_result(te.tool_output)
                        if tr.get("summary"):
                            summary_escaped = escape_puml_inline(tr["summary"][:200])
                            lines.append(f"note right of {alias}")
                            lines.append(f"  {summary_escaped}")
                            lines.append("end note")
                        lines.append(f"{alias} --> A : Done")
                    else:
                        # Fallback: shouldn't happen, but render as generic
                        label = escape_puml_inline(tl)
                        lines.append(f"A -> MCP : **{label}**")
                    lines.append("")
                    continue

                label = escape_puml_inline(tl)

                lines.append(f"A -> {target} : **{label}**")

                note = tool_note(te)
                if note:
                    # Multi-line note
                    note_lines = note.split("\n")
                    if len(note_lines) == 1:
                        lines.append(f"note right of {target} : {escape_puml_inline(note)}")
                    else:
                        lines.append(f"note right of {target}")
                        for nl in note_lines:
                            escaped = escape_puml_inline(nl)
                            if escaped.strip():
                                lines.append(f"  {escaped}")
                        lines.append("end note")

            # Process text events
            for te in step_texts:
                txt = te.text_content.strip()
                if not txt:
                    continue
                first_line = txt.split("\n")[0][:150]
                first_line_esc = escape_puml_inline(first_line)
                has_more = len(txt.split("\n")) > 1
                if has_more:
                    lines.append("note over A")
                    lines.append(f"  {first_line_esc}")
                    lines.append("  ...")
                    lines.append("end note")
                else:
                    lines.append(f"note over A : {first_line_esc}")
                lines.append("")

    lines.append("")
    lines.append("note over A, FS")
    lines.append(f"  **Final:** {before} to {after} issues / <color:{status_color}><b>{status}</b></color>")
    lines.append(f"  Duration: {fmt_duration(total_ms)}")
    lines.append("end note")
    lines.append("")
    lines.append("@enduml")

    with open(output_path, "w") as f:
        f.write("\n".join(lines))

    print(f"  Written: {output_path}")


def generate_comparison(
    results: list[tuple[str, dict, int]],
    output_path: str,
):
    """Generate comparison overview diagram."""
    lines = []
    lines.append("@startuml")
    lines.append("title Agent Session Comparison Overview")
    lines.append("")
    lines.append("skinparam backgroundColor #FEFEFE")
    lines.append("")
    lines.append('rectangle "Comparison" {')
    lines.append("  note as N")

    for name, res, dur in results:
        model = res.get("Model", "?")
        tc = res.get("TestCase", "?")
        before = res.get("BeforeIssues", 0)
        after = res.get("AfterIssues", 0)
        calls = res.get("ToolCalls", 0)
        passed = res.get("Pass", False)
        reduction = res.get("IssueReduction", 0)
        subagent_count = res.get("_subagent_count", 0)
        status = "PASS" if passed else "FAIL"
        status_color = "green" if passed else "red"
        line = (
            f"  **{escape_puml(name)}**: {escape_puml(model)} "
            f"- {tc} - {before} to {after} ({reduction}%) - "
            f"{calls} calls - {fmt_duration(dur)} - "
        )
        if subagent_count > 0:
            line += f"{subagent_count} subagents - "
        line += f"<color:{status_color}><b>{status}</b></color>"
        lines.append(line)
    lines.append("  end note")
    lines.append("}")
    lines.append("")
    lines.append("@enduml")

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

    comparison_data = []

    for fname in files:
        ndjson_path = os.path.join(ndjson_dir, fname)
        stem = fname.replace(".ndjson", "")

        result_path = os.path.join(ndjson_dir, stem + "-result.json")
        result = {}
        if os.path.exists(result_path):
            with open(result_path) as f:
                result = json.load(f)

        print(f"Processing: {fname}")
        events = parse_ndjson(ndjson_path)
        if not events:
            print("  No events found, skipping")
            continue

        total_ms = events[-1].timestamp - events[0].timestamp

        # Count subagent task calls for comparison
        task_count = sum(1 for e in events if e.event_type == "tool_use" and e.tool_name == "task")
        if task_count > 0:
            result["_subagent_count"] = task_count

        comparison_data.append((stem, result, total_ms))

        puml_path = os.path.join(out_dir, stem + ".puml")
        generate_puml(stem, events, result, puml_path)

    # Process probe trace
    probe_path = os.path.join(ndjson_dir, "probe.ndjson")
    if os.path.exists(probe_path):
        print("Processing: probe.ndjson")
        events = parse_ndjson(probe_path)
        if events:
            puml_path = os.path.join(out_dir, "probe.puml")
            generate_puml("probe", events, {}, puml_path)

    if comparison_data:
        comp_path = os.path.join(out_dir, "comparison-overview.puml")
        generate_comparison(comparison_data, comp_path)

    print(f"\nDone. {len(comparison_data)} traces processed.")


if __name__ == "__main__":
    main()
