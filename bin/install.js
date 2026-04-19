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
