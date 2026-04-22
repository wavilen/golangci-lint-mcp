'use strict';

import { describe, it } from 'node:test';
import assert from 'node:assert/strict';
import {
  isGolangciLintCommand,
  extractInnerCommand,
  stripOutputFilters,
  injectJsonOutputFlag,
  parseDiagnostics
} from './golangci-lint.js';

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
    var result = injectJsonOutputFlag('golangci-lint run --out-format tab');
    assert.ok(result.indexOf('--out-format') === -1, 'should not contain --out-format');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('strips --color flag and injects json flag', () => {
    var result = injectJsonOutputFlag('golangci-lint run --color always');
    assert.ok(result.indexOf('--color') === -1, 'should not contain --color');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('replaces existing --output.json.path with new value', () => {
    var result = injectJsonOutputFlag('golangci-lint run --output.json.path stderr');
    assert.ok(result.indexOf('stderr') === -1, 'should not contain old value');
    assert.ok(result.indexOf('--output.json.path stdout') !== -1, 'should contain new value');
  });
});

// ─── parseDiagnostics regression tests ────────────────────────────────────

describe('parseDiagnostics', () => {
  it('parses single diagnostic', () => {
    var input = JSON.stringify({ FromLinter: 'errcheck', Text: 'error return is unchecked' });
    var result = parseDiagnostics(input);
    assert.equal(result.totalUnique, 1);
    assert.equal(result.linterCounts.errcheck, 1);
  });

  it('parses multiple linters', () => {
    var line1 = JSON.stringify({ FromLinter: 'errcheck', Text: 'unchecked' });
    var line2 = JSON.stringify({ FromLinter: 'govet', Text: 'lostcancel: cancel function not used' });
    var line3 = JSON.stringify({ FromLinter: 'errcheck', Text: 'another unchecked' });
    var result = parseDiagnostics(line1 + '\n' + line2 + '\n' + line3);
    assert.equal(result.totalUnique, 2);
    assert.equal(result.linterCounts.errcheck, 1);
    assert.equal(result.linterCounts.govet, 1);
  });

  it('handles empty input', () => {
    var result = parseDiagnostics('');
    assert.equal(result.totalUnique, 0);
    assert.deepEqual(result.linterCounts, {});
  });

  it('handles invalid JSON lines', () => {
    var result = parseDiagnostics('not json\nalso not json');
    assert.equal(result.totalUnique, 0);
  });

  it('extracts rule from compound linter text', () => {
    var line = JSON.stringify({ FromLinter: 'govet', Text: 'lostcancel: cancel function is not used on deferral path' });
    var result = parseDiagnostics(line);
    assert.equal(result.totalUnique, 1);
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
    var inner = extractInnerCommand('bash -c "golangci-lint run | grep -v SA1019"');
    var stripped = stripOutputFilters(inner);
    assert.equal(stripped, 'golangci-lint run');
  });

  it('after extractInnerCommand: /bin/bash -c wrapped with redirect + grep -v', () => {
    var inner = extractInnerCommand("/bin/bash -c 'golangci-lint run ./... 2>&1 | grep -v \"SA1019\"'");
    var stripped = stripOutputFilters(inner);
    assert.equal(stripped, 'golangci-lint run ./...');
  });
});

// ─── full pipeline: extractInnerCommand → stripOutputFilters → injectJsonOutputFlag ──

describe('full pipeline: shell-wrapped + piped commands', () => {
  function processCommand(raw) {
    var cdMatch = raw.match(/^(cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*)/);
    var cdPrefix = cdMatch ? cdMatch[0] : '';
    var inner = extractInnerCommand(raw);
    var stripped = stripOutputFilters(inner);
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
