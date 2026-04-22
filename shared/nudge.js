'use strict';

var OUTPUT_FLAG_PATTERNS_VALUE = [
  '--output.text.path',
  '--output.tab.path',
  '--output.html.path',
  '--output.checkstyle.path',
  '--output.code-climate.path',
  '--output.junit-xml.path',
  '--output.teamcity.path',
  '--output.sarif.path',
  '--output.json.path',
  '--out-format',
  '--color'
];

var OUTPUT_FLAG_PATTERNS_BOOL = [
  '--output.text.print-linter-name',
  '--output.text.print-issued-lines',
  '--output.text.colors',
  '--output.tab.print-linter-name',
  '--output.tab.colors',
  '--output.junit-xml.extended',
  '--show-stats',
  '--verbose',
  '-v',
  '--print-issued-lines',
  '--print-linter-name'
];

var NON_RUN_SUBCOMMANDS = ['cache', 'completion', 'config', 'custom', 'help', 'linters', 'version'];

function extractInnerCommand(command) {
  if (!command || typeof command !== 'string') return command || '';
  var cmd = command.trimStart();
  while (/^[A-Za-z_][A-Za-z0-9_.]*=\S+/.test(cmd)) {
    cmd = cmd.replace(/^[A-Za-z_][A-Za-z0-9_.]*=\S+\s*/, '');
  }
  var shellMatch = cmd.match(/^(?:\/[\w\/]+\/)?(?:bash|sh|zsh|dash|ksh)\s+-c\s+["']?/);
  if (shellMatch) {
    cmd = cmd.substring(shellMatch[0].length);
    if (cmd.endsWith('"') || cmd.endsWith("'")) {
      cmd = cmd.slice(0, -1);
    }
  }
  var cdMatch = cmd.match(/^cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*/);
  if (cdMatch) {
    cmd = cmd.substring(cdMatch[0].length);
  }
  return cmd;
}

function isGolangciLintCommand(command) {
  if (!command || typeof command !== 'string') return false;
  var cmd = extractInnerCommand(command);
  // Extract the first token (the actual command being invoked)
  var match = cmd.match(/^(\S+)/);
  if (!match) return false;
  var firstToken = match[1];
  // Must be "golangci-lint" or a path ending with "/golangci-lint"
  if (firstToken !== 'golangci-lint' && !firstToken.endsWith('/golangci-lint')) return false;

  // Extract remaining tokens after the command itself
  var rest = cmd.substring(match[0].length).trim();
  var restTokens = rest.split(/\s+/);

  // Find the first non-flag token (not starting with '-') — that is the subcommand
  // Skip flag values: flags like --timeout take a separate value token
  var subcommand = null;
  var skipNext = false;
  for (var k = 0; k < restTokens.length; k++) {
    if (skipNext) {
      skipNext = false;
      continue;
    }
    var token = restTokens[k];
    if (!token) continue;
    if (token.charAt(0) === '-') {
      // If flag has no = sign, next token is likely its value
      if (token.indexOf('=') === -1) {
        skipNext = true;
      }
      continue;
    }
    subcommand = token;
    break;
  }

  // No subcommand (bare command or flags only) → defaults to run → intercept
  if (!subcommand) return true;
  // Explicit 'run' subcommand → intercept
  if (subcommand === 'run') return true;
  // Known non-run subcommand → don't intercept
  if (NON_RUN_SUBCOMMANDS.indexOf(subcommand) !== -1) return false;
  // Unknown subcommand → assume run-like, intercept (safer to over-intercept)
  return true;
}

function stripOutputFilters(command) {
  // Find where the golangci-lint command appears as an actual command (not as argument)
  var match = command.match(/(?:^|\s)(golangci-lint)(?:\s|$)/);
  if (!match) return command;

  var idx = command.indexOf(match[1], match.index);

  var after = command.substring(idx);
  var before = command.substring(0, idx);

  // Strip output redirects that appear after golangci-lint token
  // Handle: > file, >> file, 2>&1, 2>file, &>file
  after = after.replace(/\s*2>&1/g, '');
  after = after.replace(/\s*2>\/\S*/g, '');
  after = after.replace(/\s*2>\s*\S+/g, '');
  after = after.replace(/\s*&>\S*/g, '');
  after = after.replace(/\s*>>\s*\S+/g, '');
  after = after.replace(/\s*>\s*\S+/g, '');

  // Strip pipe segments: find first | after golangci-lint and truncate
  var pipeIdx = after.indexOf('|');
  if (pipeIdx !== -1) {
    after = after.substring(0, pipeIdx);
  }

  return (before + after).trim();
}

function injectJsonOutputFlag(command) {
  var tokens = command.split(/\s+/);
  var kept = [];
  var i = 0;

  while (i < tokens.length) {
    var token = tokens[i];
    var stripped = false;
    var j;

    for (j = 0; j < OUTPUT_FLAG_PATTERNS_VALUE.length; j++) {
      if (token === OUTPUT_FLAG_PATTERNS_VALUE[j]) {
        stripped = true;
        i++;
        if (i < tokens.length && tokens[i].charAt(0) !== '-') {
          i++;
        }
        break;
      }
      if (token.indexOf(OUTPUT_FLAG_PATTERNS_VALUE[j] + '=') === 0) {
        stripped = true;
        i++;
        break;
      }
    }

    if (stripped) continue;

    for (j = 0; j < OUTPUT_FLAG_PATTERNS_BOOL.length; j++) {
      if (token === OUTPUT_FLAG_PATTERNS_BOOL[j]) {
        stripped = true;
        i++;
        break;
      }
    }

    if (stripped) continue;

    kept.push(token);
    i++;
  }

  return kept.join(' ') + ' --output.json.path stdout';
}

var COMPOUND_LINTERS = ['gocritic', 'gosec', 'staticcheck', 'revive', 'govet'];

function extractRule(linter, text) {
  if (COMPOUND_LINTERS.indexOf(linter) === -1) return null;
  var colonIdx = text.indexOf(': ');
  if (colonIdx === -1) return null;
  var prefix = text.substring(0, colonIdx);
  return prefix;
}

function parseDiagnostics(output) {
  var lines = output.split('\n');
  var seen = {};
  var linterCounts = {};
  var totalUnique = 0;

  for (var i = 0; i < lines.length; i++) {
    var line = lines[i].trim();
    if (!line) continue;
    try {
      var issue = JSON.parse(line);
      var linter = issue.FromLinter || '';
      var text = issue.Text || '';
      var rule = extractRule(linter, text);
      var key = linter + '|' + (rule || '');

      if (!seen[key]) {
        seen[key] = true;
        totalUnique++;
        linterCounts[linter] = (linterCounts[linter] || 0) + 1;
      }
    } catch (e) {
      // Skip non-JSON lines
    }
  }

  return { totalUnique: totalUnique, linterCounts: linterCounts };
}

function buildStrategyANudge(count, rawOutput) {
  var nudge = 'golangci-lint found ' + count + ' diagnostic' + (count !== 1 ? 's' : '') + '. ';
  nudge += 'Call golangci_lint_parse with the following JSON output to get fix guidance for all issues:\n\n';
  nudge += rawOutput;
  return nudge;
}

function buildStrategyBNudge(count, linterCounts, rawOutput) {
  var linterNames = Object.keys(linterCounts).sort();
  var summaryParts = [];
  for (var i = 0; i < linterNames.length; i++) {
    summaryParts.push(linterNames[i] + ': ' + linterCounts[linterNames[i]]);
  }

  var nudge = 'golangci-lint found ' + count + ' diagnostics across ' + linterNames.length + ' linter' + (linterNames.length !== 1 ? 's' : '') + '. ';
  nudge += 'Consider splitting into per-linter fixes:\n';
  nudge += summaryParts.join(', ') + '\n';
  nudge += 'For each linter, call golangci_lint_guide(linter="{linter}") for fix guidance.\n';
  nudge += 'The raw JSON output is available to pass to golangci_lint_parse.\n';

  var maxRawLen = 2000;
  if (rawOutput.length > maxRawLen) {
    nudge += '\nTruncated output (first ' + maxRawLen + ' chars):\n';
    nudge += rawOutput.substring(0, maxRawLen);
    nudge += '\n\nOutput truncated. Call golangci_lint_parse with the full output for complete guidance.';
  } else {
    nudge += '\nRaw JSON:\n' + rawOutput;
  }

  return nudge;
}

function truncateNudge(nudge) {
  var maxLen = 8000;
  if (nudge.length <= maxLen) return nudge;

  var truncationNotice = '\n\nOutput truncated. Call golangci_lint_parse with the full output for complete guidance.';
  var budget = maxLen - truncationNotice.length;
  return nudge.substring(0, budget) + truncationNotice;
}

module.exports = {
  COMPOUND_LINTERS: COMPOUND_LINTERS,
  OUTPUT_FLAG_PATTERNS_VALUE: OUTPUT_FLAG_PATTERNS_VALUE,
  OUTPUT_FLAG_PATTERNS_BOOL: OUTPUT_FLAG_PATTERNS_BOOL,
  NON_RUN_SUBCOMMANDS: NON_RUN_SUBCOMMANDS,
  extractRule: extractRule,
  extractInnerCommand: extractInnerCommand,
  isGolangciLintCommand: isGolangciLintCommand,
  stripOutputFilters: stripOutputFilters,
  injectJsonOutputFlag: injectJsonOutputFlag,
  parseDiagnostics: parseDiagnostics,
  buildStrategyANudge: buildStrategyANudge,
  buildStrategyBNudge: buildStrategyBNudge,
  truncateNudge: truncateNudge
};
