# Stage 1: Node.js builder — produces skill/plugin/rules/command artifacts (throwaway)
FROM node:22-bookworm-slim AS node-installer
WORKDIR /workspace
COPY build/golangci-lint-guide.tgz /tmp/
RUN npm install -g /tmp/golangci-lint-guide.tgz && \
    npx golangci-lint-guide --platforms=opencode --plugin-scope=project && \
    rm -f /tmp/golangci-lint-guide.tgz

# Stage 2: Final image — Go toolchain only, NO Node.js runtime
FROM golang:1.26-bookworm

# Layer 2: golangci-lint v2.11.4 (rarely changes)
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v2.11.4

# Layer 3: Copy binaries from build context (changes per build)
COPY build/opencode /usr/local/bin/opencode
COPY build/golangci-lint-mcp /usr/local/bin/golangci-lint-mcp
RUN chmod +x /usr/local/bin/opencode /usr/local/bin/golangci-lint-mcp

# Layer 4: opencode MCP config (generated inline — file is gitignored)
RUN mkdir -p /root/.config/opencode && \
    printf '{\n  "$schema": "https://opencode.ai/config.json",\n  "mcp": {\n    "golangci-lint": {\n      "type": "local",\n      "command": ["golangci-lint-mcp"]\n    }\n  }\n}\n' > /root/.config/opencode/opencode.json

# Layer 5: Copy skill/plugin/rules from Node builder stage (no Node.js runtime needed)
COPY --from=node-installer /workspace/.opencode /workspace/.opencode
COPY --from=node-installer /root/.agents /root/.agents
COPY --from=node-installer /root/.config/opencode/commands /root/.config/opencode/commands

# Layer 6: Runtime directories + verification
WORKDIR /workspace
RUN mkdir -p /root/.local/share/opencode /workspace/testdata
RUN golangci-lint version && opencode --version || true
