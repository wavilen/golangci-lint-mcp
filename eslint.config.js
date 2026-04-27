'use strict';

const js = require('@eslint/js');
const globals = require('globals');

const extraRules = {
  'no-var': 'error',
  'prefer-const': 'error',
  'no-throw-literal': 'error',
  'eqeqeq': 'error',
  'curly': ['error', 'multi-line'],
  'no-eval': 'error',
  'no-implied-eval': 'error',
  'no-self-compare': 'error',
  'no-shadow-restricted-names': 'error',
  'no-template-curly-in-string': 'error',
  'no-async-promise-executor': 'error',
  'no-constant-binary-expression': 'error',
  'no-constructor-return': 'error',
  'no-duplicate-imports': 'error',
  'no-useless-return': 'error',
  'no-unsafe-negation': 'error',
  'no-unused-vars': ['error', { argsIgnorePattern: '^_', caughtErrorsIgnorePattern: '^_' }]
};

module.exports = [
  // Global ignores
  {
    ignores: ['.opencode/**', 'node_modules/**', '.planning/**', 'tmp/**', 'out/**', 'graphify-out/**']
  },
  // CJS files (shared/, hooks/, bin/)
  {
    files: ['shared/**/*.js', 'hooks/**/*.js', 'bin/**/*.js'],
    languageOptions: {
      sourceType: 'script',
      globals: {
        ...globals.node
      }
    },
    ...js.configs.recommended,
    rules: {
      ...js.configs.recommended.rules,
      ...extraRules
    }
  },
  // ESM files (plugins/)
  {
    files: ['plugins/**/*.js'],
    languageOptions: {
      sourceType: 'module',
      globals: {
        ...globals.node
      }
    },
    ...js.configs.recommended,
    rules: {
      ...js.configs.recommended.rules,
      ...extraRules
    }
  }
];
