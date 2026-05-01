# Installation

[← Back to README](../README.md)

## Install the MCP Server

```bash
go install github.com/wavilen/golangci-lint-mcp@latest
```

This places the `golangci-lint-mcp` binary in `$GOPATH/bin`. Make sure your Go bin directory is in your `PATH`. The version is derived automatically from the git tag via Go's built-in VCS info.

## Build from Source

```bash
git clone <repo-url>
cd golangci-lint-mcp
make install
```

This installs with the exact version from `git describe --tags` injected via ldflags. Use `make build` to build locally without installing.

## Install the OpenCode Skill

**One-command install (recommended):**

```bash
npx @wavilen/golangci-lint-guide
```

**Or install globally:**

```bash
npm install -g @wavilen/golangci-lint-guide
golangci-lint-guide
```

**Or from source:**

```bash
make install-skill
```

This copies the golangci-lint-guide skill to `~/.agents/skills/golangci-lint-guide/`, making it available in any Go project opened with opencode.

## Compatibility

This server ships guides validated against **golangci-lint v2.0+**.

**golangci-lint v1.x is incompatible** — it uses completely different CLI flags and will not work with this server.

At startup, the server checks your installed golangci-lint version and logs a warning if:
- The version is below v2.0 (incompatible — wrong CLI flags)
- The version is significantly newer (6+ minor versions ahead) — some linters may have changed behavior compared to when the guides were written

The version check is non-blocking — the server starts normally regardless of the result. Warnings appear in stderr logs, visible in MCP client debug output.

## Versioning

The server reports its own version at startup and to MCP clients. The version is derived from git tags:

- **`go install @latest`** — version comes from Go's built-in VCS info (`vcs.tag` build setting)
- **`make install`** — version injected via ldflags from `git describe --tags`
- **Development builds** — falls back to commit hash or `"dev"`

To sync `package.json` with the latest git tag:

```bash
make sync-version
```
