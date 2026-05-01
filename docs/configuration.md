# MCP Client Configuration

[← Back to README](../README.md)

Configure golangci-lint-mcp as an MCP server in your client. The basic config is just the command name — see [Quick Start](../README.md#quick-start) in the README.

## opencode

Add to your project's `opencode.json`:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "golangci-lint": {
      "type": "local",
      "command": ["golangci-lint-mcp"]
    }
  }
}
```

The plugin automatically injects `--output.json.path stdout` into any `golangci-lint` command and strips conflicting output format flags (e.g., `--output.text.*`, `--out-format`, `--verbose`, `--show-stats`) that would break JSON parsing. No manual flag management needed.

If the binary is not in PATH, use the full path:

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "golangci-lint": {
      "type": "local",
      "command": ["/path/to/golangci-lint-mcp"]
    }
  }
}
```

## Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "golangci-lint": {
      "command": "golangci-lint-mcp"
    }
  }
}
```

## Cursor

Add to `.cursor/mcp.json` in your project root:

```json
{
  "mcpServers": {
    "golangci-lint": {
      "command": "golangci-lint-mcp"
    }
  }
}
```

## gosec AI Autofix (optional)

To enable gosec AI autofix, add the flag and configure the required environment variables:

**opencode:**

```json
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {
    "golangci-lint": {
      "type": "local",
      "command": ["golangci-lint-mcp", "--gosec-ai"],
      "env": {
        "GOSEC_AI_API_PROVIDER": "gemini-2.0-flash",
        "GOSEC_AI_API_KEY": "your-api-key-here"
      }
    }
  }
}
```

**Claude Desktop / Cursor:**

```json
{
  "mcpServers": {
    "golangci-lint": {
      "command": "golangci-lint-mcp",
      "args": ["--gosec-ai"],
      "env": {
        "GOSEC_AI_API_PROVIDER": "gemini-2.0-flash",
        "GOSEC_AI_API_KEY": "your-api-key-here"
      }
    }
  }
}
```

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `GOSEC_AI_API_KEY` | Yes | API key for the AI provider. The `gosec_ai_autofix` tool is only registered when this is set. |
| `GOSEC_AI_API_PROVIDER` | No | AI provider/model (default: `gemini-2.0-flash`). Options: `gemini-2.0-flash`, `claude-sonnet-4-0`, `gpt-4o`, or a custom model name. |
| `GOSEC_AI_BASE_URL` | No | Custom base URL for the AI provider API endpoint. |
| `GOSEC_AI_SKIP_SSL` | No | Set to `"true"` to skip SSL verification for the AI provider connection. |

### How It Works

The API key is passed directly to the gosec subprocess by the MCP server — it is **never exposed in tool responses**. The `gosec_ai_autofix` tool is only available when both `--gosec-ai` and `GOSEC_AI_API_KEY` are configured. When enabled, gosec guide responses include an `<autofix>` section pointing to the `gosec_ai_autofix` MCP tool instead of hardcoded CLI commands.
