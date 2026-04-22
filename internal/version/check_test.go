package version

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        Version
		wantErr     bool
		errContains string
	}{
		{
			name:        "full output format",
			input:       "golangci-lint has version 2.11.1 built with go1.26.1 from 89a46a24 on 2026-03-06T14:04:16Z",
			want:        Version{Major: 2, Minor: 11, Patch: 1},
			wantErr:     false,
			errContains: "",
		},
		{
			name:        "short version string",
			input:       "2.11.1",
			want:        Version{Major: 2, Minor: 11, Patch: 1},
			wantErr:     false,
			errContains: "",
		},
		{
			name:        "v1 era version",
			input:       "1.64.5",
			want:        Version{Major: 1, Minor: 64, Patch: 5},
			wantErr:     false,
			errContains: "",
		},
		{
			name:        "version with v prefix",
			input:       "v2.0.0",
			want:        Version{Major: 2, Minor: 0, Patch: 0},
			wantErr:     false,
			errContains: "",
		},
		{
			name:        "empty string",
			input:       "",
			want:        Version{Major: 0, Minor: 0, Patch: 0},
			wantErr:     true,
			errContains: "",
		},
		{
			name:        "no version pattern",
			input:       "some random text without version",
			want:        Version{Major: 0, Minor: 0, Patch: 0},
			wantErr:     true,
			errContains: "",
		},
		{
			name:        "JSON output",
			input:       `{"version":"2.11.1","goVersion":"go1.26.1","commit":"89a46a24","date":"2026-03-06T14:04:16Z"}`,
			want:        Version{Major: 2, Minor: 11, Patch: 1},
			wantErr:     false,
			errContains: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := ParseVersion(testCase.input)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ParseVersion(%q) error = %v, wantErr %v", testCase.input, err, testCase.wantErr)
				return
			}
			if !testCase.wantErr && got != testCase.want {
				t.Errorf("ParseVersion(%q) = %+v, want %+v", testCase.input, got, testCase.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		left  Version
		right Version
		want  int
	}{
		{
			name:  "equal versions",
			left:  Version{Major: 2, Minor: 0, Patch: 0},
			right: Version{Major: 2, Minor: 0, Patch: 0},
			want:  0,
		},
		{
			name:  "older major",
			left:  Version{Major: 1, Minor: 64, Patch: 5},
			right: Version{Major: 2, Minor: 0, Patch: 0},
			want:  -1,
		},
		{
			name:  "newer major",
			left:  Version{Major: 3, Minor: 0, Patch: 0},
			right: Version{Major: 2, Minor: 0, Patch: 0},
			want:  1,
		},
		{
			name:  "older minor",
			left:  Version{Major: 2, Minor: 5, Patch: 0},
			right: Version{Major: 2, Minor: 11, Patch: 0},
			want:  -1,
		},
		{
			name:  "newer minor",
			left:  Version{Major: 2, Minor: 15, Patch: 0},
			right: Version{Major: 2, Minor: 11, Patch: 0},
			want:  1,
		},
		{
			name:  "older patch",
			left:  Version{Major: 2, Minor: 11, Patch: 0},
			right: Version{Major: 2, Minor: 11, Patch: 1},
			want:  -1,
		},
		{
			name:  "newer patch",
			left:  Version{Major: 2, Minor: 11, Patch: 2},
			right: Version{Major: 2, Minor: 11, Patch: 1},
			want:  1,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := Compare(testCase.left, testCase.right)
			if got != testCase.want {
				t.Errorf("Compare(%+v, %+v) = %d, want %d", testCase.left, testCase.right, got, testCase.want)
			}
		})
	}
}

func newTestChecker(execFn func() (string, error), logOutput *bytes.Buffer) *checker {
	checkerInstance := &checker{
		execFn:    execFn,
		logOutput: nil,
	}
	if logOutput != nil {
		checkerInstance.logOutput = logOutput
	}
	return checkerInstance
}

func TestCheck(t *testing.T) {
	t.Run("golangci-lint not found", func(_ *testing.T) {
		checkerInstance := newTestChecker(func() (string, error) {
			return "", errors.New("executable file not found")
		}, nil)
		checkerInstance.check()
	})

	t.Run("v1 incompatible version", func(_ *testing.T) {
		checkerInstance := newTestChecker(func() (string, error) {
			return "golangci-lint has version 1.64.5 built with go1.21.0", nil
		}, nil)
		checkerInstance.check()
	})

	t.Run("acceptable version", func(_ *testing.T) {
		checkerInstance := newTestChecker(func() (string, error) {
			return "golangci-lint has version 2.0.0 built with go1.23.0", nil
		}, nil)
		checkerInstance.check()
	})

	t.Run("significantly newer version", func(_ *testing.T) {
		checkerInstance := newTestChecker(func() (string, error) {
			return "golangci-lint has version 2.11.1 built with go1.26.1", nil
		}, nil)
		checkerInstance.check()
	})

	t.Run("unparseable output", func(_ *testing.T) {
		checkerInstance := newTestChecker(func() (string, error) {
			return "something went wrong", nil
		}, nil)
		checkerInstance.check()
	})

	t.Run("v1 era incompatible version", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "1.64.5", nil
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if !strings.Contains(output, "incompatible") {
			t.Errorf("expected 'incompatible' warning for v1.x, got: %s", output)
		}
	})
}

func TestCheckLogs(t *testing.T) {
	t.Run("too old logs incompatible warning", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "1.49.0", nil
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if !strings.Contains(output, "incompatible") {
			t.Errorf("expected 'incompatible' warning, got: %s", output)
		}
	})

	t.Run("not found logs warning", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "", errors.New("not found")
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if !strings.Contains(output, "not found in PATH") {
			t.Errorf("expected 'not found in PATH' warning, got: %s", output)
		}
	})

	t.Run("acceptable version no warning", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "2.5.0", nil
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if output != "" {
			t.Errorf("expected no warning for acceptable version, got: %s", output)
		}
	})

	t.Run("significantly newer logs warning", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "2.11.1", nil
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if !strings.Contains(output, "significantly newer") {
			t.Errorf("expected 'significantly newer' warning, got: %s", output)
		}
	})

	t.Run("unparseable logs warning", func(t *testing.T) {
		var buf bytes.Buffer
		checkerInstance := newTestChecker(func() (string, error) {
			return "no-version-here", nil
		}, &buf)
		checkerInstance.check()

		output := buf.String()
		if !strings.Contains(output, "could not parse") {
			t.Errorf("expected 'could not parse' warning, got: %s", output)
		}
	})
}
