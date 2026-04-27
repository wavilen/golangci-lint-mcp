'use strict';

const { describe, it } = require('node:test');
const assert = require('node:assert/strict');
const {
  isGolangciLintCommand,
  extractInnerCommand,
  stripOutputFilters,
  injectJsonOutputFlag,
  parseDiagnostics,
  buildStrategyANudge,
  buildStrategyBNudge,
  truncateNudge,
  extractRule,
  splitCompoundCommand,
  injectJqFilter,
  COMPOUND_LINTERS,
  OUTPUT_FLAG_PATTERNS_VALUE,
  OUTPUT_FLAG_PATTERNS_BOOL,
  NON_RUN_SUBCOMMANDS
} = require('./nudge.js');

// ─── isGolangciLintCommand: positive cases ────────────────────────────────

describe('isGolangciLintCommand — positive cases', () => {
  it('bare command: golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('golangci-lint run'), true);
  });

  it('bare command without args: golangci-lint', () => {
    assert.equal(isGolangciLintCommand('golangci-lint'), true);
  });

  it('leading whitespace:   golangci-lint run ./...', () => {
    assert.equal(isGolangciLintCommand('  golangci-lint run ./...'), true);
  });

  it('absolute path: /usr/bin/golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('/usr/bin/golangci-lint run'), true);
  });

  it('absolute path: /usr/local/bin/golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('/usr/local/bin/golangci-lint run'), true);
  });

  it('relative path: ./golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('./golangci-lint run'), true);
  });

  it('home path: ~/bin/golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('~/bin/golangci-lint run'), true);
  });

  it('single env var prefix: FOO=bar golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('FOO=bar golangci-lint run'), true);
  });

  it('go flags env var: GOFLAGS=... golangci-lint run ./...', () => {
    assert.equal(isGolangciLintCommand('GOFLAGS=... golangci-lint run ./...'), true);
  });

  it('multiple env var prefixes: VAR1=val VAR2=val golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('VAR1=val VAR2=val golangci-lint run'), true);
  });

  it('env var with extra whitespace:   FOO=bar   golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('  FOO=bar   golangci-lint run'), true);
  });

  it('cd prefix with && : cd /some/dir && golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('cd /some/dir && golangci-lint run'), true);
  });

  it('cd prefix with && and grep -v pipe', () => {
    assert.equal(isGolangciLintCommand("cd /home/user/project && golangci-lint run 2>&1 | grep -v '^tmp/\\|^level='"), true);
  });

  it('cd prefix with non-golangci-lint: cd /some/dir && go test', () => {
    assert.equal(isGolangciLintCommand('cd /some/dir && go test ./...'), false);
  });

  it('cd prefix with golangci-lint version: cd /dir && golangci-lint version', () => {
    assert.equal(isGolangciLintCommand('cd /dir && golangci-lint version'), false);
  });
});

// ─── isGolangciLintCommand: negative cases ────────────────────────────────

describe('isGolangciLintCommand — negative cases', () => {
  it('echo golangci-lint', () => {
    assert.equal(isGolangciLintCommand('echo golangci-lint'), false);
  });

  it('echo "golangci-lint"', () => {
    assert.equal(isGolangciLintCommand('echo "golangci-lint"'), false);
  });

  it('cat docs/golangci-lint.md', () => {
    assert.equal(isGolangciLintCommand('cat docs/golangci-lint.md'), false);
  });

  it('git commit -m "fix golangci-lint"', () => {
    assert.equal(isGolangciLintCommand('git commit -m "fix golangci-lint"'), false);
  });

  it('ls golangci-lint/', () => {
    assert.equal(isGolangciLintCommand('ls golangci-lint/'), false);
  });

  it('rm golangci-lint.log', () => {
    assert.equal(isGolangciLintCommand('rm golangci-lint.log'), false);
  });

  it('grep golangci-lint file.txt', () => {
    assert.equal(isGolangciLintCommand('grep golangci-lint file.txt'), false);
  });

  it('which golangci-lint', () => {
    assert.equal(isGolangciLintCommand('which golangci-lint'), false);
  });

  it('empty string', () => {
    assert.equal(isGolangciLintCommand(''), false);
  });

  it('echo (no golangci-lint at all)', () => {
    assert.equal(isGolangciLintCommand('echo'), false);
  });

  it('vim golangci-lint-config.yaml', () => {
    assert.equal(isGolangciLintCommand('vim golangci-lint-config.yaml'), false);
  });

  it('null input', () => {
    assert.equal(isGolangciLintCommand(null), false);
  });

  it('undefined input', () => {
    assert.equal(isGolangciLintCommand(undefined), false);
  });

  it('non-string input (number)', () => {
    assert.equal(isGolangciLintCommand(123), false);
  });
});

// ─── isGolangciLintCommand — only run subcommand ──────────────────────────

describe('isGolangciLintCommand — only run subcommand', () => {
  // Positive cases: should return true (run or bare = defaults to run)
  it('explicit run: golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('golangci-lint run'), true);
  });

  it('bare command: golangci-lint (defaults to run)', () => {
    assert.equal(isGolangciLintCommand('golangci-lint'), true);
  });

  it('run with args: golangci-lint run ./...', () => {
    assert.equal(isGolangciLintCommand('golangci-lint run ./...'), true);
  });

  it('flags before run: golangci-lint --timeout 5m run', () => {
    assert.equal(isGolangciLintCommand('golangci-lint --timeout 5m run'), true);
  });

  it('flags only, no subcommand: golangci-lint -E errcheck (defaults to run)', () => {
    assert.equal(isGolangciLintCommand('golangci-lint -E errcheck'), true);
  });

  it('path + run: /usr/bin/golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('/usr/bin/golangci-lint run'), true);
  });

  it('env vars + run: FOO=bar golangci-lint run', () => {
    assert.equal(isGolangciLintCommand('FOO=bar golangci-lint run'), true);
  });

  it('env vars + bare: FOO=bar golangci-lint (defaults to run)', () => {
    assert.equal(isGolangciLintCommand('FOO=bar golangci-lint'), true);
  });

  // Negative cases: non-run subcommands should return false
  it('cache subcommand: golangci-lint cache', () => {
    assert.equal(isGolangciLintCommand('golangci-lint cache'), false);
  });

  it('cache clean: golangci-lint cache clean', () => {
    assert.equal(isGolangciLintCommand('golangci-lint cache clean'), false);
  });

  it('version subcommand: golangci-lint version', () => {
    assert.equal(isGolangciLintCommand('golangci-lint version'), false);
  });

  it('linters subcommand: golangci-lint linters', () => {
    assert.equal(isGolangciLintCommand('golangci-lint linters'), false);
  });

  it('help subcommand: golangci-lint help', () => {
    assert.equal(isGolangciLintCommand('golangci-lint help'), false);
  });

  it('config subcommand: golangci-lint config', () => {
    assert.equal(isGolangciLintCommand('golangci-lint config'), false);
  });

  it('config path: golangci-lint config path', () => {
    assert.equal(isGolangciLintCommand('golangci-lint config path'), false);
  });

  it('completion subcommand: golangci-lint completion bash', () => {
    assert.equal(isGolangciLintCommand('golangci-lint completion bash'), false);
  });

  it('custom subcommand: golangci-lint custom', () => {
    assert.equal(isGolangciLintCommand('golangci-lint custom'), false);
  });

  it('path + cache: /usr/bin/golangci-lint cache', () => {
    assert.equal(isGolangciLintCommand('/usr/bin/golangci-lint cache'), false);
  });

  it('env vars + version: FOO=bar golangci-lint version', () => {
    assert.equal(isGolangciLintCommand('FOO=bar golangci-lint version'), false);
  });

  it('flags before non-run: golangci-lint --timeout 5m cache clean', () => {
    assert.equal(isGolangciLintCommand('golangci-lint --timeout 5m cache clean'), false);
  });
});

// ─── stripOutputFilters regression tests ──────────────────────────────────

describe('stripOutputFilters', () => {
  it('removes pipe segments: golangci-lint run | head', () => {
    assert.equal(stripOutputFilters('golangci-lint run | head'), 'golangci-lint run');
  });

  it('removes output redirect: golangci-lint run > out.txt', () => {
    assert.equal(stripOutputFilters('golangci-lint run > out.txt'), 'golangci-lint run');
  });

  it('removes append redirect: golangci-lint run >> out.txt', () => {
    assert.equal(stripOutputFilters('golangci-lint run >> out.txt'), 'golangci-lint run');
  });

  it('removes stderr redirect: golangci-lint run 2>&1', () => {
    assert.equal(stripOutputFilters('golangci-lint run 2>&1'), 'golangci-lint run');
  });

  it('removes combined pipe and redirect', () => {
    assert.equal(stripOutputFilters('golangci-lint run | grep err > out.txt'), 'golangci-lint run');
  });

  it('no filters needed: returns command unchanged', () => {
    assert.equal(stripOutputFilters('golangci-lint run ./...'), 'golangci-lint run ./...');
  });

  it('non-golangci-lint command: returns unchanged', () => {
    assert.equal(stripOutputFilters('echo hello'), 'echo hello');
  });

  it('removes &> redirect', () => {
    assert.equal(stripOutputFilters('golangci-lint run &>out.txt'), 'golangci-lint run');
  });
});

// ─── injectJsonOutputFlag regression tests ────────────────────────────────

describe('injectJsonOutputFlag', () => {
  it('adds --output.json.path stdout to basic command', () => {
    assert.equal(
      injectJsonOutputFlag('golangci-lint run'),
      'golangci-lint run --output.json.path stdout'
    );
  });

  it('strips conflicting --out-format tab and injects json flag', () => {
    const result = injectJsonOutputFlag('golangci-lint run --out-format tab');
    assert.ok(result.indexOf('--out-format') === -1, 'should not contain --out-format');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('strips --color flag and injects json flag', () => {
    const result = injectJsonOutputFlag('golangci-lint run --color always');
    assert.ok(result.indexOf('--color') === -1, 'should not contain --color');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('replaces existing --output.json.path with new value', () => {
    const result = injectJsonOutputFlag('golangci-lint run --output.json.path stderr');
    assert.ok(result.indexOf('stderr') === -1, 'should not contain old value');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain new value');
  });
});

// ─── parseDiagnostics regression tests ────────────────────────────────────

describe('parseDiagnostics', () => {
  it('parses single diagnostic', () => {
    const input = JSON.stringify({ FromLinter: 'errcheck', Text: 'error return is unchecked' });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.errcheck, 1);
  });

  it('parses multiple linters', () => {
    const line1 = JSON.stringify({ FromLinter: 'errcheck', Text: 'unchecked' });
    const line2 = JSON.stringify({ FromLinter: 'govet', Text: 'lostcancel: cancel function not used' });
    const line3 = JSON.stringify({ FromLinter: 'errcheck', Text: 'another unchecked' });
    const result = parseDiagnostics(line1 + '\n' + line2 + '\n' + line3);
    assert.equal(result.totalUnique, 2);
    assert.equal(result.linterCounts.errcheck, 1);
    assert.equal(result.linterCounts.govet, 1);
  });

  it('handles empty input', () => {
    const result = parseDiagnostics('');
    assert.equal(result.totalUnique, 0);
    assert.deepEqual(result.linterCounts, {});
  });

  it('handles invalid JSON lines', () => {
    const result = parseDiagnostics('not json\nalso not json');
    assert.equal(result.totalUnique, 0);
  });

  it('extracts rule from compound linter text', () => {
    const line = JSON.stringify({ FromLinter: 'govet', Text: 'lostcancel: cancel function is not used on deferral path' });
    const result = parseDiagnostics(line);
    assert.equal(result.totalUnique, 1);
  });

  // ─── envelope format tests ────────────────────────────────────────────────

  it('parses envelope with single issue', () => {
    const input = JSON.stringify({ Issues: [{ FromLinter: 'errcheck', Text: 'unchecked' }], Report: {} });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.errcheck, 1);
    assert.deepEqual(result.warnings, []);
  });

  it('parses envelope with warnings', () => {
    const input = JSON.stringify({ Issues: [], Report: { Warnings: [{ Text: 'config deprecated' }] } });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 0);
    assert.deepEqual(result.linterCounts, {});
    assert.equal(result.warnings.length, 1);
    assert.equal(result.warnings[0].Text, 'config deprecated');
  });

  it('parses envelope with issues AND warnings', () => {
    const input = JSON.stringify({
      Issues: [{ FromLinter: 'errcheck', Text: 'x' }],
      Report: { Warnings: [{ Text: 'config deprecated' }, { Text: 'slow linter' }] }
    });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.errcheck, 1);
    assert.equal(result.warnings.length, 2);
    assert.equal(result.warnings[0].Text, 'config deprecated');
    assert.equal(result.warnings[1].Text, 'slow linter');
  });

  it('parses envelope with compound linter', () => {
    const input = JSON.stringify({
      Issues: [{ FromLinter: 'govet', Text: 'lostcancel: cancel function not used' }],
      Report: {}
    });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.govet, 1);
    assert.deepEqual(result.warnings, []);
  });

  it('parses empty envelope', () => {
    const input = JSON.stringify({ Issues: [], Report: {} });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 0);
    assert.deepEqual(result.linterCounts, {});
    assert.deepEqual(result.warnings, []);
  });

  it('parses envelope with missing Report', () => {
    const input = JSON.stringify({ Issues: [{ FromLinter: 'errcheck', Text: 'x' }] });
    const result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.errcheck, 1);
    assert.deepEqual(result.warnings, []);
  });

  it('returns warnings array in line-by-line path', () => {
    const input = JSON.stringify({ FromLinter: 'errcheck', Text: 'unchecked' });
    const result = parseDiagnostics(input);
    assert.ok(Array.isArray(result.warnings));
    assert.equal(result.warnings.length, 0);
  });
});

// ─── isGolangciLintCommand — shell wrapper detection ──────────────────────

describe('isGolangciLintCommand — shell wrapper detection', () => {
  it('bash -c wrapper: bash -c "golangci-lint run"', () => {
    assert.equal(isGolangciLintCommand('bash -c "golangci-lint run"'), true);
  });

  it('/bin/bash -c wrapper: /bin/bash -c golangci-lint run ./...', () => {
    assert.equal(isGolangciLintCommand('/bin/bash -c golangci-lint run ./...'), true);
  });

  it('sh -c wrapper with single quotes: sh -c \'golangci-lint run\'', () => {
    assert.equal(isGolangciLintCommand("sh -c 'golangci-lint run'"), true);
  });

  it('zsh -c wrapper: /usr/bin/zsh -c "golangci-lint run"', () => {
    assert.equal(isGolangciLintCommand('/usr/bin/zsh -c "golangci-lint run"'), true);
  });

  it('bash -c with echo should NOT match: bash -c "echo golangci-lint"', () => {
    assert.equal(isGolangciLintCommand('bash -c "echo golangci-lint"'), false);
  });

  it('env var + shell wrapper: FOO=bar bash -c "golangci-lint run"', () => {
    assert.equal(isGolangciLintCommand('FOO=bar bash -c "golangci-lint run"'), true);
  });

  it('shell wrapper with non-run subcommand: bash -c "golangci-lint version"', () => {
    assert.equal(isGolangciLintCommand('bash -c "golangci-lint version"'), false);
  });
});

// ─── extractInnerCommand ──────────────────────────────────────────────────

describe('extractInnerCommand', () => {
  it('bash -c with double quotes + grep -v pipe', () => {
    assert.equal(
      extractInnerCommand('bash -c "golangci-lint run | grep -v SA1019"'),
      'golangci-lint run | grep -v SA1019'
    );
  });

  it('sh -c with single quotes + grep -v pipe', () => {
    assert.equal(
      extractInnerCommand("sh -c 'golangci-lint run ./... | grep -v SA1019'"),
      'golangci-lint run ./... | grep -v SA1019'
    );
  });

  it('/bin/bash -c with redirect + grep -v pipe', () => {
    assert.equal(
      extractInnerCommand('/bin/bash -c "golangci-lint run 2>&1 | grep -v SA1019"'),
      'golangci-lint run 2>&1 | grep -v SA1019'
    );
  });

  it('no wrapper passthrough', () => {
    assert.equal(
      extractInnerCommand('golangci-lint run ./...'),
      'golangci-lint run ./...'
    );
  });

  it('env var + shell wrapper', () => {
    assert.equal(
      extractInnerCommand('FOO=bar bash -c "golangci-lint run"'),
      'golangci-lint run'
    );
  });

  it('multiple env vars + shell wrapper', () => {
    assert.equal(
      extractInnerCommand('VAR1=a VAR2=b sh -c "golangci-lint run ./..."'),
      'golangci-lint run ./...'
    );
  });

  it('plain command with no change needed', () => {
    assert.equal(
      extractInnerCommand('golangci-lint run'),
      'golangci-lint run'
    );
  });

  it('cd prefix with && : cd /some/dir && golangci-lint run', () => {
    assert.equal(
      extractInnerCommand('cd /some/dir && golangci-lint run'),
      'golangci-lint run'
    );
  });

  it('cd prefix with && and 2>&1 and grep pipe', () => {
    assert.equal(
      extractInnerCommand("cd /home/wavilen/Workspace/golangcilint-mcp && golangci-lint run 2>&1 | grep -v '^tmp/\\|^level='"),
      "golangci-lint run 2>&1 | grep -v '^tmp/\\|^level='"
    );
  });

  it('cd prefix with ; separator: cd /dir ; golangci-lint run', () => {
    assert.equal(
      extractInnerCommand('cd /dir ; golangci-lint run'),
      'golangci-lint run'
    );
  });
});

// ─── stripOutputFilters — shell-wrapped pipe scenarios ─────────────────────

describe('stripOutputFilters — shell-wrapped pipe scenarios', () => {
  it('after extractInnerCommand: bash -c wrapped command with grep -v pipe', () => {
    const inner = extractInnerCommand('bash -c "golangci-lint run | grep -v SA1019"');
    const stripped = stripOutputFilters(inner);
    assert.equal(stripped, 'golangci-lint run');
  });

  it('after extractInnerCommand: /bin/bash -c wrapped with redirect + grep -v', () => {
    const inner = extractInnerCommand("/bin/bash -c 'golangci-lint run ./... 2>&1 | grep -v \"SA1019\"'");
    const stripped = stripOutputFilters(inner);
    assert.equal(stripped, 'golangci-lint run ./...');
  });
});

// ─── full pipeline: extractInnerCommand → stripOutputFilters → injectJsonOutputFlag ──

describe('full pipeline: shell-wrapped + piped commands', () => {
  function processCommand(raw) {
    const cdMatch = raw.match(/^(cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*)/);
    const cdPrefix = cdMatch ? cdMatch[0] : '';
    const inner = extractInnerCommand(raw);
    const stripped = stripOutputFilters(inner);
    return cdPrefix + injectJsonOutputFlag(stripped);
  }

  it('bash -c wrapped command with grep -v produces clean output', () => {
    assert.equal(
      processCommand('bash -c "golangci-lint run ./... | grep -v SA1019"'),
      'golangci-lint run ./... --output.json.path stdout'
    );
  });

  it('plain command with grep -v pipe still works', () => {
    assert.equal(
      processCommand('golangci-lint run | grep -v SA1019'),
      'golangci-lint run --output.json.path stdout'
    );
  });

  it('plain command with no pipe still works', () => {
    assert.equal(
      processCommand('golangci-lint run ./...'),
      'golangci-lint run ./... --output.json.path stdout'
    );
  });

  it('env var + bash -c with pipe produces clean output', () => {
    assert.equal(
      processCommand('FOO=bar bash -c "golangci-lint run ./... | grep -v SA1019"'),
      'golangci-lint run ./... --output.json.path stdout'
    );
  });

  it('cd prefix with && and grep -v pipe produces clean output', () => {
    assert.equal(
      processCommand("cd /home/wavilen/Workspace/golangcilint-mcp && golangci-lint run 2>&1 | grep -v '^tmp/\\|^level='"),
      "cd /home/wavilen/Workspace/golangcilint-mcp && golangci-lint run --output.json.path stdout"
    );
  });

  it('cd prefix preserves explicit path after golangci-lint run', () => {
    assert.equal(
      processCommand('cd /home/wavilen/Workspace/golangcilint-mcp && golangci-lint run ./pkg/... 2>&1 | grep -v SA1019'),
      'cd /home/wavilen/Workspace/golangcilint-mcp && golangci-lint run ./pkg/... --output.json.path stdout'
    );
  });

  it('cd to non-project dir is preserved in output', () => {
    assert.equal(
      processCommand("cd /tmp && golangci-lint run ./... 2>&1 | grep -v SA1019"),
      'cd /tmp && golangci-lint run ./... --output.json.path stdout'
    );
  });

  it('cd with quoted path is preserved', () => {
    assert.equal(
      processCommand('cd "/path with spaces" && golangci-lint run 2>&1 | grep -v SA'),
      'cd "/path with spaces" && golangci-lint run --output.json.path stdout'
    );
  });
});

// ─── shared module constants ──────────────────────────────────────────────

describe('shared module constants', () => {
  it('COMPOUND_LINTERS includes expected linters', () => {
    assert.ok(COMPOUND_LINTERS.indexOf('gocritic') !== -1);
    assert.ok(COMPOUND_LINTERS.indexOf('gosec') !== -1);
    assert.ok(COMPOUND_LINTERS.indexOf('staticcheck') !== -1);
    assert.ok(COMPOUND_LINTERS.indexOf('revive') !== -1);
    assert.ok(COMPOUND_LINTERS.indexOf('govet') !== -1);
  });

  it('NON_RUN_SUBCOMMANDS includes expected subcommands', () => {
    assert.ok(NON_RUN_SUBCOMMANDS.indexOf('cache') !== -1);
    assert.ok(NON_RUN_SUBCOMMANDS.indexOf('version') !== -1);
    assert.ok(NON_RUN_SUBCOMMANDS.indexOf('help') !== -1);
  });

  it('OUTPUT_FLAG_PATTERNS_VALUE includes expected flags', () => {
    assert.ok(OUTPUT_FLAG_PATTERNS_VALUE.indexOf('--output.json.path') !== -1);
    assert.ok(OUTPUT_FLAG_PATTERNS_VALUE.indexOf('--out-format') !== -1);
  });

  it('OUTPUT_FLAG_PATTERNS_BOOL includes expected flags', () => {
    assert.ok(OUTPUT_FLAG_PATTERNS_BOOL.indexOf('--verbose') !== -1);
    assert.ok(OUTPUT_FLAG_PATTERNS_BOOL.indexOf('-v') !== -1);
  });
});

// ─── nudge functions ─────────────────────────────────────────────────────

describe('buildStrategyANudge', () => {
  it('builds nudge with count and raw output', () => {
    const nudge = buildStrategyANudge(2, 'raw');
    assert.ok(nudge.indexOf('2 diagnostics') !== -1);
    assert.ok(nudge.indexOf('raw') !== -1);
  });

  it('singular: 1 diagnostic', () => {
    const nudge = buildStrategyANudge(1, 'raw');
    assert.ok(nudge.indexOf('1 diagnostic') !== -1);
    assert.ok(nudge.indexOf('diagnostics') === -1);
  });
});

describe('buildStrategyBNudge', () => {
  it('builds nudge with linter breakdown', () => {
    const nudge = buildStrategyBNudge(5, { errcheck: 3, govet: 2 }, 'raw');
    assert.ok(nudge.indexOf('5 diagnostics') !== -1);
    assert.ok(nudge.indexOf('2 linters') !== -1);
  });

  it('singular: 1 linter', () => {
    const nudge = buildStrategyBNudge(3, { errcheck: 3 }, 'raw');
    assert.ok(nudge.indexOf('1 linter') !== -1);
    assert.ok(nudge.indexOf('linters') === -1);
  });

  it('directs agents to golangci_lint_run (not golangci_lint_guide)', () => {
    const nudge = buildStrategyBNudge(5, { errcheck: 3, govet: 2 }, 'raw');
    assert.ok(nudge.indexOf('golangci_lint_run') !== -1, 'nudge should mention golangci_lint_run');
    assert.equal(nudge.indexOf('golangci_lint_guide'), -1, 'nudge should NOT mention golangci_lint_guide');
  });

  it('nudge body is concise for short raw output (under 500 chars)', () => {
    const nudge = buildStrategyBNudge(5, { errcheck: 3, govet: 2 }, 'short');
    // The nudge body (without raw output) should be under 500 chars
    // Find where the raw output section starts (after the action guidance line)
    const lines = nudge.split('\n');
    let bodyLen = 0;
    for (let i = 0; i < lines.length; i++) {
      bodyLen += lines[i].length + 1;
      if (lines[i].indexOf('recommendation') !== -1) {
        // Action guidance ends here; rest is raw output section
        break;
      }
    }
    assert.ok(bodyLen < 500, 'nudge body (without raw output) should be under 500 chars, got ' + bodyLen);
  });
});

describe('truncateNudge', () => {
  it('does not truncate short nudge', () => {
    const nudge = 'short message';
    assert.equal(truncateNudge(nudge), nudge);
  });

  it('truncates long nudge with notice', () => {
    const nudge = 'x'.repeat(9000);
    const result = truncateNudge(nudge);
    assert.ok(result.length <= 8000);
    assert.ok(result.indexOf('truncated') !== -1 || result.indexOf('Truncated') !== -1);
  });
});

describe('extractRule', () => {
  it('returns null for non-compound linter', () => {
    assert.equal(extractRule('errcheck', 'unchecked error'), null);
  });

  it('returns rule from compound linter text', () => {
    assert.equal(extractRule('govet', 'lostcancel: something'), 'lostcancel');
  });

  it('returns null for compound linter with no colon', () => {
    assert.equal(extractRule('govet', 'no colon here'), null);
  });
});

// ─── splitCompoundCommand ─────────────────────────────────────────────────

describe('splitCompoundCommand', () => {
  it('splits on semicolon: echo "x"; golangci-lint run ./...', () => {
    const result = splitCompoundCommand('echo "x"; golangci-lint run ./...');
    assert.deepEqual(result, { lintCommand: 'golangci-lint run ./...', jqSegment: '', echoSuffix: '' });
  });

  it('extracts jq pipe: golangci-lint run | jq \'.\'', () => {
    const result = splitCompoundCommand("golangci-lint run | jq '.'");
    assert.equal(result.lintCommand, 'golangci-lint run');
    assert.equal(result.jqSegment, "| jq '.'");
    assert.equal(result.echoSuffix, '');
  });

  it('extracts jq pipe with path prefix: golangci-lint run | /usr/bin/jq \'.\'', () => {
    const result = splitCompoundCommand("golangci-lint run | /usr/bin/jq '.'");
    assert.equal(result.lintCommand, 'golangci-lint run');
    assert.equal(result.jqSegment, "| /usr/bin/jq '.'");
    assert.equal(result.echoSuffix, '');
  });

  it('preserves || echo exit-code pattern', () => {
    const result = splitCompoundCommand("golangci-lint run || echo 'lint failed'");
    assert.equal(result.lintCommand, 'golangci-lint run');
    assert.equal(result.jqSegment, '');
    assert.equal(result.echoSuffix, "|| echo 'lint failed'");
  });

  it('preserves && echo exit-code pattern', () => {
    const result = splitCompoundCommand("golangci-lint run && echo 'all clear'");
    assert.equal(result.lintCommand, 'golangci-lint run');
    assert.equal(result.jqSegment, '');
    assert.equal(result.echoSuffix, "&& echo 'all clear'");
  });

  it('plain golangci-lint run returns empty jqSegment and echoSuffix', () => {
    const result = splitCompoundCommand('golangci-lint run');
    assert.deepEqual(result, { lintCommand: 'golangci-lint run', jqSegment: '', echoSuffix: '' });
  });

  it('non-jq pipe stays in lintCommand: golangci-lint run | grep -v SA1019', () => {
    const result = splitCompoundCommand('golangci-lint run | grep -v SA1019');
    assert.equal(result.lintCommand, 'golangci-lint run | grep -v SA1019');
    assert.equal(result.jqSegment, '');
    assert.equal(result.echoSuffix, '');
  });

  it('quoted semicolons NOT split: golangci-lint run --args \'a;b\'', () => {
    const result = splitCompoundCommand("golangci-lint run --args 'a;b'");
    assert.equal(result.lintCommand, "golangci-lint run --args 'a;b'");
    assert.equal(result.jqSegment, '');
    assert.equal(result.echoSuffix, '');
  });

  it('multiple semicolons: echo \'a\'; golangci-lint run; echo \'b\'', () => {
    const result = splitCompoundCommand("echo 'a'; golangci-lint run; echo 'b'");
    assert.deepEqual(result, { lintCommand: 'golangci-lint run', jqSegment: '', echoSuffix: '' });
  });

  it('jq with second pipe strips after jq: golangci-lint run | jq \'.\' | grep something', () => {
    const result = splitCompoundCommand("golangci-lint run | jq '.' | grep something");
    assert.equal(result.lintCommand, 'golangci-lint run');
    assert.equal(result.jqSegment, "| jq '.'");
    assert.equal(result.echoSuffix, '');
  });

  it('falsy/empty input returns empty fields', () => {
    const result = splitCompoundCommand('');
    assert.deepEqual(result, { lintCommand: '', jqSegment: '', echoSuffix: '' });
  });

  it('null input returns empty string lintCommand', () => {
    const result = splitCompoundCommand(null);
    assert.deepEqual(result, { lintCommand: '', jqSegment: '', echoSuffix: '' });
  });

  it('undefined input returns empty string lintCommand', () => {
    const result = splitCompoundCommand(undefined);
    assert.deepEqual(result, { lintCommand: '', jqSegment: '', echoSuffix: '' });
  });

  it('no golangci-lint segment returns passthrough', () => {
    const result = splitCompoundCommand('echo hello');
    assert.deepEqual(result, { lintCommand: 'echo hello', jqSegment: '', echoSuffix: '' });
  });
});

// ─── injectJqFilter ──────────────────────────────────────────────────────

describe('injectJqFilter', () => {
  it('replaces trivial | jq with envelope filter', () => {
    const result = injectJqFilter('| jq');
    assert.equal(result, "| jq '{Issues: .Issues, Report: .Report}'");
  });

  it('replaces trivial | jq \'.\' with envelope filter', () => {
    const result = injectJqFilter("| jq '.'");
    assert.equal(result, "| jq '{Issues: .Issues, Report: .Report}'");
  });

  it('replaces trivial | jq "." with envelope filter', () => {
    const result = injectJqFilter('| jq "."');
    assert.equal(result, "| jq '{Issues: .Issues, Report: .Report}'");
  });

  it('replaces trivial | /usr/bin/jq \'.\' preserving path prefix', () => {
    const result = injectJqFilter("| /usr/bin/jq '.'");
    assert.equal(result, "| /usr/bin/jq '{Issues: .Issues, Report: .Report}'");
  });

  it('leaves non-trivial jq filter unchanged: | jq \'.Issues[]\'', () => {
    const result = injectJqFilter("| jq '.Issues[]'");
    assert.equal(result, "| jq '.Issues[]'");
  });

  it('leaves jq with -r flag unchanged: | jq -r \'.Issues[].Text\'', () => {
    const result = injectJqFilter("| jq -r '.Issues[].Text'");
    assert.equal(result, "| jq -r '.Issues[].Text'");
  });

  it('leaves jq with -c flag unchanged: | jq -c \'.Issues | length\'', () => {
    const result = injectJqFilter("| jq -c '.Issues | length'");
    assert.equal(result, "| jq -c '.Issues | length'");
  });

  it('empty string returns empty string', () => {
    const result = injectJqFilter('');
    assert.equal(result, '');
  });

  it('falsy input returns empty string', () => {
    const result = injectJqFilter(null);
    assert.equal(result, '');
  });
});
