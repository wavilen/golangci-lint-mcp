#!/usr/bin/env node
'use strict';
// PostToolUse hook for Cursor — injects MCP nudge after golangci-lint calls. Always exits 0.
const path = require('path');
const shared = require(path.join(__dirname, '..', 'shared', 'nudge.js'));
let input = '';
process.stdin.on('data', function (chunk) { input += chunk; });
process.stdin.on('end', function () {
  try {
    const data = JSON.parse(input);
    if (data.tool_name !== 'Shell') process.exit(0);
    const command = (data.tool_input && data.tool_input.command) || '';
    if (!shared.isGolangciLintCommand(command)) process.exit(0);
    // Skip nudge when golangci-lint-mcp intercept was used (per D-11)
    if (shared.hasGolangciLintMcp()) process.exit(0);
    // Fallback: v1.3 behavior — inject nudge into raw output
    let output = data.tool_output || '';
    if (typeof output === 'string' && output.length > 0) {
      try {
        const parsed = JSON.parse(output);
        // Cursor may wrap as { exitCode, stdout } — extract stdout
        if (parsed && typeof parsed === 'object' && typeof parsed.stdout === 'string') {
          output = parsed.stdout;
        }
        // If not a wrapper, output stays as-is (may be raw golangci-lint JSON)
      } catch (_e) {
        // Not valid JSON — use as raw output
      }
    }
    if (!output || !output.trim()) process.exit(0);
    const result = shared.parseDiagnostics(output.trim());
    if (result.totalUnique === 0 && (!result.warnings || result.warnings.length === 0)) process.exit(0);
    let nudge = result.totalUnique <= 10
      ? shared.buildStrategyANudge(result.totalUnique, output.trim())
      : shared.buildStrategyBNudge(result.totalUnique, result.linterCounts, output.trim());
    nudge = shared.truncateNudge(nudge);
    if (result.warnings && result.warnings.length > 0) {
      const warnText = result.warnings.map(function(w) { return w.Text || w.text || String(w); }).join('; ');
      nudge += '\n\n⚠ Config warnings: ' + warnText;
    }
    // Cursor output format: { additional_context: nudge }
    console.log(JSON.stringify({ additional_context: nudge }));
    process.exit(0);
  } catch (_err) { process.exit(0); }
});
process.stdin.on('error', function () { process.exit(0); });
