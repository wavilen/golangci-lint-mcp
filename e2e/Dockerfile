# Layer 1: Go toolchain (rarely changes)
FROM golang:1.26-bookworm

# Layer 2: golangci-lint v2.11.4 (rarely changes)
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v2.11.4

# Layer 3: Copy binaries from build context (changes per build)
# The Makefile copies opencode and golangci-lint-mcp into e2e/build/ before docker build
COPY build/opencode /usr/local/bin/opencode
COPY build/golangci-lint-mcp /usr/local/bin/golangci-lint-mcp
RUN chmod +x /usr/local/bin/opencode /usr/local/bin/golangci-lint-mcp

# Layer 4: opencode MCP config (static per test run)
RUN mkdir -p /root/.config/opencode
COPY opencode.json /root/.config/opencode/opencode.json

# Layer 4b: opencode skills and rules (enables golangci-lint-guide skill + golang rules)
RUN mkdir -p /root/.agents/skills/golangci-lint-guide /root/.config/opencode/rules
COPY build/skills/golangci-lint-guide/SKILL.md /root/.agents/skills/golangci-lint-guide/SKILL.md
COPY build/rules/golang.md /root/.config/opencode/rules/golang.md
COPY build/rules/golangci-lint.md /root/.config/opencode/rules/golangci-lint.md

# Layer 5: opencode auth + testdata mounted at runtime via testcontainers
RUN mkdir -p /root/.local/share/opencode /workspace/testdata

# Verify installations
RUN golangci-lint version && opencode --version || true

WORKDIR /workspace
