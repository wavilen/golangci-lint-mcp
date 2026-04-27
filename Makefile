.PHONY: build install install-skill install-commands install-rules install-hook install-claude-shared install-all install-agent test clean npm-pack npm-publish update-golden-config crosscheck crosscheck-clean install-plugin verify-plugin verify-shared sync-deployed sync-version lint-js

BINARY := golangci-lint-mcp
VERSION := $(shell git describe --tags --always 2>/dev/null | sed 's/^v//')
SKILL_SRC := skills/golangci-lint-guide/SKILL.md
SKILL_DEST := $(HOME)/.agents/skills/golangci-lint-guide

build:
	go build -ldflags "-X github.com/wavilen/golangci-lint-mcp/internal/version.Server=$(VERSION)" -o $(BINARY) .

install:
	go install -ldflags "-X github.com/wavilen/golangci-lint-mcp/internal/version.Server=$(VERSION)" .

install-skill: ## Copy the golangci-lint-guide skill to ~/.agents/skills/
	@mkdir -p $(SKILL_DEST)
	@cp $(SKILL_SRC) $(SKILL_DEST)/SKILL.md
	@echo "Skill installed to $(SKILL_DEST)/SKILL.md"

test:
	go test ./...

clean:
	rm -f $(BINARY)

npm-pack: ## Preview npm package contents (dry run)
	npm pack --dry-run 2>&1 | head -30

npm-publish: ## Publish to npm (requires npm login first)
	npm publish --access public

update-golden-config: ## Fetch latest maratori golden config and vendor it
	@LATEST_TAG=$$(curl -sL https://api.github.com/repos/maratori/golangci-lint-config/tags | grep -m1 '"name"' | sed 's/.*"name": "\([^"]*\)".*/\1/') && \
	echo "Fetching maratori/golangci-lint-config $${LATEST_TAG}" && \
	curl -sL "https://raw.githubusercontent.com/maratori/golangci-lint-config/refs/tags/$${LATEST_TAG}/.golangci.yml" -o golden-config/.golangci.yml && \
	TODAY=$$(date -u +"%Y-%m-%d") && \
	printf '# Vendored from github.com/maratori/golangci-lint-config\n# Tag: %s\n# Updated: %s\n# Update: make update-golden-config\n\n' "$${LATEST_TAG}" "$${TODAY}" > golden-config/.golangci.yml.tmp && \
	cat golden-config/.golangci.yml >> golden-config/.golangci.yml.tmp && \
	mv golden-config/.golangci.yml.tmp golden-config/.golangci.yml && \
	sed -i '/^      local-prefixes:$$/{N;s/^      local-prefixes:\n        -.*/      local-prefixes: []/}' golden-config/.golangci.yml && \
	sed -i '/^        "non-main files":$$/,/^              desc: Use log\/slog instead/c\        # \"non-main files\" rule disabled — extracted examples may use log/log.Fatal as valid demonstration code' golden-config/.golangci.yml && \
	awk 'BEGIN{s=0}/^linters:/{s=1}s==1&&/## you may want to enable/{s=2;sub(/## you may want to enable/,"## you may want to enable (auto-enabled by update-golden-config)");print;next}s==2&&/## disabled/{s=1;print;next}s==2&&/^    #- /{sub(/^    #- /,"    - ")}{print}' golden-config/.golangci.yml > golden-config/.golangci.yml.tmp && mv golden-config/.golangci.yml.tmp golden-config/.golangci.yml && \
	echo "Vendored $${LATEST_TAG} to golden-config/.golangci.yml"

crosscheck: ## Run golden config cross-check on all guide Good examples
	go run ./cmd/crosscheck/

crosscheck-clean: ## Remove extracted crosscheck files
	rm -rf tmp/crosscheck/

install-plugin: ## Install opencode plugin + shared module to local .opencode/
	@mkdir -p .opencode/plugins .opencode/shared
	cp plugins/golangci-lint.js .opencode/plugins/golangci-lint.js
	cp shared/nudge.js .opencode/shared/nudge.js
	@echo "✓ Plugin + shared module installed to .opencode/"

verify-plugin: install-plugin ## Install plugin and verify it loads
	@node -e "import('./.opencode/plugins/golangci-lint.js').then(m => { \
		if (typeof m.GolangciLintPlugin !== 'function') { console.error('ERROR: GolangciLintPlugin is not a function'); process.exit(1); } \
		if (typeof m.isGolangciLintCommand !== 'function') { console.error('ERROR: isGolangciLintCommand not exported'); process.exit(1); } \
		console.log('✓ Plugin loads and exports are correct'); \
	})"

verify-shared: ## Verify shared nudge module loads and exports correctly
	@node -e "var m = require('./shared/nudge.js'); \
		if (typeof m.isGolangciLintCommand !== 'function') { console.error('ERROR: isGolangciLintCommand missing'); process.exit(1); } \
		if (typeof m.parseDiagnostics !== 'function') { console.error('ERROR: parseDiagnostics missing'); process.exit(1); } \
		if (typeof m.truncateNudge !== 'function') { console.error('ERROR: truncateNudge missing'); process.exit(1); } \
		console.log('✓ Shared nudge module loads and exports are correct')"

install-commands: ## Install golangci-lint command to opencode and crush
	@mkdir -p $(HOME)/.config/opencode/commands $(HOME)/.config/crush/commands
	@cp commands/golangci-lint.md $(HOME)/.config/opencode/commands/golangci-lint.md
	@cp commands/golangci-lint.md $(HOME)/.config/crush/commands/golangci-lint.md
	@echo "✓ Commands installed to opencode and crush"

install-rules: ## Install all platform rules files
	@mkdir -p .claude/rules .cursor/rules .opencode/rules
	@cp rules/claude-code.md .claude/rules/golangci-lint.md
	@cp rules/cursor.mdc .cursor/rules/golangci-lint.mdc
	@cp rules/opencode.md .opencode/rules/golangci-lint.md
	@echo "✓ Rules installed to Claude Code, Cursor, and opencode"

install-hook: ## Install Claude Code PostToolUse hook
	@mkdir -p .claude/hooks
	@cp hooks/golangci-lint-post.js .claude/hooks/golangci-lint-post.js
	@echo "✓ Claude Code hook installed"

install-claude-shared: ## Install shared nudge module for Claude Code hook
	@mkdir -p .claude/shared
	@cp shared/nudge.js .claude/shared/nudge.js
	@echo "✓ Shared nudge module installed to .claude/shared/"

install-agent: ## Deploy pre-publish agent to .opencode/agents/
	@mkdir -p .opencode/agents
	cp agents/pre-publish.md .opencode/agents/pre-publish.md
	@echo "✓ Agent installed to .opencode/agents/"

install-all: ## Install all opencode resources (commands, rules, hook, shared, agents)
	@$(MAKE) install-commands
	@$(MAKE) install-rules
	@$(MAKE) install-hook
	@$(MAKE) install-claude-shared
	@$(MAKE) install-plugin
	@$(MAKE) install-skill
	@$(MAKE) install-agent
	@echo "✓ All resources installed"

sync-deployed: ## Sync all source files to deployed locations (.claude/, .cursor/, .opencode/)
	@# Rules files
	@mkdir -p .claude/rules .cursor/rules .opencode/rules
	@cp rules/claude-code.md .claude/rules/golangci-lint.md
	@cp rules/cursor.mdc .cursor/rules/golangci-lint.mdc
	@cp rules/opencode.md .opencode/rules/golangci-lint.md
	@# Claude hook + shared module
	@mkdir -p .claude/hooks .claude/shared
	@cp hooks/golangci-lint-post.js .claude/hooks/golangci-lint-post.js
	@cp shared/nudge.js .claude/shared/nudge.js
	@# opencode plugin + shared module
	@mkdir -p .opencode/plugins .opencode/shared
	@cp plugins/golangci-lint.js .opencode/plugins/golangci-lint.js
	@cp shared/nudge.js .opencode/shared/nudge.js
	@echo "✓ All deployed files synced from source"

sync-version: ## Update package.json version from git tag
	@V=$$(git describe --tags --always 2>/dev/null | sed 's/^v//'); \
	sed -i '0,/"version": "[^"]*"/s//"version": "'$$V'"/' package.json && \
	echo "✓ package.json version updated to $$V"

lint-js: ## Run ESLint on JavaScript source files
	npx eslint plugins/ shared/ hooks/ bin/install.js
