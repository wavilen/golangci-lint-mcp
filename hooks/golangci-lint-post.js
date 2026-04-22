#!/usr/bin/env node
'use strict';
// PostToolUse hook for Claude Code — injects MCP nudge after golangci-lint calls. Always exits 0.
var path = require('path');
var shared = require(path.join(__dirname, '..', 'shared', 'nudge.js'));
var input = '';
process.stdin.on('data', function (chunk) { input += chunk; });
process.stdin.on('end', function () {
  try {
    var data = JSON.parse(input);
    if (data.tool_name !== 'Bash') process.exit(0);
    var command = (data.tool_input && data.tool_input.command) || '';
    if (!shared.isGolangciLintCommand(command)) process.exit(0);
    var output = data.tool_response ? (data.tool_response.stdout || data.tool_response.output || '') : '';
    if (!output || !output.trim()) process.exit(0);
    var result = shared.parseDiagnostics(output);
    if (result.totalUnique === 0) process.exit(0);
    var nudge = result.totalUnique <= 10
      ? shared.buildStrategyANudge(result.totalUnique, output.trim())
      : shared.buildStrategyBNudge(result.totalUnique, result.linterCounts, output.trim());
    nudge = shared.truncateNudge(nudge);
    console.log(JSON.stringify({ hookSpecificOutput: { hookEventName: 'PostToolUse', additionalContext: nudge } }));
    process.exit(0);
  } catch (err) { process.exit(0); }
});
process.stdin.on('error', function () { process.exit(0); });
