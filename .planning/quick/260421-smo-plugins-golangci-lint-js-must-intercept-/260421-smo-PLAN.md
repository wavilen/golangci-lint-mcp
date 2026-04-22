---
phase: quick
plan: 01
type: execute
wave: 1
depends_on: []
files_modified:
  - plugins/golangci-lint.js
  - plugins/golangci-lint.test.js
autonomous: true
requirements: [QUICK-smo]
must_haves:
  truths:
    - "golangci-lint run is intercepted and gets JSON output injected"
    - "golangci-lint (bare, no subcommand) is intercepted — defaults to run"
    - "golangci-lint with flags but no subcommand is intercepted — defaults to run"
    - "golangci-lint cache/version/linters/help/config/completion/custom are NOT intercepted"
  artifacts:
    - path: "plugins/golangci-lint.js"
      provides: "isGolangciLintCommand returns true only for run commands"
    - path: "plugins/golangci-lint.test.js"
      provides: "Tests covering run vs non-run subcommand discrimination"
  key_links:
    - from: "plugins/golangci-lint.js"
      to: "isGolangciLintCommand"
      via: "subcommand check after first token"
      pattern: "NON_RUN_SUBCOMMANDS"
---

<objective>
Make `isGolangciLintCommand` return `true` only for `golangci-lint run` (or bare `golangci-lint` which defaults to run), not for other subcommands like `cache`, `version`, `linters`, etc.

Purpose: The plugin currently intercepts ALL golangci-lint commands (cache clean, version, linters list, etc.) and forces JSON output. This breaks non-run subcommands — `golangci-lint version --output.json.path stdout` is nonsensical, and `golangci-lint cache clean` gets corrupted with output flags stripped.

Output: Modified `isGolangciLintCommand` that discriminates between `run` and non-run subcommands, plus updated tests.
</objective>

<execution_context>
@$HOME/.config/opencode/get-shit-done/workflows/execute-plan.md
@$HOME/.config/opencode/get-shit-done/templates/summary.md
</execution_context>

<context>
@plugins/golangci-lint.js
@plugins/golangci-lint.test.js
</context>

<tasks>

<task type="auto" tdd="true">
  <name>task 1: restrict isGolangciLintCommand to run subcommand only</name>
  <files>plugins/golangci-lint.test.js, plugins/golangci-lint.js</files>
  <behavior>
    Positive cases (must return true):
    - `golangci-lint run` → true (explicit run)
    - `golangci-lint` → true (bare command, defaults to run)
    - `golangci-lint run ./...` → true (run with args)
    - `golangci-lint --timeout 5m run` → true (flags before run)
    - `golangci-lint -E errcheck` → true (flags, no subcommand = defaults to run)
    - `/usr/bin/golangci-lint run` → true (path + run)
    - `FOO=bar golangci-lint run` → true (env vars + run)
    - `FOO=bar golangci-lint` → true (env vars, bare = run)

    Negative cases (must return false):
    - `golangci-lint cache` → false
    - `golangci-lint cache clean` → false
    - `golangci-lint version` → false
    - `golangci-lint linters` → false
    - `golangci-lint help` → false
    - `golangci-lint config` → false
    - `golangci-lint config path` → false
    - `golangci-lint completion bash` → false
    - `golangci-lint custom` → false
    - All existing negative cases remain false (echo, grep, etc.)
  </behavior>
  <action>
    **RED phase:** Add new test cases to `plugins/golangci-lint.test.js`:

    1. Add a new describe block `'isGolangciLintCommand — only run subcommand'` with:
       - Positive tests: `golangci-lint run`, bare `golangci-lint`, `golangci-lint --timeout 5m run`, `golangci-lint -E errcheck`, env-var-prefixed versions
       - Negative tests: `golangci-lint cache`, `golangci-lint cache clean`, `golangci-lint version`, `golangci-lint linters`, `golangci-lint help`, `golangci-lint config`, `golangci-lint completion bash`, `golangci-lint custom`

    2. Run tests — the new negative cases for `cache`, `version`, `linters` etc. will FAIL because current `isGolangciLintCommand` returns `true` for all golangci-lint invocations.

    **GREEN phase:** Modify `isGolangciLintCommand` in `plugins/golangci-lint.js`:

    1. After stripping env vars and identifying first token as `golangci-lint`/path ending in `/golangci-lint`:
    2. Collect remaining tokens after the first token (skip the command itself).
    3. Find the first non-flag token (not starting with `-`). This is the subcommand.
       - If no non-flag token exists → bare command with only flags → return `true` (defaults to run)
       - If the first non-flag token is `run` → return `true`
       - If the first non-flag token is a known non-run subcommand → return `false`
    4. Define `NON_RUN_SUBCOMMANDS` array: `['cache', 'completion', 'config', 'custom', 'help', 'linters', 'version']`
    5. Keep the function in ES5 style (var, indexOf — no includes/startsWith/set/const) to match existing code style.
    6. Export `NON_RUN_SUBCOMMANDS` for testability.

    Implementation approach for the subcommand check:
    ```
    // After identifying firstToken as golangci-lint...
    // Extract remaining tokens after the command
    var rest = cmd.substring(match[0].length).trim();
    var restTokens = rest.split(/\s+/);
    // Find first non-flag token
    var subcommand = null;
    for (var k = 0; k < restTokens.length; k++) {
      if (restTokens[k] && restTokens[k].charAt(0) !== '-') {
        subcommand = restTokens[k];
        break;
      }
    }
    // No subcommand or subcommand is 'run' → intercept
    if (!subcommand || subcommand === 'run') return true;
    // Known non-run subcommand → don't intercept
    if (NON_RUN_SUBCOMMANDS.indexOf(subcommand) !== -1) return false;
    // Unknown subcommand → assume it might be run-like, intercept
    return true;
    ```

    Note: The "unknown subcommand defaults to true" is intentional — it's safer to over-intercept than to miss a `golangci-lint run` variant.

    Run all tests — must pass (existing + new).
  </action>
  <verify>
    <automated>cd /home/wavilen/Workspace/golangcilint-mcp && node --test plugins/golangci-lint.test.js</automated>
  </verify>
  <done>
    - `isGolangciLintCommand('golangci-lint cache')` returns `false`
    - `isGolangciLintCommand('golangci-lint version')` returns `false`
    - `isGolangciLintCommand('golangci-lint run')` returns `true`
    - `isGolangciLintCommand('golangci-lint')` returns `true` (bare = run)
    - `isGolangciLintCommand('golangci-lint -E errcheck')` returns `true` (flags only = run)
    - All 30+ existing tests still pass
  </done>
</task>

</tasks>

<verification>
node --test plugins/golangci-lint.test.js — all tests pass
</verification>

<success_criteria>
- `isGolangciLintCommand` only returns true for `run` or bare golangci-lint (defaults to run)
- Non-run subcommands (cache, version, linters, help, config, completion, custom) return false
- All existing positive and negative test cases continue to pass
- New test cases cover the subcommand discrimination
- Code style remains ES5 (var, indexOf) matching existing conventions
</success_criteria>

<output>
After completion, create `.planning/quick/260421-smo-plugins-golangci-lint-js-must-intercept-/260421-smo-SUMMARY.md`
</output>
