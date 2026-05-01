//go:build e2e

package e2e

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var sharedContainer testcontainers.Container

func createTestContainer(ctx context.Context) func() {
	testdataAbs, err := filepath.Abs("testdata")
	if err != nil {
		panic(fmt.Sprintf("resolve testdata path: %v", err))
	}

	binds := []string{
		testdataAbs + ":/workspace/testdata",
	}

	authPath := filepath.Join(os.Getenv("HOME"), ".local", "share", "opencode", "auth.json")
	_, statErr := os.Stat(authPath)
	if statErr == nil {
		binds = append(binds, authPath+":/root/.local/share/opencode/auth.json")
	}

	c, err := testcontainers.Run(ctx, "golangci-lint-mcp-e2e",
		testcontainers.WithCmd("sh", "-c", "echo 'container-ready' && tail -f /dev/null"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("container-ready").WithStartupTimeout(10*time.Second),
		),
		testcontainers.WithHostConfigModifier(func(hc *container.HostConfig) {
			hc.Binds = binds
		}),
	)
	if err != nil {
		panic(fmt.Sprintf("create container: %v", err))
	}

	sharedContainer = c
	cleanup := func() {
		testcontainers.TerminateContainer(c)
		sharedContainer = nil
	}

	return cleanup
}

func containerExec(ctx context.Context, cmd []string) (int, string, error) {
	if sharedContainer == nil {
		return -1, "", errors.New("no active container")
	}

	exitCode, reader, err := sharedContainer.Exec(ctx, cmd)
	if err != nil {
		return -1, "", fmt.Errorf("exec: %w", err)
	}

	var stdout, stderr bytes.Buffer
	_, copyErr := stdcopy.StdCopy(&stdout, &stderr, reader)
	if copyErr != nil {
		return exitCode, stdout.String() + stderr.String(), fmt.Errorf("demux: %w", copyErr)
	}

	return exitCode, stdout.String(), nil
}

func containerExportToFile(ctx context.Context, cmd []string) (string, error) {
	if sharedContainer == nil {
		return "", errors.New("no active container")
	}

	tmpPath := "/tmp/gsd-export-tmp.json"
	shellCmd := []string{"sh", "-c", fmt.Sprintf("%s > %s", joinCmd(cmd), tmpPath)}

	exitCode, _, err := sharedContainer.Exec(ctx, shellCmd)
	if err != nil {
		return "", fmt.Errorf("exec to file: %w", err)
	}
	if exitCode != 0 {
		return "", fmt.Errorf("export command exited with code %d", exitCode)
	}

	reader, copyErr := sharedContainer.CopyFileFromContainer(ctx, tmpPath)
	if copyErr != nil {
		return "", fmt.Errorf("copy from container: %w", copyErr)
	}
	defer reader.Close()

	var buf bytes.Buffer
	if _, ioErr := io.Copy(&buf, reader); ioErr != nil {
		return "", fmt.Errorf("read exported file: %w", ioErr)
	}

	return buf.String(), nil
}

func joinCmd(cmd []string) string {
	escaped := make([]string, len(cmd))
	for i, c := range cmd {
		escaped[i] = shellescape(c)
	}
	return joinWithSpaces(escaped)
}

func shellescape(s string) string {
	if s == "" {
		return "''"
	}
	needsQuote := false
	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c == '/') {
			needsQuote = true
			break
		}
	}
	if !needsQuote {
		return s
	}
	return "'" + string([]rune(s)) + "'"
}

func joinWithSpaces(parts []string) string {
	result := parts[0]
	for _, p := range parts[1:] {
		result += " " + p
	}
	return result
}
