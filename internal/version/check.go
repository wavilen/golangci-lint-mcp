package version

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

const ReferenceVersion = "2.0.0"
const MinVersion = "2.0.0"
const MaxMinorDrift = 6

type Version struct {
	Major, Minor, Patch int
}

var semverRe = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)

const semverMatchCount = 4

func ParseVersion(output string) (Version, error) {
	matches := semverRe.FindStringSubmatch(output)
	if len(matches) != semverMatchCount {
		return Version{}, fmt.Errorf("no semver found in %q", output)
	}
	major, err := strconv.Atoi(matches[1])
	if err != nil {
		return Version{}, fmt.Errorf("invalid major version %q: %w", matches[1], err)
	}
	minor, err := strconv.Atoi(matches[2])
	if err != nil {
		return Version{}, fmt.Errorf("invalid minor version %q: %w", matches[2], err)
	}
	patch, err := strconv.Atoi(matches[3])
	if err != nil {
		return Version{}, fmt.Errorf("invalid patch version %q: %w", matches[3], err)
	}
	return Version{Major: major, Minor: minor, Patch: patch}, nil
}

func Compare(left, right Version) int {
	if left.Major != right.Major {
		if left.Major < right.Major {
			return -1
		}
		return 1
	}
	if left.Minor != right.Minor {
		if left.Minor < right.Minor {
			return -1
		}
		return 1
	}
	if left.Patch != right.Patch {
		if left.Patch < right.Patch {
			return -1
		}
		return 1
	}
	return 0
}

type checker struct {
	logOutput io.Writer
	execFn    func() (string, error)
}

func newDefaultChecker() *checker {
	return &checker{
		logOutput: nil,
		execFn: func() (string, error) {
			out, err := exec.CommandContext(context.Background(), "golangci-lint", "version").Output()
			if err != nil {
				return "", fmt.Errorf("running golangci-lint version: %w", err)
			}
			return string(out), nil
		},
	}
}

func (checkerInstance *checker) check() {
	var logger *log.Logger
	if checkerInstance.logOutput != nil {
		logger = log.New(checkerInstance.logOutput, "", 0)
	} else {
		logger = log.New(log.Writer(), "warning: ", 0)
	}

	output, err := checkerInstance.execFn()
	if err != nil {
		logger.Printf("golangci-lint not found in PATH, version check skipped")
		return
	}

	installed, err := ParseVersion(output)
	if err != nil {
		logger.Printf("could not parse golangci-lint version from: %s", output)
		return
	}

	minV, _ := ParseVersion(MinVersion)
	refV, _ := ParseVersion(ReferenceVersion)

	if Compare(installed, minV) < 0 {
		logger.Printf(
			"golangci-lint version %d.%d.%d is incompatible, minimum required is %s — v1.x uses different CLI flags and will not work with this server",
			installed.Major,
			installed.Minor,
			installed.Patch,
			MinVersion,
		)
		return
	}

	if installed.Major == refV.Major && installed.Minor > refV.Minor+MaxMinorDrift {
		logger.Printf(
			"golangci-lint version %d.%d.%d is significantly newer than the validated version %s — some linters may have changed behavior",
			installed.Major,
			installed.Minor,
			installed.Patch,
			ReferenceVersion,
		)
	}
}

func Check() {
	newDefaultChecker().check()
}
