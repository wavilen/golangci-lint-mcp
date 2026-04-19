.PHONY: build install install-skill test clean npm-pack npm-publish

BINARY := golangci-lint-mcp
SKILL_SRC := skills/golangci-lint-guide/SKILL.md
SKILL_DEST := $(HOME)/.agents/skills/golangci-lint-guide

build:
	go build -o $(BINARY) .

install:
	go install .

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
