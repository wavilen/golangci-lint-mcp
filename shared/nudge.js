'use strict';

const OUTPUT_FLAG_PATTERNS_VALUE = [
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

const OUTPUT_FLAG_PATTERNS_BOOL = [
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

const NON_RUN_SUBCOMMANDS = ['cache', 'completion', 'config', 'custom', 'help', 'linters', 'version'];

function extractInnerCommand(command) {
  if (!command || typeof command !== 'string') return command || '';
  let cmd = command.trimStart();
  while (/^[A-Za-z_][A-Za-z0-9_.]*=\S+/.test(cmd)) {
    cmd = cmd.replace(/^[A-Za-z_][A-Za-z0-9_.]*=\S+\s*/, '');
  }
  const shellMatch = cmd.match(/^(?:\/[\w/]+\/)?(?:bash|sh|zsh|dash|ksh)\s+-c\s+["']?/);
  if (shellMatch) {
    cmd = cmd.substring(shellMatch[0].length);
    if (cmd.endsWith('"') || cmd.endsWith("'")) {
      cmd = cmd.slice(0, -1);
    }
  }
  const cdMatch = cmd.match(/^cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*/);
  if (cdMatch) {
    cmd = cmd.substring(cdMatch[0].length);
  }
  return cmd;
}

function isGolangciLintCommand(command) {
  if (!command || typeof command !== 'string') return false;
  const cmd = extractInnerCommand(command);
  // Extract the first token (the actual command being invoked)
  const match = cmd.match(/^(\S+)/);
  if (!match) return false;
  const firstToken = match[1];
  // Must be "golangci-lint" or a path ending with "/golangci-lint"
  if (firstToken !== 'golangci-lint' && !firstToken.endsWith('/golangci-lint')) return false;

  // Extract remaining tokens after the command itself
  const rest = cmd.substring(match[0].length).trim();
  const restTokens = rest.split(/\s+/);

  // Find the first non-flag token (not starting with '-') — that is the subcommand
  // Skip flag values: flags like --timeout take a separate value token
  let subcommand = null;
  let skipNext = false;
  for (let k = 0; k < restTokens.length; k++) {
    if (skipNext) {
      skipNext = false;
      continue;
    }
    const token = restTokens[k];
    if (!token) continue;
    if (token.charAt(0) === '-') {
      // If flag has no = sign, next token is likely its value (unless next token is a flag itself)
      if (token.indexOf('=') === -1) {
        if (k + 1 < restTokens.length && restTokens[k + 1].charAt(0) !== '-') {
          skipNext = true;
        }
      }
      continue;
    }
    // Handle compound commands: strip semicolon boundary from subcommand token
    // e.g. "version;" from "golangci-lint version; echo 'done'"
    subcommand = token.split(';')[0];
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
  const match = command.match(/(?:^|\s)(golangci-lint)(?:\s|$)/);
  if (!match) return command;

  const idx = command.indexOf(match[1], match.index);

  let after = command.substring(idx);
  const before = command.substring(0, idx);

  // Strip output redirects that appear after golangci-lint token
  // Handle: > file, >> file, 2>&1, 2>file, &>file
  after = after.replace(/\s*2>&1/g, '');
  after = after.replace(/\s*2>\/\S*/g, '');
  after = after.replace(/\s*2>\s*\S+/g, '');
  after = after.replace(/\s*&>\S*/g, '');
  after = after.replace(/\s*>>\s*\S+/g, '');
  after = after.replace(/\s*>\s*\S+/g, '');

  // Strip pipe segments: find first | after golangci-lint and truncate
  const pipeIdx = after.indexOf('|');
  if (pipeIdx !== -1) {
    after = after.substring(0, pipeIdx);
  }

  return (before + after).trim();
}

function injectJsonOutputFlag(command) {
  const tokens = command.split(/\s+/);
  const kept = [];
  let i = 0;

  while (i < tokens.length) {
    const token = tokens[i];
    let stripped = false;
    let j;

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

const COMPOUND_LINTERS = ['gocritic', 'gosec', 'staticcheck', 'revive', 'govet'];

function extractRule(linter, text) {
  if (COMPOUND_LINTERS.indexOf(linter) === -1) return null;
  const colonIdx = text.indexOf(': ');
  if (colonIdx === -1) return null;
  const prefix = text.substring(0, colonIdx);
  return prefix;
}

function splitOnUnquotedSemicolons(command) {
  const segments = [];
  let current = '';
  let inSingle = false;
  let inDouble = false;

  for (let i = 0; i < command.length; i++) {
    const ch = command[i];
    if (ch === "'" && !inDouble) {
      inSingle = !inSingle;
      current += ch;
    } else if (ch === '"' && !inSingle) {
      inDouble = !inDouble;
      current += ch;
    } else if (ch === ';' && !inSingle && !inDouble) {
      const trimmed = current.trim();
      if (trimmed) segments.push(trimmed);
      current = '';
    } else {
      current += ch;
    }
  }
  const trimmed = current.trim();
  if (trimmed) segments.push(trimmed);
  return segments;
}

function findUnquotedPipe(str) {
  let inSingle = false;
  let inDouble = false;
  for (let i = 0; i < str.length; i++) {
    const ch = str[i];
    if (ch === "'" && !inDouble) {
      inSingle = !inSingle;
    } else if (ch === '"' && !inSingle) {
      inDouble = !inDouble;
    } else if (ch === '|' && !inSingle && !inDouble) {
      return i;
    }
  }
  return -1;
}

function splitCompoundCommand(command) {
  if (!command || typeof command !== 'string') {
    return { lintCommand: command || '', jqSegment: '', echoSuffix: '' };
  }

  const segments = splitOnUnquotedSemicolons(command);

  // Find the first segment containing golangci-lint as a word
  let lintSegment = null;
  for (let s = 0; s < segments.length; s++) {
    if (/\bgolangci-lint\b/.test(segments[s])) {
      lintSegment = segments[s];
      break;
    }
  }

  if (lintSegment === null) {
    return { lintCommand: command, jqSegment: '', echoSuffix: '' };
  }

  // Detect exit-code echo patterns: || echo ... or && echo ...
  const echoMatch = lintSegment.match(/\s*(\|\|\s*echo\s+\S.*|&&\s*echo\s+\S.*)$/);
  let echoSuffix = '';
  if (echoMatch) {
    echoSuffix = echoMatch[1].trim();
    lintSegment = lintSegment.substring(0, echoMatch.index).trim();
  }

  // Detect jq pipe
  let jqSegment = '';
  const pipeIdx = findUnquotedPipe(lintSegment);
  if (pipeIdx !== -1) {
    const afterPipe = lintSegment.substring(pipeIdx + 1).trimStart();
    // Check if pipe target is jq (with optional path prefix)
    const jqMatch = afterPipe.match(/^(?:\/[\w/]+\/)?jq\b/);
    if (jqMatch) {
      // Find the next unquoted pipe (or end of string) to delimit the jq segment
      const secondPipeIdx = findUnquotedPipe(afterPipe);
      if (secondPipeIdx !== -1) {
        jqSegment = '| ' + afterPipe.substring(0, secondPipeIdx).trim();
        // lintSegment is everything before the first pipe
        lintSegment = lintSegment.substring(0, pipeIdx).trim();
      } else {
        jqSegment = '| ' + afterPipe.trim();
        lintSegment = lintSegment.substring(0, pipeIdx).trim();
      }
    }
    // If pipe target is NOT jq, leave the pipe in lintSegment for stripOutputFilters
  }

  return { lintCommand: lintSegment.trim(), jqSegment: jqSegment.trim(), echoSuffix: echoSuffix.trim() };
}

function injectJqFilter(jqSegment) {
  if (!jqSegment || typeof jqSegment !== 'string') return '';

  // Detect trivial jq: | jq, | jq '.', | jq ".", | jq ., with optional path prefix
  const trivialMatch = jqSegment.match(/^\|\s*(\/[\w/]+\/)?jq(?:\s+(?:'\.'|"\."|\.))?\s*$/);
  if (trivialMatch) {
    const pathPrefix = trivialMatch[1] || '';
    return "| " + pathPrefix + "jq '{Issues: .Issues, Report: .Report}'";
  }

  // Non-trivial — leave unchanged
  return jqSegment;
}

function parseDiagnostics(output) {
  // Try envelope format first: {Issues: [...], Report: {...}}
  try {
    const parsed = JSON.parse(output.trim());
    if (parsed && Array.isArray(parsed.Issues)) {
      const seen = {};
      const linterCounts = {};
      let totalUnique = 0;

      for (let i = 0; i < parsed.Issues.length; i++) {
        const issue = parsed.Issues[i];
        const linter = issue.FromLinter || '';
        const text = issue.Text || '';
        const rule = extractRule(linter, text);
        const key = linter + '|' + (rule || '');

        if (!seen[key]) {
          seen[key] = true;
          totalUnique++;
          linterCounts[linter] = (linterCounts[linter] || 0) + 1;
        }
      }

      const warnings = (parsed.Report && parsed.Report.Warnings) ? parsed.Report.Warnings : [];
      return { totalUnique: totalUnique, linterCounts: linterCounts, warnings: warnings };
    }
  } catch (_e) {
    // Not a single JSON object — fall through to line-by-line parsing
  }

  // Line-by-line parsing (bare JSON lines / NDJSON)
  const lines = output.split('\n');
  const seen = {};
  const linterCounts = {};
  let totalUnique = 0;

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();
    if (!line) continue;
    try {
      const issue = JSON.parse(line);
      const linter = issue.FromLinter || '';
      const text = issue.Text || '';
      const rule = extractRule(linter, text);
      const key = linter + '|' + (rule || '');

      if (!seen[key]) {
        seen[key] = true;
        totalUnique++;
        linterCounts[linter] = (linterCounts[linter] || 0) + 1;
      }
    } catch (_e) {
      // Skip non-JSON lines
    }
  }

  return { totalUnique: totalUnique, linterCounts: linterCounts, warnings: [] };
}

function buildStrategyANudge(count, rawOutput) {
  let nudge = 'golangci-lint found ' + count + ' diagnostic' + (count !== 1 ? 's' : '') + '. ';
  nudge += 'Call golangci_lint_parse with the following JSON output to get fix guidance for all issues:\n\n';
  nudge += rawOutput;
  return nudge;
}

function buildStrategyBNudge(count, linterCounts, rawOutput) {
  const linterNames = Object.keys(linterCounts).sort();
  const summaryParts = [];
  for (let i = 0; i < linterNames.length; i++) {
    summaryParts.push(linterNames[i] + ': ' + linterCounts[linterNames[i]]);
  }

  let nudge = 'golangci-lint found ' + count + ' diagnostics across ' + linterNames.length + ' linter' + (linterNames.length !== 1 ? 's' : '') + '. ';
  nudge += 'Linters: ' + summaryParts.join(', ') + '. ';
  nudge += 'Call golangci_lint_run with a full-project path to get package breakdown and strategy recommendation.\n';

  const maxRawLen = 2000;
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
  const maxLen = 8000;
  if (nudge.length <= maxLen) return nudge;

  const truncationNotice = '\n\nOutput truncated. Call golangci_lint_parse with the full output for complete guidance.';
  const budget = maxLen - truncationNotice.length;
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
  splitCompoundCommand: splitCompoundCommand,
  injectJqFilter: injectJqFilter,
  parseDiagnostics: parseDiagnostics,
  buildStrategyANudge: buildStrategyANudge,
  buildStrategyBNudge: buildStrategyBNudge,
  truncateNudge: truncateNudge
};
