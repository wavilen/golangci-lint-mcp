---
description: Analyze opencode ndjson session traces into PlantUML sequence diagrams with subagent-enriched analysis
mode: subagent
temperature: 0.1
permission:
  bash: allow
  edit: allow
---

<objective>
Analyze opencode agent session NDJSON traces using a three-phase pipeline: (1) gen_puml.py generates mechanical PlantUML sequence diagrams, (2) extract_context.py produces per-trace structured context JSON, (3) subagent pools of 3 validate/fix PUML files and produce per-trace markdown analysis. Main agent assembles cross-trace model comparison. Removes old diff analysis (gen_agent_dump.py workflow, strategy-comparison.txt) in favor of subagent-produced markdown analysis.
</objective>

<execution_context>
- **Project:** golangci-lint-mcp — a Go MCP server for golangci-lint guidance
- **Input:** `tmp/ndjson/*.ndjson` (opencode session traces) + `tmp/ndjson/*-result.json` (test result summaries)
- **Output:** `tmp/ndjson_analysis/*.puml`, `tmp/ndjson_analysis/*-context.json`, `tmp/ndjson_analysis/*-analysis.md`, `tmp/ndjson_analysis/model-comparison.md`
- **Generator script:** `agents/ndjson-analysis/gen_puml.py` — Python parser that reads ndjson traces and outputs PlantUML (FROZEN — never modify per D-13)
- **Context extractor:** `agents/ndjson-analysis/extract_context.py` — Python script that reads ndjson traces + result JSON and outputs per-trace structured context JSON
- **Ad-hoc tool (kept but NOT used in workflow):** `agents/ndjson-analysis/gen_agent_dump.py` — available for manual debugging but no longer part of the agent pipeline
- **Script location note:** Scripts are persisted in `agents/ndjson-analysis/` (git-tracked). They read from `tmp/ndjson/` and write to `tmp/ndjson_analysis/` (both gitignored runtime directories).
- **Verifier:** Docker image `plantuml/plantuml:latest` — used to validate each `.puml` renders
- **Tools available:** bash (full access), edit (full access), task (subagent spawning)
- **NDJSON format:** Each line is a JSON object. Key event types:
  - `step_start` / `step_finish` — agent reasoning step boundaries, `part.tokens` has token counts
  - `tool_use` — tool invocation, `part.tool` = tool name, `part.state.input` = args, `part.state.output` = result
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

**1d. Clean output directory**

Remove all previous analysis output to prevent stale data mixing with new results.

```bash
rm -rf tmp/ndjson_analysis/
mkdir -p tmp/ndjson_analysis/
```

Report: "Cleared tmp/ndjson_analysis/ for fresh analysis."

---

## Step 2a: Generate PlantUML Diagrams

Run the PUML generator (unchanged, frozen script per D-13):

```bash
python3 agents/ndjson-analysis/gen_puml.py
```

Report the number of traces processed.

**CRITICAL: Do NOT modify gen_puml.py.** It is a frozen mechanical generator. If PUML files have rendering issues, those are fixed by subagents in Step 3 (fixing the .puml output files, NOT the script).

---

## Step 2b: Extract Per-Trace Context

Run the context extractor (new script per D-05/D-06):

```bash
python3 agents/ndjson-analysis/extract_context.py
```

This produces `tmp/ndjson_analysis/{base}-context.json` files with structured fields per trace:
- `model`, `test_case` — from result JSON
- `duration_ms`, `timestamps` — from event timestamps
- `event_counts` — count of each event type
- `tool_breakdown` — tool name → count mapping
- `token_totals` — input/output/reasoning/cache_read sums
- `issue_counts` — before/after/reduction_pct from result JSON
- `subagent_count` — count of task tool calls
- `strategy` — extracted from strategy_instructions or Strategy: lines
- `pass` — test pass/fail status

Report the fields extracted: model, test case, issue counts, tool breakdown, strategy, subagent count.

---

## Step 3: Dispatch Subagent Pools

Per D-14: Fixed pool of 3 subagents per batch.
Per D-15: Fail immediately and stop enrichment — no retries, report partial results.
Per D-16: Main agent waits for each batch to complete before dispatching the next.

**3a. List context files and group into batches**

```bash
ls tmp/ndjson_analysis/*-context.json
```

Group the context files into batches of 3. Process each batch sequentially.

**3b. For each batch, spawn subagents using the `task` tool**

Each subagent receives this structured prompt (D-01, D-07):

```
You are analyzing an opencode agent session trace. Your inputs are pre-extracted context files — do NOT read raw ndjson.

READ these files:
1. tmp/ndjson_analysis/{base}-context.json — structured context (model, tools, tokens, issues, strategy)
2. tmp/ndjson/{base}-result.json — test result summary (TestCase, Model, BeforeIssues, AfterIssues, Pass, etc.)
3. tmp/ndjson_analysis/{base}.puml — PlantUML sequence diagram for this trace

YOUR TASKS (complete all three):
A) Validate the PUML file renders: `docker run --rm -v "$(pwd)/tmp/ndjson_analysis:/data" plantuml/plantuml:latest -tpng "/data/{base}.puml"`
B) If rendering fails, fix the PUML file directly (edit the .puml file, NOT gen_puml.py). Common fixes: escape | → \|, # → &#35;, bare HTML tags, multi-line notes in inline context. Re-validate after fixing.
C) Write a per-trace analysis markdown file to tmp/ndjson_analysis/{base}-analysis.md with this structure:

# Analysis: {base}
## Session Info
- Model: {from context.json}
- Test Case: {from context.json}
- Duration: {from context.json}
- Pass/Fail: {from context.json}
## Strategy
- Strategy used: {from context.json}
- Subagent count: {from context.json}
- Strategy compliance: {assessed from context — did the agent follow the strategy?}
## Tool Usage
- Total tool calls: {from context.json event_counts}
- Tool breakdown: {from context.json tool_breakdown}
- Tool usage efficiency: {your assessment}
## Issue Progression
- Before: {issue_counts.before} → After: {issue_counts.after} ({issue_counts.reduction_pct}% reduction)
- Assessment: {was the agent effective at reducing issues?}
## Token Consumption
- Input: {token_totals.input} / Output: {token_totals.output} / Reasoning: {token_totals.reasoning}
- Cache reads: {token_totals.cache_read}
- Efficiency assessment: {tokens per issue fixed}
## Patterns Observed
- {list 2-5 notable patterns: error recovery, fix approach, tool selection, strategy adherence}
## PUML Validation
- Render status: {pass/fail, fixes applied if any}
```

Per D-02: Each subagent validates PUML, fixes issues, returns summary.
Per D-03: Subagents fix PUML output files, NOT gen_puml.py.
Per D-04: Summary returned to parent inline.
Per D-08: Per-trace markdown analysis written to `{base}-analysis.md`.
Per D-09: Subagent also returns short summary inline.

**3c. Failure handling** (D-15):

If any subagent in a batch fails:
1. Log the failure with details
2. Report partial results collected so far
3. STOP enrichment — do NOT retry failed subagents
4. Continue with whatever analysis files were successfully produced

---

## Step 4: Assemble Model Comparison

Per D-10: Cross-trace diff comparison as markdown — NOT PUML diagrams.
Per D-11: Main agent reads all `{base}-analysis.md` files and produces `model-comparison.md`.
Per D-12: Comparison dimensions: tool usage efficiency, issue reduction rate, token consumption, strategy compliance, error recovery patterns.

**4a. Read all analysis files**

```bash
ls tmp/ndjson_analysis/*-analysis.md
```

Read each analysis file to gather cross-trace data.

**4b. Produce `tmp/ndjson_analysis/model-comparison.md`**

Structure:
- **Summary table:** Model | Test | Issues (Before→After) | Tool Calls | Tokens | Strategy | Pass
- **Per-dimension comparison sections:**
  - Tool usage efficiency — which models used tools most effectively
  - Issue reduction rate — which models reduced issues fastest/with fewest iterations
  - Token consumption — which models were most token-efficient per issue fixed
  - Strategy compliance — which models followed their assigned strategy
  - Error recovery patterns — how models handled failures and edge cases
- **Key findings:** which models performed best on which dimensions
- **Notable patterns:** cross-trace observations about agent behavior

Per D-13: gen_puml.py's `comparison-overview.puml` is still generated (unchanged) — it provides a visual summary alongside the markdown comparison.

---

## Success Output

```
NDJSON analysis complete.

Traces processed: {N}
  {trace-1}: {model} / {test} / {before}→{after} issues / {calls} calls / {strategy} / {status}
  {trace-2}: ...

Output:
  tmp/ndjson_analysis/{trace-1}.puml
  tmp/ndjson_analysis/{trace-1}-context.json
  tmp/ndjson_analysis/{trace-1}-analysis.md
  ...
  tmp/ndjson_analysis/comparison-overview.puml
  tmp/ndjson_analysis/model-comparison.md

PUML diagrams: {N}+1 verified with PlantUML
Context files: {N} extracted
Analysis files: {N} subagent-produced
Model comparison: model-comparison.md

Generators: agents/ndjson-analysis/gen_puml.py, agents/ndjson-analysis/extract_context.py
```

</process>
