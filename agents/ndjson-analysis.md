---
description: Analyze opencode ndjson session traces into PlantUML sequence diagrams
mode: subagent
temperature: 0.1
permission:
  bash: allow
  edit: allow
---

<objective>
Analyze opencode agent session NDJSON traces and produce PlantUML sequence diagrams for every trace, plus a comparison overview. The agent reads raw `.ndjson` + `-result.json` files from `tmp/ndjson/`, generates `.puml` files into `tmp/ndjson_analysis/`, and verifies each renders cleanly with Docker PlantUML.
</objective>

<execution_context>
- **Project:** golangci-lint-mcp — a Go MCP server for golangci-lint guidance
- **Input:** `tmp/ndjson/*.ndjson` (opencode session traces) + `tmp/ndjson/*-result.json` (test result summaries)
- **Output:** `tmp/ndjson_analysis/*.puml`
- **Generator script:** `tmp/ndjson_analysis/gen_puml.py` — Python parser that reads ndjson traces and outputs PlantUML
- **Verifier:** Docker image `plantuml/plantuml:latest` — used to syntax-check each `.puml`
- **Tools available:** bash (full access), edit (full access)
- **NDJSON format:** Each line is a JSON object. Key event types:
  - `step_start` / `step_finish` — agent reasoning step boundaries, `part.tokens` has token counts
  - `tool_use` — tool invocation, `part.tool` = tool name, `part.state.input` = args, `part.state.output` = result (truncated to 300 chars)
  - `text` — agent text output, `part.text` = content
- **Result JSON format:** `{TestCase, Model, BeforeIssues, AfterIssues, BuildsClean, LintClean, NolintCount, IssueReduction, ToolCalls, Retries, Pass, Error, ConfigModified}`
</execution_context>

<process>

## Step 1: Verify Input Data

**1a. Check ndjson data exists**

```bash
ls tmp/ndjson/*.ndjson 2>/dev/null
```

If no `.ndjson` files found, ABORT with:
> "No ndjson traces found in tmp/ndjson/. Run e2e tests first to generate traces."

**1b. Check Docker + PlantUML available**

```bash
docker run --rm plantuml/plantuml:latest -version
```

If Docker or PlantUML image is unavailable, ABORT with:
> "Docker + plantuml/plantuml:latest image required. Pull with: docker pull plantuml/plantuml:latest"

**1c. List discovered traces**

Report the list of `.ndjson` files and matching `-result.json` files found.

---

## Step 2: Update Generator Script

Read `tmp/ndjson_analysis/gen_puml.py`. If it doesn't exist, create it. The script must:

1. **Parse ndjson** — extract `step_start`, `step_finish`, `tool_use`, `text` events with timestamps
2. **Generate per-trace PlantUML** — sequence diagram with 3 participants (Agent, MCP Tools, FileSystem):
   - Header note with session metadata (model, test case, issue counts, duration, pass/fail)
   - Tool calls as arrows from Agent to target participant, labeled with tool name + args
   - Agent text as `note over A` blocks (first line only, truncated at 150 chars)
   - Tool output notes (for MCP calls: extract strategy + issue summary; for edits: "Edit applied"; for bash: first line of output)
   - Footer note with final issue count and status
3. **Generate comparison overview** — one diagram summarizing all sessions
4. **Escape PlantUML special characters** — `|` → `\|`, `<` → `&lt;`, `>` → `&gt;`, `"` → `'`
5. **Handle missing result files** — use empty dict defaults if `-result.json` doesn't exist
6. **Output paths** — write `.puml` to `tmp/ndjson_analysis/`

**Key escape rules for PlantUML notes:**
- No bare `|` characters in notes (breaks PlantUML table parsing)
- No multi-line markdown tables in single-line notes
- Agent text with multiple lines: show first line + "..." in a multi-line note block
- Truncate all text fields to reasonable lengths (150 chars for agent text, 80 for bash output, 100 for inline notes)

Run the script:

```bash
python3 tmp/ndjson_analysis/gen_puml.py
```

Report the number of traces processed.

---

## Step 3: Verify PlantUML Renders

For each `.puml` file in `tmp/ndjson_analysis/`:

```bash
docker run --rm -v "$(pwd)/tmp/ndjson_analysis:/data" plantuml/plantuml:latest -tpng "/data/{filename}.puml"
```

**This generates PNGs as a render verification step.** If any file has syntax errors, fix the escape logic in `gen_puml.py` and regenerate:

Common causes: unescaped `|`, `#`, or newlines in note text; markdown tables; bare HTML tags.

After fixing, re-run Step 2 then re-verify only the failing files.

---

## Step 4: Verify Output

```bash
ls -lh tmp/ndjson_analysis/*.puml
```

Count must equal: number of `.ndjson` traces + 1 (comparison overview). Report the final file list.

---

## Step 5: Clean Old Files

Remove stale `.puml` or `.png` files not backed by a current `.ndjson` trace.

```bash
find tmp/ndjson_analysis/ -maxdepth 1 -name '*.png' -delete
```

Keep only `gen_puml.py` and `*.puml`.

---

## Success Output

```
NDJSON analysis complete.

Traces processed: {N}
  {trace-1}: {model} / {test} / {before}->{after} issues / {calls} calls / {duration} / {status}
  {trace-2}: ...

Output:
  tmp/ndjson_analysis/{trace-1}.puml
  tmp/ndjson_analysis/{trace-2}.puml
  ...
  tmp/ndjson_analysis/comparison-overview.puml

All {N}+1 diagrams verified with PlantUML.

Generator: tmp/ndjson_analysis/gen_puml.py
```

</process>
