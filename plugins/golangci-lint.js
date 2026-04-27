'use strict';
import {
  isGolangciLintCommand, extractInnerCommand, stripOutputFilters,
  injectJsonOutputFlag, parseDiagnostics, buildStrategyANudge,
  buildStrategyBNudge, truncateNudge,
  splitCompoundCommand, injectJqFilter
} from '../shared/nudge.js';

export const GolangciLintPlugin = async function ({ project: _project, client: _client, $: _$, directory: _directory, worktree: _worktree }) {
  return {
    'tool.execute.before': async function (input, output) {
      try {
        if (input.tool !== 'bash') return;
        let command = (output.args && output.args.command) || '';
        if (!isGolangciLintCommand(command)) return;
        const cdMatch = command.match(/^(cd\s+(?:"[^"]+"|'[^']+'|\S+)\s*(?:&&|;)\s*)/);
        const cdPrefix = cdMatch ? cdMatch[0] : '';
        command = extractInnerCommand(command);
        const parts = splitCompoundCommand(command);
        let lintCommand = stripOutputFilters(parts.lintCommand);
        lintCommand = injectJsonOutputFlag(lintCommand);
        const jqPart = injectJqFilter(parts.jqSegment);
        output.args.command = cdPrefix + lintCommand + jqPart + (parts.echoSuffix ? ' ' + parts.echoSuffix : '');
      } catch (_err) { /* never break tool execution */ }
    },
    'tool.execute.after': async function (input, output) {
      try {
        if (input.tool !== 'bash') return;
        const command = (input.args && input.args.command) || '';
        if (!isGolangciLintCommand(command)) return;
        let rawOutput = '';
        if (output) {
          rawOutput = output.output || output.stdout || output.result || '';
          if (typeof rawOutput !== 'string') rawOutput = rawOutput.stdout || rawOutput.output || rawOutput.result || '';
        }
        if (!rawOutput || typeof rawOutput !== 'string' || !rawOutput.trim()) return;
        const result = parseDiagnostics(rawOutput.trim());
        if (result.totalUnique === 0 && (!result.warnings || result.warnings.length === 0)) return;
        let nudge = result.totalUnique <= 10
          ? buildStrategyANudge(result.totalUnique, rawOutput.trim())
          : buildStrategyBNudge(result.totalUnique, result.linterCounts, rawOutput.trim());
        nudge = truncateNudge(nudge);
        if (result.warnings && result.warnings.length > 0) {
          const warnText = result.warnings.map(function(w) { return w.Text || w.text || String(w); }).join('; ');
          nudge += '\n\n⚠ Config warnings: ' + warnText;
        }
        if (typeof output.output === 'string') output.output = output.output + '\n\n' + nudge;
        else if (typeof output.result === 'string') output.result = output.result + '\n\n' + nudge;
        else if (typeof output.stdout === 'string') output.stdout = output.stdout + '\n\n' + nudge;
      } catch (_err) { /* never break tool execution */ }
    }
  };
};

export { isGolangciLintCommand, extractInnerCommand, stripOutputFilters, injectJsonOutputFlag, parseDiagnostics, splitCompoundCommand, injectJqFilter };
export default GolangciLintPlugin;
