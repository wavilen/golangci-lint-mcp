#!/usr/bin/env bash
# install-lint.sh — Download and install golangci-lint v2.11.4
# Installs to ./bin/golangci-lint relative to project root.
# Skips download if the correct version is already installed.
set -euo pipefail

VERSION="v2.11.4"
INSTALL_DIR="$(git rev-parse --show-toplevel)/bin"
BINARY="${INSTALL_DIR}/golangci-lint"

# --- Skip if correct version already installed ---
if [ -x "$BINARY" ]; then
    INSTALLED=$("$BINARY" version 2>&1 | grep -oP 'version \K[0-9]+\.[0-9]+\.[0-9]+' || true)
    if [ "$INSTALLED" = "${VERSION#v}" ]; then
        echo "✓ golangci-lint ${VERSION} already installed"
        exit 0
    fi
fi

# --- Detect OS and architecture ---
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
    linux)  OS="linux" ;;
    darwin) OS="darwin" ;;
    *) echo "Error: Unsupported OS: $OS" >&2; exit 1 ;;
esac

case "$ARCH" in
    x86_64)  ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Error: Unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

# --- Download and install ---
URL="https://github.com/golangci/golangci-lint/releases/download/${VERSION}/golangci-lint-${VERSION#v}-${OS}-${ARCH}.tar.gz"

echo "Downloading golangci-lint ${VERSION} for ${OS}-${ARCH}..."
TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

curl -fsSL "$URL" | tar -xz -C "$TMPDIR" --strip-components=1

mkdir -p "$INSTALL_DIR"
mv "$TMPDIR/golangci-lint" "$BINARY"
chmod +x "$BINARY"

echo "✓ golangci-lint ${VERSION} installed to ${BINARY}"
