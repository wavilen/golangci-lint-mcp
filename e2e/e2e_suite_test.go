//go:build e2e

package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/moby/moby/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	testContainer testcontainers.Container
	containerCtx  context.Context
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Integration Suite")
}

var _ = BeforeSuite(func() {
	containerCtx = context.Background()

	// Clean tmp/ndjson/ to ensure fresh state on re-runs
	// (files are kept after suite ends for post-analysis)
	ndjsonDir := filepath.Join("..", "tmp", "ndjson")
	err := os.RemoveAll(ndjsonDir)
	if err != nil {
		fmt.Fprintf(GinkgoWriter, "Warning: failed to clean %s: %v\n", ndjsonDir, err)
	}

	// Verify Docker image exists (built by `make integration-test`)
	cmd := exec.Command("docker", "inspect", "golangci-lint-mcp-e2e")
	err = cmd.Run()
	if err != nil {
		Skip("Docker image golangci-lint-mcp-e2e not found. Run: make integration-test")
	}

	// Resolve absolute path for testdata (Docker bind mounts require absolute host paths)
	testdataAbs, absErr := filepath.Abs("testdata")
	Expect(absErr).ToNot(HaveOccurred())

	// Start long-lived container (tail -f keeps it running)
	binds := []string{
		testdataAbs + ":/workspace/testdata",
	}

	// Mount opencode auth.json if available (required for LLM API access)
	authPath := filepath.Join(os.Getenv("HOME"), ".local", "share", "opencode", "auth.json")
	_, statErr := os.Stat(authPath)
	if statErr == nil {
		binds = append(binds, authPath+":/root/.local/share/opencode/auth.json")
	}

	testContainer, err = testcontainers.Run(containerCtx, "golangci-lint-mcp-e2e",
		testcontainers.WithCmd("sh", "-c", "echo 'container-ready' && tail -f /dev/null"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("container-ready").WithStartupTimeout(10*time.Second),
		),
		testcontainers.WithHostConfigModifier(func(hc *container.HostConfig) {
			hc.Binds = binds
		}),
	)
	Expect(err).ToNot(HaveOccurred(), "Failed to start test container")

	DeferCleanup(func() {
		testcontainers.TerminateContainer(testContainer)
	})
})
