#!/usr/bin/env node
'use strict';

const fs = require('fs');
const path = require('path');
const os = require('os');

const src = path.join(__dirname, '..', 'skills', 'golangci-lint-guide', 'SKILL.md');
const destDir = path.join(os.homedir(), '.agents', 'skills', 'golangci-lint-guide');
const dest = path.join(destDir, 'SKILL.md');

if (!fs.existsSync(src)) {
  console.error('Error: SKILL.md not found at ' + src);
  console.error('The npm package may be corrupted. Try reinstalling.');
  process.exit(1);
}

fs.mkdirSync(destDir, { recursive: true });
fs.copyFileSync(src, dest);
console.log('\u2713 golangci-lint-guide skill installed to ' + dest);

// Install custom command file to opencode and crush commands directories
const cmdSrc = path.join(__dirname, '..', 'commands', 'golangci-lint.md');

if (fs.existsSync(cmdSrc)) {
  const cmdTargets = [
    { name: 'opencode', dir: path.join(os.homedir(), '.config', 'opencode', 'commands') },
    { name: 'crush',    dir: path.join(os.homedir(), '.config', 'crush', 'commands') }
  ];

  for (let i = 0; i < cmdTargets.length; i++) {
    const t = cmdTargets[i];
    try {
      fs.mkdirSync(t.dir, { recursive: true });
      const cmdDest = path.join(t.dir, 'golangci-lint.md');
      fs.copyFileSync(cmdSrc, cmdDest);
      console.log('\u2713 golangci-lint command installed to ' + cmdDest);
    } catch (err) {
      console.warn('Warning: could not install command to ' + t.name + ' (' + err.message + ')');
    }
  }
} else {
  console.warn('Warning: commands/golangci-lint.md not found at ' + cmdSrc + ' — skipping command installation');
}

// --- Platform rules installation ---
const rulesDir = path.join(__dirname, '..', 'rules');

const rulesSources = [
  { platform: 'claude', src: 'claude-code.md', destDir: '.claude/rules', destFile: 'golangci-lint.md' },
  { platform: 'cursor', src: 'cursor.mdc', destDir: '.cursor/rules', destFile: 'golangci-lint.mdc' },
  { platform: 'opencode', src: 'opencode.md', destDir: '.opencode/rules', destFile: 'golangci-lint.md' }
];

let platforms;
let pluginScopeFlag;

if (!fs.existsSync(rulesDir)) {
  console.warn('Warning: rules/ directory not found at ' + rulesDir + ' — skipping rules installation');
} else {
  // Parse --platforms= CLI flag
  let platformFlag = null;
  // Parse --plugin-scope= CLI flag (for opencode plugin install scope)
  pluginScopeFlag = null;
  for (let a = 2; a < process.argv.length; a++) {
    if (process.argv[a].startsWith('--platforms=')) {
      platformFlag = process.argv[a].split('=')[1].split(',');
    }
    if (process.argv[a].startsWith('--plugin-scope=')) {
      pluginScopeFlag = process.argv[a].split('=')[1];
      if (['project', 'user', 'both'].indexOf(pluginScopeFlag) === -1) {
        console.warn('Warning: --plugin-scope must be project, user, or both. Ignoring.');
        pluginScopeFlag = null;
      }
    }
  }

  // Resolve target platforms
  if (platformFlag) {
    platforms = platformFlag;
  } else {
    // Auto-detect installed platforms
    const cwd = process.cwd();
    const home = os.homedir();
    const detected = [];

    for (let i = 0; i < rulesSources.length; i++) {
      const p = rulesSources[i];
      const projectDir = path.join(cwd, p.destDir.split('/')[0]);
      let userDir = null;
      if (p.platform === 'claude') { userDir = path.join(home, '.claude'); }
      else if (p.platform === 'cursor') { userDir = path.join(home, '.cursor'); }
      else if (p.platform === 'opencode') { userDir = path.join(home, '.config', 'opencode'); }

      if (fs.existsSync(projectDir) || (userDir && fs.existsSync(userDir))) {
        detected.push(p.platform);
      }
    }

    platforms = detected;
  }

  if (platforms.length === 0) {
    console.log('No AI coding platforms detected. Skipping rules installation.');
    console.log('Run with --platforms=claude,cursor,opencode to install for specific platforms.');
  } else {
    console.log('Detected platforms: ' + platforms.join(', '));

    const installed = [];
    const skipped = [];

    for (let i = 0; i < rulesSources.length; i++) {
      const r = rulesSources[i];
      if (platforms.indexOf(r.platform) === -1) {
        skipped.push(r.platform);
        continue;
      }

      const ruleSrc = path.join(rulesDir, r.src);
      if (!fs.existsSync(ruleSrc)) {
        console.warn('Warning: rules/' + r.src + ' not found — skipping ' + r.platform);
        skipped.push(r.platform);
        continue;
      }

      const destPath = path.join(process.cwd(), r.destDir);
      try {
        fs.mkdirSync(destPath, { recursive: true });
        fs.copyFileSync(ruleSrc, path.join(destPath, r.destFile));
        console.log('\u2713 ' + r.platform + ' rules installed to ' + r.destDir + '/' + r.destFile);
        installed.push(r.platform);
      } catch (err) {
        console.warn('Warning: could not install ' + r.platform + ' rules (' + err.message + ')');
        skipped.push(r.platform);
      }
    }

    if (installed.length > 0) {
      console.log('Installed rules for: ' + installed.join(', '));
    }
    if (skipped.length > 0) {
      console.log('Skipped: ' + skipped.join(', '));
    }
  }
}

// --- Claude Code hook + config installation ---
// Only runs when 'claude' is in the platforms list
if (platforms && platforms.indexOf('claude') !== -1) {
  const projectDir = process.cwd();

  // 1. Copy hook script to .claude/hooks/
  try {
    const hookSrc = path.join(__dirname, '..', 'hooks', 'golangci-lint-post.js');
    if (fs.existsSync(hookSrc)) {
      const hookDestDir = path.join(projectDir, '.claude', 'hooks');
      fs.mkdirSync(hookDestDir, { recursive: true });
      fs.copyFileSync(hookSrc, path.join(hookDestDir, 'golangci-lint-post.js'));
      console.log('\u2713 Claude Code hook installed to .claude/hooks/golangci-lint-post.js');
    } else {
      console.warn('Warning: hooks/golangci-lint-post.js not found — skipping hook installation');
    }
  } catch (err) {
    console.warn('Warning: could not install Claude Code hook (' + err.message + ')');
  }

  // 1.5. Copy shared nudge module to .claude/shared/
  try {
    const sharedSrc = path.join(__dirname, '..', 'shared', 'nudge.js');
    if (fs.existsSync(sharedSrc)) {
      const sharedDestDir = path.join(projectDir, '.claude', 'shared');
      fs.mkdirSync(sharedDestDir, { recursive: true });
      fs.copyFileSync(sharedSrc, path.join(sharedDestDir, 'nudge.js'));
      console.log('\u2713 Shared nudge module installed to .claude/shared/nudge.js');
    }
  } catch (err) {
    console.warn('Warning: could not install shared nudge module (' + err.message + ')');
  }

  // 2. Merge hooks block into .claude/settings.json
  try {
    const settingsDir = path.join(projectDir, '.claude');
    fs.mkdirSync(settingsDir, { recursive: true });
    const settingsPath = path.join(settingsDir, 'settings.json');

    let settings = {};
    if (fs.existsSync(settingsPath)) {
      settings = JSON.parse(fs.readFileSync(settingsPath, 'utf8'));
    }

    const hookConfig = {
      type: 'command',
      'if': 'Bash(golangci-lint*)',
      command: '"$CLAUDE_PROJECT_DIR"/.claude/hooks/golangci-lint-post.js',
      timeout: 30
    };

    if (!settings.hooks) { settings.hooks = {}; }
    if (!settings.hooks.PostToolUse) {
      settings.hooks.PostToolUse = [{ matcher: 'Bash', hooks: [hookConfig] }];
    } else {
      // Find existing Bash matcher entry
      let bashEntry = null;
      for (let i = 0; i < settings.hooks.PostToolUse.length; i++) {
        if (settings.hooks.PostToolUse[i].matcher === 'Bash') {
          bashEntry = settings.hooks.PostToolUse[i];
          break;
        }
      }
      if (bashEntry) {
        // Update existing entry — replace or add our hook
        if (!bashEntry.hooks) { bashEntry.hooks = []; }
        let found = false;
        for (let j = 0; j < bashEntry.hooks.length; j++) {
          if (bashEntry.hooks[j].command && bashEntry.hooks[j].command.indexOf('golangci-lint-post.js') !== -1) {
            bashEntry.hooks[j] = hookConfig;
            found = true;
            break;
          }
        }
        if (!found) { bashEntry.hooks.push(hookConfig); }
      } else {
        settings.hooks.PostToolUse.push({ matcher: 'Bash', hooks: [hookConfig] });
      }
    }

    fs.writeFileSync(settingsPath, JSON.stringify(settings, null, 2) + '\n');
    console.log('\u2713 Claude Code hooks configured in .claude/settings.json');
  } catch (err) {
    console.warn('Warning: could not configure Claude Code hooks (' + err.message + ')');
  }

  // 3. Merge MCP server into .mcp.json
  try {
    const mcpPath = path.join(projectDir, '.mcp.json');
    let mcpConfig = {};
    if (fs.existsSync(mcpPath)) {
      mcpConfig = JSON.parse(fs.readFileSync(mcpPath, 'utf8'));
    }

    if (!mcpConfig.mcpServers) { mcpConfig.mcpServers = {}; }
    mcpConfig.mcpServers['golangci-lint-mcp'] = { command: 'golangci-lint-mcp', args: [] };

    fs.writeFileSync(mcpPath, JSON.stringify(mcpConfig, null, 2) + '\n');
    console.log('\u2713 MCP server configured in .mcp.json');
  } catch (err) {
    console.warn('Warning: could not configure MCP server (' + err.message + ')');
  }
}

// --- OpenCode plugin + MCP config installation ---
// Only runs when 'opencode' is in the platforms list
if (platforms && platforms.indexOf('opencode') !== -1) {
  const readline = require('readline');

  // 1. Resolve plugin scope (project, user, or both)
  let scopes;
  if (pluginScopeFlag === 'both') {
    scopes = ['project', 'user'];
  } else if (pluginScopeFlag === 'project' || pluginScopeFlag === 'user') {
    scopes = [pluginScopeFlag];
  } else {
    // Auto-detect and prompt interactively
    const projectOpenCodeDir = path.join(process.cwd(), '.opencode');
    const userOpenCodeDir = path.join(os.homedir(), '.config', 'opencode');
    const projectExists = fs.existsSync(projectOpenCodeDir);
    const userExists = fs.existsSync(userOpenCodeDir);

    if (projectExists && userExists) {
      // Both exist — ask user which scope
      const rl = readline.createInterface({ input: process.stdin, output: process.stdout });
      console.log('\nBoth project-level and user-level opencode directories detected.');
      console.log('Install plugin to:');
      console.log('  1) Project (.opencode/plugins/)');
      console.log('  2) User (~/.config/opencode/plugins/)');
      console.log('  3) Both');
      rl.question('Enter choice [1-3]: ', function(answer) {
        rl.close();
        const choice = (answer || '').trim();
        if (choice === '2') {
          scopes = ['user'];
        } else if (choice === '3') {
          scopes = ['project', 'user'];
        } else {
          scopes = ['project'];
        }
        runOpenCodeInstall(scopes);
      });
      // Prevent rest of script from running synchronously
      scopes = null;
    } else if (userExists && !projectExists) {
      scopes = ['user'];
    } else {
      // Default to project scope (create .opencode/plugins/)
      scopes = ['project'];
    }
  }

  if (scopes) {
    runOpenCodeInstall(scopes);
  }
}

function runOpenCodeInstall(scopes) {
  // 2. Copy plugin file to each scope
  const pluginSrc = path.join(__dirname, '..', 'plugins', 'golangci-lint.js');

  if (!fs.existsSync(pluginSrc)) {
    console.warn('Warning: plugins/golangci-lint.js not found — skipping plugin installation');
  } else {
    for (let i = 0; i < scopes.length; i++) {
      const scope = scopes[i];
      try {
        let pluginDestDir;
        if (scope === 'project') {
          pluginDestDir = path.join(process.cwd(), '.opencode', 'plugins');
        } else {
          pluginDestDir = path.join(os.homedir(), '.config', 'opencode', 'plugins');
        }
        fs.mkdirSync(pluginDestDir, { recursive: true });
        const pluginDest = path.join(pluginDestDir, 'golangci-lint.js');
        fs.copyFileSync(pluginSrc, pluginDest);
        console.log('\u2713 opencode plugin installed to ' + pluginDest);
      } catch (err) {
        console.warn('Warning: could not install opencode plugin to ' + scope + ' scope (' + err.message + ')');
      }
    }
  }

  // 2.5. Copy shared nudge module to each scope's shared/ directory
  const sharedSrc = path.join(__dirname, '..', 'shared', 'nudge.js');
  if (fs.existsSync(sharedSrc)) {
    for (let s = 0; s < scopes.length; s++) {
      try {
        let sharedDestDir;
        if (scopes[s] === 'project') {
          sharedDestDir = path.join(process.cwd(), '.opencode', 'shared');
        } else {
          sharedDestDir = path.join(os.homedir(), '.config', 'opencode', 'shared');
        }
        fs.mkdirSync(sharedDestDir, { recursive: true });
        fs.copyFileSync(sharedSrc, path.join(sharedDestDir, 'nudge.js'));
        console.log('\u2713 Shared nudge module installed to ' + sharedDestDir + '/nudge.js');
      } catch (err) {
        console.warn('Warning: could not install shared nudge module to ' + scopes[s] + ' scope (' + err.message + ')');
      }
    }
  }

  // 3. Merge MCP server config into opencode.json for each scope
  for (let i = 0; i < scopes.length; i++) {
    const scope = scopes[i];
    try {
      let configPath;
      if (scope === 'project') {
        configPath = path.join(process.cwd(), 'opencode.json');
      } else {
        configPath = path.join(os.homedir(), '.config', 'opencode', 'opencode.json');
      }

      // Backup user-level config before modifying
      if (scope === 'user' && fs.existsSync(configPath)) {
        const backupPath = configPath + '.bak.' + Date.now();
        fs.copyFileSync(configPath, backupPath);
        console.log('\u2713 Backup created: ' + backupPath);
      }

      // Read existing config or start fresh
      let openCodeConfig = {};
      if (fs.existsSync(configPath)) {
        openCodeConfig = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      }

      // Merge MCP config — only add/overwrite mcp.golangci-lint key
      if (!openCodeConfig.mcp) { openCodeConfig.mcp = {}; }
      openCodeConfig.mcp['golangci-lint'] = {
        "type": "local",
        "command": ["golangci-lint-mcp"],
        "enabled": true
      };

      // Ensure parent directory exists
      const configDir = path.dirname(configPath);
      fs.mkdirSync(configDir, { recursive: true });

      fs.writeFileSync(configPath, JSON.stringify(openCodeConfig, null, 2) + '\n');
      console.log('\u2713 MCP server configured in ' + configPath);
    } catch (err) {
      console.warn('Warning: could not configure MCP server for ' + scope + ' scope (' + err.message + ')');
    }
  }

  // 4. Binary prerequisite checks (informational warnings only)
  const execSync = require('child_process').execSync;

  try {
    execSync('which golangci-lint-mcp', { stdio: 'pipe' });
    console.log('\u2713 golangci-lint-mcp binary found on PATH');
  } catch (_e) {
    console.warn('\u26A0 golangci-lint-mcp not found on PATH.');
    console.warn('  Install with: go install github.com/wavilen/golangci-lint-mcp@latest');
  }

  try {
    execSync('which golangci-lint', { stdio: 'pipe' });
    console.log('\u2713 golangci-lint binary found on PATH');
  } catch (_e) {
    console.warn('\u26A0 golangci-lint not found on PATH.');
    console.warn('  Install with: https://golangci-lint.run/usage/install/');
  }
}
