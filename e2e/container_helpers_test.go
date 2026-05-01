package e2e_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
