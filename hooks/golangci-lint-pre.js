#!/usr/bin/env node
'use strict';
// PreToolUse hook for Claude Code — modifies golangci-lint Bash commands before execution to ensure JSON output format. Always exits 0.
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
    // Route through golangci-lint-mcp intercept when binary is available (per D-07)
    if (shared.hasGolangciLintMcp()) {
      const interceptCmd = shared.buildInterceptCommand(command);
      console.log(JSON.stringify({ hookSpecificOutput: { permissionDecision: 'allow', updatedInput: { command: interceptCmd } } }));
      process.exit(0);
    }
    // Fallback: v1.3 behavior — inject JSON output flag for nudge pipeline
    const cdMatch = command.match(/^(cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*)/);
    const cdPrefix = cdMatch ? cdMatch[0] : '';
    const inner = shared.extractInnerCommand(command);
    const parts = shared.splitCompoundCommand(inner);
    let lintCmd = shared.stripOutputFilters(parts.lintCommand);
    lintCmd = shared.injectJsonOutputFlag(lintCmd);
    const jqPart = shared.injectJqFilter(parts.jqSegment);
    // Collect trailing echo-like segments from semicolon-separated compound commands
    // splitCompoundCommand only handles || echo / && echo within the lint segment.
    // We also need to preserve plain semicolon-separated echo/true commands after lint.
    let trailingSegments = '';
    if (parts.echoSuffix) {
      trailingSegments = ' ' + parts.echoSuffix;
    } else if (inner.indexOf(';') !== -1) {
      // Re-split on unquoted semicolons to find trailing echo/true segments
      const segs = inner.split(';').map(function(s) { return s.trim(); }).filter(Boolean);
      // Find the lint segment and collect everything after it
      let foundLint = false;
      const trailing = [];
      for (let si = 0; si < segs.length; si++) {
        if (!foundLint && /\bgolangci-lint\b/.test(segs[si])) {
          foundLint = true;
          continue;
        }
        if (foundLint) {
          trailing.push(segs[si]);
        }
      }
      if (trailing.length > 0) {
        trailingSegments = '; ' + trailing.join('; ');
      }
    }
    const modified = cdPrefix + lintCmd + jqPart + trailingSegments;
    console.log(JSON.stringify({ hookSpecificOutput: { permissionDecision: 'allow', updatedInput: { command: modified } } }));
    process.exit(0);
  } catch (_err) { process.exit(0); }
});
process.stdin.on('error', function () { process.exit(0); });
