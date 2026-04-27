#!/usr/bin/env node
'use strict';
// PostToolUse hook for Claude Code — injects MCP nudge after golangci-lint calls. Always exits 0.
const path = require('path');
const shared = require(path.join(__dirname, '..', 'shared', 'nudge.js'));
let input = '';
process.stdin.on('data', function (chunk) { input += chunk; });
process.stdin.on('end', function () {
  try {
    const data = JSON.parse(input);
    if (data.tool_name !== 'Bash') process.exit(0);
    const command = (data.tool_input && data.tool_input.command) || '';
    if (!shared.isGolangciLintCommand(command)) process.exit(0);
    let output = '';
    if (data.tool_response) {
      output = data.tool_response.output || data.tool_response.stdout || data.tool_response.result || '';
      if (typeof output !== 'string') output = output.stdout || output.output || output.result || '';
    }
    if (!output || !output.trim()) process.exit(0);
    const result = shared.parseDiagnostics(output);
    if (result.totalUnique === 0 && (!result.warnings || result.warnings.length === 0)) process.exit(0);
    let nudge = result.totalUnique <= 10
      ? shared.buildStrategyANudge(result.totalUnique, output.trim())
      : shared.buildStrategyBNudge(result.totalUnique, result.linterCounts, output.trim());
    nudge = shared.truncateNudge(nudge);
    if (result.warnings && result.warnings.length > 0) {
      const warnText = result.warnings.map(function(w) { return w.Text || w.text || String(w); }).join('; ');
      nudge += '\n\n⚠ Config warnings: ' + warnText;
    }
    console.log(JSON.stringify({ hookSpecificOutput: { hookEventName: 'PostToolUse', additionalContext: nudge } }));
    process.exit(0);
  } catch (_err) { process.exit(0); }
});
process.stdin.on('error', function () { process.exit(0); });
