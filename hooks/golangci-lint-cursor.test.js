'use strict';

const { describe, it } = require('node:test');
const assert = require('node:assert/strict');
const { spawn } = require('child_process');
const path = require('path');

const PRE_HOOK_PATH = path.join(__dirname, 'golangci-lint-cursor-pre.js');
const POST_HOOK_PATH = path.join(__dirname, 'golangci-lint-cursor-post.js');

/**
 * Spawns a hook as a child process, pipes input JSON to stdin,
 * captures stdout + stderr, and returns { exitCode, stdout, stderr }.
 */
function runHook(hookPath, inputJson) {
  return new Promise((resolve, reject) => {
    const child = spawn('node', [hookPath], { stdio: ['pipe', 'pipe', 'pipe'] });
    let stdout = '';
    let stderr = '';
    child.stdout.on('data', (chunk) => { stdout += chunk; });
    child.stderr.on('data', (chunk) => { stderr += chunk; });
    child.on('error', reject);
    child.on('close', (code) => {
      resolve({ exitCode: code, stdout: stdout.trim(), stderr: stderr.trim() });
    });
    if (inputJson !== undefined && inputJson !== null) {
      child.stdin.write(JSON.stringify(inputJson));
    }
    child.stdin.end();
  });
}

/**
 * Spawns a hook with raw (non-JSON) stdin input.
 */
function runHookRaw(hookPath, rawInput) {
  return new Promise((resolve, reject) => {
    const child = spawn('node', [hookPath], { stdio: ['pipe', 'pipe', 'pipe'] });
    let stdout = '';
    let stderr = '';
    child.stdout.on('data', (chunk) => { stdout += chunk; });
    child.stderr.on('data', (chunk) => { stderr += chunk; });
    child.on('error', reject);
    child.on('close', (code) => {
      resolve({ exitCode: code, stdout: stdout.trim(), stderr: stderr.trim() });
    });
    if (rawInput !== undefined && rawInput !== null) {
      child.stdin.write(rawInput);
    }
    child.stdin.end();
  });
}

/**
 * Builds a golangci-lint JSON output object with the given issues.
 */
function makeLintJson(issues) {
  return JSON.stringify({
    Issues: issues.map(function (iss) {
      return {
        FromLinter: iss.linter || 'errcheck',
        Text: iss.text || 'test issue',
        Pos: { Filename: iss.file || 'test.go', Line: iss.line || 1 },
        RuleName: ''
      };
    }),
    Report: {}
  });
}

// ─── PreToolUse hook — passthrough cases ──────────────────────────────────

describe('Cursor PreToolUse hook — passthrough cases', () => {
  it('non-Shell tool_name → exits 0 with no output', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Read', tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('Shell with non-golangci-lint command → exits 0 with no output', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'go test ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('malformed stdin JSON → exits 0 gracefully', async () => {
    const result = await runHookRaw(PRE_HOOK_PATH, 'not-valid-json{{{');
    assert.equal(result.exitCode, 0);
  });

  it('empty stdin → exits 0 gracefully', async () => {
    const result = await runHookRaw(PRE_HOOK_PATH, '');
    assert.equal(result.exitCode, 0);
  });

  it('missing tool_name → exits 0 gracefully', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });
});

// ─── PreToolUse hook — command modification ────────────────────────────────

describe('Cursor PreToolUse hook — command modification', () => {
  it('"golangci-lint run ./..." → outputs JSON with decision "allow" and updated_input containing --output.json.path stdout', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    assert.equal(json.decision, 'allow');
    assert.ok(json.updated_input.command.indexOf('--output.json.path stdout') !== -1);
  });

  it('"golangci-lint run --out-format json ./..." → strips --out-format, injects --output.json.path stdout', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run --out-format json ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.updated_input.command;
    assert.ok(modified.indexOf('--out-format') === -1, 'should not contain --out-format');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"cd /dir && golangci-lint run ./..." → preserves cd prefix, injects JSON flag', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'cd /dir && golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.updated_input.command;
    assert.ok(modified.indexOf('cd /dir &&') !== -1, 'should preserve cd prefix');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"golangci-lint run ./...; echo done" → preserves echo suffix, injects JSON flag', async () => {
    const result = await runHook(PRE_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...; echo done' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.updated_input.command;
    assert.ok(modified.indexOf('echo done') !== -1, 'should preserve echo suffix');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });
});

// ─── PostToolUse hook — passthrough cases ──────────────────────────────────

describe('Cursor PostToolUse hook — passthrough cases', () => {
  it('non-Shell tool_name → exits 0 with no output', async () => {
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Read', tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('Shell with non-golangci-lint command → exits 0 with no output', async () => {
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'go test ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('Shell with golangci-lint command but empty tool_output → exits 0 with no output', async () => {
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' }, tool_output: '' });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('Shell with golangci-lint command but zero diagnostics → exits 0 with no output', async () => {
    const emptyOutput = makeLintJson([]);
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' }, tool_output: emptyOutput });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });
});

// ─── PostToolUse hook — nudge injection ────────────────────────────────────

describe('Cursor PostToolUse hook — nudge injection', () => {
  it('Shell with golangci-lint JSON output containing issues → outputs JSON with additional_context containing "golangci-lint" and "MCP"', async () => {
    const lintOutput = makeLintJson([
      { linter: 'errcheck', text: ' Error return value is not checked' }
    ]);
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' }, tool_output: lintOutput });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    assert.ok(json.additional_context, 'should have additional_context field');
    assert.ok(json.additional_context.indexOf('golangci-lint') !== -1, 'nudge should mention golangci-lint');
  });

  it('tool_output as JSON-stringified wrapper → parses to extract stdout, outputs nudge', async () => {
    const lintOutput = makeLintJson([
      { linter: 'errcheck', text: ' Error return value is not checked' }
    ]);
    const wrappedOutput = JSON.stringify({ exitCode: 0, stdout: lintOutput });
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' }, tool_output: wrappedOutput });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    assert.ok(json.additional_context, 'should have additional_context field');
    assert.ok(json.additional_context.indexOf('golangci-lint') !== -1, 'nudge should mention golangci-lint');
  });

  it('tool_output as raw golangci-lint JSON (not wrapped) → parses directly, outputs nudge', async () => {
    const lintOutput = makeLintJson([
      { linter: 'errcheck', text: ' Error return value is not checked' },
      { linter: 'staticcheck', text: 'SA1000: invalid argument' }
    ]);
    const result = await runHook(POST_HOOK_PATH, { tool_name: 'Shell', tool_input: { command: 'golangci-lint run ./...' }, tool_output: lintOutput });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    assert.ok(json.additional_context, 'should have additional_context field');
    assert.ok(json.additional_context.indexOf('golangci-lint') !== -1, 'nudge should mention golangci-lint');
  });
});
