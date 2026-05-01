'use strict';

const { describe, it } = require('node:test');
const assert = require('node:assert/strict');
const { spawn } = require('child_process');
const path = require('path');

const HOOK_PATH = path.join(__dirname, 'golangci-lint-pre.js');

/**
 * Spawns the PreToolUse hook as a child process, pipes input JSON to stdin,
 * captures stdout + stderr, and returns { exitCode, stdout, stderr }.
 */
function runHook(inputJson) {
  return new Promise((resolve, reject) => {
    const child = spawn('node', [HOOK_PATH], { stdio: ['pipe', 'pipe', 'pipe'] });
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

// ─── Passthrough cases (no modification) ────────────────────────────────

describe('PreToolUse hook — passthrough cases', () => {
  it('non-Bash tool_name → exits 0 with no output', async () => {
    const result = await runHook({ tool_name: 'Read', tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('Bash with non-golangci-lint command → exits 0 with no output', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'go test ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });

  it('malformed stdin JSON → exits 0 gracefully', async () => {
    const result = await new Promise((resolve, reject) => {
      const child = spawn('node', [HOOK_PATH], { stdio: ['pipe', 'pipe', 'pipe'] });
      let stdout = '';
      let stderr = '';
      child.stdout.on('data', (chunk) => { stdout += chunk; });
      child.stderr.on('data', (chunk) => { stderr += chunk; });
      child.on('error', reject);
      child.on('close', (code) => {
        resolve({ exitCode: code, stdout: stdout.trim(), stderr: stderr.trim() });
      });
      child.stdin.write('not-valid-json{{{');
      child.stdin.end();
    });
    assert.equal(result.exitCode, 0);
  });

  it('empty stdin → exits 0 gracefully', async () => {
    const result = await new Promise((resolve, reject) => {
      const child = spawn('node', [HOOK_PATH], { stdio: ['pipe', 'pipe', 'pipe'] });
      let stdout = '';
      let stderr = '';
      child.stdout.on('data', (chunk) => { stdout += chunk; });
      child.stderr.on('data', (chunk) => { stderr += chunk; });
      child.on('error', reject);
      child.on('close', (code) => {
        resolve({ exitCode: code, stdout: stdout.trim(), stderr: stderr.trim() });
      });
      child.stdin.end();
    });
    assert.equal(result.exitCode, 0);
  });

  it('tool_name is missing → exits 0 gracefully', async () => {
    const result = await runHook({ tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    assert.equal(result.stdout, '');
  });
});

// ─── Command modification cases ─────────────────────────────────────────

describe('PreToolUse hook — command modification', () => {
  it('simple "golangci-lint run ./..." → injects --output.json.path stdout', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    assert.equal(json.hookSpecificOutput.permissionDecision, 'allow');
    assert.ok(json.hookSpecificOutput.updatedInput.command.indexOf('--output.json.path stdout') !== -1);
  });

  it('"golangci-lint run --out-format json ./..." → strips --out-format, injects --output.json.path stdout', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'golangci-lint run --out-format json ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.hookSpecificOutput.updatedInput.command;
    assert.ok(modified.indexOf('--out-format') === -1, 'should not contain --out-format');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"cd /dir && golangci-lint run ./..." → preserves cd prefix, injects JSON flag', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'cd /dir && golangci-lint run ./...' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.hookSpecificOutput.updatedInput.command;
    assert.ok(modified.indexOf('cd /dir &&') !== -1, 'should preserve cd prefix');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"golangci-lint run ./... 2>&1 | grep -v something" → strips pipe and redirect, injects JSON flag', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'golangci-lint run ./... 2>&1 | grep -v something' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.hookSpecificOutput.updatedInput.command;
    assert.ok(modified.indexOf('2>&1') === -1, 'should strip 2>&1');
    assert.ok(modified.indexOf('grep') === -1, 'should strip pipe segment');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"golangci-lint run ./...; echo done" → preserves echo suffix, injects JSON flag', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: 'golangci-lint run ./...; echo done' } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.hookSpecificOutput.updatedInput.command;
    assert.ok(modified.indexOf('echo done') !== -1, 'should preserve echo suffix');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });

  it('"golangci-lint run ./... | jq \'.Issues\'" → injects structured jq filter', async () => {
    const result = await runHook({ tool_name: 'Bash', tool_input: { command: "golangci-lint run ./... | jq '.Issues'" } });
    assert.equal(result.exitCode, 0);
    const json = JSON.parse(result.stdout);
    const modified = json.hookSpecificOutput.updatedInput.command;
    // Non-trivial jq filter (not just '.' or empty) should be left unchanged
    assert.ok(modified.indexOf("jq '.Issues'") !== -1, 'should preserve non-trivial jq filter');
    assert.ok(modified.indexOf('--output.json.path stdout') !== -1, 'should contain json flag');
  });
});
