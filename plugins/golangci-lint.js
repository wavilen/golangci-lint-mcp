'use strict';
import {
  isGolangciLintCommand, extractInnerCommand, stripOutputFilters,
  injectJsonOutputFlag, parseDiagnostics, buildStrategyANudge,
  buildStrategyBNudge, truncateNudge, COMPOUND_LINTERS,
  OUTPUT_FLAG_PATTERNS_VALUE, OUTPUT_FLAG_PATTERNS_BOOL, NON_RUN_SUBCOMMANDS
} from '../shared/nudge.js';

export const GolangciLintPlugin = async function ({ project, client, $, directory, worktree }) {
  return {
    'tool.execute.before': async function (input, output) {
      try {
        if (input.tool !== 'bash') return;
        var command = (output.args && output.args.command) || '';
        if (!isGolangciLintCommand(command)) return;
        var cdMatch = command.match(/^(cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*)/);
        var cdPrefix = cdMatch ? cdMatch[0] : '';
        command = extractInnerCommand(command);
        command = stripOutputFilters(command);
        output.args.command = cdPrefix + injectJsonOutputFlag(command);
      } catch (err) { /* never break tool execution */ }
    },
    'tool.execute.after': async function (input, output) {
      try {
        if (input.tool !== 'bash') return;
        var command = (input.args && input.args.command) || '';
        if (!isGolangciLintCommand(command)) return;
        var rawOutput = '';
        if (output) {
          rawOutput = output.output || output.stdout || output.result || '';
          if (typeof rawOutput !== 'string') rawOutput = rawOutput.stdout || rawOutput.output || rawOutput.result || '';
        }
        if (!rawOutput || typeof rawOutput !== 'string' || !rawOutput.trim()) return;
        var result = parseDiagnostics(rawOutput.trim());
        if (result.totalUnique === 0) return;
        var nudge = result.totalUnique <= 10
          ? buildStrategyANudge(result.totalUnique, rawOutput.trim())
          : buildStrategyBNudge(result.totalUnique, result.linterCounts, rawOutput.trim());
        nudge = truncateNudge(nudge);
        if (typeof output.output === 'string') output.output = output.output + '\n\n' + nudge;
        else if (typeof output.result === 'string') output.result = output.result + '\n\n' + nudge;
        else if (typeof output.stdout === 'string') output.stdout = output.stdout + '\n\n' + nudge;
      } catch (err) { /* never break tool execution */ }
    }
  };
};

export { isGolangciLintCommand, extractInnerCommand, stripOutputFilters, injectJsonOutputFlag, parseDiagnostics };
export default GolangciLintPlugin;
